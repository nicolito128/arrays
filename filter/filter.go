package filter

// CallbackInt ... Iterate a []int
type CallbackInt func(element int, index int, arr []int) bool

// CallbackFloat ... Iterate a []float64
type CallbackFloat func(element float64, index int, arr []float64) bool

// CallbackByte ... Iterate a []byte
type CallbackByte func(element byte, index int, arr []byte) bool

// CallbackString ... Iterate a []string
type CallbackString func(element string, index int, arr []string) bool

// Int ... Evaluates a boolean to enter an element in a new []int
func Int(arr []int, callback CallbackInt) []int {
	var filteredArr []int
	var element int
	var condition bool

	for index := range arr {
		element = arr[index]
		condition = callback(element, index, arr)

		if condition {
			filteredArr = append(filteredArr, element)
		}
	}

	return filteredArr
}

// Float ... Evaluates a boolean to enter an element in a new []float64
func Float(arr []float64, callback CallbackFloat) []float64 {
	var filteredArr []float64
	var element float64
	var condition bool

	for index := range arr {
		element = arr[index]
		condition = callback(element, index, arr)

		if condition {
			filteredArr = append(filteredArr, element)
		}
	}

	return filteredArr
}

// Byte ... Evaluates a boolean to enter an element in a new []byte
func Byte(arr []byte, callback CallbackByte) []byte {
	var filteredArr []byte
	var element byte
	var condition bool

	for index := range arr {
		element = arr[index]
		condition = callback(element, index, arr)

		if condition {
			filteredArr = append(filteredArr, element)
		}
	}

	return filteredArr
}

// String ... Evaluates a boolean to enter an element in a new []string
func String(arr []string, callback CallbackString) []string {
	var filteredArr []string
	var element string
	var condition bool

	for index := range arr {
		element = arr[index]
		condition = callback(element, index, arr)

		if condition {
			filteredArr = append(filteredArr, element)
		}
	}

	return filteredArr
}
