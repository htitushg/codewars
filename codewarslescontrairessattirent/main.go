package main

import "fmt"

func LoveFunc(flower1, flower2 int) bool {
	if ((flower1%2 == 0) && !(flower2%2 == 0)) || ((flower2%2 == 0) && !(flower1%2 == 0)) {
		return true
	}
	return false
}

func CheckForFactor(base int, factor int) bool {
	return base%factor == 0
}
func GetGrade2(a, b, c int) rune {
	switch (a + b + c) / 30 {
	case 10:
		return 'A'
	case 9:
		return 'A'
	case 8:
		return 'B'
	case 7:
		return 'C'
	case 6:
		return 'D'
	default:
		return 'F'
	}
}
func GetGrade(a, b, c int) rune {
	n := (a + b + c) / 3
	if n >= 90 && n <= 100 {
		return 'A'
	} else if n >= 80 && n <= 90 {
		return 'B'
	} else if n >= 70 && n <= 80 {
		return 'C'
	} else if n >= 60 && n <= 70 {
		return 'D'
	} else if n >= 0 && n <= 60 {
		return 'F'
	}
	return 'F'
}

//var busStops [][2]int

func Number(busStops [][2]int) int {
	nombre := 0
	for i := 0; i < len(busStops); i++ {
		nombre += busStops[i][0] - busStops[i][1]
		fmt.Printf("i = %d, busStops[i][0] = %d, busStops[i][1] = %d, Nombre = %d\n", i, busStops[i][0], busStops[i][1], nombre)
	}
	return nombre // Good Luck!
}
func Number2(busStops [][2]int) (inBus int) {
	for _, stopInfo := range busStops {
		inBus += stopInfo[0] - stopInfo[1]
	}
	return // Good Luck!
}
func main() {
	//fmt.Printf("LoveFunc = %t\n", LoveFunc(1, 4))
	//fmt.Printf("LoveFunc = %t\n", LoveFunc(2, 2))
	//fmt.Printf("LoveFunc = %t\n", LoveFunc(0, 1))
	//fmt.Printf("LoveFunc = %t\n", LoveFunc(0, 0))
	//fmt.Println(CheckForFactor(10, 2))
	//fmt.Println(CheckForFactor(63, 7))
	//fmt.Println(CheckForFactor(2450, 5))
	//fmt.Println(CheckForFactor(24612, 3))
	//fmt.Println(CheckForFactor(9, 2))
	//fmt.Println(CheckForFactor(653, 7))
	//fmt.Println(CheckForFactor(2453, 5))
	//fmt.Println(CheckForFactor(24617, 3))
	//fmt.Printf("%s\n", string(GetGrade2(95, 90, 93)))
	//fmt.Printf("%s\n", string(GetGrade2(100, 85, 96)))
	//fmt.Printf("%s\n", string(GetGrade2(92, 93, 94)))
	//fmt.Printf("%s\n", string(GetGrade2(70, 70, 100)))
	//fmt.Printf("%s\n", string(GetGrade2(82, 85, 87)))
	fmt.Printf("Nombre = %v\n", Number2([][2]int{{10, 0}, {3, 5}, {5, 8}}))
}
