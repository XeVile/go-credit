package main

import (
	"bytes"
	"regexp"
	"strconv"
	"time"
)

func stripSeperator(card_number string) []byte {
	data := bytes.Split([]byte(card_number), []byte(" "))
	stripped := bytes.Join(data, []byte(""))

	return stripped
}

func Card(card_number string) string {
	/*
	  1st = Major Industry Identifier (MII)
	  2nd - 6th = Issuer Identification Number (IIN) {5}
	  7th - Second Last = Uniquely Generated per Individual account {9-12}
	  Last = Payload

	  System checks IIN first to get Issuer

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
	validation := Validator(data)

	/*
	  Compile filter of Regular Expressions:
	  * Amex: For both (34 & 37)
	  * Mastercard: For Corporate 2017 updated (2221-2720)
	  				and Normal (51-55)
	  * Visa: Normal (4xxxx)
	*/
	filterMap := map[string]*regexp.Regexp{
		"amex": regexp.MustCompile(`^3[47][0-9]{13}$`),
		"mc":   regexp.MustCompile(`^2(?:2[2-9][1-9]|[3-6][0-9]{2}|7[0-1][0-9]|720)[0-9]{12}$|^5[1-5][0-9]{14}$`),
		"visa": regexp.MustCompile(`^4[0-9]{12}(?:[0-9]{3})?$`),
	}

	networkMap := map[string]string{
		"amex": "American Express",
		"mc":   "Mastercard",
		"visa": "Visa",
	}

	/*
	 	Filter the Card as per the map and print the network name
	    only IF the validation passes
	    ...
	*/
	for network, filter := range filterMap {
		if filter.Match(data) {
			if validation {
				return networkMap[network]
			}
		}
	}

	// ...otherwise return empty string
	return ""
}

func Date(card_type string, expiryM, expiryY int) string {
	/*
	   Use Regular Expressions to transfer all logic to Backend.
	   Frontend will not have to deal with any input validation

	   Validation will be done regardless of wrong type input,
	   only type of int will come out to be verified on the
	   conditions being met.
	*/
	filterY := regexp.MustCompile("^20[0-9]{2}$")
	filterM := regexp.MustCompile("^(?:[1-9]|1[012])$")
	curY, curM, _ := time.Now().Date()

	if filterM.Match([]byte(strconv.Itoa(expiryM))) &&
		filterY.Match([]byte(strconv.Itoa(expiryY))) {

		if len(card_type) > 1 {
			if curY < expiryY {
				return "Verified (" + card_type + ")"
			} else if curY == expiryY {
				if int(curM) <= expiryM {
					return "Verified (" + card_type + ")"
				}
			}

			return "Expired (" + card_type + ")"
		}

		return "False Card"
	}

	return "Please check Expiry Date"
}
