package cql

import (
	"bytes"
	"sync"
)

// KeyspaceBinder - simple thread safe reusable buffer for string generation
type KeyspaceBinder struct {
	defaultKeyspace    *bytes.Buffer
	defaultkeyspaceLen int
	dynamicKeyspace    *bytes.Buffer
	delimiter          string
	m                  *sync.Mutex
}

// NewKeyspaceBinder - new instance of keyspaceBinder
func NewKeyspaceBinder(keyspace string) *KeyspaceBinder {
	buffer := bytes.NewBufferString(keyspace)
	if !bytes.ContainsAny(buffer.Bytes(), ".") {
		buffer.WriteString(".")
	}
	return &KeyspaceBinder{defaultKeyspace: buffer, defaultkeyspaceLen: buffer.Len(), dynamicKeyspace: bytes.NewBuffer(make([]byte, 0, 50)), delimiter: ".", m: new(sync.Mutex)}
}

// Table - returns keyspace.table_name string
func (k *KeyspaceBinder) Table(table string) string {
	k.m.Lock()
	k.defaultKeyspace.WriteString(table)
	str := k.defaultKeyspace.String()
	k.defaultKeyspace.Truncate(k.defaultkeyspaceLen)
	k.m.Unlock()
	return str
}
