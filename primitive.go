package randomizer

import (
	"github.com/nu7hatch/gouuid"
	"math/rand"
	"time"
)

var (
	letters     = []rune("abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ")
	letters_len = len(letters)

	chars     = append(letters, []rune("123456789")...)
	chars_len = len(chars)
)

// INIT ===========================================================================================

func init() {
	rand.Seed(time.Now().UnixNano())
}

// FUNCTIONS ======================================================================================

func ID() (string, error) {
	u4, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return u4.String(), nil
}

func String(l int) string {
	runes := make([]rune, l)
	for i := 0; i < l; i++ {
		runes[i] = chars[Int(0, chars_len-1)]
	}
	return string(runes)
}

func Letter() rune {
	return letters[Int(0, letters_len-1)]
}

func Bool() bool {
	return Int(0, 1) == 1
}

func Percent() float32 {
	return rand.Float32()
}

func Int(min, max int) int {
	if min > max {
		tmp := min
		min = max
		max = tmp
	}
	return min + rand.Intn(max-min)
}
