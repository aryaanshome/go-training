package main

import (
	"fmt"
	"sort"
	"strings"
)

func wordsextracter() {

	text := "As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far overhead the multiple transparency of huge leaves filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever he felt able he ran again; the ground continued soft and springy, covered with the same resilient weed which was the first thing his hands had touched in Malacandra. Once or twice a small red creature scuttled across his path, but otherwise there seemed to be no life stirring in the wood; nothing to fear—except the fact of wandering unprovisioned and alone in a forest of unknown vegetation thousands or millions of miles beyond the reach or knowledge of man."

	punctuation := []string{",", ".", ";"}

	for i := range punctuation {
		text = strings.ReplaceAll(text, punctuation[i], "")
	}

	text = strings.ToLower(text)

	words := strings.Split(text, " ")

	counter := make(wordcounter)

	for word := range words {
		counter[words[word]]++
	}

	keys_arr := make([]string, 0, len(words))

	for keys := range counter {
		keys_arr = append(keys_arr, keys)
	}

	sort.Strings(keys_arr)

	for _, k := range keys_arr {
		fmt.Printf("%v: %v\n", k, counter[k])
	}

}
