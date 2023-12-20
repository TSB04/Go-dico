package main

import (
	"dico/dicorep"
)

func main() {

	d := dicorep.New("dico.txt")

	d.Add("golang", "Le meilleur langage de programmation du monde")
	d.Add("rust", "Un langage de programmation système")
	d.Add("html", "Un langage de programmation de balisage utilisé pour créer des pages Web")
	d.Add("java", "Un langage de programmation orienté objet")
	println("---------------------------------------------------------")

	// entrez un mot et affichez sa définition.
	d.Get("rust")
	println("---------------------------------------------------------")

	// Supprimez un mot du dictionnaire.
	d.Remove("java")
	println("---------------------------------------------------------")

	// Affichez tous les mots et définitions du dictionnaire.
	d.Display()
	println("---------------------------------------------------------")
}
