package utils

func StringSliceContains(sl []string, str string) bool {
	for _, item := range sl {
		if str == item {
			return true
		}
	}
	return false
}
