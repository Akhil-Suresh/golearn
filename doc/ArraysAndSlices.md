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

Arrays are considered values in go. so when you copy an array we are actually copying value, Not underlying values, unlike other languages.

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

The problem with above implementation arises when we have millions of data, This is then going to slow our program down.
This behavior can be changed so that the copying array is a reference to main array. This can be done using address of operator.

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

// OUTPUTS
// [1, 10, 3]
// [1, 10, 3]
```

## Slices

Slices are very similar to array, so everything we do with array can be done with slices as well with some exceptions. Slices can be initialized as below.

```go
a := []int{1, 2, 3}
```

But there are some slight changes as well.

For slices we have `len` function to get the length
Also slices have `cap` function which defines its capacity.

```go
package main

import "fmt"

func main() {
  a := []int{1, 2, 3}

  fmt.Printf("Length: %v\n", len(a))
  fmt.Printf("Capacity: %v\n", cap(a))
}

// OUTPUT
// Length: 3
// Capacity: 3
```

Here capacity and length are same, But it can be different as well. Slice is just a projection of array

```go
package main

import "fmt"

func main() {
  a := []int{0, 1, 2, 3, 4}  
  b := a[2:4] // Slice of all elements

  fmt.Printf("b: %v\n", b)
  fmt.Printf("Length: %v\n", len(b))
  fmt.Printf("Capacity: %v\n", cap(b))
}
// OUTPUTS
// b: [2 3]
// Length: 2
// Capacity: 3

```

Unlike arrays slices are reference type data. Which means if we assign a slice to another both will point to same underlying data.

```go
package main

import "fmt"

func main() {
  a := []int{1, 2, 3}
  fmt.Printf("%v\n", a)

  b := a
  b[0] = 0

  fmt.Printf("%v\n", a)
  fmt.Printf("%v\n", b)
}

// OUTPUT
// [1 2 3]
// [0 2 3]
// [0 2 3]
```


```go
package main

import (
  "fmt"
)

func main() {
  a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
  b := a[:]
  c := a[3:]
  d := a[:6]
  e := a[3:6]

  fmt.Printf("a : %v\n", a)
  fmt.Printf("b : %v\n", b)
  fmt.Printf("c : %v\n", c)
  fmt.Printf("d : %v\n", d)
  fmt.Printf("d : %v\n", e)

}

/*
OUTPUT

a : [0 1 2 3 4 5 6 7 8 9]
b : [0 1 2 3 4 5 6 7 8 9]
c : [3 4 5 6 7 8 9]
d : [0 1 2 3 4 5]
d : [3 4 5]

*/
```

### The make function

Slices can be created by make function also

```go
package main

import "fmt"

func main() {
  a := make([]int, 3)
  fmt.Println(a)
  fmt.Printf("Length: %v\n", len(a))
  fmt.Printf("Capacity: %v\n", cap(a))
}
```

We know that slice is just __slice__ of an __array__. So by using make function we can actually define the size of the main array itself. It can be defined as third parameter to the array

```go
package main

import "fmt"

func main() {
  a := make([]int, 3, 100)
  fmt.Println(a)
  fmt.Printf("Length: %v\n", len(a))
  fmt.Printf("Capacity: %v\n", cap(a))
}
```

### Appending data to an array

Once a slice is created it can grow dynamically. You can check at the below program and see how the size of the underlying array get increased. The value of **cap** shows the size of the underlying array.


```go
package main

import "fmt"

func main() {
  a := []int{}
  fmt.Println(a)
  fmt.Printf("Length: %v\n", len(a))
  fmt.Printf("Capacity: %v\n\n", cap(a))

  a = append(a, 1)
  fmt.Println(a)
  fmt.Printf("Length: %v\n", len(a))
  fmt.Printf("Capacity: %v\n\n", cap(a))

  a = append(a, 2, 3, 4, 5)
  fmt.Println(a)
  fmt.Printf("Length: %v\n", len(a))
  fmt.Printf("Capacity: %v\n\n", cap(a))

  a = append(a, []int{6, 7, 8, 9}...)   // Slices can be appended by using ... operator
  fmt.Println(a)  
  fmt.Printf("Length: %v\n", len(a))
  fmt.Printf("Capacity: %v\n", cap(a))
}
```

**OUTPUTS**
```
[]
Length: 0
Capacity: 0

[1]
Length: 1
Capacity: 1

[1 2 3 4 5]
Length: 5
Capacity: 6

[1 2 3 4 5 6 7 8 9]
Length: 9
Capacity: 12
```