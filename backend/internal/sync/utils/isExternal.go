package utils

func IsExternal(external string) int {
	if external == "true" {
		return 1
	}
	return 0

}
