package ds_hw_0

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strings"
)

func topWords(path string, numWords int, charThreshold int) []WordCount {
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	wordCounts := make(map[string]int)
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")

	for _, word := range strings.Fields(string(fileBytes)) {
		word = strings.ToLower(word)
		word = reg.ReplaceAllString(word, "")
		if len(word) >= charThreshold {
			wordCounts[word]++
		}
	}

	var wc []WordCount
	for k, v := range wordCounts {
		wc = append(wc, WordCount{k, v})
	}
	sortWordCounts(wc)

	if len(wc) < numWords {
		numWords = len(wc)
	}

	return wc[:numWords]
}

type WordCount struct {
	Word  string
	Count int
}

func (wc WordCount) String() string {
	return fmt.Sprintf("%v: %v", wc.Word, wc.Count)
}

func sortWordCounts(wordCounts []WordCount) {
	sort.Slice(wordCounts, func(i, j int) bool {
		wc1 := wordCounts[i]
		wc2 := wordCounts[j]
		if wc1.Count == wc2.Count {
			return wc1.Word < wc2.Word
		}
		return wc1.Count > wc2.Count
	})
}
