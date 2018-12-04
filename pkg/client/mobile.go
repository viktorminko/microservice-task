package client

import (
	"regexp"
	"strings"
)

const defaultCountryCode = "44"

//Mobile represents client mobile phone number
type Mobile struct {
	CountryCode string
	Number      string
}

//NewMobile creates new mobile struct with provided number
func NewMobile(number string) *Mobile {
	return &Mobile{Number: number}
}

//SetCountryCode sets country code for current mobile
func (m *Mobile) SetCountryCode(c string) {
	m.CountryCode = c
}

//GetFormattedNumber returns formatted mobile number with country code as prefix
func (m *Mobile) GetFormattedNumber() (string, error) {
	//leave only numbers in the phone
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		return "", err
	}
	//remove spaces
	r := strings.Replace(m.Number, " ", "", -1)

	//remove non digits
	r = reg.ReplaceAllString(r, "")

	//Add country code
	c := defaultCountryCode
	if len(m.CountryCode) > 0 {
		c = m.CountryCode
	}

	r = "+" + c + r
	return r, nil
}
