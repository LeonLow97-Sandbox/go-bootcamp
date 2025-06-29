## Cards Project Instruction

|    function     |                              Description                              |
| :-------------: | :-------------------------------------------------------------------: |
|     newDeck     | Create a list of playing cards. <br> Essentially an array of strings. |
|      print      |               Log out the contents of a deck of cards.                |
|     shuffle     |                   Shuffles all the cards in a deck.                   |
|      deal       |                       Create a 'hand' of cards.                       |
|   saveToFile    |         Save a list of cards to a file on the local machine.          |
| newDeckFromFile |             Load a list of cards from the local machine.              |

## Basic Go Types (static language)

|  Type   |          Example           |
| :-----: | :------------------------: |
|  bool   |         true false         |
| string  |  "Hi!" "How's it going?"   |
|   int   |       0 -10000 99999       |
| float64 | 10.000001 0.00009 -100.003 |

## Initialization of variables

- Can initialize and then later assign a value to a variable.
- Can initialize a variable outside of a function, just cannot assign a value to it.
- Initializing variable: `card := "Ace of Diamonds"` or `var card string = "Ace of Diamonds"`
- Reassign variables: `card = "Two of Hearts"`

## Array and Slice

- **Array**: Fixed length list of things.
- **Slice**: An array that can grow or shrink. (add or subtract cards out of slice).
- Array and slices must have the **same data type**.
- Declare a slice: `cards := []string{"Ace of Diamonds", newCard()}`
- Add an element at the end of slice: `cards = append(cards, "Six of Spades")`

## For Loop

```
for i, card := range cards {
  fmt.Println(i, card)
}
```

- `index`: index of this element in the array
- `card`: current card we are iterating over
- `range cards`: take the slice of 'cards' and loop over it.

## Additional Information

- Files in the same package do not have to be imported into each other.

## Custom Type and Function

|                 code                  |                                                       description                                                       |
| :-----------------------------------: | :---------------------------------------------------------------------------------------------------------------------: |
|         `type deck []string`          | Tell Go we want to create an <br> array of strings and add a bunch <br> of functions specifically made to work with it. |
| Functions with 'deck' as a 'receiver' |             A function with a receiver is like a "method" <br> - a function that belongs to an "instance".              |

- `go run main.go deck.go`

## Receiver on Function

```
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
```

- Receiver: `d` first letter of `deck` (by convention).
- `func (d deck) print()`: Any variable of type "deck" now gets access to the "print" method.
- `d`: "cards" variable. The actual copy of the deck we are working with is avaible in the fuction as a variable called "d".
- `deck`: Every variable of type 'deck' can call this function on itself.
- By creating a new type with a function that has a receiver, we are adding a "method" to any value of that type.

## To skip the 'index' on for loop

```
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}
```

## Slicing

- `array[startIndexIncluding : upToExcluding]`
- `array[:2]`
- `array[index]`

## Write file

`func WriteFile(filename string, data []byte, perm fs.FileMode) error`

|   string    |             byte slice             |
| :---------: | :--------------------------------: |
| "Hi there!" | [72 105 32 116 104 101 114 101 33] |

- To turn byte to string, visit <a href=https://www.asciitable.com>Ascii Table</a>

#### Type Conversion in Go

Code: `[]byte("Hi There!")`

- `[]byte`: type we want
- `("Hi There!")`: value we have
- Basically converting the type from string to byte slice.

## Join Function on `Strings`

- Join concatenates the elements of its first argument to create a single string.

  `func Join(elems []string, sep string) string`

- The separator string `sep` is placed between elements in the resulting string.

## Read File

`func ReadFile(filename string) ([]byte, error)` <br>
E.g., bs, err := ioutil.ReadFile(filename)

- ReadFile reads the file named by filename and returns the contents.
- A successful call returns err == nil

| name  |                 code                 |                                                                                         description                                                                                          |
| :---: | :----------------------------------: | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------: |
| Exit  |        `func Exit(code int)`         | Exit causes the program to exit with the given status code. <br> Conventionally, code zero indicates success, non-zero an error. <br> The program terminates immediately. E.g., `os.Exit(1)` |
| Split | `func Split(s, sep string) []string` |                                   Split slices s into all substrings separated by sep and <br> returns a slice of the substrings between those separators.                                   |
| Intn  |        `func Intn(n int) int`        |                                                                                 returns a number from 0 to n                                                                                 |

## Testing with Go

- To make a test, create a new file ending in "\_test.go"
  `deck_test.go`
- To Run all tests in a package, run the command...
  `go test`
- test files also need to include `package main`

## To mod Go with git init

`go mod init cards`
