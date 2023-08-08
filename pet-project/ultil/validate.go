package ultil

import (
	"net/mail"
	"regexp"
)

func EmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func PhoneValid(phone string) bool {
	match, _ := regexp.MatchString("`(84|0[3|5|7|8|9])+([0-9]{8})\\b`gm", phone)
	return match
}
