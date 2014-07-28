package randgen

import (
	"math/rand"
	"time"

	"github.com/nu7hatch/gouuid"
)

var (
	letters     = []rune("abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ")
	letters_len = len(letters)

	chars     = append(letters, []rune("123456789")...)
	chars_len = len(chars)
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// UUID returns a Universally Unique Identifier.
func UUID() (string, error) {
	u4, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return u4.String(), nil
}

// String returns a random string with the passed-in length.
func String(len int) string {
	runes := make([]rune, len)
	for i := 0; i < len; i++ {
		runes[i] = chars[Int(0, chars_len-1)]
	}
	return string(runes)
}

// Letter returns a random letter, uppercase or lowercase.
func Letter() rune {
	return letters[Int(0, letters_len-1)]
}

// Bool returns a random boolean.
func Bool() bool {
	return Int(0, 1) == 1
}

// Percent returns a random floating number between 0.0 and 1.0.
func Percent() float32 {
	return rand.Float32()
}

// Int returns a random non-negative integer between the passed-in
// min (inclusive) and max (exclusive).
func Int(min, max int) int {
	if min > max {
		min, max = max, min
	}
	return min + rand.Intn(max-min)
}

// Element returns a random element from the passed-in slice of string.
func Element(source []string) string {
	return source[rand.Intn(len(source))]
}
