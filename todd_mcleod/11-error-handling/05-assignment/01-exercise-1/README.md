- Using the following code, instead of using the blank identiifer, make sure the code is checking and handling the error

```go
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "Leon",
		Last:    "Low",
		Sayings: []string{"Shaken, not stirred", "Any alst wishes?", "Never say never"},
	}

	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))
}
```