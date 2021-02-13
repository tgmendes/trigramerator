package trigram_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tgmendes/trigramerator/business/trigram"
)

func TestLearnTrigram(t *testing.T) {
	in := "To be or not to be, that is the question"

	expResult := map[string]map[string]int{
		"to be":   {"or": 1, "that": 1},
		"be or":   {"not": 1},
		"or not":  {"to": 1},
		"not to":  {"be,": 1},
		"be that": {"is": 1},
		"that is": {"the": 1},
		"is the":  {"question": 1},
	}

	trigram.Learn(in)

	assert.Equal(t, expResult, trigram.GetStore())
}
