# Arrays

Arrays can be declared as

```go
// Defining array and not initializing
var device  [10]string
var ages    [10]int16

// Defining and initializing array
device := [3]string{ "mouse", "monitor", "keyboard" }

// Defining and initializing array without defining array size
device := [...]string{ "mouse", "monitor", "keyboard" }
```

We can assign value to array like

```go
device[0] = "speaker"
```

When we create an array without initializing it. It will be initialized with zero(default) value.

One of the intresting fact based on it is that we can do

```go
var students [3]string
students[2] = "Akhil"
students[0] = "Jithesh"
fmt.Printf("%v", students)

// OUTPUTS
// [Jithesh  Akhil]
```

We can get the length of the array by `len` function.
We already know an array's size, but there may be situation,
where we need to revaluate it.

```go
len(students)

//Outputs
// 3
```

Arrays are considered values in go. so when you copy an array we are actually copying value, Not undlining values, just like other languages

```go
package main

import "fmt"

func main() {
  myArray := [...]int{1, 2, 3}
  tempArray := myArray
  tempArray[1] = 10
  fmt.Println(myArray)
  fmt.Println(tempArray)
}

// OUTPUT
// [1 2 3]
// [1 10 3]
```

The problem with above implementaion arises when we have millions of data, This is then going to slow our program down.

We can change this behaviour. This can be done by

```go
package main

import "fmt"

func main() {
  myArray := [...]int{1, 2, 3}
  tempArray := &myArray
  tempArray[1] = 10
  fmt.Println(myArray)
  fmt.Println(tempArray)
}
```

## Slices

Slices are very similar to array, so everything we do with array can be done with slices as well with some exceptions. Slices can be initialized by code below.

```go
a := []int{1, 2, 3}
```

But there are some slight changes as well.

For slices we have `len` function to get the length
Also slices have `cap` function which defines its capacity.

Unlike arrays slices are reference type data. Which means if we assign a slice to another both will point to same underlying data.

```go

```