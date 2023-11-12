package dp

import (
	"sort"
)

func LongestStrChain(words []string) int {
	m := make(map[string]int)
	mxLen := 1

	sort.SliceStable(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	for _, word := range words {
		m[word] = 1
		for i := range word {
			cmp := word[:i] + word[i+1:]
			if v, ok := m[cmp]; ok {
				m[word] = Max(m[word], v+1)
				mxLen = Max(mxLen, m[word])
			}
		}
	}
	return mxLen
}

/*
"a","b","ba","bca","bda","bdca"

a
b
*/
