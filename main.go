package main

import "fmt"

func main() {
	var (
		a int
		b int
	)
	scan, err := fmt.Scanf("%d%d", &a, &b)
	if err != nil {
		println(scan)
		return
	} else {
		println(a + b)
	}

}
