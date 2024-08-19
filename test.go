package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var romeToArab = map[string]int{
	"I":     1,
	"II":    2,
	"III":   3,
	"IV":    4,
	"V":     5,
	"VI":    6,
	"VII":   7,
	"VIII":  8,
	"IX":    9,
	"X":     10,
	"XI":    11,
	"XII":   12,
	"XIII":  13,
	"XIV":   14,
	"XV":    15,
	"XVI":   16,
	"XVII":  17,
	"XVIII": 18,
	"XIX":   19,
	"XX":    20,
} //map to convert Romes numeral into Arabic

var arabToRome = map[int]string{
	1:  "I",
	2:  "II",
	3:  "III",
	4:  "IV",
	5:  "V",
	6:  "VI",
	7:  "VII",
	8:  "VIII",
	9:  "IX",
	10: "X",
	11: "XI",
	12: "XII",
	13: "XIII",
	14: "XIV",
	15: "XV",
	16: "XVI",
	17: "XVII",
	18: "XVIII",
	19: "XIX",
	20: "XX",
} //Arabic to Romes numeral

func Read() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	Text := scanner.Text()
	Text = strings.TrimSpace(Text)
	return Text
}

var digitCheck = regexp.MustCompile(`^[0-9]+$`)
var r1, r2, a, b int
var err error
var rome1, rome2, arab1, arab2, rome, arab, errcond bool

func main() {
	InputStr := Read()
	SplitStr := strings.Split(InputStr, " ")
	if len(SplitStr) != 3 {
		errcond = true
		fmt.Println("Error")
		os.Exit(0)
	}
	arab1 = digitCheck.MatchString(SplitStr[0])
	arab2 = digitCheck.MatchString(SplitStr[2])
	if arab1 == true && arab2 == true {
		arab = true
	}

	if arab == false {
		r1, rome1 = romeToArab[SplitStr[0]]
		r2, rome2 = romeToArab[SplitStr[2]]
	}
	if rome1 == true && rome2 == true {
		rome = true
	} else if arab1 != true || arab2 != true {
		errcond = true
	}
	if errcond == true {
		fmt.Println("Error")
		os.Exit(0)
	}
	if rome == true {
		a = romeToArab[SplitStr[0]]
		b = romeToArab[SplitStr[2]]
	} else if arab == true {
		a, err = strconv.Atoi(SplitStr[0])
		if err != nil {
			log.Fatal(err)
		}
		b, err = strconv.Atoi(SplitStr[2])
		if err != nil {
			log.Fatal(err)
		}
	}
	var c int
	if a > 10 || b > 10 {
		fmt.Println("Error")
		os.Exit(0)
	}
	switch SplitStr[1] {
	case "+":
		c = a + b
	case "-":
		c = a - b
	case "/":
		c = a / b
	case "*":
		c = a * b
	default:
		fmt.Println("Error")
		os.Exit(0)
	}
	if rome == true && c <= 0 {
		fmt.Println("Error")
		os.Exit(0)
	}
	if rome == true {
		fmt.Println(arabToRome[c])
	} else {
		fmt.Println(c)
	}
}
