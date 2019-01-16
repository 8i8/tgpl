package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

const (
	IN = iota
	IS
	ISCONTROL
	ISDIGIT
	ISGRAPHIC
	ISLETTER
	ISLOWER
	ISMARK
	ISNUMBER
	ISONEOF
	ISPRINT
	ISPUNCT
	ISSPACE
	ISSYMBOL
	ISTITLE
	ISUPPER
	LEN // Used to obtain the length of this list of constants.
)

var (
	utfstr = []string{
		"Is",
		"Is",
		"Control",
		"Digit",
		"Graphic",
		"Letter",
		"Lower",
		"Mark",
		"Number",
		"Oneof",
		"Print",
		"Punct",
		"Space",
		"Symbol",
		"Title",
		"Upper",
	}
)

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	var utftyp [LEN]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		utftyp = runeCountType(r, utftyp)
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Print("\ntype\tcount\n")
	for i, n := range utftyp {
		if n > 0 {
			fmt.Printf("%s\t%d\n", utfstr[i], n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

func runeCountType(r rune, utftyp [LEN]int) [LEN]int {

	//if In(r , ranges ...*RangeTable) {
	//}
	//if Is(rangeTab *RangeTable, r ) {
	//}
	if unicode.IsControl(r) {
		utftyp[ISCONTROL]++
	}
	if unicode.IsDigit(r) {
		utftyp[ISDIGIT]++
	}
	if unicode.IsGraphic(r) {
		utftyp[ISGRAPHIC]++
	}
	if unicode.IsLetter(r) {
		utftyp[ISLETTER]++
	}
	if unicode.IsLower(r) {
		utftyp[ISLOWER]++
	}
	if unicode.IsMark(r) {
		utftyp[ISMARK]++
	}
	if unicode.IsNumber(r) {
		utftyp[ISNUMBER]++
	}
	//if unicode.IsOneOf(ranges []*RangeTable, r ) {
	//}
	if unicode.IsPrint(r) {
		utftyp[ISPRINT]++
	}
	if unicode.IsPunct(r) {
		utftyp[ISPUNCT]++
	}
	if unicode.IsSpace(r) {
		utftyp[ISSPACE]++
	}
	if unicode.IsSymbol(r) {
		utftyp[ISSYMBOL]++
	}
	if unicode.IsTitle(r) {
		utftyp[ISTITLE]++
	}
	if unicode.IsUpper(r) {
		utftyp[ISUPPER]++
	}
	return utftyp
}
