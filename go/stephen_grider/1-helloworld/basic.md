## Go CLI

|     Code     |                                                 Description                                                  |
| :----------: | :----------------------------------------------------------------------------------------------------------: |
|   `go run`   |                                     Compiles and executes 1 or 2 files.                                      |
|  `go build`  | Compiles a bunch of go source code files. <br> Creates a `main.exe*` file. Use `./main.exe` to run the file. |
|   `go fmt`   |                         Formats all the code in each file in the current directory.                          |
| `go install` |                                      Compiles and "installs" a package.                                      |
|   `go get`   |                           Downloads the raw source code of someone else's package.                           |
|  `go test`   |                             Runs any tests associated with the current project.                              |

`go build main.go`

## What does "package main" mean?

Package == Project == Workspace

- Declare the package on the first line in every file.

|            Executable            |                            Reusable                            |
| :------------------------------: | :------------------------------------------------------------: |
| Generates a file that we can run | Code used as 'helpers'. <br> Good place to put reusable logic. |

- `package main` with the word "main" means that we want to create a file that we can **run**. (Executable Package)
- must create a function called main

- `package anothername` with another name will not produce a run-able file.

## What does 'import "fmt"' mean?

- import all the functions inside the package fmt (format).
- import standard library packages <a href="golang.org/pkg">Click Here</a>
- import reusable package.

## What does 'func' mean?

- **func**: function.

## How is the main.go file organized?

1. Package Declaration
2. Import other packages that we need
3. Declare function, tell "Go" to do things.

# Quiz
1. What is the purpose of a package in Go? To group together code with a similar purpose.
2. What is the special name we use for a package to tell Go that we want it to be turned into a file that can be executed? main.
3. The one requirement of packages named "main" is that we... Define a function named "main", which is ran automatically when the program runs.
4. Why do we use "import" statements? To give our package access to code written in another package.
