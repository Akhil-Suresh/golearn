# Operations

## Arithmetic Operations

```go
package main

import (
    "fmt"
)

func main() {
    a := 10
    b:= 7
    fmt.Println(a + b)  // Addition
    fmt.Println(a - b)  // Subtraction
    fmt.Println(a * b)  // Multiplication
    fmt.Println(a / b)  // Division
    fmt.Println(a % b)  // Modulus i.e Reminder
}
```

**OUTPUTS**
```
17
3
70
1
3
```

Just like other languages we can't do operations on different variable types. 

e.g:
* float and int operations, 
* unsignedint and signedint operation, etc...

Typecasting doesn't work on the go. Meaning if division is done between two integers, it doesn't be converted to float.

## Bitwise Operations

```go
package main

import "fmt"

func main() {
    a := 10 // 1010
    b := 7  // 0111

    fmt.Println(a & b)  // and operator                                 0010
    fmt.Println(a | b)  // or operator                                  1111
    fmt.Println(a ^ b)  // exclusive or operator (either of bit set)    1101
    fmt.Println(a &^ b) // and not operator (neither of the bit set)    0000
    fmt.Println(a << 3) // left bit shifting                            01010000
    fmt.Println(a >> 3) // right bit shifting                           00000001
}

```

One of the very intresting fact about bit shifting is that 
A single left shift multiplies a binary number by 2

```
0010 is 2
0100 is 4
```

So five left shift of a value 3 means
3 x 2^5^


## Complex Numbers

```go
var n complex64 = 2i        // Recognized as 0 + 2i
fmt.Printf("%v, %T\n", n, n) // OUTPUTS: (0+2i), complex64
```

To get the real part and imaginary part we can use `real` and `imag` functions

```go
var n complex64 = 1 + 2i
fmt.Printf("%v, %T", real(n), real(n))
fmt.Printf("%v, %T", imag(n), imag(n))
```

Complex number can also be created by using complex function

```go
var g complex128 = complex(5, 10)  // complex(Real, Imaginary)
```
