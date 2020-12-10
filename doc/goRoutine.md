## GO Routines

We can turn normal function into go routine by putting keyword **go** infront of the method call 

For instance 

```go
package main


import "fmt"

func main() {
	go sayHello()
}

func sayHello() {
	fmt.Println("Hello Everyone !")
}
```

Here *sayHello* function will execute as go Routine. So basically putting a *go* keyword infront of function call will make go to spin off a **green thread** and run the **sayHello** inside that green thread.

### Threads
In most programming language use os threads what that means is they got individual function call stack  dedicated to execution of whatever code is handed to that thread. Traditionally these tend to be very very large they have around 1MB of RAM. They take quite a bit of time for the application to setup so you want to be conservative about how you use your threads.

Green threads
Instead of creating these very massive heavy overhead threads we are gonna create abstraction of the thread called go routine. Inside go runtime we got a scheduler to map these go routine on to these os threads for period of time and the scheduler will take turns with evey cpu threads that is available and assign different go routines a certain amount of processing time on the threads so we don't have to interact with low level threads directly we are interacting with high level go routine. HTe advantage of that is since we have the abstraction of go routine we can start with  very very small stack spaces since they can be reallocated very very quickly and so they are very cheap to create and destroy. So its not uncommon in a go application to see thousands or tens of thousands of go routine running at same time and the application will have no problem with it at all. SO when we compare with other language whick uses os tnreads that have 1MD overheads there is no way we can go along with ten thousands threads in a environment like that.


## The Catch
In the above program though sayHello function will execute in a go routine, It doesn't gonna print anything. Hmm... that's a bit disappointing the reason why that happens is that before the io operation our main go routine ends.

so how to make our print function visible? obviouslyy.... put that horrible method sleep....

```go
package main


import (
	"fmt"
	"time"
)

func main() {
	go sayHello()
	// Implementing sleep method since sayHello won't get output
	// Since it is now running a different thread from our 
	// main go routine. So yeah you got it main function is 
	// running inside a go routine.

	msg := "hello"
	go func() {
		fmt.Println(msg)
	}()
	time.Sleep(1 * time.Millisecond)
}

func sayHello() {
	fmt.Println("Hello Everyone !")
}
```



## Closures

Consider the program below.

```go
package main

import "fmt"
import "time"

func main() {
    msg := "Hello"
    go func () {
        fmt.Println(msg)
    }()
    time.Sleep(1 * time.Millisecond)
}

```

**OUTPUTS**
```
Hello
```

You may be wondering how the go routine running inside main got the value from main function which is running entirely in seperate call stack. The answer to that is what we call **closure**. So go routine make use of **closures**. The anonymous function has access to variable outside its outer scope.


### The problem arises

The problem with the above approach is that we create a dependency with the go routine in main and our anonymous function. Though doing so is not very harmful while we look at the code above but consider what happens when the value of *msg* variable changes 


```go
package main

import "fmt"

func main() {
    msg := "Hello"
    go func () {
        fmt.Println(msg)
    }()
    msg = "GoodBye!"
    time.Sleep(1 * time.Millisecond)
}
```

**OUTPUTS**
```
GoodBye!
```

huh?? what happened? Well what happened is that even though our go routine is called when value of msg is Hello when the go routine is about to print those message into io the value got changed by our main to GoodBye!

So effectively a race condition is happening here.

So what we can do to avoid this is pass the msg value to anonymous function so that the go routine doesn't need to look for value of msg in outer scope.

```go
package main

import "fmt"

func main() {
    msg := "Hello"
    go func (msg string) {
        fmt.Println(msg)
    }(msg)
    msg = "GoodBye!"
    time.Sleep(1 * time.Millisecond)
}
```

**OUTPUTS**
```
Hello!
```

That way every thing work as expected.

## Problem with sleep

When we are using sleep method like above what we are doing is we are making our cpu clock to get in sync with real time clock. So this is gonna slow down our system. 

## Wait Group

The  above problem can be solved by making use of **WaitGroup** inside **sync** package. So here we will wait for the go routine to complete and get back to us.

So its a group. What we does here is each time we spawn a go routine we'll add 1 to the group and inside our go routine we call WaitGroup.Done() so that if our function got executed nicely it subtracts 1 from wait group. Then finally we will call WaitGroup.Wait in our parent(main) function so that main function can get terminated effecive immediately all go routines get executed.


see program below

```go
package main

import "fmt"
import "sync"

var wg = sync.WaitGroup{}

func main () {
    wg.Add(1)
    go func () {
        fmt.Println("hello")
        wg.Done()
    }()
    wg.Wait()
}

```
**OUTPUTS**
```
hello
```

Lets see an intresting phenomenon

```go
package main

import "fmt"
import "sync"

var wg = sync.WaitGroup {}
var counter = 0

func main() {

	for i:=0; i<10; i++ {
		wg.Add(2)
		go sayHello()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello #%v \n", counter)
	wg.Done()
}

func increment() {
	counter++
	wg.Done()
}
```

**OUTPUTS**
```
Hello #0 
Hello #4 
Hello #5 
Hello #3 
Hello #4 
Hello #1 
Hello #8 
Hello #9 
Hello #6 
Hello #5 
```

Each time we execute above program the value will get changed. The reason for such a behaviour is that each go routine is rasing to finish its task since its all run in parallel.

So the solution to this is to synchronise these two functions together.

This can be acheived by Mutex. Mutex is basically a lock the application is going to honour. A simple Mutex is locked or unlocked. So if a Mutex is locked and something is trying to manipulate its value, the something has to wait till the Mutex is in unlocked state so that the something can obtain the Mutex lock itself and perform the alteration. So what we can do with it is we can protect parts of our code so that only one entity can perfom alteration to our code at a time and typically why we use it is to protect out data.

with RWMutex as many things can read our data but only one thing can write to that data, and if anything is reading we can't write to it at all.

## GOMAXPROCS

```go
package main

import "fmt"
import "runtime"

func main () {
    fmt.Printf("Threads: %v", runtime.GOMAXPROCS)
}
```

What this prints out is the number of threads availabe to process. By default it will be equal to operating system cores. We can change those values by passing in as parameter to `GOMAXPROCS`. So effectively `runtime.GOMAXPROCS(1)` will make the system run in single threaded mode.


## Checking for race condition with help of Compliler

We can make use of compiler to check for race condition. What we can do here is run the program as

```
go run -race src/main.go
```