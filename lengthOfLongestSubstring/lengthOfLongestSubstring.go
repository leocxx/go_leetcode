func lengthOfLongestSubstring(s string) int {
	i := 0
	max := 0
	a := []rune(s)
	for m, c := range a {
		for n := i; n &lt; m; n++ {
			if a[n] == c {
				i = n + 1
			}
		}
		if m-i+1 &gt; max {
			max = m - i + 1
		}
	}
	return max
}