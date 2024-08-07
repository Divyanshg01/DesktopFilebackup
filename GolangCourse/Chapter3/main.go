package main

import (
	"fmt"
	"structs"
)

func main() {
	// var x [3]int; // initialized to zero value
	// var x = [3]int {1,2,3}
	// var x = [12]int {1,5:4,6,10:100,15} // {1,0,0,0,0,4,6,0,0,0,100,15}
	// 	var x = [...]int{2, 3, 4}
	// // array are comparable like x==y

	// 	// Slices
	// 	var y = []int{1, 2, 3}
	// 	var y1 = []int{1, 3: 2, 10: 0}
	// 	var x1 [][]int
	// 	var x2 []int // nil
	//     // a slice is not comparable , you can only do x == nil
	//
	//     //to check
	//     a := slices.Equal(x1,x2) //true or false

	// Append Slices

	// var x = []int{1, 2, 3}
	// x = append(x, 4, 5)
	// var y = []int{6, 7, 8}
	// x = append(x, y...)
	// fmt.Println(x)

	// Make Function

	// x := make([]int, 5, 10)
	// data type , len , cap

	// x = append(x, 3, 6)
	// here the length of x is 5 so the append with place 3,6 at index 5,6 respectively
	//you can also define a slice of 0 length and a capacity

	// Clearning slice
	// clear(x) // initializes all the elements to 0 , but length remains the same

	// []int{} this is not a nill slice where as this is []int

	//  Slicing Slices

	// x := []string{"a", "b"}
	//
	// y := x[:2]
	// z := x[:]

	// full slice expression
	// x := y[1:2:3]
	// starting , ending , position of the orignal slice upto which modification
	// so capacity of x is 3-1 = 2

	// Copy of the slices

	// x := []int{1, 2, 3, 4}
	// y := make([]int, 4)
	// num := copy(y, x)
	// fmt.Println(y, num)
	// copy function returns no. of elements copied

	//convert array to slice
	// x := arr[:]

	// convert slice to array
	// xSlice := []int{1, 2, 3, 4}
	// xArray := [4]int(xSlice)
	// smallArray := [2]int(xSlice)
	// xSlice[0] = 10

	//     Go also has a type that represents a single code point. The rune type is an alias for
	// the int32 type, just as byte is an alias for uint8. As you could probably guess, a rune
	// literal’s default type is a rune, and a string literal’s default type is a string.

	//string are made up of bytes(not runes)

	// strings are immutable
	// len(s) ==> returns the length in no. of bytes

	// String , rune , byte conversion
	// var a rune = 'x'
	// var s string = string(a)
	// var b byte = 'y'
	// var s2 string = string(b)
	//
	//     var s string = "Hello, "
	// var bs []byte = []byte(s)
	// var rs []rune = []rune(s)
	// fmt.Println(bs)
	// fmt.Println(rs)
	//

	// Maps
	// var nilmaps map[string]int // nil map
	//
	// totalWinds := map[string]int{}
	//
	// teams := map[string][]string{
	// 	"Orcas":   []string{"Fred", "Ralph", "Bijou"},
	// 	"Lions":   []string{"Sarah", "Peter", "Billie"},
	// 	"Kittens": []string{"Waldo", "Raul", "Ze"},
	// }

	//     There are some restrictions on the types of the keys that I’ll discuss in a bit.
	// If you know how many key-value pairs you intend to put in the map but don’t know
	// the exact values, you can use make to create a map with a default size:
	// ages := make(map[int][]string, 10)

	// Comma , ok idiom

	// m := map[string]int{
	//     "hello" : 5,
	//     "world" : 0,
	// }
	//
	// v, ok := m["hello"]
	// fmt.Println(v , ok)
	//
	// delete(m , "hello")
	//
	// clear(m)
	//
	// maps.Equal(m , n)

	// Struct

	// type person struct {
	// 	name string
	// 	age  int
	// 	pet  string
	// }
	//
 //    var fred person;
	//
 //    bob := person{}
	//
 //    julia := person{
 //        age : 30,
 //        name : "Beth",
 //    }
	//
 //    // julia.age
	//
 //    //Anonymous structs
	//
 //    pet := struct{
 //        name string
 //        kind string
 //    }{
 //        name : "Fido",
 //        kind : "dog",
 //    }


//     type firstPerson struct {
// name string
// age int
// }
// you can use a type conversion to convert an instance of firstPerson to
// secondPerson, but you can’t use == to compare an instance of firstPerson and
// an instance of secondPerson, because they are different types:
// type secondPerson struct {
// name string
// age int
// }


}
