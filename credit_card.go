package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func stripSeperator(card_number string) []byte {
	data := bytes.Split([]byte(card_number), []byte(" "))
	stripped := bytes.Join(data, []byte(""))

	return stripped
}

func CheckCard(card_number string) string {
	/*
	  1st = Network Of Visa
	  2nd - 6th = Bank Branch
	  6th - Last = Uniquely Generated
	  Last = Payload

	  Visa:
	  * Begins with 4
	  * 13 - 16 Digits

	  Amex:
	  * Begins with 3 then 4 or 7
	  * 15 Digits

	  Mastercard:
	  * Begins with 5 then range of 1-5
	  * Corporate within 2221-2720
	  * 16 Digits
	*/
	data := stripSeperator(card_number)

	filterMap := map[string]*regexp.Regexp{
		"amex": regexp.MustCompile("^3[47][0-9]{13}$"),
		"mc":   regexp.MustCompile("^2(?:2[2-9][1-9]|[3-6][0-9]{2}|7[0-1][0-9]|720)[0-9]{12}$|^5[1-5][0-9]{14}$"),
		"visa": regexp.MustCompile("^4[0-9]{12}(?:[0-9]{3})?$"),
	}

	networkMap := map[string]string{
		"amex": "American Express",
		"mc":   "Mastercard",
		"visa": "Visa",
	}

	for network, filter := range filterMap {
		if filter.Match(data) {
			return "Verified (" + networkMap[network] + ")"
		}
	}

	return "False Card"
}

func main() {
	fmt.Println(CheckCard("4001 5900 0000 0001"))
}
