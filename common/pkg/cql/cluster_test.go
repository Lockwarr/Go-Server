package cql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCluster(t *testing.T) {
	assert := assert.New(t)
	session := NewCluster("", "denislav", "0.0.0.0")
	assert.NotNil(session)
}
