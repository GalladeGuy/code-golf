package routes

import (
	"math/rand"
	"strings"
)

var morseMap = map[rune]string{
	'A': "▄ ▄▄▄",
	'B': "▄▄▄ ▄ ▄ ▄",
	'C': "▄▄▄ ▄ ▄▄▄ ▄",
	'D': "▄▄▄ ▄ ▄",
	'E': "▄",
	'F': "▄ ▄ ▄▄▄ ▄",
	'G': "▄▄▄ ▄▄▄ ▄",
	'H': "▄ ▄ ▄ ▄",
	'I': "▄ ▄",
	'J': "▄ ▄▄▄ ▄▄▄ ▄▄▄",
	'K': "▄▄▄ ▄ ▄▄▄",
	'L': "▄ ▄▄▄ ▄ ▄",
	'M': "▄▄▄ ▄▄▄",
	'N': "▄▄▄ ▄",
	'O': "▄▄▄ ▄▄▄ ▄▄▄",
	'P': "▄ ▄▄▄ ▄▄▄ ▄",
	'Q': "▄▄▄ ▄▄▄ ▄ ▄▄▄",
	'R': "▄ ▄▄▄ ▄",
	'S': "▄ ▄ ▄",
	'T': "▄▄▄",
	'U': "▄ ▄ ▄▄▄",
	'V': "▄ ▄ ▄ ▄▄▄",
	'W': "▄ ▄▄▄ ▄▄▄",
	'X': "▄▄▄ ▄ ▄ ▄▄▄",
	'Y': "▄▄▄ ▄ ▄▄▄ ▄▄▄",
	'Z': "▄▄▄ ▄▄▄ ▄ ▄",
	'1': "▄ ▄▄▄ ▄▄▄ ▄▄▄ ▄▄▄",
	'2': "▄ ▄ ▄▄▄ ▄▄▄ ▄▄▄",
	'3': "▄ ▄ ▄ ▄▄▄ ▄▄▄",
	'4': "▄ ▄ ▄ ▄ ▄▄▄",
	'5': "▄ ▄ ▄ ▄ ▄",
	'6': "▄▄▄ ▄ ▄ ▄ ▄",
	'7': "▄▄▄ ▄▄▄ ▄ ▄ ▄",
	'8': "▄▄▄ ▄▄▄ ▄▄▄ ▄ ▄",
	'9': "▄▄▄ ▄▄▄ ▄▄▄ ▄▄▄ ▄",
	'0': "▄▄▄ ▄▄▄ ▄▄▄ ▄▄▄ ▄▄▄",
	' ': "    ",
}

func morse(reverse bool) (args []string, out string) {
	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	// Shuffle the digits.
	for i := range digits {
		j := rand.Intn(i + 1)
		digits[i], digits[j] = digits[j], digits[i]
	}

	words := []string{
		"BUD",
		"FOR",
		"JIGS",
		"NYMPH",
		"QUICK",
		"VEX",
		"WALTZ",
		strings.Join(digits, ""),
	}

	// Shuffle the words.
	for i := range words {
		j := rand.Intn(i + 1)
		words[i], words[j] = words[j], words[i]
	}

	args = []string{strings.Join(words, " ")}

	for _, char := range args[0] {
		out += morseMap[char] + "   "
	}

	// Knock off the trailing three spaces.
	out = out[:len(out)-3]

	if reverse {
		args[0], out = out, args[0]
	}

	return
}
