package dp

func NumberOfWays1(str string) int {
	dp := map[string]int{
		"0":   0,
		"1":   0,
		"01":  0,
		"10":  0,
		"010": 0,
		"101": 0,
	}

	for i := 0; i < len(str); i++ {
		if str[i] == '0' {
			dp["0"]++
			dp["10"] += dp["1"]
			dp["010"] += dp["01"]
		}
		if str[i] == '1' {
			dp["1"]++
			dp["01"] += dp["0"]
			dp["101"] += dp["10"]
		}
	}

	return dp["010"] + dp["101"]
}
