package filter

// CallbackInt ... Iterate a []int
type CallbackInt func(element int, index int, arr []int) bool

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
