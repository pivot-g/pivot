package utility

import (
	"github.com/pivot-g/pivot/pivot/log"
	"github.com/sethvargo/go-password/password"
)

func GenPassword(length *int, numDigits *int, numSymbols *int, noUpper *bool, allowRepeat *bool) (string, error) {
	gen, err := password.NewGenerator(&password.GeneratorInput{
		Symbols: "!@#$%^()",
	})
	if err != nil {
		log.Warn(err)
		return "", err
	}

	passwd, err := gen.Generate(*length, *numDigits, *numSymbols, *noUpper, *allowRepeat)
	return passwd, err
}
