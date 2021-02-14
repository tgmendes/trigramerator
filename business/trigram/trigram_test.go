package trigram_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tgmendes/trigramerator/business/trigram"
)

func TestLearnTrigram(t *testing.T) {
	mockDB := new(storerMock)
	mockDB.On("Append", mock.Anything, mock.Anything).Return(nil)
	in := "To be or not to be, that is the question"

	err := trigram.Learn(mockDB, in)
	mockDB.AssertNumberOfCalls(t, "Append", 8)
	mockDB.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestGenerate(t *testing.T) {
	// this is a simple predictable dataset to test our generate
	testData := map[string][]string{
		"to be":        {"or"},
		"be or":        {"not"},
		"or not":       {"to,"},
		"not to":       {"that"},
		"to that":      {"is"},
		"that is":      {"the\n"},
		"is the":       {"question."},
		"the question": {"is"},
		"question is":  {"that"},
	}
	mockDB := new(storerMock)
	mockDB.On("Get", mock.Anything).Return(nil)
	// this guarantees repeatibility by always returning the same "seed" text
	mockDB.On("RandomKey").Return("to be")
	mockDB.data = testData

	text, err := trigram.Generate(mockDB, "", "")

	assert.Equal(t, "To be or not to, that is the\n Question. Is that", text)
	assert.NoError(t, err)
}

type storerMock struct {
	mock.Mock

	data map[string][]string
}

func (m *storerMock) Append(key, value string) {
	m.Called(key, value)
}

func (m *storerMock) Get(key string) []string {
	m.Called(key)
	return m.data[key]
}

func (m *storerMock) RandomKey() string {
	args := m.Called()
	return args.String(0)
}
