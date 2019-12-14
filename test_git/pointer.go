// pointer
package main

import (
	"fmt"
)

func pointers() {
	var x int16 = 1
	var p *int16 = &x
	fmt.Print("%d", *p)

	fmt.Println(p == &x)
}
