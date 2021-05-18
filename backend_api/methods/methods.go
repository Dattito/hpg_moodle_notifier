package methods

import (
	"strings"

	"github.com/ttacon/libphonenumber"
)

func IsValidGermanNumber(number string) bool {
	num, err := libphonenumber.Parse(number, "DE")
	if err != nil {
		return false
	}
	return num.GetCountryCode() == 49
}

func ToInternationalNumber(number string) string {
	num, err := libphonenumber.Parse(number, "DE")
	if err != nil {
		panic(err)
	}
	return strings.ReplaceAll(libphonenumber.Format(num, libphonenumber.INTERNATIONAL), " ", "")
}

func ValidateMoodleToken(token string) bool {
	return true
}
