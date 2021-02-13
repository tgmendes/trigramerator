package trigram

import (
	"log"
	"regexp"
	"strings"
)

var store map[string]map[string]int

// Learn generates new trigram maps for a given text.
func Learn(text string) {
	if store == nil {
		store = map[string]map[string]int{}
	}

	words := strings.Split(strings.ToLower(text), " ")
	for i := 0; i < len(words)-2; i++ {
		trigramKey := cleanText(strings.Join(words[i:i+2], " "))

		if val, ok := store[trigramKey]; ok {
			if newVal, ok := val[words[i+2]]; ok {
				newVal++
			} else {
				val[words[i+2]] = 1
			}
		} else {
			newMapVal := map[string]int{
				words[i+2]: 1,
			}
			store[trigramKey] = newMapVal
		}
	}
}

// GetStore returns the trigram store
func GetStore() map[string]map[string]int {
	return store
}

// cleanText will remove special characters from a given text.
// Trigram keys should not include special characters for simplicity of search.
func cleanText(text string) string {
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^A-Za-zÀ-ÿ0-9? ]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(text, "")
}
