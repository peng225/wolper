package dictionary

import (
	"math"
	"sort"
	"strings"

	"github.com/peng225/wolper/util"
)

type wordAndEntropy struct {
	word    string
	entropy float64
}

type Color int

const (
	WHITE Color = iota
	GREEN
	YELLOW
	GRAY
)

func convToBucketIndex(colorPattern []Color) int {
	index := 0
	for _, color := range colorPattern {
		index *= 3
		switch color {
		case GREEN:
			index += 2
		case YELLOW:
			index += 1
		}
	}
	return index
}

func getBucketIndex(target, answer string) int {
	if len(target) != len(answer) {
		panic("len(target) != len(answer)")
	}

	colorPattern := make([]Color, 5)
	// Search for GREEN positions
	for i := 0; i < len(target); i++ {
		if target[i] == answer[i] {
			colorPattern[i] = GREEN
		}
	}

	// Get remained characters in the answer string
	remained := ""
	for i := 0; i < len(target); i++ {
		if colorPattern[i] != GREEN {
			remained += string(answer[i])
		}
	}

	// Search for YELLOW and GRAY positions
	for i := 0; i < len(target); i++ {
		if colorPattern[i] == GREEN {
			continue
		}

		if strings.Contains(remained, string(target[i])) {
			colorPattern[i] = YELLOW
			remained = string(util.DeleteChar([]byte(remained), byte(target[i])))
		} else {
			colorPattern[i] = GRAY
		}
	}

	// Calc the index value
	return convToBucketIndex(colorPattern)
}

func getEntropy(target string, words []string) float64 {
	bucket := make(map[int]int, 0)
	for _, word := range words {
		index := getBucketIndex(target, word)
		if _, ok := bucket[index]; ok {
			bucket[index]++
		} else {
			bucket[index] = 1
		}
	}

	entropy := 0.0
	for _, value := range bucket {
		p := float64(value) / float64(len(words))
		if p != 0 {
			entropy -= p * math.Log2(p)
		}
	}
	return entropy
}

func sortWithEntropy(words []string) {
	wae := make([]wordAndEntropy, len(words))

	for i, word := range words {
		wae[i] = wordAndEntropy{word, getEntropy(word, words)}
	}

	sort.Slice(wae, func(i, j int) bool {
		return wae[i].entropy > wae[j].entropy ||
			(wae[i].entropy == wae[j].entropy && wae[i].word < wae[j].word)
	})

	for i, val := range wae {
		words[i] = val.word
	}
}
