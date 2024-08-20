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
	21: "XXI",
	22: "XXII",
	23:"XXIII",
	24: "XXIV",
	25: "XXV",
	26: "XXVI",
	27: "XXVII",
	28: "XXVIII",
	29: "XXIX",
	30: "XXX",
	31: "XXXI",
	32: "XXXII",
	33: "XXXIII",
	34: "XXXIV",
	35: "XXXV",
	36: "XXXVI",
	37: "XXXVII",
	38: "XXXVIII",
	39: "XXXIX",
	40: "XL",
	41: "XLI",
	42: "XLII",
	43: "XLIII",
	44: "XLIV",
	45: "XLV",
	46: "XLVI",
	47: "XLVII",
	48: "XLVIII",
	49: "XLIX",
	50: "L",
	51: "LI",
	52: "LII",
	53: "LIII",
	54: "LIV",
	55: "LV",
	56: "LVI",
	57: "LVII",
	58: "LVIII",
	59: "LIX",
	60: "LX",
	61: "LXI",
	62: "LXII",
	63: "LXIII",
	64: "LXIV",
	65: "LXV",
	66: "LXVI",
	67: "LXVII",
	68: "LXVIII",
	69: "LXIX",
	70: "LXX",
	71: "LXXI",
	72: "LXXII",
	73: "LXXIII",
	74: "LXXIV",
	75: "LXXV",
	76: "LXXVI",
	77: "LXXVII",
	78:"LXXVIII",
	79: "LXXIX",
	80: "LXXX",
	81: "LXXXI",
	82: "LXXXII",
	83: "LXXXIII",
	84: "LXXXIV",
	85: "LXXXV",
	86: "LXXXVI",
	87: "LXXXVII",
	88: "LXXXVIII",
	89: "LXXXIX",
	90: "XC",
	91: "XCI",
	92: "XCII",
	93: "XCIII",
	94: "XCIV",
	95: "XCV",
	96: "XCVI",
	97: "XCVII",
	98: "XCVIII",
	99: "XCIX",
	100: "C",
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
		panic("Формат операции не удовлетворяет условию")
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
		panic("Используются арабские и римские цифры одновременно")
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
	if a>10 || a<1 || b>10 || b<1 {
		panic("Входные значения выходят за диапазон разрешенных значений")
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
		panic("Строка не является математической операцией")
	}
	if rome == true && c <= 0 {
		panic("В римской системе счисления нет отрицательных чисел")
	}
	if rome == true {
		fmt.Println(arabToRome[c])
	} else {
		fmt.Println(c)
	}
}
