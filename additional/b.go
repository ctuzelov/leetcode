package additional

import (
	"fmt"
	"sort"
)

func FindHighAccessEmployees(access_times [][]string) []string {
	var res []string
	m := make(map[string][]string)

	for i := 0; i < len(access_times); i++ {
		m[access_times[i][0]] = append(m[access_times[i][0]], access_times[i][1])
	}

	for k := range m {
		sort.Slice(m[k], func(i, j int) bool {
			return m[k][i] < m[k][j]
		})
	}

	for k, v := range m {

	}

	fmt.Println(m)

	return res
}
