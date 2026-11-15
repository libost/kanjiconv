# kanjiconv
[![Go Reference](https://pkg.go.dev/badge/github.com/libost/kanjiconv.svg)](https://pkg.go.dev/github.com/libost/kanjiconv)  

A Go library for converting Japanese text to Romanized form.

## Example Usage

```bash
go get github.com/libost/kanjiconv
```

```go
package main

import (
	"fmt"
	"github.com/libost/kanjiconv"
)

func main() {
	t := kanjiconv.InitTokenizer()
	opts := &kanjiconv.ConvertOptions{
		WithSpace: true,
		CleanNN:   true,
	}
	japaneseText := "Hello, 世界! こんにちは"
	romanizedText := kanjiconv.ConvertJapaneseToRoman(t, japaneseText, opts)
	fmt.Println(romanizedText) // Output: Hello, sekai! konnichiwa
}
```