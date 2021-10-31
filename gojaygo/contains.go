package gojaygo

func StringArrayContains(strArray []string, toFind string) bool {
	for _, a := range strArray {
		if a == toFind {
			return true
		}
	}

	return false
}

func IntArrayContains(intArray []int, toFind int) bool {
	for _, a := range intArray {
		if a == toFind {
			return true
		}
	}

	return false
}

func Float32ArrayContain(floatArray []float32, toFind float32) bool {
	for _, a := range floatArray {
		if a == toFind {
			return true
		}
	}

	return false
}

func Float64ArrayContains(floatArray []float64, toFind float64) bool {
	for _, a := range floatArray {
		if a == toFind {
			return true
		}
	}

	return false
}
