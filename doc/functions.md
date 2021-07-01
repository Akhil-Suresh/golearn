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

## Returning pointer
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

Actually we may be in a bit surprise here as the variable we are passing out is needed to end its life on the local stack.
In other language its not a safe operation as we are now returning a address of the memory which just got free so you have no idea what value is gonna be there.
Well in the go when it recognizes we are returning the value that is generated in the  local stack, Its automatically gonna promote the variable on the shared memory which is also called heap memory. So every thing works as fine.


## Named return values

Its not used very common but there is this technique in go. 
In go we can define the return value to be specified way up on the function signature itself

```go
package main

import (
    "fmt"
)

func main() {
    result := sum(1, 2, 3, 4, 5)
    fmt.Println(result)
}

func sum(values ...int) (result int) {
    for _, value := range values {
        result += value
    }
    return
}
```

## Returning Error

As we have already discussed there may arise scenarios where error occur in our custom functions, Which are not very deadly. we also told that we can do handle it by raising panic and recovering it.

There is a standard way in go to handle such non deadly issues, By non deadly I mean the errors that doesn't put the executing function to halt.

That is by returning the error from our custom function.

```go
package main

import (
    "fmt"
)

func main() {
    result, err := divide(1.0, 0.0)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(result)
}

func divide(a, b float64) (float64, error) {
    if b == 0.0 {
        return 0.0, fmt.Errorf("Cannot divide by zero")
    }
    return a/b, nil
}


// OUTPUT
// Cannot divide by zero
```


# Playing with functions.

So far we have been treating the function as the special thing. As its defined with func keyword and has return value and should be defined in global scope and so..

The function in go are much more powerful than that. Because function itself can be treated as *type* in go. Meaning function can be passed around just like variables.

## Anonymous function

```go
package main

import (
  "fmt"
)

func main() {
  func() {
    fmt.Println("Hello World!")
  }()
}
```

We can see in the above code that the function has no name and its called just after its definition. This is an immediately invoked function.

```go
package main

import (
  "fmt"
)

func main() {
  for i:=0; i<5; i++ {
    func(){
      fmt.Println(i)
  }()
  }
}
```

See the above example its working fine. As the global value inside the for loop is getting inside the anonymous function and everything is fine. But note that when the anonymous function is made asynchronous the whole behaviour is gonna change. So its better to pass in the value of variable inside for loop into the anonymous function as below.

```go
package main

import (
  "fmt"
)

func main() {
  for i:=0; i<5; i++ {
    func(i int){
      fmt.Println(i)
  }(i)
  }
}
```

### Assigning function to a variable

```go
package main

import (
  "fmt"
)

func main() {
  var f func() = func(){
      fmt.Println("Hello World!")
  }
  f()
  // Can be simplified as
  newF := func(){
    fmt.Println("A simple way to defnine func")
  }
  newF()
}
```

We can pass in parameters as well

```go
package main

import (
  "fmt"
)

func main() {
  var divide func(a,b float64)(float64, error)
  divide = func(a, b float64)(float64, error){
    if b == 0.0 {
      return 0.0, fmt.Errorf("Can't divide by zero")
    } else {
      return a/b, nil
    }
  }

  result, err := divide(1.0, 0.0)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(result)
}
```

Or simply

```go
package main

import (
  "fmt"
)

func main() {
  var f func() = func(){
      fmt.Println("Hello World!")
  }
  f()
  // Can be simplified as
  newF := func(a float64)(string, error){
    if a == 0.0 {
        return "", fmt.Errorf("Test Error")
    }
    return "A simple way to define func", nil
  }  
  _, err := newF(0.0)

  if err != nil {
    fmt.Println(err)
    return
  }
```

## Methods

Method is basically a function that's executing in a known context. Known context in go is any type.

```go
package main

import (
  "fmt"
)

func main() {
  g := greeter{
    greeting: "Hello",
    name: "Go",
  }
  g.greet()
  g.greetMeAgain()
}

type greeter struct {
  greeting string
  name string
}

// This is the method on the struct above
func (f greeter) greet() {
  fmt.Println(f.greeting, f.name)
}

func (k greeter) greetMeAgain(){
  fmt.Println("Greet me again", k.greeting, k.name)
}

// OUTPUTS
// Hello Go
// Greet me again Hello Go
```

The struct object receiving inside the greet function is a copy of the main struct. So changing values of the struct inside the method has no effect on the global scope.

```go
package main

import (
  "fmt"
)

func main() {
  g := greeter{
    greeting: "Hello",
    name: "Go",
  }
  g.greet()
  g.greetMeAgain()

  // This is to prove that g is not passed in as reference.
  fmt.Println(g.name)
}

type greeter struct {
  greeting string
  name string
}

// This is the method on the struct above
func (f greeter) greet() {
  fmt.Println(f.greeting, f.name)
}

func (k greeter) greetMeAgain(){
  fmt.Println("Greet me again", k.greeting, k.name)
  k.name = ""
}

// OUTPUTS
// Hello Go
// Greet me again Hello Go
// Go
```

well we can make go pass the same struct from global scope to method scope by making use of pointers. This is highly useful when dealing with large struct data. As each time method gets invoked we can avoid the overhead of creating copy of it.

```go
package main

import (
  "fmt"
)

func main() {
  g := greeter{
    greeting: "Hello",
    name: "Go",
  }
  g.greet()
  g.greetMeAgain()

  // Making use of pass by reference
  fmt.Println(g.name)
}

type greeter struct {
  greeting string
  name string
}

// This is the method on the struct above
func (f greeter) greet() {
  fmt.Println(f.greeting, f.name)
}

func (k *greeter) greetMeAgain(){
  fmt.Println("Greet me again", k.greeting, k.name)
  k.name = "oh boy! its changed"
}

// OUTPUTS
// Hello Go
// Greet me again Hello Go
// oh boy! its changed
```

