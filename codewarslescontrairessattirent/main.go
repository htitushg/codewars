package main

import (
	"fmt"
	"regexp"
	"slices"
	"sort"
	"strings"
	"time"
	"unicode"
)

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
func CalculateYears(years int) (result [3]int) {
	//age_cat := []int{15, 9, 4}
	//age_dog := []int{15, 9, 5}
	switch years {
	case 1:
		result = [3]int{1, 15, 15}
	case 2:
		result = [3]int{2, 24, 24}
	default:
		result = [3]int{years, 24 + 5*(years-2), 24 + 5*(years-2)}
	}
	return result
}
func contains(s rune) bool {
	voyelles := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	for _, a := range voyelles {
		if a == s {
			return true
		}
	}
	return false
}

func Disemvowel(comment string) string {

	resultat := ""
	for i := 0; i < len(comment); i++ {

		if !contains(rune(comment[i])) {
			resultat += string(comment[i])
		}
	}
	return resultat
}
func Disemvowel2(comment string) string {
	return regexp.MustCompile(`(?i)[aeiou]`).ReplaceAllString(comment, "")
}

// Recherche la 1ère occurence d'une chaine dans s et renvoie l'index de cette lettre
func Index(s string, tofind string) int {
	phrase := []rune(s)
	mot := []rune(tofind)
	trouve := 0
	motencours := false
	index := -1
	// on recherche une lettre de mot dans phrase
	y := 0
	for i := 0; i < len(phrase); i++ {
		if mot[y] == phrase[i] {
			// On a trouvé la lettre du mot dans la phrase
			if !motencours {
				// si on a pas déjà trouvé une lettre on mémorise l'index de la lettre trouvée
				index = i
			}
			motencours = true
			trouve++
			if trouve == len(mot) && motencours {
				return index
			} else {
				y++
			}
		} else {
			trouve = 0
			motencours = false
			index = -1
			y = 0
		}

	}
	return index
}

// Renvoie un tableau de chaines qui contient les chaines contenues dans s et qui sont séparées par la chaine sep
func DNAtoRNA2(dna string) string {
	phrase := []rune(dna)
	if len(dna) >= 4 {
		j := 0
		for i := 0; i < len(phrase); i++ {
			if phrase[i+j] == 'G' {
				if phrase[i+1+j] == 'C' {
					if phrase[i+2+j] == 'A' {
						if phrase[i+3+j] == 'T' {
							phrase[i+3+j] = 'U'
						}
					}
				}
			}
			j++
			if i+j+3 >= len(phrase) {
				break
			}
		}

	} else {
		return ""
	}
	return string(phrase)
}
func DNAtoRNA(s string) string {
	//Replace(s, old, new string, n int) string
	return strings.Replace(s, "T", "U", -1)
}
func DNAtoRNA3(dna string) string {
	var response string

	for _, letter := range dna {
		if letter == 'T' {
			letter = 'U'
		}
		response += string(letter)
	}

	return response
}

// Renvoie -1 si a<b, 0 si a=b, 1 si a>b
func Compare(a, b string) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}
func SamallestIntegerFinder2(numbers []int) int {
	smallest := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < smallest {
			smallest = numbers[i]
		}
	}
	return smallest
}
func SamallestIntegerFinder(numbers []int) int {
	smallest := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < smallest {
			smallest = numbers[i]
		}
	}
	return smallest
}
func Rps2(p1, p2 string) string {
	if p1 == p2 {
		return "Draw!"
	} else if p1 == "" {
		return "<string>: Player 2 won!"
	} else if p2 == "" {
		return "<string>: Player 1 won!"
	} else if p1 == "rock" && p2 == "scissors" {
		return "Player 1 won!"
	} else if p1 == "rock" && p2 == "paper" {
		return "Player 2 won!"
	} else if p1 == "scissors" && p2 == "paper" {
		return "Player 1 won!"
	} else if p1 == "scissors" && p2 == "rock" {
		return "Player 2 won!"
	} else if p1 == "paper" && p2 == "scissors" {
		return "Player 2 won!"
	} else if p1 == "paper" && p2 == "rock" {
		return "Player 1 won!"
	} else {
		return "Draw!"
	}
}

var m = map[string]string{"rock": "paper", "paper": "scissors", "scissors": "rock"}

func Rps(a, b string) string {
	if a == b {
		return "Draw!"
	}
	if m[a] == b {
		return "Player 2 won!"
	}
	return "Player 1 won!"
}
func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
func ReverseWords(str string) string {
	var rev string
	var word string

	for _, i := range str {
		if i == ' ' {
			rev = rev + word + " " // Adds word and space to result
			word = ""              // Empties word variable
		} else {
			word = string(i) + word // Adds letter to temporary word variable
		}
	}

	return rev + word // reverse those words
}
func ReverseWords2(str string) string {
	new := strings.Split(str, " ")
	for i := 0; i < len(new); i++ {
		new[i] = Reverse(new[i])
	}
	str = ""
	for i := 0; i < len(new)-1; i++ {
		str += new[i] + " "
	}
	str += new[len(new)-1]
	return str // reverse those words
}
func mod(a, b int) int {
	return (a%b + b) % b
}
func RoundToNext51(a int) int {
	fmt.Printf("a modulo 5 = %d\n", a%5)
	return (a%5 + 5) % 5
}
func RoundToNext52(n int) int {
	negatif := false
	fmt.Printf("n modulo 5 = %d\n", n%5)
	if n < 0 {
		n = -n
		negatif = true

	}
	if n == 0 {
		return 0
	}
	if negatif {
		if n%5 == 0 {
			return -n
		} else {
			fmt.Printf("Negatif : n %d modulo 5 = %d\n", n, n%5)
			return (n%5 - n)
		}
	} else {
		if n%5 == 0 {
			return n
		} else {
			fmt.Printf("Positif : n %d modulo 5 = %d\n", n, n%5)
			return n + 5 - n%5
		}

	}
}
func RoundToNext53(n int) int {
	for n%5 != 0 {
		n++
	}
	return n
}
func RoundToNext5(n int) int {
	return n + (5-n%5)%5
}

func SortString(s string) string {
	r := []rune(s)
	// Tri à bulle
	tableautrie := false
	for i := len(r) - 1; i > -1; i-- {
		tableautrie = true
		for j := 0; j < i; j++ {
			//fmt.Printf("Avant swap i= %d :r[%d]:%v,r[%d]:%v\n", i, j, string(r[j]), j+1, string(r[j+1]))
			if r[j+1] < r[j] {
				r[j], r[j+1] = r[j+1], r[j]
				tableautrie = false
			}
		}
		if tableautrie {
			return string(r)
		}
	}
	return string(r)
}
func SplitString(s string) string {
	memelettre := false
	for i := 0; i < len(s)-1; i++ {
		fmt.Printf("Avant: memelettre=%t, s[%d]= %v, s[%d]= %v, %v\n", memelettre, i, string(s[i]), i+1, string(s[i+1]), s)
		if memelettre {
			i--
		}
		if s[i] == s[i+1] {
			s = s[:i+1] + s[(i+2):]
			memelettre = true
		} else {
			memelettre = false
		}
		//fmt.Printf("Après: s[%d]= %v, s[%d]= %v, %v\n", i, string(s[i]), i+1, string(s[i+1]), s)
		time.Sleep(300 * time.Millisecond)
	}
	if s[len(s)-1] == s[len(s)-2] {
		s = s[:len(s)-1]
	}
	return s
}

func TwoToOne2(s1 string, s2 string) string {
	s1 += s2
	s1 = SortString(s1)
	fmt.Printf("S1= %s\n", s1)
	s1 = SplitString(s1)
	return s1
}
func TwoToOne(s1 string, s2 string) string {
	chars := strings.Split(s1+s2, "")
	fmt.Printf("chars=%v\n", chars)
	slices.Sort(chars)
	fmt.Printf("chars=%v\n", chars)
	result := ""
	for _, s := range chars {
		chr := string(s)
		if !strings.Contains(result, chr) {
			result = result + chr
		}
	}
	return result
}
func TwoToOne1(s1 string, s2 string) string {
	result := ""
	for _, char := range strings.Split("abcdefghijklmnopqrstuvwxyz", "") {
		if strings.Contains(s1+s2, char) {
			result += char
		}
	}
	return result
}

type MyString string

func (s MyString) IsUpperCase2() bool {
	SenRune := []rune(s)
	for i := 0; i < len(SenRune); i++ {
		if !((SenRune[i] >= 'A') && (SenRune[i] <= 'Z')) && (SenRune[i] != ' ') {
			return false
		}
	}
	return true
}
func (s MyString) IsUpperCase1() bool {
	// Your code here!
	for _, char := range s {
		if unicode.IsLower(char) {
			return false
		}
	}

	return true
}
func (s MyString) IsUpperCase() bool {
	return string(s) == strings.ToUpper(string(s))
}
func PositiveSum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		if number > 0 {
			sum += number
		}
	}
	return sum
}

// Renvoie un entier résultat de nb puissance power, si nb<0 renvoie 0
func RecursivePower(nb int, power int) int {
	if nb == 0 || power == 0 {
		return 1
	}
	if power < 0 {
		return 0
	}
	if power > 1 {
		return nb * RecursivePower(nb, power-1)
	}
	return nb
}
func PowersOfTwo2(n int) []uint64 {
	result := make([]uint64, 0)
	for i := 0; i <= n; i++ {
		fmt.Printf("n=%d RecursivePower(2,%d)= %v\n", n, i, RecursivePower(2, i))
		result = append(result, uint64(RecursivePower(2, i)))
	}
	return result
}
func PowersOfTwo(n int) (arr []uint64) {
	for i := 0; i <= n; i++ {
		arr = append(arr, 1<<i)
	}
	return
}
func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	result := (dadYearsOld - 2*sonYearsOld)
	if result < 0 {
		return -result
	}
	return result
}

//a := []int{5, 3, 4}
//sort.SliceStable(a, func(i, j int) bool {
//    return a[i] < a[j]
//})

func Gimme2(array [3]int) int {
	if array[0] < array[1] { //0<1
		if array[1] < array[2] { //0<1<2
			return 1
		} else if array[0] < array[2] { //0<1 et 2> 1 et 0<2
			return 2

		} else { //0<1 et 0>2 et 2<1
			return 0
		}
	} else { //1 <0
		if array[0] < array[2] { //1<0 et 0<2
			return 0
		} else if array[1] < array[2] { //1<0 et 2<0 et 1<2
			return 2
		}

	}
	return 1
}
func Gimme1(array [3]int) int {
	left, middle, right := array[0], array[1], array[2]
	if (left > middle && left < right) || (left < middle && left > right) {
		return 0
	}
	if (middle > left && middle < right) || (middle < left && middle > right) {
		return 1
	}
	return 2
}
func Gimme(array [3]int) int {
	arrayCopy := array

	arrSlice := arrayCopy[:]
	sort.Ints(arrSlice)
	val := arrSlice[1]

	for i, v := range array {
		if v == val {
			return i
		}
	}

	return 0
}
func MinMax2(arr []int) [2]int {
	sort.Ints(arr)
	return [2]int{arr[0], arr[len(arr)-1]}
}
func MinMax(arr ...int) []int {
	sort.Ints(arr)
	return append(arr[:1], arr[len(arr)-2:]...)
}
func FindNextSquare(sq int64) int64 {
	return -1
}

func main() {
	fmt.Printf("%v\n", MinMax(1, 2, 3, 4, 5))
	fmt.Printf("%v\n", MinMax(2334454, 5))
	//fmt.Printf("%v\n", Gimme([3]int{2, 3, 1}))
	//fmt.Printf("%v\n", Gimme([3]int{5, 10, 14}))
	//fmt.Printf("%v\n", Gimme([3]int{1, 3, 4}))
	//fmt.Printf("%v\n", Gimme([3]int{15, 10, 14}))
	//fmt.Printf("%v\n", Gimme([3]int{-4, -23, 4}))
	//fmt.Printf("%v\n", Gimme([3]int{-15, -10, 14}))
	//fmt.Printf("%v\n", TwiceAsOld(36, 7))
	//fmt.Printf("%v\n", TwiceAsOld(55, 30))
	//fmt.Printf("%v\n", RecursivePower(2, 0))
	//fmt.Printf("%v\n", RecursivePower(2, 1))
	//fmt.Printf("%v\n", RecursivePower(2, 4))
	//fmt.Printf("%v\n", PowersOfTwo(0))
	//fmt.Printf("%v\n", PowersOfTwo(1))
	//fmt.Printf("%v\n", PowersOfTwo(4))
	//fmt.Printf("%v\n", PositiveSum([]int{1, 2, 3, 4, 5}))
	//fmt.Printf("%v\n", PositiveSum([]int{-1, -2, -3, -4, -5}))
	//fmt.Printf("%v\n", PositiveSum([]int{}))
	//fmt.Printf("%v\n", MyString("aretheyhere").IsUpperCase())
	//fmt.Printf("%v\n", MyString("hello I AM DONALD").IsUpperCase())
	//fmt.Printf("%v\n", MyString("HELLO I AM DONALD").IsUpperCase())
	//fmt.Printf("%v\n", TwoToOne("aretheyhere", "yestheyarehere"))
	//fmt.Printf("%v\n", TwoToOne("aabcdefghilnoprstu", ""))
	//fmt.Printf("%v\n", TwoToOne("aehrstyy", ""))

	//fmt.Printf("%v\n", RoundToNext5(-5585))
	//fmt.Printf("%v\n", RoundToNext5(0))
	//fmt.Printf("%v\n", RoundToNext5(5585))
	//fmt.Printf("%v\n", RoundToNext5(-33601))
	//fmt.Printf("%v\n", RoundToNext5(33601))
	//fmt.Printf("%v\n", Reverse("lazy"))
	//fmt.Printf("%v\n", ReverseWords("The quick brown fox jumps over the lazy dog."))
	//fmt.Printf("%v\n", Rps("rock", "scissors"))
	//fmt.Printf("%v\n", Rps("rock", "paper"))
	//fmt.Printf("%v\n", Rps("scissors", "rock"))
	//fmt.Printf("%v\n", Rps("scissors", "paper"))
	//fmt.Printf("%v\n", Rps("paper", "rock"))
	//fmt.Printf("%v\n", Rps("paper", "scissors"))
	//fmt.Printf("%v\n", Rps("rock", "rock"))

	//fmt.Printf("%v\n", SamallestIntegerFinder([]int{34, 15, 88, 2}))
	//fmt.Printf("%v\n", SamallestIntegerFinder([]int{34, -345, -1, 100}))

	//fmt.Printf("%s\n", (DNAtoRNA("GCATGCTA")))
	//fmt.Printf("%s\n", (DNAtoRNA("TTCAGTGACCCATGCCCTCCTGCTA"))) //UUCAGUGACCCAUGCCCUCCUGCUA

	//fmt.Printf("%v\n", Disemvowel("This website is for losers LOL!"))
	//fmt.Printf("%v\n", Disemvowel("XYz?"))
	//fmt.Printf("%v\n", Disemvowel("wgtwrJjKznbGLKRFRscnbST XQDzdfnjkSsBphtHFYtRLPvVBLFVm"))
	//fmt.Printf("%v\n", (CalculateYears(1)))
	//fmt.Printf("%v\n", (CalculateYears(2)))
	//fmt.Printf("%v\n", (CalculateYears(10)))
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
	//fmt.Printf("Nombre = %v\n", Number2([][2]int{{10, 0}, {3, 5}, {5, 8}}))

}
