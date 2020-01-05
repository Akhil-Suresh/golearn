

In a normal go application the control flows from top to bottom. We can alter that by introducing branching logic so that we can skip a little  bit

```go
package main

import "fmt"

func main() {
    fmt.Println("start")
    fmt.Println("middle")
    fmt.Println("end")
}

/*
    OUTPUTS

    start
    middle
    end
*/

```

# Defer

In golang we can alter the control flow behaviour using defer. It works in LIFO order

```go
package main

import "fmt"

func main() {
    fmt.Println("start")
    defer fmt.Println("middle")
    fmt.Println("end")
}

// OUTPUTS
// start
// end
// middle
```

```go
package main

import "fmt"

func main() {
    defer fmt.Println("start")
    defer fmt.Println("middle")
    defer fmt.Println("end")
}

// OUTPUTS
// end
// middle
// start

```

