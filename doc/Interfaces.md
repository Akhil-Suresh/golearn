# Interface

Interfaces are type just like struct or maps or...

`type <name> interface`

Interfaces don't describe data just like struct, But it defines behaviour. So instead of describing a bunch of data types inside writer interface below code, we actually going to store method definitions.

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
}

type Writer inteface {
	Write([]byte)(int, error)
}

type ConsoleWriter struct {}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
```

These are the definition of the interface. So the obvious question running in our mind is that what is the use case of this?

Well this can be seen as, Inside the main function we can define an interface and Its the responsibility of the interface to decide on what to do with the value we provide into the method.


```go
package main

import (
	"fmt"
)

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello World!"))
}

type Writer interface {
	Write([]byte)(int, error)
}

type ConsoleWriter struct {}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

//OUTPUT
// Hello World!
```

## Naming Convention

Obviously the name of the interface should represent what that interface is going to do for you. There is one special case, If you got single method interface, which are very common in the go lang then the convention is to name the interface with the *method name* + *er*. Like *Writer*

We used struct in the above example which is very common way to implement interface, But we don't need to. Any type that has a method associated with it can implement an interface.

See the below example to understand this.

```go
package main

import (
  "fmt"
)

func main() {
  myInt := IntCounter(0)
  var counter Incrementer = &myInt
  for i:=0; i<10;i++ {
    fmt.Println(counter.Increment())
  }
}

type Incrementer interface {
  Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
  *ic++
  return int(*ic)
}

```

## Composing Interfaces Together (Interface Embedding)

Composing Interfaces Together is one of the powerful concept in golang. Its a key to scalability.
We mentioned above that single method interfaces are very powerful and common in the golang because they define a very specific behavior.

The composing of Interfaces comes in situation where we need more that one method and we can't decompose the interfaces down.

```go
package main

import (
  "fmt"
)

type Area interface {
    area() (float64, error)
}

type Perimeter interface {
    perimeter() (float64, error)
}

type Shape interface {
    Area
    Perimeter
}

type Rectangle struct {
    length  float64
    breadth float64
}

func (rect *Rectangle) area() (float64, error){
    return rect.length * rect.breadth, nil
}

func (rect *Rectangle) perimeter() (float64, error){
    perimeter := 2 * (rect.length + rect.breadth)
    return perimeter, nil
}

func main() {
  var shape Shape = NewRectangle()
  area, _ := shape.area()

  fmt.Println(area)
  perimeter, _ := shape.perimeter()
  fmt.Println(perimeter)
}


func NewRectangle() *Rectangle {
  return &Rectangle{
    length: 2.0,
    breadth: 3.0,
  }
}


// OUTPUTS
// 6
// 10
```

# Empty interface

