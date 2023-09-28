package main

import (
	"fmt"
	"math/rand"
)

type c []int

func Min(x c) int {
	min := x[0]
	for i := 1; i < len(x); i++ {
		if x[i] < min {
			min = x[i]
		}
	}
	if min == 0 {
		return 1
	}
	return min
}

func main() {
	var cordes c = make(c, 0)
	// génération aléatoire du tableau de cordes
	longueur := rand.Intn(10)
	//fmt.Printf("longueur= %d\n", longueur)
	for i := 0; i < longueur; i++ {
		nombre := rand.Intn(10)
		//fmt.Printf("Nombre = %d\n", nombre)
		cordes = append(cordes, nombre)
		//fmt.Printf("chiffre n° %d = %d\n", i, cordes[i])
	}
	/// Affichage du tableau de cordes
	for i := 0; i < longueur; i++ {
		//fmt.Printf("%d", cordes[i])
	}
	//fmt.Printf("\n")
	// Calcul de chaque itération pour conduire le tableau cordes à 0
	termine := false
	nombredetours := 0
	// boucle d'itération
	for !termine {
		termine = true
		chiffremin := Min(cordes)
		nombredetours++
		for i := 0; i < len(cordes); i++ {
			if cordes[i] >= chiffremin {
				cordes[i] -= chiffremin
				termine = false
			} else {
				termine = true
			}
		}
		/// Affichage du tableau de cordes
		//for j := 0; j < len(cordes); j++ {
		//	fmt.Printf("%d", cordes[j])
		//}
		//fmt.Printf("\n")
		//time.Sleep(10 * time.Second)

	}

	// Affichage du nombre de tours
	fmt.Printf("Nombre de tours = %d\n", nombredetours)

}
