package rpackage

func ContainsString(list []string, val string) bool {
	for _, str := range list {
		if val == str {
			return true
		}
	}
	return false
}
