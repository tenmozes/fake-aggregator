package server

import (
	"sort"
)

type numberResponse struct {
	Numbers []int `json:"numbers"`
}

func (n *numberResponse) sort()  {
	sort.Ints(n.Numbers)
}

func (n *numberResponse) unique()  {
	if len(n.Numbers) <= 1 {
		return
	}
	unique := []int{n.Numbers[0]}
	for i := 1; i < len(n.Numbers); i++ {
		if unique[len(unique)-1] != n.Numbers[i] {
			unique = append(unique, n.Numbers[i])
		}
	}
	n.Numbers = unique
}

func (n *numberResponse) Compile()  {
	n.sort()
	n.unique()
}
