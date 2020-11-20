/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package idios

func unique(arr []string) []string {
	uniqueWords := make(map[string]struct{})

	for _, k := range arr {
		if _, ok := uniqueWords[k]; !ok {
			uniqueWords[k] = struct{}{}
		}
	}

	keys := make([]string, len(uniqueWords))
	i := 0

	for k := range uniqueWords {
		keys[i] = k
		i++
	}

	return keys
}

func count(arr []string, val string) int {
	ret := 0

	for _, v := range arr {
		if v == val {
			ret++
		}
	}

	return ret
}

func mean(arr []float64) float64 {
	ret := float64(0)

	for _, v := range arr {
		ret += v
	}

	return ret / float64(len(arr))
}