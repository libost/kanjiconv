package kanjiconv

import (
	"log"
	"regexp"
	"strings"

	"github.com/gojp/kana"
	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

// InitTokenizer initializes and returns a new tokenizer instance using the IPA dictionary.
func InitTokenizer() *tokenizer.Tokenizer {
	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		log.Fatalf("Error creating tokenizer: %v", err)
	}
	return t
}

// ConvertJapaneseToRoman converts a Japanese string to its Romanized form using the provided tokenizer.
//
// Use InitTokenizer() to create a tokenizer instance before calling this function.
func ConvertJapaneseToRoman(t *tokenizer.Tokenizer, str string) string {
	tokens := t.Tokenize(str)
	var result []string
	for _, token := range tokens {
		pron, ok := token.Pronunciation()
		if !ok {
			pron = token.Surface
		}
		romaji := toCleanRomaji(pron)
		result = append(result, romaji)
	}
	return strings.Join(result, " ")
}

var (
	nnBeforeVowelOrY = regexp.MustCompile(`(^|[^n])nn([aiueoy])`)
	remainingNN      = regexp.MustCompile(`nn`)
)

// toCleanRomaji converts a Japanese string to its Romanized form, cleaning up any double 'n' sounds.
func toCleanRomaji(katakana string) string {
	r := kana.KanaToRomaji(katakana)
	r = nnBeforeVowelOrY.ReplaceAllString(r, "${1}n'${2}")
	r = remainingNN.ReplaceAllString(r, "n")
	return r
}

// stringEqualFoldNoSpaces compares two strings for equality, ignoring case and spaces.
func stringEqualFoldNoSpaces(a, b string) bool {
	return strings.EqualFold(strings.ReplaceAll(a, " ", ""), strings.ReplaceAll(b, " ", ""))
}
