package main

import (
	"testing"
)

func TestVisa(t *testing.T) {
	want := "Visa"

	t.Run("Visa: 4111 1111 4555 1142", func(t *testing.T) {
		got := Card("4111 1111 4555 1142")
		assert(t, got, want)
	})
	t.Run("Visa: 4111 1120 1426 7661", func(t *testing.T) {
		got := Card("4111 1120 1426 7661")
		assert(t, got, want)
	})
	t.Run("Visa: 4001 5900 0000 0001", func(t *testing.T) {
		got := Card("4001 5900 0000 0001")
		assert(t, got, want)
	})
}

func TestAmex(t *testing.T) {
	want := "American Express"

	t.Run("Amex: 3700 0000 0000 002", func(t *testing.T) {
		got := Card("3700 0000 0000 002")
		assert(t, got, want)
	})
	t.Run("Amex: 3700 0000 0100 018", func(t *testing.T) {
		got := Card("3700 0000 0100 018")
		assert(t, got, want)
	})
	t.Run("Amex: 3774 0011 1111 115", func(t *testing.T) {
		got := Card("3774 0011 1111 115")
		assert(t, got, want)
	})
}

func TestMastercard(t *testing.T) {
	want := "Mastercard"

	t.Run("Mastercard: 2222 4000 7000 0005", func(t *testing.T) {
		got := Card("2222 4000 7000 0005")
		assert(t, got, want)
	})
	t.Run("Mastercard: 5555 5555 5555 4444", func(t *testing.T) {
		got := Card("5555 5555 5555 4444")
		assert(t, got, want)
	})
	t.Run("Mastercard: 2222 4000 5000 0009", func(t *testing.T) {
		got := Card("2222 4000 5000 0009")
		assert(t, got, want)
	})
}

func TestExpiryDate(t *testing.T) {
	t.Run("Verified Mastercard", func(t *testing.T) {
		want := "Verified (Mastercard)"
		got := Date(Card("2222 4000 5000 0009"), 03, 2030)
		assert(t, got, want)
	})

	t.Run("Expired Visa", func(t *testing.T) {
		want := "Expired (Visa)"
		got := Date(Card("4001 5900 0000 0001"), 01, 2022)
		assert(t, got, want)
	})

	t.Run("False Card No.", func(t *testing.T) {
		want := "False Card"
		got := Date(Card("3774 0011 1121 115"), 01, 2030)
		assert(t, got, want)
	})

	t.Run("Wrong Expiry input", func(t *testing.T) {
		want := "Please check Expiry Date"
		got := Date(Card("3774 0011 1111 115"), 13, 2022)
		assert(t, got, want)
	})
}
