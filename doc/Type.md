# Types in GO

## Primitive Types

* Boolean Types
* Numeric Types
  * Integers
  * Floating point
  * Complex numbers
* Text Types
    * String
    * Rune

## Integer Types

### First types are signed integers

* int8  --> (-128-127)
* int16 --> (-32 768 - 32 767)
* int32 --> (-2 147 483 648 - 2 147 483 647)
* int64 --> (-9 223 372 036 854 775 808 - 9 223 372 036 854 775 807)

### Unsigned Integers

* unint8  --> (0-255)
* unit16  --> (0 - 65 535)
* unit32  --> (0 - 4 294 967 295)

## Float types

* float32 --> (+-1.18E-38 - +-3.4E38)
* float64 --> (+-2.23E-308 - +-1.8E308)

## Complex Type

* complex64
* complex32

## Text Type

### String literal

String in go stands for UTF8 character. So string cannot encode every type of character that's available.

```go
package main
import "fmt"

func main() {
    s := "I am a string"
    fmt.Printf("%v, %T\n", s, s)
}
```

**OUTPUTS**
```
I am a string, string
```

String can be treated like an array, or a collection of letters. That way we can get value of string at any index position starting from 0. ie.
```go 
fmt.Printf("%v, %T", s[5], s[5])    // OUTPUTS: 97, unit8
```

Yes the above code is little confusing and here is its explanation. The string is treated as array of bytes and is immutable. We can get back the value by **type casting** it.

```go
string(s[5])    // OUTPUTS: a
```

Strings can be concatenated together by plus sign.

```go
string1 := "this is a string"
string2 := "this is another string"

fmt.Println("%v, %T", string1 + string2, string1 + string2)
```

String can also be converted to collections of bytes

```go
package main

import "fmt"

func main() {
    s := "I am a string"
    b := []byte(s)
    fmt.Printf("%v, %T", b, b)
}
```

**OUTPUTS**
```
// [73 32 97 109 32 97 32 115 116 114 105 110 103], []uint8
```

This feature in go can be utilized when we are sending data from server to another server, or to a filesystem.
So that we can less worry about line endings and other string related issues

### Rune literal (Single quote Vs Double quote)

```go
package main

import "fmt"

func main() {
    var r rune = 'a'
    b := 'z'
    fmt.Printf("%v, %T", r, r)
}
```

**OUTPUTS**
```
97, int32
122, int32
```


**So double quotes and single quotes matters in go code**. 
Single quotes string are treated as **rune**. 

**Unlike a string literal which is a byte. The rune is of type int32.**