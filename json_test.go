package trier_test

import (
	"encoding/json"
	"testing"

	"github.com/boundedinfinity/go-trier"
	"github.com/stretchr/testify/assert"
)

func Test_Try_Succeded_serialization_success(t *testing.T) {
	input := trier.Success("success")
	expected := []byte(`{"result":"success","error":null}`)
	actual, err := json.Marshal(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Try_Succeded_serialization_failure(t *testing.T) {
	input := trier.Failuref[string]("an error")
	expected := []byte(`{"result":"","error":"an error"}`)
	actual, err := json.Marshal(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
