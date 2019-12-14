// test_git project main.go
package main

import (
	"fmt"
	"reflect"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {

	var c Celsius = 50
	var f Fahrenheit = 100
	var t float64 = 12
	fmt.Println(reflect.TypeOf(t))
	fmt.Println(reflect.TypeOf(Celsius(t)))

	fmt.Println(FtoC(f))
	fmt.Println(CtoF(c))

	/*
				short variable declaration을 이용한 multi 선언
				shortDec1, shortDec2 := 'a', 3
				shortDec1, shortDec3 := 'b', 4
				fmt.Printf("%c %d %d\n", shortDec1, shortDec2, shortDec3)

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

func CtoF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FtoC(f Fahrenheit) Celsius { return (Celsius(f-32) * 5 / 9) }
