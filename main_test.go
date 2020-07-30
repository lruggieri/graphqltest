package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessQuery(t *testing.T){
	assert.Equal(t, `{"data":{"jobs":[{"id":1},{"id":2},{"id":3},{"id":4},{"id":5}]}}`,processQuery("{jobs{id}}"))
	assert.Equal(t, `{"data":{"companies":[{"id":1},{"id":2},{"id":3},{"id":4}]}}`,processQuery("{companies{id}}"))
	assert.Equal(t, `{"data":{"people":[{"id":1},{"id":2},{"id":3},{"id":4}]}}`,processQuery("{people{id}}"))

	assert.Equal(t, `{"data":{"job":{"position":"Software Engineer"}}}`,processQuery("{job(id:1){position}}"))

	assert.Equal(t, `{"data":{"jobs":[{"company":{"id":1},"id":1},{"company":{"id":2},"id":2},{"company":{"id":3},"id":3},{"company":{"id":4},"id":4},{"company":{"id":4},"id":5}]}}`,processQuery("{jobs{id, company{id}}}"))
}
