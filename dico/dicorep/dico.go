package dicorep

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// dico structure
type Dico struct {
	filename string
	couple   map[string]string
}

func New(filename string) *Dico {
	d := &Dico{
		filename: filename,
		couple:   make(map[string]string),
	}
	d.loadFromFile()
	return d
}

func (d *Dico) loadFromFile() {
	file, err := os.Open(d.filename)
	if err != nil {
		return // File doesn't exist or error reading, so an empty map will be initialized
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			d.couple[key] = value
		}
	}
}

func (d *Dico) saveToFile() error {
	file, err := os.Create(d.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for key, value := range d.couple {
		_, err := fmt.Fprintf(writer, "%s : %s\n", key, value)
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

func (d *Dico) Add(key string, value string) {
	d.couple[key] = value
	err := d.saveToFile()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(key + " has been added")
}

// Other methods (Get, Remove, Display) remain unchanged...

// get a word and its definition from the dictionary with the key word
func (d *Dico) Get(word string) {
	if d.couple[word] != "" {
		fmt.Println(word + " : " + d.couple[word])
	} else {
		fmt.Println("This word doesn't exist")
	}
}

// remove a word and its definition from the dictionary
func (d *Dico) Remove(word string) {
	if d.couple[word] != "" {
		delete(d.couple, word)
		err := d.saveToFile()
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Println(word + " has been removed")
	} else {
		fmt.Println("This word doesn't exist")
	}
}

// display all words and definitions from the dictionary
func (d *Dico) Display() {
	for key, value := range d.couple {
		fmt.Println(key + " : " + value)
	}
}
