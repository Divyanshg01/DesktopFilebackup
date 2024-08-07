Type of Literal

1. Integer Literal => base 10 by default , 0b , 0o , 0x
   you can also write numbers like this 1_234_42
2. Floating point literal => 23.323e32
3. Rune literal => represents a character surrounded by single quotes , single Unicode characters
   there are also several escaped rune literal like '\n'
4. String literals => most of the time use double quotes to create an interpreted string literal like "Hello there"
   these contain zero or more rune literal

int/uint8->64
byte is an alias of uint8
int will be defined as int32 or 64 depending on the machine
It is a compile time error to comapre int with int32

float32 or float64

comlex-> float32
complex128-> float64

var x ,y = 10 ,20
var x,y int = 10 ,20
var (
x int
y = 20
z int = 2
d , y = 40 , "hello"
)
x := 12
b := 'A' // Equivalent to byte(65)
fmt.Printf("%T %d\n", b, b) // Output: uint8 65

Rune
r := 'ä' // Represents the Unicode character 'ä'
fmt.Printf("%T %d\n", r, r) // Output: int32 228
