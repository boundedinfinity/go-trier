package trier_test

import (
	"errors"
	"testing"

	"github.com/boundedinfinity/go-trier"
	"github.com/stretchr/testify/assert"
)

func Test_Try_Succeded(t *testing.T) {
	actual := trier.Success("success")

	assert.True(t, actual.Succeded())
	assert.False(t, actual.Failed())
	assert.Nil(t, actual.Error)
	assert.Equal(t, actual.Result, "success")
}

func Test_Try_Failure(t *testing.T) {
	actual := trier.Failure[string](errors.New("failure"))

	assert.False(t, actual.Succeded())
	assert.True(t, actual.Failed())
	assert.NotNil(t, actual.Error)
	assert.Equal(t, actual.Error.Error(), "failure")
	assert.Equal(t, actual.Result, "")
}

func Test_Try_Failuref(t *testing.T) {
	actual := trier.Failuref[string]("%v", "failure")

	assert.False(t, actual.Succeded())
	assert.True(t, actual.Failed())
	assert.NotNil(t, actual.Error)
	assert.Equal(t, actual.Error.Error(), "failure")
	assert.Equal(t, actual.Result, "")
}
