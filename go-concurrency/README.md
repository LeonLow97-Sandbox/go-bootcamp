# Course Links

|                   Course Name                   |  Instructor   |                                         Course Link                                         |
| :---------------------------------------------: | :-----------: | :-----------------------------------------------------------------------------------------: |
| Udemy - Working with Concurrency in Go (Golang) | Trevor Sawler |         [Link](https://www.udemy.com/course/working-with-concurrency-in-go-golang/)         |
|            YouTube - Go Concurrency             | Kantan Coding | [Link](https://www.youtube.com/watch?v=qyM8Pi1KiiM&list=PL7g1jYj15RUNqJStuwE9SCmeOKpgxC0HP) |

## Golang Currency (Trevor Sawler) Course Outline

- Basic types in the `sync` package: mutexes (semaphores), and wait groups.
- 3 computer science programs:
  - The Producer/Consumer problem
  - The Dining Philosopher problem
  - The Sleeping Barber problem
- Real-world scenario,
  - Where we build a subset of larger (imaginary) problem where a user of a service wants to subscribe and buy one of a series of available subscriptions.
  - When they purchase a subscription, we will generate invoice, send an email, and generate PDF manual and send that to the user.
  - We will do these things concurrently.
- Writing tests for Go Concurrency
