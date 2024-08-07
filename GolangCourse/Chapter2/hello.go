package main

import "fmt"

const x int64 = 19
const (
	idKey   = "idKey"
	nameKey = "name"
)
const z = 20 * 10

func main() {
	var flag bool //default is set to true
	var isAwsome = true
	var x = complex(2.32, 232.2)
	var y rune = 'J'
	fmt.Print(x, isAwsome, y)
	fmt.Print(flag, isAwsome)

}
