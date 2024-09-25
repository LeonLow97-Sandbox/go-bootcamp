package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

var entryCount = 0

// Journal has a single responsibility for storing and manipulating journal entries.
// It does not deal with whether those entries are saved or loaded.
type Journal struct {
	entries []string
}

// Encapsulation of Functionality
// Methods like AddEntry, RemoveEntry, String are specific to the Journal class,
// keeping entry management encapsulated
func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)

	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	// ...
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// separation of concerns for persistence of journal
// different responsibility, best not to use methods, pass Journal as argument to func
func (j *Journal) Save(filename string) {
	_ = os.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {

}

func (j *Journal) LoadFromWeb(url *url.URL) {

}

var LineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = os.WriteFile(filename, []byte(strings.Join(j.entries, LineSeparator)), 0644)
}

// Persistence struct handles the saving of the journal entries to a file,
// separating this concern from the journal
type Persistence struct {
	lineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = os.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func main() {
	// package Journal has a single primary responsibility of entries and manipulating entries
	j := Journal{}
	j.AddEntry("I cried today")
	j.AddEntry("I got promoted!")

	fmt.Println(j.String())

	// Persistence of journal should be stored in another package because it is another responsibility
	// that does not deal with entries
	//
	SaveToFile(&j, "journal.txt")

	//
	p := Persistence{" ****** \n"}
	p.SaveToFile(&j, "journal2.txt")
}
