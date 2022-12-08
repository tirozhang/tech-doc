package main

import "fmt"

func main() {
	res, ok := reverString("abcd")
	fmt.Println(res)
	fmt.Println(ok)

}

func reverString(s string) (string, bool) {

	r := []rune(s)
	l := len(r)

	if l > 5000 {
		return s, false
	}
	for i := 0; i < l/2; i++ {
		r[i], r[l-i-1] = r[l-i-1], r[i]
	}

	//start, end := 0, len(r)-1
	//for start < end {
	//	tmp := r[start]
	//	r[start] = r[end]
	//	r[end] = tmp
	//	start++
	//	end--
	//}

	return string(r), true
}
