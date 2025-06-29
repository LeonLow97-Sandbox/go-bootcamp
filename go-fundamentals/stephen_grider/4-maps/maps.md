## Maps

- Keys must all be of the same type.
- Values must all be of the same type.
- A key does not have to be of the same type as value.

|     Maps      |
| :-----------: |
| key --> value |

- **Note**: Maps in Go are similar to 'dict' in Python and 'object' in JavaScript.

## Creating a Map

```
func main() {
	// declaring a map where keys are strings and values are strings
	colors := map[string]string {
		"red": "#ff000",
		"green": "#4bf745",
	}

	fmt.Println(colors)
}
```

## Manipulating maps

|             Declare Map             |           Value           |
| :---------------------------------: | :-----------------------: |
|   `var colors map[string]string`    |      Zero Value Map       |
| `colors := make(map[string]string)` |      Zero Value Map       |
|    `colors["white"] = "#ffffff"`    | Adding a key-value to map |
|      `delete(colors, "white")`      |  Delete key-value in map  |

## Iterating over Maps

```
func main() {
	colors := map[string]string {
		"red": "#ff000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
```

# Maps vs Structs

|                         Map                          |                            Struct                             |
| :--------------------------------------------------: | :-----------------------------------------------------------: |
|           All keys must be the same type.            |                               -                               |
|          All values must be the same type.           |               Values can be of different type.                |
|     Keys are indexed - we can iterate over them.     |                 Keys don't support indexing.                  |
| Use to represent a collection of related properties. | Use to represent a "thing" with a lot of differnt properties. |
|   Don't need to know all the keys at compile time.   |  You need to know all the different fields at compile time.   |
|                    Reference Type                    |                          Value Type                           |

# Quiz

1. Can some of the keys in a single map be of type `int` and other of type `string`?
- No

2. Can some of the keys in a single map be of type `int` and other of type `string`?
- No

3. Take a look at the following code.  What would the print statement log out?
```
package main
import "fmt"
 
func main() {
 m := map[string]string{
   "dog": "bark",
 }
 
 changeMap(m)
 
 fmt.Println(m)
}
 
func changeMap(m map[string]string) {
 m["cat"] = "purr"
}
```
- `map[dog: bark cat:purr]`


