package errorer_test

import (
	"encoding/json"
	"testing"

	"github.com/boundedinfinity/go-trier/errorer"
	"github.com/stretchr/testify/assert"
)

func Test_JSON_serialize_error(t *testing.T) {
	input := errorer.Errorf("this is an error")
	expected := []byte(`"this is an error"`)
	actual, err := json.Marshal(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_JSON_deserialize_error(t *testing.T) {
	input := []byte(`"this is an error"`)
	expected := errorer.Errorf("this is an error")
	var actual errorer.Error

	err := json.Unmarshal(input, &actual)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_JSON_deserialize_nil(t *testing.T) {
	input := []byte(`null`)
	expected := errorer.None()
	var actual errorer.Error

	err := json.Unmarshal(input, &actual)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
