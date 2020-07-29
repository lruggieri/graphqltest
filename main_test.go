package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFailCI(t *testing.T){
	assert.Equal(t, 3, 1+1)
}
