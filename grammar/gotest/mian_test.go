package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Log("test start")
	result := add(1, 2)
	assert.Equal(t, result, 3)
}

//go test ./... -v 运行测试
