# IF

A very  basic if statement is as below
NB:
    The go lang design discussion ended that there won't be single if block. Meaning curly braces are necessary.

```go
package main

import (
    "fmt"
)

func main() {

    if true {
        fmt.Println("True case")
    }

    if false {
        fmt.Println("This statement will never execute")
    }

}
```

## The initializer syntax

```go
package main

import "fmt"

func main() {
  shapeEdgeMap := map[string]int{
    "Triangle": 3,
    "Square":   4,
    "Pentagon": 5,
    "Hexagon":  6,
    "Heptagon": 7,
  }
  fmt.Println(shapeEdgeMap)
  shapeEdgeMap["Octagon"] = 8
  fmt.Println(shapeEdgeMap)
  if value, ok := shapeEdgeMap["Line"]; ok {
    fmt.Println(value)
  }
}

```

The first part of the if statement delimited by semicolon is the initializer

```go
value, ok := shapeEdgeMap["Line"];
```
The initializer allows us to run a statement and generate some information that's gonna set us up to work with the if block

Here we are generating the boolean result by the initializer and using that(__ok__) as the check for the if block to execute.

We also have access to the value created from the initializer that can be used inside if block.

### Operators along with if

```go
package main

import "fmt"

func main() {
  var number int = 30
  var guess int = 50
  if number < 30 {
    fmt.Println("Number is less than 30")
  }
  if number > 30 {
    fmt.Println("Number is greater than 30")
  }
  if number == 30 {
    fmt.Println("Number is equal to 30")
  }
  fmt.Println(number <= guess, number >= guess, number != guess)
}

/*
    OUTPUTS
    Number is equal to 30
    true false true
*/
```

## IF ELSE

```go

// Greatest of three numbers

package main

import "fmt"

func main() {

    var (
        a int = 10
        b int = 20
        c int = 30
    )
  
    if a > b {
        if a > c {
            fmt.Println("a is greater")
        }
    } else if b > c {
        fmt.Println("b is greater")
    } else {
        fmt.Println("c is greater")
    }

}
```

## switch

```go
package main

import "fmt"

func main() {
    var day int = 10
  
    switch day {
        case 1:
            fmt.Println("Monday")
        case 2:
            fmt.Println("Tuesday")
        case 3:
            fmt.Println("Wednesday")
        case 4:
            fmt.Println("Thursday")
        case 5:
            fmt.Println("Friday")
        case 6:
            fmt.Println("Saturday")
        case 7:
            fmt.Println("Sunday")
        case 8, 9, 10:                  // Note the matching of multiple test cases
            fmt.Println("Specially reserved for holidays and stuff")
        default:
            fmt.Println("Please select a valid day [1-7]")
    }
}
```

Switch is interesting in a sense that if one of the case is validated to true the others won't. This behavior can be altered by using fallthrough

```go
package main

import "fmt"

func main() {
    var day int = 4
  
    switch {
        case day <= 5 :
            fmt.Println("Is a valid day")
        case day <= 10:
            fmt.Println("Also a valid day")
    }
}

/*
OUTPUTS
Is a valid day
*/

```

### fallthrough

```go
package main

import "fmt"

func main() {
    var day int = 4
  
    switch {
        case day <= 5 :
            fmt.Println("Is a valid day")
            fallthrough
        case day <= 10:
            fmt.Println("Also a valid day")
    }
}

/*
OUTPUTS
Is a valid day
Also a valid day
*/

```

fallthrough is not used very common in golang because of the availability of multiple tag

## Type switch and break early

```go
package main

import "fmt"

func main() {

    var i interface{} = 1
    var j interface{} = 1.0
    var k interface{} = "string"
    var l interface{} = [5]int{}

    printType(i)
    printType(j)
    printType(k)
    printType(l)

}

func printType(value interface{}) {

    switch value.(type) {
        case int:
            fmt.Println("i is an int")
            break
            fmt.Println("This will never execute")
        case float64:
            fmt.Println("i is float 64")
        case string:
            fmt.Println("i is an string")
        default:
            fmt.Println("i is another type")
    }

}
/*
    OUTPUTS
    ./prog.go:237:13: unreachable code
    Go vet exited.

    i is an int
    i is float 64
    i is an string
    i is another type
*/
```

interface can have any type of value we can.

Just like other languages we can use __break__ statement to break out of execution.
