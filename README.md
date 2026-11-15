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
	japaneseText := "こんにちは"
	romanizedText := kanjiconv.ConvertJapaneseToRoman(t, japaneseText)
	fmt.Println(romanizedText) // Output: konnichiwa
}
```