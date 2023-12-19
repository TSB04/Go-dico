package dicorep

import "fmt"


// dico structure
type Dico struct {
	couple map[string]string
}

// create a new empty dictionary
func New() *Dico {
	return &Dico{couple: make(map[string]string)}
}

// add a word and its definition to the dictionary
func (d *Dico) Add(key string, value string) {
	d.couple[key] = value
	fmt.Println(key + " has been added")
}

// get a word and its definition from the dictionary with the key word

func (d *Dico) Get(word string) {
	if d.couple[word] != "" {
		fmt.Println(word +" : " + d.couple[word])
	} else {
		fmt.Println("This word doesn't exist")
	}
	
}

// remove a word and its definition from the dictionary 
func (d *Dico) Remove(word string) {
	if d.couple[word] != "" {
		delete(d.couple, word)
		fmt.Println(word + " has been deleted")
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

