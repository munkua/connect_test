// pointer
package main

import (
	"fmt"
)

func main() {
	var x int16 = 1
	var p *int16 = &x
	fmt.Print("%d", *p)

	fmt.Println(p == &x)
}
