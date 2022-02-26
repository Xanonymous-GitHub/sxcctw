package internal

import (
	"errors"
	"math"
	"strings"
)

const (
	base62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length = uint64(len(base62))
)

func Encode(number uint64) string {
	// initialized a new string builder of size n=11
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(11)

	// write a letter to encoded string by number.
	for ; number > 0; number /= length {
		encodedBuilder.WriteByte(base62[(number % length)])
	}

	return encodedBuilder.String()
}

func Decode(encoded string) (uint64, error) {
	var number uint64

	for i, symbol := range encoded {
		// find the char position of each letter in encoded string in base62 array.
		alphabeticPosition := strings.IndexRune(base62, symbol)

		// if you find a non-base62 char, throw error
		if alphabeticPosition == -1 {
			return uint64(alphabeticPosition), errors.New("invalid character: " + string(symbol))
		}

		// calculate the origin number data and add to number.
		number += uint64(alphabeticPosition) * uint64(math.Pow(float64(length), float64(i)))
	}

	return number, nil
}
