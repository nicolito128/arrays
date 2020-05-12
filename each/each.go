package each

// HandlerInt ... A callback for []int
type HandlerInt func(element int, index int, array []int) int

// HandlerFloat ... A callback for []float64
type HandlerFloat func(element float64, index int, array []float64) float64

// HandlerString ... A callback for []string
type HandlerString func(element string, index int, array []string) string

// HandlerByte ... A callback for []byte
type HandlerByte func(element byte, index int, array []byte) byte

// Int ... Iterate a []int
func Int(arr []int, handler HandlerInt) []int {
	var newArr []int
	var element int

	for index := range arr {
		element = arr[index]
		newArr = append(newArr, handler(element, index, arr))
	}

	return newArr
}

// Float ... Iterate a []float64
func Float(arr []float64, handler HandlerFloat) []float64 {
	var newArr []float64
	var element float64

	for index := range arr {
		element = arr[index]
		newArr = append(newArr, handler(element, index, arr))
	}

	return newArr
}

// String ... Iterate a []string
func String(arr []string, handler HandlerString) []string {
	var newArr []string
	var element string

	for index := range arr {
		element = arr[index]
		newArr = append(newArr, handler(element, index, arr))
	}

	return newArr
}

// Byte ... Iterate a []byte
func Byte(arr []byte, handler HandlerByte) []byte {
	var newArr []byte
	var element byte

	for index := range arr {
		element = arr[index]
		newArr = append(newArr, handler(element, index, arr))
	}

	return newArr
}
