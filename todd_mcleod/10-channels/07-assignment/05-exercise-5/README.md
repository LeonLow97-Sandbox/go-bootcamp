- With the following code, show the comma ok idiom.

```golang
package main

import "fmt"

func main() {
	c := make(chan int)

	v, ok := 
	fmt.Println(v, ok)

	close(c)

	v, ok = 
	fmt.Println(v, ok)
}
```
