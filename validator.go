package main

/*
 * Luhn Algorithm implementation
 *
 * LuhnByte() -> uses List of byte and returns byte
 * Validator() -> Converts each byte in the list first before usage
 */

var m2 = [...]int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}

func LuhnByte(data []byte) bool {
	n := len(data)
	odd := n & 1
	var p byte

	for i, v := range data[:n-1] {
		if i&1 == odd {
			p += byte(m2[int(v-48)])
		} else {
			p += v - 48
		}
	}

	valid := (10 - p%10) == (data[n-1] - 48)

	return valid
}

func Validator(data []byte) bool {
	/*
	  	1. Get length of byte-list
	    2. Check if Odd
	    3. Bitwise & op for each index
	    4. Increment payload based of condition
	    5. Match "check digit"
	*/
	n := len(data)
	odd := n & 1
	var payload int

	for i, v := range data[:n-1] {
		if i&1 == odd {
			payload += m2[(v - 48)]
		} else {
			payload += int((v - 48))
		}
	}

	verf := (10 - payload%10) == int(data[n-1]-48)

	return verf
}
