package alphabet

type (
	Alphabet []rune
)

var EnLower = New(
	'a', 'b', 'c', 'd', 'e',
	'f', 'g', 'h', 'i', 'j',
	'k', 'l', 'm', 'n', 'o',
	'p', 'q', 'r', 's', 't',
	'u', 'v', 'w', 'x', 'y',
	'z',
	)

var EnUpper = New(
	'A', 'B', 'C', 'D', 'E',
	'F', 'G', 'H', 'I', 'J',
	'K', 'L', 'M', 'N', 'O',
	'P', 'Q', 'R', 'S', 'T',
	'U', 'V', 'W', 'X', 'Y',
	'Z',
	)

func New(runes ...rune) Alphabet {
	return runes
}

func (a Alphabet) Pos(r rune) int {
	for i, ar := range a {
		if ar == r {
			return i
		}
	}
	return -1
}

func (a Alphabet) LetterAt(i int) rune {
	return a[i]
}

func (a Alphabet) Len() int {
	return len(a)
}