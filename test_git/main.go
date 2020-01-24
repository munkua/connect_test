// test_git project main.go
package main

import "fmt"

func main() {

	var A, B int
	n, _ := fmt.Scan(&A, &B)
	fmt.Println(A - B)
	fmt.Println(n)

	/*			short variable declaration을 이용한 multi 선언
	shortDec1, shortDec2 := 'a', 3
	shortDec1, shortDec3 := 'b', 4
	fmt.Printf("%c %d %d\n", shortDec1, shortDec2, shortDec3)
	/*

				shortFunc1, shortFunc2 := severalReturn()
				fmt.Printf("%d %d\n", shortFunc1, shortFunc2)


			var x, y int16
			fmt.Println(&x == &x, &x == &y, &x == nil)

			var p = f()
			fmt.Println(p)

		v := 1
		increment(&v)
		fmt.Println(increment(&v))

		fmt.Println(exported)
	*/

}
