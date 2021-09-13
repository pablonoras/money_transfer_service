package validation

import "regexp"

var patternAlphaNumericAndUnder = regexp.MustCompile(`^[A-Za-z\d_]*$`)

//valAlphaNumericAndUnderScore this function validates if a string has characters alphanumeric and/or underscore
func ValAlphaNumericAndUnderScore(valor string) bool {
	if patternAlphaNumericAndUnder.MatchString(valor) && valor != "" {
		return true
	}
	return false
}
