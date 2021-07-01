# Init Function

The init function is a function that takes no argument and returns nothing. This function executes after the package is imported and maintains the order of execution. That means multiple init functions can be defined in a file and they will be called one after another maintaining the order.

## Usage
The init function in Go should be used to do the tasks we want to do before doing all other things in that file. That means it is more suited to tasks like initializations.

```go
package main
 
import (
    "fmt"
)
 
func init() {
    fmt.Println("Runs first")
}
 
func main() {
    fmt.Println("Hello, World")
     
    // Runs first
    // Hello, World
}
```

## Order of execution
The order of the init functions matter. The way they are declared in the same order they are executed. Here is an example showing how that works.

```go
package main
 
import (
    "fmt"
)
 
func init() {
    fmt.Println("Runs first")
}
 
func init() {
    fmt.Println("Runs second")
}
 
func main() {
    fmt.Println("Hello, World")     
}
 
func init() {
    fmt.Println("Runs last")
}


/*
 OUTPUT

 Runs first
 Runs second
 Runs last
 Hello, World
*/
```

Observe that even though one of the init function has been declared after the main, it still runs before main and maintains the order.