# idios

Go package for extracting "common" or "uncommon" features from a piece of text, given a training corpus. This tool is hypothetically language-agnostic, and is efficient to run on test documents after "training".

## install

Install the Porter2 implementation with:
`$ go get -u github.com/dchest/stemmer/porter2`

Then, while in your project directory:
`$ git clone https://github.com/haydenhigg/idios`

Finally, import it as:
```go
import "./idios"
```

## use

### modelling

- `StemExample(text string) []string`: Tokenizes a string.
- `NewModel(text string) Idios`: Returns an Idios struct with the probability per unique token of the training string `text`.

### extracting

- `(model Idios) Common(example string) []string`: Extracts the relatively-common words (***not*** stemmed tokens, but the original words from the input string) from `example`.
- `(model Idios) Uncommon(example string) []string`: Extracts the relatively-uncommon words (i.e. the words that contribute the most to the semantics of the input string, according to their probabilities in the training corpus) from `example`.

### example

```go
package main

import (
  "fmt"
  "./idios"
)

func main() {
  trainingCorpus := "This is preferably a long text that contains plenty of words..."

  model := idios.NewModel(trainingCorpus)
  
  fmt.Println(model.Common("And this is a test text, from which the commmon words will be extracted"))
  fmt.Println(model.Uncommon("This is another test text, but only the uncommon words will be returned"))
}
```
