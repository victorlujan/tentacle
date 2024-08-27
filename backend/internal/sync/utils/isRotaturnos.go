package utils

func IsRotaturnos(isRotaturnos string) int {
	if isRotaturnos == "true" {
		return 1
	}
	return 0

}

func IsActivo(isInactivo string) int {
	if isInactivo == "true" {
		return 0
	}
	return 1

}
