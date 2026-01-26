package stringx

import (
	"fmt"
	"math/rand"
)

func genRandomChar128() rune {
	return rune(rand.Intn(96) + 32)
}

func genRandomSpecialSymbol() rune {
	i := rand.Intn(4)
	switch i {
	case 0:
		return rune(rand.Intn(14) + 33)
	case 1:
		return rune(rand.Intn(6) + 57)
	case 2:
		return rune(rand.Intn(6) + 91)
	case 3:
		return rune(rand.Intn(4) + 123)
	default:
		return 0
	}
}

func genRandomUpperCaseLetter() rune {
	return rune(rand.Intn(26) + 65)
}

func genRandomLowerCaseLetter() rune {
	return rune(rand.Intn(26) + 97)
}
func genRandomNumber() rune {
	return rune(rand.Intn(10) + 48)
}

func TestBasicGen() {
	for i := 0; i < 10; i++ {
		fmt.Print(string(genRandomChar128()))
	}
	fmt.Printf("\n")
	for i := 0; i < 10; i++ {
		fmt.Print(string(genRandomLowerCaseLetter()))
	}
	fmt.Printf("\n")
	for i := 0; i < 10; i++ {
		fmt.Print(string(genRandomUpperCaseLetter()))
	}
	fmt.Printf("\n")
	for i := 0; i < 10; i++ {
		fmt.Print(string(genRandomNumber()))
	}
	fmt.Printf("\n")
}
