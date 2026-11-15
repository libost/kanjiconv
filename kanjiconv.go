// Package kanjiconv provides functions to convert Japanese text to Romanized form (Romaji).
package kanjiconv

import (
	"regexp"
	"strings"

	"github.com/gojp/kana"
	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

// ConvertOptions defines options for the ConvertJapaneseToRoman function.
type ConvertOptions struct {
	// WithSpace indicates whether to include spaces between words in the output.
	WithSpace bool
	// CleanNN indicates whether to clean up double 'n' sounds in the output.
	CleanNN bool
}

// InitTokenizer initializes and returns a new tokenizer instance using the IPA dictionary.
//
// If the tokenizer cannot be initialized, this function will panic.
//
// Use this function to create a tokenizer instance before calling ConvertJapaneseToRoman.
func InitTokenizer() *tokenizer.Tokenizer {
	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		panic("Failed to initialize tokenizer: " + err.Error())
	}
	return t
}

// ConvertJapaneseToRoman converts a Japanese string to its Romanized form using the provided tokenizer.
//
// Use InitTokenizer() to create a tokenizer instance before calling this function.
func ConvertJapaneseToRoman(t *tokenizer.Tokenizer, str string, opts *ConvertOptions) string {
	tokens := t.Tokenize(str)
	var result []string
	for _, token := range tokens {
		pron, ok := token.Pronunciation()
		if !ok {
			pron = token.Surface
		}
		var romaji string
		if opts != nil && opts.CleanNN {
			romaji = toCleanRomaji(pron)
		} else {
			romaji = kana.KanaToRomaji(pron)
		}
		result = append(result, romaji)
	}
	if opts != nil && !opts.WithSpace {
		return strings.Join(result, "")
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
