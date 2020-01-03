# Maps

Map as the name suggests is the map between the key value pair.

```go
package main

import "fmt"

func main() {
  shapeEdgeMap := map[string]int{
    "Triangle": 3,
    "Square":   4,
    "Pentagon": 5,
    "Hexagon":  6,
    "Heptagon": 7,
  }
  fmt.Println(shapeEdgeMap)
  fmt.Printf("Hexagon has %v sides", shapeEdgeMap["Hexagon"])
}

/*
  map[Heptagon:7 Hexagon:6 Pentagon:5 Square:4 Triangle:3]
  Hexagon has 6 sides
*/
```

We can make use of the __make__ keyword

```go
package main

import "fmt"

func main() {
  shapeEdgeMap = make(map[string]int)
  shapeEdgeMap = map[string]int{
    "Triangle": 3,
    "Square":   4,
    "Pentagon": 5,
    "Hexagon":  6,
    "Heptagon": 7,
  }
}
```

We can add data to the Map with similar syntax

```go
shapeEdgeMap["Octagon"] = 8
```

The slice or array return the data in the order we add data to it. But Map is little different. Map doesn't guarantee any order.

```go
package main

import "fmt"

func main() {
  shapeEdgeMap := map[string]int{
    "Triangle": 3,
    "Square":   4,
    "Pentagon": 5,
    "Hexagon":  6,
    "Heptagon": 7,
  }
  fmt.Println(shapeEdgeMap)
  shapeEdgeMap["Octagon"] = 8
  fmt.Println(shapeEdgeMap)
  
}

/*
  OUTPUT
  map[Heptagon:7 Hexagon:6 Pentagon:5 Square:4 Triangle:3]
  map[Heptagon:7 Hexagon:6 Octagon:8 Pentagon:5 Square:4 Triangle:3]
*/
```

We can delete the data from the map using __delete__ method

```go
delete(shapeEdgeMap, "Pentagon")
```
There is something little interesting here. What if we try to get the value of the deleted key?

```go
shapeEdgeMap["Pentagon"]

// OUTPUTS
// 0
```
This is what happens with maps, By using maps when we try to return value for the key thats not in the map it returns 0,

This can be checked by

comma ok syntax

```go
checkValue, ok := shapeEdgeMap["Enneagon"]
fmt.Printf("%v, %v", checkValue, ok)

/*
OUTPUTS
0, false
*/
```

So by checking the value of the ok we can determine the value exits in map or not.

We can get the length of the map by the len function.

```go
fmt.Println(len(shapeEdgeMap))

// OUTPUTS
// 7
```

Unlike arrays maps are referenced. So if we assign it to another variable and edit the data its gonna effect all referenced values.

```go
package main

import "fmt"

func main() {
  shapeEdgeMap := map[string]int{
    "Triangle": 3,
    "Square":   4,
    "Pentagon": 5,
    "Hexagon":  6,
    "Heptagon": 7,
  }
  polygonMap := shapeEdgeMap
  fmt.Println(shapeEdgeMap)
  delete(polygonMap, "Square")
  fmt.Println(shapeEdgeMap)
}

/*
OUTPUTS
map[Heptagon:7 Hexagon:6 Pentagon:5 Square:4 Triangle:3]
map[Heptagon:7 Hexagon:6 Pentagon:5 Triangle:3]
*/
```


# Structs

Struct is a collection type.

```go
package main

import "fmt"

type Employee struct {
  id int
  name string
  department []string
}


func main() {
  anEmployee := Employee{
    id: 1001,
    name: "Akhil",
    department: []string{
      "Open Source",
      "Enterprise",
    },
  }
  fmt.Println(anEmployee)
}

/*
OUTPUTS
{1001 Akhil [Open Source Enterprise]}
*/
```
It gathers the information together that are related to one concept, in a very flexible way.

So any types of data can be mixed together using struct. All other collections had to have a consistent datatype, unlike struct.
