package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func readData(path string) (map[string]int, error) {
	// Construct the path using the current working directory
	file, err := os.Open(path) // Use the relative path directly
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}

	return result, nil
}

// loading capitals.txt into memory once

type Database interface {
	GetPopulation(name string) int
}

// singleton, using lowercase so it can only be called in the defined package, prevent people from calling it more than once
type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

// sync.Once init() -- thread safety
// laziness, guaranteed using sync.Once

var once sync.Once // ensure only one call
var instance *singletonDatabase

func GetSingletonDatabase() *singletonDatabase {
	// function to be called only once
	once.Do(func() {
		fmt.Println("Reading from capitals.txt ...")
		caps, err := readData("capitals.txt")
		if err != nil {
			panic(err)
		}

		db := singletonDatabase{}
		db.capitals = caps
		instance = &db
	})
	return instance
}

func GetTotalPopulation(cities []string) int {
	result := 0
	for _, city := range cities {
		// depending on concrete instance of singleton database, violates DIP (should depend on abstractions - interface), unable to unit test
		result += GetSingletonDatabase().GetPopulation(city)
	}
	return result
}

// Singleton and Dependency Inversion Principle
func GetTotalPopulationEx(db Database, cities []string) int {
	result := 0
	for _, city := range cities {
		// depending on concrete instance of singleton database, violates DIP (should depend on abstractions - interface)
		result += db.GetPopulation(city)
	}
	return result
}

type DummyDatabase struct {
	dummyData map[string]int
}

func (d *DummyDatabase) GetPopulation(name string) int {
	if len(d.dummyData) == 0 {
		d.dummyData = map[string]int{
			"alpha": 1,
			"beta":  2,
			"gamma": 3,
		}
	}
	return d.dummyData[name] // now we can write a proper unit test, instead of accessing an actual database
}

func main() {
	db := GetSingletonDatabase()
	pop := db.GetPopulation("Seoul")

	fmt.Println("Population of Seoul =", pop)

	pop = db.GetPopulation("Mexico City")
	fmt.Println("Population of Mexico City =", pop)

	// cities := []string{"Seoul", "Mexico City"}
	// tp := GetTotalPopulation(cities)
	// ok := tp == (17500000 + 17400000)
	// fmt.Println(ok)

	names := []string{"alpha", "gamma"}
	tp := GetTotalPopulationEx(&DummyDatabase{}, names)
	fmt.Println(tp == 4)

	// singleton is not scary
	// best to use abstraction (interface) together with singleton, instead of depending on singleton instance so we can write proper unit tests
}
