package kanjiconv

import (
	"testing"
)

func TestConvertJapaneseToRoman(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "Test 1",
			str:  "君はロックを聴かない",
			want: "Kimi wa rokku wo kika nai",
		},
		{
			name: "Test 2",
			str:  "愛は勝つ",
			want: "Ai wa katsu",
		},
		{
			name: "Test 3",
			str:  "ハナミズキ",
			want: "Hanamizuki",
		},
		{
			name: "Test 4",
			str:  "Hello, World! こんにちは",
			want: "Hello, World! konnichiwa",
		},
		{
			name: "Test 5",
			str:  "カンイ",
			want: "Kan'i",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertJapaneseToRoman(InitTokenizer(), tt.str)
			if !stringEqualFoldNoSpaces(got, tt.want) {
				t.Errorf("ConvertJapaneseToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringEqualFoldNoSpaces(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want bool
	}{
		{
			name: "Test 1",
			a:    "Hello, World!",
			b:    "hello,world!",
			want: true,
		},
		{
			name: "Test 2",
			a:    "こんにちは",
			b:    "こんにちは",
			want: true,
		},
		{
			name: "Test 3",
			a:    "Kimi wa rokku wo kika nai",
			b:    "Kimiwarokkuwokikanai",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stringEqualFoldNoSpaces(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("stringEqualFoldNoSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToCleanRomaji(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "n before a vowel",
			str:  "ンア",
			want: "n'a",
		},
		{
			name: "n before ne",
			str:  "ンネ",
			want: "nne",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toCleanRomaji(tt.str)
			if !stringEqualFoldNoSpaces(got, tt.want) {
				t.Errorf("toCleanRomaji() = %v, want %v", got, tt.want)
			}
		})
	}
}
