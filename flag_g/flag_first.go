package main

import "fmt"
import "flag"

var userParams = flag.String("worker", "", "worker name")

func main() {
	flag.Parse()
	var worker = map[string]func(){
		"rick": func() {
			fmt.Println("This is good man")
		},
		"guoml": func() {
			fmt.Println("Good choice")
		},
		"hh": func() {
			fmt.Println("Bad")
		},
	}
	f, ok := worker[*userParams]
	if ok {
		f()
	} else {
		fmt.Println("params error")
	}

}
