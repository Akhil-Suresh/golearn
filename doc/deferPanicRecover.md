

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

One of the use case of the defer is as below

```go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func deferCode(){
    res, err := http.Get("http://www.google.com/robots.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer  res.Body.Close()
    robots, err := ioutil.ReadAll(res.Body)

    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s", robots)
}
```

Here we can see how defer is used efficiently to close the response Body. There arise time where we forget to close the response after we open it, means in this example itself we can see that ioutil.ReadAll() is done before res.Body.Close(). This is one of the efficient use of defer. But note that this can't be used in for loop as we already mentioned defer works in LIFO pattern. So by the end of for loop there will be many open connection if we used defer.


When we defer the fucntion it actually take the argument when the defered is called
```go 
package main

import (
	"fmt"
)

func main() {
	a := "start"
	defer fmt.Println(a)
	a = "end
}
```

# Panicing

In go there is no such thing as exception, Just like other languages. In go its quite ok to know that when we try to open a file that didn't even exist, Its returned as error with a code unlike other language which try to raise it as an exception.

But there are some situation where the execution can't move forward because of the error that happened. go prefer no to call it as exception as the word exception is used in many other language with a meaning. So go doesn't want to mess up with that.

go prefer to call it __Panic__

```go
package main

import (
    "fmt"
)

func main() {
    a, b := 1, 0
    ans := a/b
    fmt.Println(ans)

}

// OUTPUT
// panic: runtime error: integer divide by zero

// goroutine 1 [running]:
// main.main()
// /tmp/sandbox669096995/prog.go:9 +0x20

```

So by this point the execution will comes to a dead end. Just like we raise an exception manually, we can raise a panic in go as

```go
panic("message to be displayed")
```

## Use of defer and panic

There is a difference in the execution flow when we use defer and panic. We already said that the defer statement will gets executed by the end of the program in LIFO model.

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("start")
    defer fmt.Println("this was defered")
    panic("something bad happened")
    fmt.Println("end")

}
```

So what you guess the output of the above code will be?

```go
// OUTPUT
// start
// this was defered
// panic: something bad happened
```

See the defer will get executed just before the panic got raised.

So the order here is that first `main function` then `defer statement` then `panic` that occur, then the return value

So why is this important? Well this is important because the defer statement that are going to close the resources are going to succeed even if the execution is going to panic. So if somewhere up the call stack if we recover from the panic we don't have to worry about the resources being left out there and left open

## recover

what if we are using anonymous function call ? anonymous function is a function that doesn't have a name

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("start")
    defer func() {
        if err := recover(); err != nil {
            log.Println("Error:", err)
        }
    }()
    panic("something bad happened")
    fmt.Println("end")
}
```

__This is one of the important thing to note about the defer statement. defer statement doesn't take the function itself but a function call__

We defined a recover function inside the anonymous function. What the recover function do is that its gonna return nil if the application isn't panicing and will return error if the application is panic.

So the above line of code is going to output

```go
// OUTPUT
// start
// 2009/11/10 23:00:00 Error: something bad happened
```

Still we doesn't get the `end` printed out. But this is highly useful if the call stack is deep.

See the below code

```go
package main

import (
    "fmt"
    "log"
)

func main() {
    fmt.Println("start")
    iamGoingToPanic()
    fmt.Println("end")
}


func iamGoingToPanic() {
    fmt.Println("I am going to panic")
    defer func() {
        if err := recover(); err != nil {
            log.Println("Error: ", err)
        }
    }()
    panic("Oh! shit happend")
    fmt.Println("done panicing")
}

// OUTPUTS
// start
// I am going to panic
// 2009/11/10 23:00:00 Error:  Oh! shit happend
// end
```

So what happened here is that the function `iamGoingToPanic` got panic and recovered nicely. We can see that the function higher up in the call stack presume its execution without any problem.

But what if the panic done by the `iamGoingToPanic` is critical? What if our main function is not supposed to get executed because of the fatal error from the panic?

Well in such cases we just need to rethrow the panic to main function as below

By doing so we can see on the output that main function got stopped execution on the panic.

```go
package main

import (
    "fmt"
    "log"
)

func main() {
    fmt.Println("start")
    iamGoingToPanic()
    fmt.Println("end")
}


func iamGoingToPanic() {
    fmt.Println("I am going to panic")
    defer func() {
        if err := recover(); err != nil {
            log.Println("Error: ", err)
            panic(err)
        }
    }()
    panic("Oh! shit happend")
    fmt.Println("done panicing")
}

// OUTPUTS
// start
// I am going to panic
// 2009/11/10 23:00:00 Error:  Oh! shit happend
// panic: Oh! shit happend [recovered]
// panic: Oh! shit happend

// goroutine 1 [running]:
// main.iamGoingToPanic.func1()
//      /tmp/sandbox454074194/prog.go:20 +0xe0
// panic(0x101da0, 0x15ea30)
//      /usr/local/go/src/runtime/panic.go:679 +0x240
// main.iamGoingToPanic()
//      /tmp/sandbox454074194/prog.go:23 +0xc0
// main.main()
//      /tmp/sandbox454074194/prog.go:10 +0xa0
```

