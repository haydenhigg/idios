# bengal

Easy-to-use Go implementation of the multinomial Naive Bayes for text classification, optimized for multiple output features.

## install

Install the Porter2 implementation with:
`$ go get -u github.com/dchest/stemmer/porter2`

Then, while in your project directory:
`$ git clone https://github.com/haydenhigg/bengal`

Finally, import it as:
```go
import "./bengal"
```

## use

### modelling

- `StemExample(text string) []string`: Tokenizes a string.
- `NewModel(examples []string, output [][]string) MultinomialNB`: Tokenizes the inputs using `StemExamples` and creates a model from them.
- `NewModelFromVectors(input [][]string, output [][]string) MultinomialNB`: Creates a model from tokenized inputs.
- `(model MultinomialNB) Predict(example string) []string`: Predicts the classes of the input example using `StemExample`.
- `(model MultinomialNB) PredictVector(input []string) []string`: Predicts the classes of the tokenized input.

### extracting

-

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