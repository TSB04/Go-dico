package main

import (
	"dico/dicorep"
)

func main() {
	// call the New function to create a new dictionary.
	d := dicorep.New("dico.txt")

	d.Add("golang", "Le meilleur langage de programmation du monde")
	d.Add("rust", "Un langage de programmation système")
	d.Add("html", "Un langage de programmation de balisage utilisé pour créer des pages Web")
	d.Add("java", "Un langage de programmation orienté objet")
	println("---------------------------------------------------------")

	// call the Get method to get a word from the dictionary.
	d.Get("rust")
	println("---------------------------------------------------------")

	// call the Remove method to remove a word from the dictionary.
	d.Remove("java")
	println("---------------------------------------------------------")

	// call the Display method to display all words from the dictionary.
	d.Display()
	println("---------------------------------------------------------")

	// call the OpenInNotpad method to open the dictionary file in notepad.
	d.OpenInNotpad()
}
