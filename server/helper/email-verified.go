package helper

import "regexp"

func IsValidEmail(email string) bool {
	req := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@gmail\.com$`)
	return req.MatchString(email)
}
