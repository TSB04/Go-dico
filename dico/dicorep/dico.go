package dicorep

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// dico structure
type Dico struct {
	Filename string
	Couple   map[string]string
}

func New(filename string) *Dico {
	d := &Dico{
		Filename: filename,
		Couple:   make(map[string]string),
	}
	d.loadFromFile()
	return d
}

func (d *Dico) loadFromFile() {
	file, err := os.Open(d.Filename)
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
			d.Couple[key] = value
		}
	}
}

func (d *Dico) saveToFile() error {
	file, err := os.Create(d.Filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for key, value := range d.Couple {
		_, err := fmt.Fprintf(writer, "%s : %s\n", key, value)
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

func (d *Dico) Add(key string, value string) {
	d.Couple[key] = value
	err := d.saveToFile()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(key + " has been added")
}

// get a word and its definition from the dictionary with the key word
func (d *Dico) Get(word string) {
	if d.Couple[word] != "" {
		fmt.Println(word + " : " + d.Couple[word])
	} else {
		fmt.Println("This word doesn't exist")
	}
}

// remove a word and its definition from the dictionary
func (d *Dico) Remove(word string) {
	if d.Couple[word] != "" {
		delete(d.Couple, word)
		err := d.saveToFile()
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Println(word + " has been removed")
	} else {
		fmt.Println("This word doesn't exist")
	}
}

// get all words and definitions from the dictionary
func (d *Dico) GetAll() {
	for key, value := range d.Couple {
		fmt.Println(key + " : " + value)
	}
}

// Open the file and read it in notepad
func (d *Dico) OpenInNotpad() {
	cmd := exec.Command("notepad.exe", d.Filename)
	cmd.Run()
}
