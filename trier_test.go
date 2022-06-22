package trier_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/boundedinfinity/go-trier"
	"github.com/stretchr/testify/assert"
)

func Test_Try_Success(t *testing.T) {
	actual := trier.Success("success")

	assert.True(t, actual.Success())
	assert.False(t, actual.Failure())
	assert.Nil(t, actual.Error)
	assert.Equal(t, actual.Result, "success")
}

func Test_Try_Success_serialization_success(t *testing.T) {
	input := trier.Success("success")
	expected := []byte(`{"result":"success","error":null}`)
	actual, err := json.Marshal(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Try_Success_serialization_failure(t *testing.T) {
	input := trier.Failuref[string]("an error")
	expected := []byte(`{"result":"","error":"an error"}`)
	actual, err := json.Marshal(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Try_Failure(t *testing.T) {
	actual := trier.Failure[string](errors.New("failure"))

	assert.False(t, actual.Success())
	assert.True(t, actual.Failure())
	assert.NotNil(t, actual.Error)
	assert.Equal(t, actual.Error.Error(), "failure")
	assert.Equal(t, actual.Result, "")
}

func Test_Try_Failuref(t *testing.T) {
	actual := trier.Failuref[string]("%v", "failure")

	assert.False(t, actual.Success())
	assert.True(t, actual.Failure())
	assert.NotNil(t, actual.Error)
	assert.Equal(t, actual.Error.Error(), "failure")
	assert.Equal(t, actual.Result, "")
}
