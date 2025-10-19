package main

import (
	"fmt"
	"os"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	j.entries = append(j.entries[:index], j.entries[index+1:]...)
}

// Separation of Concerns have to reside in another construct else you violate SRP
// Anti-Pattern: God Object
// This breaks SRP because Journal now has two reasons to change:
// 1. Changes to how journal entries are managed
// 2. Changes to how journals are persisted (should be handled by a different class)
// func (j *Journal) Save(filename string) {
// 	_ = os.WriteFile(filename, []byte(j.String()), 0644)
// }
// func (j *Journal) Load(filename string) {
// 	// ...
// }
// func (j *Journal) LoadFromWeb(url *url.URL) {
// 	// ...
// }

// Separating persistence logic into its own function
var LineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = os.WriteFile(filename, []byte(strings.Join(j.entries, LineSeparator)), 0644)
}

type Persistence struct {
	lineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = os.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func SRP() {
	// Follows SRP because Journal only handles journal entries
	j := Journal{}
	j.AddEntry("I cried today.")
	j.AddEntry("I ate a bug.")
	fmt.Println(j.String())

	// Persistence is handled separately
	SaveToFile(&j, "journal.txt")

	// Using Persistence struct
	p := Persistence{lineSeparator: "\r\n"}
	p.SaveToFile(&j, "journal2.txt")

	// Cleanup
	_ = os.Remove("journal.txt")
	_ = os.Remove("journal2.txt")
}
