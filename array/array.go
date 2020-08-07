/*
Copyright (c) 2020 ypankaj007

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package array

// Uniq
func Uniq(array []interface{}) []interface{} {
	var result []interface{}
	for _, v := range array {
		if Contains(result, 0, func(val interface{}) bool {
			return v == val
		}) {
			continue
		}
		result = append(result, v)
	}
	return result
}

func Intersection(array []interface{}, elems ...[]interface{}) []interface{} {
	var result []interface{}
	for _, v := range array {
		if Contains(result, 0, func(val interface{}) bool {
			return v == val
		}) {
			continue
		}
		var j int
		for j = 0; j < len(elems); j++ {
			if !Contains(elems[j], 0, func(val interface{}) bool {
				return v == val
			}) {
				break
			}
		}
		if j+1 == len(elems) {
			result = append(result, v)
		}
	}
	return result
}

func Union(elems ...[]interface{}) []interface{} {
	rest := flatten(elems)
	return Uniq(rest)
}

func Difference(array []interface{}, elems ...[]interface{}) []interface{} {
	rest := flatten(elems)
	return Filter(array, func(v1 interface{}) bool {
		return !Contains(rest, 0, func(v2 interface{}) bool {
			return v1 == v2
		})
	})
}

func Contains(array []interface{}, fromIndex int, predicate func(interface{}) bool) bool {
	return IndexOf(array, fromIndex, predicate) >= 0
}

func IndexOf(array []interface{}, fromIndex int, predicate func(interface{}) bool) int {

	if len(array) == 0 {
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

	var result [][]interface{}
	if count < 1 {
		return nil
	}
	i, l := 0, len(array)

	for i < l {
		lastIdx := func() int {
			if i+count > l {
				return l - 1
			}
			return i + count - 1
		}
		result = append(result, array[i:lastIdx()])
		i += count
	}
	return result
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
