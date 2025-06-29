## Zero Value

|  Type  | Zero Value |
| :----: | :--------: |
| string |    " "     |
|  int   |     0      |
|  int   |     0      |
|  bool  |   false    |

## %+v

`fmt.Printf("%+v", alex)`: Prints out the struct properties.

## Defining a custom struct

```
type person struct {
	firstName string
	lastName string
}
```

## Assigning values to struct

- First method

```
leon := person{firstName: "Leon", lastName: "Low"}
fmt.Println(leon)
```

- Second method

```
var alex person
alex.firstName = "Alex"
alex.lastName = "Anderson"

fmt.Println(alex)
fmt.Printf("%+v", alex)
```

## Embedded Struct

- Creating another embedded data struct

```
type contactInfo struct {
	email string
	zipCode int
}
type person struct {
	firstName string
	lastName string
	contact contactInfo
}
```

- Printing out the struct

```
	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contact: contactInfo {
			email: "jim@gmail.com",
			zipCode: 94000,
		},
	}

	fmt.Printf("%+v", jim)
```

## RAM

- Go makes a copy and made the copy available to the update function, but it did not update.

| Address |           Value            |               Code               |
| :-----: | :------------------------: | :------------------------------: |
|  0000   |                            |                                  |
|  0001   | person{firstName:"Jim"...} |        `jim := person()`         |
|  0002   |                            |                                  |
|  0003   | person{firstName:"Jim"...} | `func (p person) updateName() {` |

## Pointers

- `&variable`: Give me the **memory address** of the value this variable is pointing at. (0001)

```
	jimPointer := &jim
	jimPointer.updateName("Jimmy")
```

- `*person`: Type description - it means we are working with a pointer to a person.
- `*pointerToPerson`: Give me the **value** this memory address is pointing at.

```
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
```

- Turn value into address with **&value**
- Turn address into value with **\*address**

- Modifying structs requires pointers.
- Modifying slice does not require pointers. (points to the same Array in the same memory). These are reference types.

|                       Value                       |                Reference Types                |
| :-----------------------------------------------: | :-------------------------------------------: |
|        `int, float, string, bool, structs`        | `slices, maps, channels, pointers, functions` |
| Use pointers to change these things in a function |     Don't worry about pointers with these     |
