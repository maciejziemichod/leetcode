package main

// https://leetcode.com/problems/isomorphic-strings/
func isIsomorphic(s string, t string) bool {
	m, k := make(map[byte]byte), make(map[byte]struct{})

	for i := 0; i < len(s); i++ {
		a, b := s[i], t[i]

		v, ok1 := m[a]
		_, ok2 := k[b]

		if !ok1 && ok2 {
			return false
		}

		if !ok1 {
			m[a] = b
			k[b] = struct{}{}
			continue
		}

		if ok1 && v != b {
			return false
		}
	}

	return true
}
