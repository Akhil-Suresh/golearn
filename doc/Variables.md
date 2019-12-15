# Variables

There are three main ways to declare a variable inside go

## Declare variable by type

```go
var i int
i = 5
```

## Declaring variable and assigning

```go
var i int = 7
```

## Variable type auto by compiler

```go
i := 10
```

## Declaring variable in package level

While declaring variable at package level, we can't let compiler figure out the type of variable. So we need to stick with second way of declaring variable.

```go
package main

var j float32 = 65

func main() {

}
```

## Declaring variable as block at package level

```go
package main
var (
    studentFname string
    studentLname string
    studentClass int
    studentAge   int
)
func main() {
    studentFname = "Akhil"
}
```

## Variable shadowing

Suppose we initialize a variable at package level. The we reinitialize it in our function. In that case the variable in our local scope will be having more priority.

```go
package main

var testThis int 0

func main() {
    var testThis int 1
    fmt.Println(testThis)
}

```

The above line of code print 1

But compiler gives an error if we try to reinitialize it using compiler auto detect feature

```go
package main

var testThis int 0

func main () {
    testThis := 1
}

```

## Beware of Unused variables

If there is unused variable in go, When the code gets compiled, Its gonna throw an error.

## Some naming conventions

* When you declare a variable at package level try to use upper case

```go
var CLASSNAME string
```

* When using acronym use caps

```go
var thisURL = "https://www.google.com"
var thisHTTP = "https://www.youtube.com"
```

## Typecasting

```
    int(<variable>)
```

```go
func main(){
    var i float32 = 5.6
    fmt.Printf("%v, %T", i, i)

    var j int
    j = int(i)
    fmt.Printf("%v, %T", j, j)
}
```

### Typecasting to string

typecasting to string can't be acheived that eaisly

```go
func main() {
    var i int = 10
    var str string
    str = string(i)
    // This is gonna print pretty weird result
    fmt.Printf("%v, %T", str, str)
}

```

This is because of difference in working of string. Sting is treated as alias of string of bytes

In order to achieve proper result we can use o package __strconv__.

```go
import (
    "fmt"
    "strconv"
)

func main() {
    var i int = 10
    var str string
    str = strconv.Itoa(i)
    fmt.Printf("%v, %T", str, str)
}

```
