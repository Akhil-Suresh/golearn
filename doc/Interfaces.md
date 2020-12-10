# Interface

Interfaces are type just like struct or maps or...

* Basics
* Composing Interfaces
* Type conversion
  - The empty interface
  - Type Switches
* Implementing with values vs. Pointers
* Best practices

# Basics

Interfaces are type just like struct or other type aliases.

Interface is defined as:

`type <InterfaceName> interface`

For instance:

```go
type Writer inteface {
	Write([]byte)(int, error)
}
```

One thing we can notice in the definition is that unlike struct (where we define data types inside) we define methods here. So Basically what a interface does is that it defines a medium saying if you are going to use this interface for sure you can have this method within it.

So basically interfaces don't describe data just like struct, But it defines behaviour. 

Inside the main function we can define an interface and It is the responsibility of the interface to decide on what to do with the value we provide into the method.


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
```

**OUTPUT**
```
Hello World!
```

Here we have created an interface **Writer** whose sole purpose is to have method **Write**. So inside main function we created a variable of type Writer and initialised a data holder(struct) for it. Then without knowing what the interface has under it sleeve we call its write method by passing in the value we got. So this can be utilised in surpisingly useful ways like by now we can implement TCPWriter or FileWriter or any on the dataholder(struct) we have.

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


```go
package main

import (
	"fmt"
	"bytes"
)

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

func main() {
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello Youtube Listners"))
	wc.Close()	
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}


func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}
	
	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}
	
  ```
  
# Empty interface

Empyt interface is interface that doesn't have a method on it. Its created on fly.