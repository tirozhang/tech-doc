package main

import "fmt"

type Param map[string]interface{}

type Show struct {
	Param
}

func main() {
	//panic: assignment to entry in nil map
	//s := new(Show)
	//s.Param["RMB"] = 1000

	s := new(Show)
	s = &Show{
		Param{
			"RMB": 1000,
		},
	}
	fmt.Println(s)
}
