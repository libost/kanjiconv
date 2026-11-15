# kanjiconv

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
		CleanNN:     true,
	}
	japaneseText := "こんにちは"
	romanizedText := kanjiconv.ConvertJapaneseToRoman(t, japaneseText, opts)
	fmt.Println(romanizedText) // Output: konnichiwa
}
```