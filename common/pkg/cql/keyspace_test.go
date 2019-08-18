package cql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewKeyspaceBinder(t *testing.T) {
	assert := assert.New(t)
	kb := NewKeyspaceBinder("test")
	assert.NotNil(kb)
}

func TestKeyspaceBinderTable(t *testing.T) {
	assert := assert.New(t)
	kb := NewKeyspaceBinder("test")
	assert.NotNil(kb)
	actual := kb.Table("example_table")
	expected := "test.example_table"
	assert.EqualValues(expected, actual)
}
