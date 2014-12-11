package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	inputData := "Input data"
	encodedValue := encode(inputData)

	assert.True(t, true, isBase64(encodedValue))
}

func TestDecode(t *testing.T) {
	inputData := "Some data"
	encodedValue := encode(inputData)

	decodedValue := decode(encodedValue)

	assert.False(t, isBase64(decodedValue))
	assert.Equal(t, inputData, decodedValue)
}
