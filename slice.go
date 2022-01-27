package common

func InSlice(s []string, target string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] == target || s[j] == target {
			return true
		}
	}
	return false
}
