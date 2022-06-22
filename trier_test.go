package trier_test

import (
	"errors"
	"testing"

	"github.com/boundedinfinity/go-trier"
	"github.com/stretchr/testify/assert"
)

func Test_Try_Success(t *testing.T) {
	actual := trier.Success("success")

	assert.True(t, actual.Success())
	assert.False(t, actual.Failure())
	assert.Nil(t, actual.Err)
	assert.Equal(t, actual.Result, "success")
}

func Test_Try_Failure(t *testing.T) {
	actual := trier.Failure[string](errors.New("failure"))

	assert.False(t, actual.Success())
	assert.True(t, actual.Failure())
	assert.NotNil(t, actual.Err)
	assert.Equal(t, actual.Err.Error(), "failure")
	assert.Equal(t, actual.Result, "")
}

func Test_Try_Failuref(t *testing.T) {
	actual := trier.Failuref[string]("%v", "failure")

	assert.False(t, actual.Success())
	assert.True(t, actual.Failure())
	assert.NotNil(t, actual.Err)
	assert.Equal(t, actual.Err.Error(), "failure")
	assert.Equal(t, actual.Result, "")
}
