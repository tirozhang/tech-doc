package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	name string `json:"name"`
}

func main() {
	js := `{"name":"zhangsan"}`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		panic(err)
	}
	fmt.Println("people:", p)
}
