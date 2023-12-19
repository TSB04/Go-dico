package main

import (
	"dico/dicorep"
)

func main() {

	// // Créez un nouveau dictionnaire vide.
	// m := make(map[string]string)

	// // Ajoutez quelques mots et définitions au dictionnaire.
	// m["golang"] = "Le meilleur langage de programmation du monde"
	// m["rust"] = "Un langage de programmation système"
	// m["html"] = "Un langage de programmation de balisage utilisé pour créer des pages Web"
	// m["java"] = "Un langage de programmation orienté objet"
	// m["c"] = "Un langage de programmation procédural"

	// //call function ---------------------------------------------------------
	// // add a word and its definition to the dictionary
	// add(m, "python", "Un langage de programmation interprété")

	// // entrez un mot et affichez sa définition.
	// get(m, "golang")

	// // Supprimez un mot du dictionnaire.
	// remove(m, "java")

	// // Affichez tous les mots et définitions du dictionnaire.
	// display(m)
	// //---------------------------------------------------------

	//call method ---------------------------------------------------------
	// add a word and its definition to the dictionary
	
	d := dicorep.New()

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

// // function ---------------------------------------------------------
// // add a word and its definition to the dictionary
// func add(m map[string]string, key string, value string) {
// 	m[key] = value
// }

// func get(m map[string]string, key string) {
// 	// print key value
// 	fmt.Println(m[key])
// }

// func remove(m map[string]string, key string) {
// 	delete(m, key)
// }

// func display(m map[string]string) {
// 	//list
// 	l := list.New()
// 	for key, value := range m {
// 		l.PushBack(key)
// 		l.PushBack(value)
// 	}
// 	for e := l.Front(); e != nil; e = e.Next() {
// 		fmt.Println(e.Value)
// 	}
// }