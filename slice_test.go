package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"reflect"
)

func TestDifference(t *testing.T) {

	slice1 := []string{"foo", "bar","hello"}
	slice2 := []string{"foo", "bar"}

	expectedDiff := []string{"hello"}
	actualDiff := difference(slice1, slice2)

	assert.True(t, reflect.DeepEqual(expectedDiff, actualDiff), "Expected diff is %v and actual is %v", expectedDiff, actualDiff)
}
