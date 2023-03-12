package ui

import (
	"math/rand"
	"time"
)

const idLen = 7
const alphabetLen = 64

// `A-Za-z0-9_-`
var alphabet = [alphabetLen]byte{
	'A', 'B', 'C', 'D', 'E',
	'F', 'G', 'H', 'I', 'J',
	'K', 'L', 'M', 'N', 'O',
	'P', 'Q', 'R', 'S', 'T',
	'U', 'V', 'W', 'X', 'Y',
	'Z', 'a', 'b', 'c', 'd',
	'e', 'f', 'g', 'h', 'i',
	'j', 'k', 'l', 'm', 'n',
	'o', 'p', 'q', 'r', 's',
	't', 'u', 'v', 'w', 'x',
	'y', 'z', '0', '1', '2',
	'3', '4', '5', '6', '7',
	'8', '9', '-', '_',
}

func init() {
	idRand = rand.New(rand.NewSource(time.Now().Unix()))
}

var idRand *rand.Rand

func NewId() string {
	id := make([]byte, idLen)
	for i := 0; i < idLen; i++ {
		n := idRand.Intn(alphabetLen)
		id[i] = alphabet[n]
	}
	return string(id)
}
