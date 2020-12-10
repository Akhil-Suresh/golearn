# Pointers

Lets start very simple

We said before that the variables inside go are treated as value types.

So when we assign a value a to b, underneath the go is copying the value of a and assigning to b. So that both have seperate memory location. So if we change one the other is not going to be changed.

This can be verified by __address of__ __&__ operator.

```go
package main

import (
    "fmt"
)

func main() {
    var a int = 42
    var b int = a
    fmt.Println(&a, &b)
}

// OUTPUTS
// 0x40e020 0x40e024
```

Here we can see that both carry different address, Which means one doesn't depends on other.

## Pointers syntax

We can define a pointer by using `*` preceding the data type

```go
var pointerVar *int
```

Now this variable can hold the memory location of other variables.

for example

```go
package main

import (
    "fmt"
)

func main() {
    var a int = 42
    var b *int = &a
    fmt.Println(a, b, &a)
}

// OUTPUTS
// 42 0x40e020 0x40e020
```

So what if we need to print the value at the address stored in b

Here also we can use `*` to get the value. So when `*` is used infront of a variable it acts as deferencing operator. And infront of datatype it acts as a pointer.

```go
package main

import (
    "fmt"
)

func main() {
    var a int = 42
    var b *int = &a
    fmt.Println(a, b,*b)
}

// OUTPUTS
// 42 0x40e020 42

```

## Pointer Arithmetic ahhh.. no no

One of the highly used useful stuff in c and c++ languages is pointer arithmetic. Lets see it in action

```go
package main

import (
    "fmt"
)

func main() {
    var a = [4]int{1, 2, 3, 4}

    b := &a[1]  // Shorthand for pointer
    c := &a[2]

    fmt.Println(a, b, c)

    // So by now we know that it takes 4 bytes to save int
    // One of the very intresting thing that can be done is
    // subtract 4 from address of b so that we can get value of a[0]

    d := &a[1] - 4    // But as far as go is considered this is unsafe. Thus go lang doesn't allow this.
    fmt.Println(d)
}

// OUTPUTS
// invalid operation: &a[1] - 4 (mismatched types *int and int)
```

But if we are absolutely in a dilemma we can do such arithmetic by making use of `unsafe` package in go packages


## Pointer Struct

Normal struct declaration is as

```go
package main

import (
    "fmt"
)

type Config struct {
    port int
    name string
}

func main() {
    var c Config
    c = Config{port: 80, name: "http"}
    fmt.Println(c)
}
// OUTPUT
// {80 http}
```

We can actually make use of pointers to get the same result as

```go
package main

import (
    "fmt"
)

type Config struct {
    port int
    name string
}

func main() {
    var c *Config
    c = &Config{port: 80, name: "http"}
    fmt.Println(c)
}

// OUTPUT
// &{80 http}
```

Notice the `&` It says the value c is **not actually a variable** but a **pointer** which has a value of `{80 http}`

## making use of new

```go
package main

import (
    "fmt"
)

type Config struct {
    port int
    name string
}

func main() {
    var c *Config
    c = new(Config)
    fmt.Println(c)
}

// OUTPUTS
// &{0 }
```

So the obvious question in here is how will we access port and name? Since by now we know that `c` is not a variable but a pointer.

We can access the fields inside the stuct as

```go
package main

import (
    "fmt"
)

type Config struct {
    port int
    name string
}

func main() {
    var c *Config
    c = new(Config)
    (*c).port = 8000
    (*c).name = "Web Server"
    fmt.Println((*c).port, (*c).name)
}

// OUTPUTS
// 8000 Web Server
```

eeh but...

Ya that's a bit tough I agree , Well due to the limitation put on pointers compilers are smart enough to add syntactical sugar to this. Thus the params and star can be avoided...

```go
package main

import (
    "fmt"
)

type Config struct {
    port int
    name string
}

func main() {
    var c *Config
    c = new(Config)
    c.port = 8000
    c.name = "Web Server"
    fmt.Println(c.port, c.name)
}

// OUTPUTS
// 8000 Web Server
```

Well This is crazy. AHH  But it does the job...

