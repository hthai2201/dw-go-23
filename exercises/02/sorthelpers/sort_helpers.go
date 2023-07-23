package sorthelpers

import (
	"sort"
	"strconv"
)

func SortStrings(data []string){
	sort.Slice(data, func(i, j int) bool {
		num1, err1 := strconv.ParseFloat(data[i], 64)
		num2, err2 := strconv.ParseFloat(data[j], 64)
		if err1 == nil && err2 == nil {
			return num1 < num2
		}

		return data[i] < data[j]
	})
}
