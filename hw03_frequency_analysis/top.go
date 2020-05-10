package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"sort"
	"strings"
	"sync"
)

type wordDict struct {
	Word      string
	Frequency int
}

func Top10(inpText string) (result []string) {
	const numberOfWords int = 10

	var dictMutex sync.RWMutex
	freqDict := map[string]int{}
	splitedText := strings.Fields(inpText)
	for _, v := range splitedText {
		dictMutex.Lock()
		freqDict[v]++
		dictMutex.Unlock()
	}

	wd := make([]wordDict, 0)
	for k, v := range freqDict {
		wd = append(wd, wordDict{k, v})
	}

	sort.Slice(wd, func(i, j int) bool {
		return wd[i].Frequency > wd[j].Frequency
	})

	if len(wd) > numberOfWords {
		wd = wd[:numberOfWords]
	}

	for _, v := range wd {
		result = append(result, v.Word)
	}

	return
}

// func main() {
// 	text1 := "раз три два три два три"

// 	fmt.Println(Top10(text1))
// }
