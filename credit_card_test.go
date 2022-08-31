package main

import (
	"testing"
)

func TestCreditCard(t *testing.T) {
	wantV := "Verified (Visa)"
	wantA := "Verified (American Express)"
	wantMC := "Verified (Mastercard)"

	t.Run("Card: 4111 1111 4555 1142\nEXP: 03/2030", func(t *testing.T) {
		got := CheckCard("4111 1111 4555 1142", "03/2030")
		assert(t, got, wantV)
	})

	t.Run("Card: 4001 5900 0000 0001\nEXP: 03/2030", func(t *testing.T) {
		got := CheckCard("4001 5900 0000 0001", "03/2030")
		assert(t, got, wantV)
	})

	t.Run("Card: 3700 0000 0000 002\nEXP: 03/2030", func(t *testing.T) {
		got := CheckCard("3700 0000 0000 002", "03/2030")
		assert(t, got, wantA)
	})

	t.Run("Card: 3774 0011 1111 115\nEXP: 03/2030", func(t *testing.T) {
		got := CheckCard("3774 0011 1111 115", "03/2030")
		assert(t, got, wantA)
	})

	t.Run("Card: 2222 4000 7000 0005\nEXP: 03/2030", func(t *testing.T) {
		got := CheckCard("2222 4000 7000 0005", "03/2030")
		assert(t, got, wantMC)
	})
	t.Run("Card: 5555 5555 5555 4444\nEXP: 03/2030", func(t *testing.T) {
		got := CheckCard("5555 5555 5555 4444", "03/2030")
		assert(t, got, wantMC)
	})
	t.Run("Card: 2222 4000 5000 0009\nEXP: 03/2030", func(t *testing.T) {
		got := CheckCard("2222 4000 5000 0009", "03/2030")
		assert(t, got, wantMC)
	})
}

func assert(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Got: %q\nWant: %q", got, want)
	}
}
