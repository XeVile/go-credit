package main

import (
	"strconv"
	"strings"
)

func CheckCard(card_number, expiry string) string {
	card_type := Card(card_number)
	dateString := strings.Split(expiry, "/")

	date := map[string]int{
		"month": 00,
		"year":  2021,
	}

	date["month"], _ = strconv.Atoi(dateString[0])
	date["year"], _ = strconv.Atoi(dateString[1])

	return Date(card_type, date["month"], date["year"])
}
