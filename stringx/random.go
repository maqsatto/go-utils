package stringx

import "math/rand"

func GenRandomString(len int) string {

	randString := make([]rune, 0, len)

	for i := 0; i < len; i++ {
		randString = append(randString, rune(rand.Intn(96)+32))
	}

	return string(randString)

}

func GenRandomPassword(upperCaseLetters, lowerCaseLetters, specialSymbols, numbers int) string {
	randPassword := make([]rune, 0, (upperCaseLetters + lowerCaseLetters + specialSymbols + numbers))
	for i := 0; i < specialSymbols; i++ {
		randPassword = append(randPassword, genRandomSpecialSymbol())
	}
	for i := 0; i < upperCaseLetters; i++ {
		randPassword = append(randPassword, genRandomLowerCaseLetter())
	}
	for i := 0; i < upperCaseLetters; i++ {
		randPassword = append(randPassword, genRandomUpperCaseLetter())
	}
	for i := 0; i < upperCaseLetters; i++ {
		randPassword = append(randPassword, genRandomNumber())
	}
	rand.Shuffle(len(randPassword), func(i, j int) {
		randPassword[i], randPassword[j] = randPassword[j], randPassword[i]
	})
	return string(randPassword)
}
