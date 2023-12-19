package main

import (
	"fmt"
)

func main() {

	// Créez un nouveau dictionnaire vide.
	m := make(map[string]string)

	// Ajoutez quelques mots et définitions au dictionnaire.
	m["golang"] = "Le meilleur langage de programmation du monde"
    m["rust"] = "Un langage de programmation système"
	m["html"] = "Un langage de programmation de balisage utilisé pour créer des pages Web"
	m["java"] = "Un langage de programmation orienté objet"
	m["c"] = "Un langage de programmation procédural"
	
	// affichez le dictionnaire.
	fmt.Println(m)

	//Utilisez la méthode Get pour afficher la définition d'un mot spécifique.
	fmt.Println(m["golang"])

	//Utilisez la méthode Remove pour supprimer un mot du dictionnaire.
	delete(m, "c")

	//Appelez la méthode List pour obtenir la liste triée des mots et de leurs définitions.
	fmt.Println(m)

}
