package romantoint

func romanToInt(s string) int {
	m := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}
	var ret int
	for i := 0; i < len(s); {
		// 先尝试读两个字符，注意索引不要越界
		if i+2 <= len(s) && m[s[i:i+2]] != 0 {
			ret += m[s[i:i+2]]
			i += 2
			continue
		} else {
			ret += m[s[i:i+1]]
			i++
			continue
		}
	}
	return ret
}
