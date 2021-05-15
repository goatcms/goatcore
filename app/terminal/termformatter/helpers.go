package termformatter

import (
	"fmt"
	"math"
	"strings"
)

func SeparateWords(in string) []string {
	return strings.Fields(in)
}

func SeparateLines(words []string, lineMax int) (line []string, rest []string) {
	if lineMax > maxLineLength {
		panic(fmt.Errorf("max line length is %d", maxLineLength))
	}
	i := 0
	count := 0
	if len(words) > 0 && len(words[0]) > lineMax {
		word := words[0]
		return []string{word[:lineMax-1] + "-"}, append([]string{"-" + word[lineMax-1:]}, words[1:]...)
	}
	for i = 0; i < len(words) && count+i < lineMax; i++ {
		count += len(words[i])
	}
	if count+i-1 > lineMax {
		i--
	}
	return words[:i], words[i:]
}

func ToLeft(words []string, lineMax int) string {
	if lineMax > maxLineLength {
		panic(fmt.Errorf("max line length is %d", maxLineLength))
	}
	line := strings.Join(words, " ")
	if len(line) > lineMax {
		panic(fmt.Errorf("'%s' is longer then maximum line length (%d)", line, lineMax))
	}
	return line + emptyLine[:lineMax-len(line)]
}

func ToRight(words []string, lineMax int) string {
	if lineMax > maxLineLength {
		panic(fmt.Errorf("max line length is %d", maxLineLength))
	}
	line := strings.Join(words, " ")
	if len(line) > lineMax {
		panic(fmt.Errorf("'%s' is longer then maximum line length (%d)", line, lineMax))
	}
	return emptyLine[:lineMax-len(line)] + line
}

func Justify(words []string, lineMax int) (result string) {
	var extraSpacePer int
	if len(words) == 0 {
		return emptyLine[:lineMax]
	}
	if len(words) == 1 {
		return words[0] + emptyLine[:lineMax-len(words[0])]
	}
	length := 0
	for _, word := range words {
		length += len(word)
	}
	spacesAreas := len(words) - 1
	spacesCounter := lineMax - length - spacesAreas
	if spacesCounter < 0 {
		panic("Not enought spaces")
	}
	//spacesCounter -= 1 // prevent skip space at the end
	//words[len(words)-1] = " " + words[len(words)-1]
	spaceLength := spacesCounter / spacesAreas
	extraSpace := spacesCounter % spacesAreas
	if extraSpace != 0 {
		extraSpacePer = int(math.Ceil(float64(spacesAreas) / float64(extraSpace)))
	}
	result = words[0]
	lestWord := words[len(words)-1]
	for i, word := range words[1 : len(words)-1] {
		var wordSpaces = 0
		if extraSpacePer != 0 && (i+1)%extraSpacePer == 0 {
			wordSpaces = spaceLength + 2
		} else {
			wordSpaces = spaceLength + 1
		}
		result += emptyLine[:wordSpaces] + word
	}
	endSpaces := (lineMax - len(result)) - len(lestWord)
	result += emptyLine[:endSpaces] + lestWord
	return result
}
