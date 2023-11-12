package dp

func MinFlipsMonoIncr(s string) int {

	res, flip := 0, 0

	for i := 0; i <= len(s)-1; i++ {
		if s[i] == '0' {
			res = Min(res+1, flip)
		} else {
			flip++
		}
	}

	return res
}

/*

10011111110010111011
flip := 8
000000000000

010110
i = 5
otoi = 1,2
itoo =

010101011011
i = 9
otoi = 1,1,2,3,4
itoo = 1,2,3,4,
*/
