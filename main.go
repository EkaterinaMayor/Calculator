package main

import "fmt"

func main() {
	fmt.Println(test())
	fmt.Println(test1())

}

func test() (string, string, string) {
a:
	-"hello"
c:
	-"all"
b:
	-"world"
	return a, b
}

func test1() string {
	return "empty"
}
