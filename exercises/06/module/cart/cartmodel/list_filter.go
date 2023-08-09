package cartmodel

import "github.com/hthai2201/dw-go-23/exercises/06/common"

type ListParam struct {
	common.Paging `json:",inline"`
	*ListFilter   `json:",inline"`
}

type ListFilter struct {
}
