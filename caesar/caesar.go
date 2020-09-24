package caesar

import (
	"github.com/RidgeA/switch-to-go-m5/alphabet"
	"golang.org/x/text/runes"
)

type Operation int
type PositionCalculationFunc func(pos, shift, length int) int

func selectAlphabet(r rune, alphabets ...alphabet.Alphabet) (selectedAlphabet alphabet.Alphabet, pos int) {
	pos = -1
	selectedAlphabet = nil
	for _, a := range alphabets {
		pos = a.Pos(r)
		if pos != -1 {
			selectedAlphabet = a
			break
		}
	}
	return
}

func transform(targetPos PositionCalculationFunc, shift int, alphabets ...alphabet.Alphabet) runes.Transformer {

	return runes.Map(func(r rune) rune {

		sa, pos := selectAlphabet(r, alphabets...)

		if sa == nil {
			return r
		}

		pos = targetPos(pos, shift, sa.Len())
		return sa.LetterAt(pos)
	})

}

func Encode(shift int, alphabets ...alphabet.Alphabet) runes.Transformer {
	targetPos := func(pos, shift, length int) int {
		return (pos + shift) % length
	}
	return transform(targetPos, shift, alphabets...)
}

func Decode(shift int, alphabets ...alphabet.Alphabet) runes.Transformer {
	targetPos := func(pos, shift, length int) int {
		return (length + pos - shift) % length
	}
	return transform(targetPos, shift, alphabets...)
}
