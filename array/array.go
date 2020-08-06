package array

type Function func(interface{}) bool

func Uniq(array []interface{}) []interface{} {
	return nil
}

func Intersection() []interface{} {
	return nil
}

func Union(elems ...[]interface{}) []interface{} {
	rest := flatten(elems)
	return Uniq(rest)
}

func Difference(array []interface{}, elems ...[]interface{}) []interface{} {
	var result []interface{}
	rest := flatten(elems)
	for _, v := range array {
		if !Contains(rest, 0, func(e interface{}) bool {
			return e == v
		}) {
			result = append(result, v)
		}
	}
	return nil
}

func Contains(array []interface{}, fromIndex int, predicate func(interface{}) bool) bool {
	return IndexOf(array, fromIndex, predicate) >= 0
}

func IndexOf(array []interface{}, fromIndex int, predicate func(interface{}) bool) int {

	if len(array) == 0 || predicate == nil {
		return -1
	}

	if fromIndex < 0 {
		fromIndex = 0
	}

	for ; fromIndex < len(array); fromIndex++ {
		if predicate(array[fromIndex]) {
			return fromIndex
		}
	}
	return -1

}

func Chunk(array []interface{}, count int) [][]interface{} {
	return nil
}

func Without(array []interface{}, otherArray []interface{}) []interface{} {
	return Difference(array, otherArray)
}

func Filter(array []interface{}, predicate func(interface{}) bool) []interface{} {
	var result []interface{}
	if array != nil && len(array) > 0 {
		for _, v := range array {
			if predicate(v) {
				result = append(result, v)
			}
		}
	}
	return result
}

func flatten(array [][]interface{}) []interface{} {

	var result []interface{}

	for _, row := range array {
		for _, col := range row {
			result = append(result, col)
		}
	}
	return result
}
