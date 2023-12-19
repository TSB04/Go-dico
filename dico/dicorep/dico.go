package dicorep

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
}

// get a word and its definition from the dictionary with the key word

func (d *Dico) Get(word string) {
	if d.couple[word] != "" {
		println("value :" + d.couple[word])
	}
}

// remove a word and its definition from the dictionary with the key word

func (d *Dico) Remove(word string) {
	if d.couple[word] != "" {
		d.couple[word] = ""
	}
}

// display all words and definitions from the dictionary

func (d *Dico) Display() {
	for key, value := range d.couple {
		println(key + " : " + value)
	}
}

