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

Obviously the name of the interface should represent what that interface is going to do for you. There is one special case, If you got single method interface, which are very common in the go lang then the convention is to name the interface with the method name + er. Like *Writer*

