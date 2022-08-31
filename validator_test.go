package main

/*
 * Testing & Benchmarking :: Luhn
 *
 * Benchmarking of using byte vs int conversion
 * Int converted is performant - negligibly
 *
 * Testing 4 types of Card, 3 Networks
 */

import (
	"testing"
)

func TestLuhn(t *testing.T) {
	want := true
	t.Run("4001 5900 0000 0001", func(t *testing.T) {
		got := Validator(stripSeperator("4001 5900 0000 0001"))
		assertLuhn(t, got, want)
	})

	t.Run("3700 0000 0100 018", func(t *testing.T) {
		got := Validator(stripSeperator("3700 0000 0100 018"))
		assertLuhn(t, got, want)
	})

	t.Run("2222 4000 7000 0005", func(t *testing.T) {
		got := Validator(stripSeperator("2222 4000 7000 0005"))
		assertLuhn(t, got, want)
	})

	t.Run("5555 5555 5555 4444", func(t *testing.T) {
		got := Validator(stripSeperator("5555 5555 5555 4444"))
		assertLuhn(t, got, want)
	})
}

func assertLuhn(t testing.TB, got, want bool) {
	t.Helper()

	if got != want {
		t.Errorf("Got: %t\nWant: %t", got, want)
	}
}

func BenchmarkLuhn(b *testing.B) {
	data := stripSeperator("4001590000000001")

	b.Run("Byte Luhn", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			LuhnByte(data)
		}
	})

	b.Run("Int Luhn", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Validator(data)
		}
	})
}
