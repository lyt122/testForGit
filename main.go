package main

import "fmt"

func main() {
	var (
		a int
		b int
	)
	scan, err := fmt.Scanln(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(a+b, scan)
}
