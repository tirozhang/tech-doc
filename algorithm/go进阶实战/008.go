package main

import "fmt"

type student struct {
	Name string
}

func zhoujielun(v interface{}) {
	switch msg := v.(type) {
	case student, *student:
		//fmt.Println(msg.Name) error
		fmt.Println(msg)
	}

	switch msg := v.(type) {
	case *student:
		fmt.Println(msg.Name)
	}

}
