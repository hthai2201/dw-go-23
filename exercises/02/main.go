// main.go

package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/hthai2201/dw-go-23/exercises/02/sorthelpers"
)



func main() {
	var isNumber bool
	var isString bool
	var isMix bool
	flag.BoolVar(&isNumber, "int", false, "Sort as numbers")
	flag.BoolVar(&isString, "string", false, "Sort as strings")
	flag.BoolVar(&isMix, "mix", false, "Sort as mixed array")
	flag.Parse()
	data := flag.Args()
 switch  {
 case isNumber:
	sorthelpers.SortStrings((data))
 case isString:
	sort.Strings(data)
 case isMix:
	sorthelpers.SortStrings(data)
 default:
	sorthelpers.SortStrings(data)
 }
 
	fmt.Println("Output: ",data)
}
