# Variables

There are mainly three ways to declare a variable inside go

1. Declaring variable by type
2. Declaring and assigning variable in single line
3. Compiler auto detect variable type

For example

1. **Declare variable by type**

    ```go
    var i int
    i = 5
    ```

2. **Declaring and assigning variable in single line**

    ```go
    var i int = 7
    ```

3. **Compiler auto detect variable type**

    ```go
    i := 10
    ```

# Declaring variables

In go there are two ways to declare a variable
1. Inside a function
2. In package level

Till now what we saw above is decalring variable inside function. Now lets see how its done in a package level.

## Declaring variable in package level

While decalring variable at package level, we can't make use of `:=` syntax. So we have to use the full declaration syntax.

```go
package main

import "fmt"

var j float32 = 65

func main() {
    fmt.Printf("%v, %T", j, j)
}

// OUTPUTS
// 65, float32
```

**Note:**: *While declaring variable at package level, we can't let compiler figure out the type of variable.*


## Declaring variable as block at package level

```go
package main

// Declaring variable by not cluttering up code with too may var keywords!
var (
    studentFname string
    studentLname string
    studentClass int
    studentAge   int
)

// Declaring and Initializing variable at package level.
var (
    name string = "Akhil"
    age int = 16
    class string = 10
    division   string = "A"
)

func main() {
    studentFname = "Akhil"
}
```

## Initializing variable without value

Unlike other languages in go if we initialize a variable without value. It gets defaulted to 0. For instance in C the value will be a random integer.

```go
package main
import "fmt"
func main() {
    var i int
    fmt.Printf("%v, %T\n", i, i)  //OUTPUTS: 0, int

    var j string
    fmt.Printf("%v, %T\n", j, j)  //OUTPUTS: , string

    var k bool
    fmt.Printf("%v, %T\n", k, k)  //OUTPUTS: false, bool
}
```

### Unused variables :X
    If there is unused variable in go, When the code gets compiled, Its gonna throw an error.

## Variable shadowing

Suppose we initialize a variable at package level. The we reinitialize it in our function. In that case the variable in our local scope will be having more priority.

```go
package main

var testThis int = 0

func main() {
    var testThis int = 1
    fmt.Println(testThis)
}

// OUTPUTS
// 1
```

# Some naming conventions

There are mainly two things to be taken care of.
1. Variable visibility
2. Naming convention.

## Variable Visibility

Variables at package level can be exported. Till now you must have noticed that we are only using lowercase letters for variable naming. We can also use upper case letter for variable name as well. Doing so will make it exportable.

Consider the scenario below, here the scop of i is limited to package level and not outside the package.

```go
package main
import "fmt"
var age int = 32

func main() {
    fmt.Println(age)
}
```

We can make the value exportable by using uppercase for first letter

```go
package main
import "fmt"
var Age int = 32

func main() {
    fmt.Println(Age)
}
```
The variable `Age` can be used inside other packages now.

### Naming Convention

* When you declare a variable at package level try to use upper case

    ```go
    var CLASSNAME string
    ```

* When using acronym use caps

    ```go
    var thisURL = "https://www.google.com"
    var thisHTTP = "https://www.youtube.com"
    ```

# Typecasting

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

Typecasting to string can't be achieved that easily

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

In order to achieve proper result we can make use of package __strconv__.

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

## Blank Identifier  _ (Underscore)

The real use of *Blank Identifier* comes when a function returns multiple values, but we need only a few values and want to discard some values. Basically it tells the compiler that this variable is not needed and ignore it without any error. It hides the variable’s values and makes the program readable. **So whenever you will assign a value to Blank Identifier it becomes unusable.**
