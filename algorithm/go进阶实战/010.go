package main

import "fmt"

type People2 struct {
	Name string `json:"name"`
}

// 循环调用
func (p *People2) String() string {
	return fmt.Sprintf("print: %v", p)
}

func main() {
	p := &People2{"zhangsan"}
	p.String()
}
