# Constants

* Naming convention
* Typed constants
* Untyped constants
* Enumerated Constants
* Enumeration expressions

## How we name our constants

Unlike other language we can't name our constants in all Upper case. The problem being like discussed in variable section, Giving first letter of the word in caps will make the variable exportable. So if we are using the constant inside our function locally. This is going to create trouble.

So we can use camel casing.

```go
const testConst int = 05
```

Another important thing is we can't set value to const at runtime. Meaning we can't have value to this constant as a result of function. The below code is gonna throw an error.

```go
import "math/rand"
const pi float64 = rand.Intn(100)
```

Constants can be made from primitive types. But can't be created from mutable type.
Meaning it can't be created from an array, map or any collection type which are inherently mutable.

Also Other important characteristics is __constants can be shadowed__.

```go
package main

import "fmt"

const myConst int = 56

func main() {
  const myConst int = 26
  fmt.Printf("%v, %T", myConst, myConst) // The output will be 26
}
```

We can do usual addition, subtraction, multiplication, division on constant and variable of same type

__We can initialize a constant by letting compiler figure out its type as well.__

```go
const a = 64
```

One of the intresting thing about this is we can do

```go
const a = 64
var b int16 = 32
fmt.Printf(a + b)
```

The above program will execute without an issue.

## Enumerated Constants

Enumerated constants are mostly used in package level

### iota

iota is a counter that can be used with enumerated constant.

```go
package main

import "fmt"

const a = iota

func main() {
  fmt.Printf("%v, %T", a, a)
}
```

If we are defining multiple iota values inside const block that's gonna increment value for each constants

```go
package main

import "fmt"

const (
  a = iota
  b = iota
  c = iota
)


func main() {
  fmt.Printf("%v, %T\n", a, a)
  fmt.Printf("%v, %T\n", b, b)
  fmt.Printf("%v, %T\n", c, c)
}

// OUTPUTS
// 0, int
// 1, int
// 2, int
```

An intesting feature of iota is even if we don't give value for preceding variable inside const block, complier is gonna infer its value.

```go
package main

import "fmt"

const (
  a = iota
  b
  c
)


func main() {
  fmt.Printf("%v, %T\n", a, a)
  fmt.Printf("%v, %T\n", b, b)
  fmt.Printf("%v, %T\n", c, c)
}

// OUTPUTS
// 0, int
// 1, int
// 2, int
```

And if we are creating seperate constant block. The value will reinitaize

```go
package main

import "fmt"

const (
  a = iota
  b
  c
)

const (
  d = iota
  e
  f
)

func main() {
  fmt.Printf("%v, %T\n", a, a)
  fmt.Printf("%v, %T\n", b, b)
  fmt.Printf("%v, %T\n", c, c)
  fmt.Printf("%v, %T\n", d, d)
  fmt.Printf("%v, %T\n", e, e)
  fmt.Printf("%v, %T\n", f, f)
}

// OUTPUTS
// 0, int
// 1, int
// 2, int
// 0, int
// 1, int
// 2, int
```

## Offsetting iota variable

```go
const (
  a = iota + 5
  b
)

fmt.Printf("%v, %T", a, a)
fmt.Printf("%v, %T", b, b)

// OUTPUTS
// 5, int
// 6, int
```

