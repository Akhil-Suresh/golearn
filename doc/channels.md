# Channels

Most language were designed with born with single processor core in mind. When concurrency and parallesinm come in play the really can't bolted down. So you  need to use third party apps to go with concurrency and all.

## BASICS

when we are working with channels we are almost possibly working in context of go routines the reason being, Channels are designed to synchronize data flow between go routines. Channels are define as `ch := make(chan int)`. 
So when we create a channel it is stictly typed so int channel can recive only int data and likewise.


```go
package main

import "fmt"
import "sync"

var wg = sync.WaitGroup {}
var counter = 0

func main() {
	ch := make(chan int)
	wg.Add(2)
	go func() {
		i := <- ch
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		ch <- 42
		wg.Done()
	}()
	wg.Wait()
}
```

**OUTPUT**
```
42
```

## RESTRICTING DATA FLOW
Sometimes we want to restrict data flow between channels. So we need a channel that can only receive and one that can only send data. So by that way 

```go
package main

import "fmt"
import "sync"

var wg = sync.WaitGroup {}
var counter = 0

func main() {
	ch := make(chan int)
	wg.Add(2)
	go func(ch <- chan int) {
		i := <- ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	go func(ch chan <- int) {
		ch <- 42
		wg.Done()
	}(ch)
	wg.Wait()
}
// OUTPUTS
// 42
```

## BUFFERED CHANNELS

Buffered channels are implemented by adding a buffer value to channel creation.
`ch := make(chan int, 50)` So now we can possibly do the following program which otherwise have paniced saying deadlock.

```go

```

## CLOSING CHANNELS
## FOR...RANGE LOOPS WITH CHANNELS
## SELECT STATEMENTS

