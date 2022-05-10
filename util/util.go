package util

func DeleteChar(s []byte, c byte) []byte {
	foundIndex := -1
	for i, v := range s {
		if v == c {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		return s
	} else {
		s[foundIndex] = s[len(s)-1]
		return s[:len(s)-1]
	}
}
