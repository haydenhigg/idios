/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package idios

import (
	"strings"
	"regexp"
	"github.com/dchest/stemmer/porter2"
	
	"math"
)

func StemExample(text string) []string {
	unallowed := regexp.MustCompile("[^0-9a-z ]")
	stem := porter2.Stemmer.Stem

	words := strings.Fields(strings.ToLower(strings.Trim(text, "\t\n\f\r ")))
	ret := make([]string, len(words))

	for i, w := range words {
		ret[i] = stem(unallowed.ReplaceAllLiteralString(w, ""))
	}
	
	return ret
}

type Idios struct {
	Vocabulary	map[string]float64
	Threshold 	float64
}

func NewModel(text string) Idios {
	tokens := StemExample(text)
	wordCount := float64(len(tokens))

	vocabulary := make(map[string]float64)
	var probs []float64

	for _, word := range unique(tokens) {
		prob := math.Log1p(float64(count(tokens, word)) / float64(wordCount))

		vocabulary[word] = prob
		probs = append(probs, prob)
	}

	return Idios{
		Vocabulary: vocabulary,
		Threshold: median(probs)}
}

func (model Idios) isCommonFilter(prob float64, ok bool) bool {
	return ok && prob > model.Threshold
}

func (model Idios) isUncommonFilter(prob float64, ok bool) bool {
	return !ok || prob <= model.Threshold
}

func (model Idios) filterTokens(example string, persists func(float64, bool) bool) []string {
	format := regexp.MustCompile(`^[0-9a-zA-Z]+('s){0,1}`)

	tokens, words := StemExample(example), strings.Fields(example)

	var ret []string

	for i, token := range tokens {
		if prob, ok := model.Vocabulary[token]; persists(prob, ok) {
			ret = append(ret, format.FindStringSubmatch(strings.Trim(words[i], "\t\n\f\r "))[0])
		}
	}

	return ret
}

func (model Idios) Common(example string) []string {
	return model.filterTokens(example, model.isCommonFilter)
}

func (model Idios) Uncommon(example string) []string {
	return model.filterTokens(example, model.isUncommonFilter)
}
