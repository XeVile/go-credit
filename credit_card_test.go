package main

import (
	"testing"
)

func TestVisa(t *testing.T) {
	want := "Verified (Visa)"

	t.Run("Visa: 4111 1111 4555 1142", func(t *testing.T) {
		got := CheckCard("4111 1111 4555 1142")
		assert(t, got, want)
	})
	t.Run("Visa: 4111 1120 1426 7661", func(t *testing.T) {
		got := CheckCard("4111 1120 1426 7661")
		assert(t, got, want)
	})
	t.Run("Visa: 4001 5900 0000 0001", func(t *testing.T) {
		got := CheckCard("4001 5900 0000 0001")
		assert(t, got, want)
	})
}

func TestAmex(t *testing.T) {
	want := "Verified (American Express)"

	t.Run("Amex: 3700 0000 0000 002", func(t *testing.T) {
		got := CheckCard("3700 0000 0000 002")
		assert(t, got, want)
	})
	t.Run("Amex: 3700 0000 0100 018", func(t *testing.T) {
		got := CheckCard("3700 0000 0100 018")
		assert(t, got, want)
	})
	t.Run("Amex: 3774 0011 1111 115", func(t *testing.T) {
		got := CheckCard("3774 0011 1111 115")
		assert(t, got, want)
	})
}

func TestMastercard(t *testing.T) {
	want := "Verified (Mastercard)"

	t.Run("Mastercard: 2222 4000 7000 0005", func(t *testing.T) {
		got := CheckCard("2222 4000 7000 0005")
		assert(t, got, want)
	})
	t.Run("Mastercard: 5555 5555 5555 4444", func(t *testing.T) {
		got := CheckCard("5555 5555 5555 4444")
		assert(t, got, want)
	})
	t.Run("Mastercard: 2222 4000 5000 0009", func(t *testing.T) {
		got := CheckCard("2222 4000 5000 0009")
		assert(t, got, want)
	})
}

func assert(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Got: %q\nWant: %q", got, want)
	}
}
