func isValid(s string) bool {
	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	a := make([]byte, 0, len(s)/2)
	for _, b := range []byte(s) {
		if b == '(' || b == '{' || b == '[' {
			a = append(a, b)
			continue
		}
		if b == ')' || b == '}' || b == ']' {
			if len(a) > 0 && m[b] == a[len(a)-1] {
				a = a[:len(a)-1]
				continue
			} else {
				return false
			}
		}
	}
	if len(a) == 0 {
		return true
	} else {
		return false
	}
}
