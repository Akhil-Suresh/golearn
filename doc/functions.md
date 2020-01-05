# Functions

The way go is structured is we always have to have an entry point. That is the main function inside package main

The naming conventions are much like the variables.

The upper case beginnings are for the public and lower case are for the private.

The actual body are contained inside the curly braces

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello, world")
}

```

## parameters

```go
package main

import (
    "fmt"
)

func main() {
    greet("Hello", "Akhil")
}

func greet(message string, name string) {
    fmt.Println(message, name)
}

// OUTPUTS
// Hello Akhil

```

Same data type params can be added together as below

```go
package main

import (
    "fmt"
)

func main() {
    greet("Hello", "Akhil", 34)
}

func greet(message, name string, count int) {
    fmt.Println(message, name)
    fmt.Printf("Count %v", count)
}

// OUTPUTS
// Hello Akhil
// Count 34
```

## Passing By Value

By default the value passed to a function act as *passing by value*

In the below function we can see that the value of *name* doesn't get changed even after we changed it inside *greet* function.

```go
package main

import (
    "fmt"
)

func main() {
    name := "Akhil"
    greet("Hello", name, 34)
    fmt.Println(name)

}

func greet(message, name string, count int) {
    fmt.Println(message, name)
    fmt.Printf("Count %v\n", count)
    name = "No Effect"
}

// OUTPUTS
// Hello Akhil
// Count 34
// Akhil
```

## Passing by Pointers

```go
package main

import (
    "fmt"
)

func main() {
    name := "Akhil"
    greet("Hello", &name, 34)
    fmt.Println(name)

}

func greet(message string, name *string, count int) {
    fmt.Println(message, *name)
    fmt.Printf("Count %v\n", count)
    *name = "This changes Everything"
}

// OUTPUTS
// Hello Akhil
// Count 34
// This changes Everything
```

So from the above we can see that the value of the variable *name* has changed not only on the scope of the function *greet* but also on the scope of *main* function.

so why we need to do this??

There are a couple of scenarios. 

There may arise scenarios such as the calling function need to act on the values that are passed into them. This is one of the safest way to achieve it.

The other reason is *passing by pointer* is much efficient than *passing by value*. Though passing by copying on strings has no much of performance improvement, However when dealing with large data structures *passing by value* causes entire data structure to gets up copied every single time which eats up a lot of the memory space. At such points passing by pointers is much efficient.

Be really need to aware when playing with maps and slices because they are internally using pointers for this by default.

## Variadic parameters

In the below code *sum function call* is not creating five different values. Instead its wrapping up the values passed in to slice.


```go
package main

import (
    "fmt"
)

func main() {
    sum(1, 2, 3, 4, 5)
}

func sum(values ...int) {
    fmt.Println(values)
    result := 0
    for _, value := range values {
        result += value
    }
    fmt.Println(result)
}

OUTPUTS
[1 2 3 4 5]
15
```

One important thing to be note on variadic parameters is that it has to be one and it has to be at the end.

Which means we can't have two variadic parameters as arguments in one function call. Also its definition needs to be at the end.

```go
package main

import (
    "fmt"
)

func main() {
    sum("The sum is", 1, 2, 3, 4, 5)
}

func sum(msg string,values ...int) {
    fmt.Println(values)
    result := 0
    for _, value := range values {
        result += value
    }
    fmt.Println(msg, result)
}

// OUTPUTS
// [1 2 3 4 5]
// The sum is 15
```

## Returning the value

Note that we need to define the return type as well when we are returning a data from the function.

```go
package main

import (
    "fmt"
)

func main() {
    result := sum(1, 2, 3, 4, 5)
    fmt.Println(result)
}

func sum(values ...int) int {
    result := 0
    for _, value := range values {
        result += value
    }
    return result
}

```

The other feature go has is the ability to return a local variable as pointer.

```go
package main

import (
    "fmt"
)

func main() {
    result := sum(1, 2, 3, 4, 5)
    fmt.Println(*result)
}

func sum(values ...int) *int {
    result := 0
    for _, value := range values {
        result += value
    }
    return &result
}

// OUTPUTS
// 15
```

Actually we may be in a bit suprise here as the variable we are passing out is needed to end on the 