# For loops

A simple for loop can be implemented as

```go
package main

import (
    "fmt"
)

func main() {
    for i:= 0; i < 5; i ++ {
        fmt.Println(i)
    }
}
```

We can increment multiple value in for loop like

```go
package main

import (
    "fmt"
)

func main() {
    for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
        fmt.Println(i, j)
    }
}

/*
    OUTPUTS
    0 0
    1 2
    2 4
    3 6
    4 8
*/

```

We can have nested for loops as

```go
package main

import (
 "fmt"
)

func main() {

    name := "AKHIL"

    for i := 0; i<len(name); i++ {
        for j:= 0; j <= i; j++ {
            fmt.Printf("%v", string(name[j]))
        }
        fmt.Printf("\n")
    }
}
/*
    OUTPUTS
    A
    AK
    AKH
    AKHI
    AKHIL
```

We can also loop by having a variable from main scope

```go
package main

import (
    "fmt"
)

func main() {
    i := 0
    for ; i < 5; i++ {
        fmt.Println(i)
    }
}
```

Just like other language if we forget the increment operator its gonna run in an infinite loop

one interesting thing we can do with for loop in go is we can omit both the semicolons and let the condition be checked on the go

Its just a syntactic sugar

```go
package main

import (
    "fmt"
)

func main() {
    i := 0
    for i < 5 {
        fmt.Println(i)
        i++
    }
}
```

Just kinda implementation of while loop we can use for loop in such a fashion in go

```go
package main

import "fmt"

func main() {
    i := 0
    for {
        if i == 5 {
            break
        }
        fmt.Println(i)
        i++
    }
}

/*
    OUTPUTS
    0
    1
    2
    3
    4
*/

```

## break

Another important feature for go is label. There arise a scenario when we have multiple nested loops and we need to break out our execution to a certain point. This can be acheived by using labels as below

```go
package main

import (
 "fmt"
)

func main() {

    for i := 1; i<10; i++ {
        for j:= 1; j <= i; j++ {
        if i == 5 {
            break
        }
            fmt.Printf("*")
        }
        fmt.Printf("\n")
    }
}

/*
    OUTPUTS
    *
    **
    ***
    ****

    ******
    *******
    ********
    *********
*/
```
as we can see, We clearly needed to break out of the code when iteration reached value 5 but what happened is not expected result.

This can be acheived by

```go
package main

import (
 "fmt"
)

func main() {

    BreakToOuterLoop:
    for i := 1; i<10; i++ {
        for j:= 1; j <= i; j++ {
            if i == 5 {
                break BreakToOuterLoop
            }
            fmt.Printf("*")
        }
        fmt.Printf("\n")
    }
}

/*
    OUTPUTS
    *
    **
    ***
    ****
*/

```


## continue

Just like other languages continue can be used to skip the execution of the function.

```go

// Function to print odd numbers

package main

import "fmt"

func main() {
    for i := 0; i<10; i++ {
        if i % 2 == 0 {
            continue
        }
        fmt.Println(i)
    }
}

```

## Working with collections - range

```go
package main

import (
 "fmt"
)

func main() {

    i := []int{1, 2, 3, 4}
    for key, value := range i {
        fmt.Println(key, value)
    }
    printMap()
}

func printMap() {
  var shapeEdgeMap = make(map[string]int)
  shapeEdgeMap = map[string]int{
    "Triangle": 3,
    "Square":   4,
    "Pentagon": 5,
    "Hexagon":  6,
    "Heptagon": 7,
  }
  for key, value := range shapeEdgeMap {
    fmt.Println(key, value)
  }

}
```
