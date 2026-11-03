package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CasingStyle string

const (
	CaseLower CasingStyle = "lower"
	CaseUpper CasingStyle = "upper"
	CaseTitle CasingStyle = "title"
)

type ConvertOptions struct {
	Currency  bool
	CaseStyle CasingStyle
}

var units = []string{"", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan", "sepuluh", "sebelas"}

func FromInt(n int64) string {
	if n == 0 {
		return "nol"
	}
	if n < 0 {
		return "minus " + FromInt(-n)
	}

	var result string
	switch {
	case n < 12:
		result = units[n]
	case n < 20:
		result = FromInt(n-10) + " belas"
	case n < 100:
		result = FromInt(n/10) + " puluh"
		if n%10 != 0 {
			result += " " + FromInt(n%10)
		}
	case n < 200:
		result = "seratus"
		if n-100 != 0 {
			result += " " + FromInt(n-100)
		}
	case n < 1000:
		result = FromInt(n/100) + " ratus"
		if n%100 != 0 {
			result += " " + FromInt(n%100)
		}
	case n < 2000:
		result = "seribu"
		if n-1000 != 0 {
			result += " " + FromInt(n-1000)
		}
	case n < 1000000:
		result = FromInt(n/1000) + " ribu"
		if n%1000 != 0 {
			result += " " + FromInt(n%1000)
		}
	case n < 1000000000:
		result = FromInt(n/1000000) + " juta"
		if n%1000000 != 0 {
			result += " " + FromInt(n%1000000)
		}
	case n < 1000000000000:
		result = FromInt(n/1000000000) + " miliar"
		if n%1000000000 != 0 {
			result += " " + FromInt(n%1000000000)
		}
	default:
		result = FromInt(n/1000000000000) + " triliun"
		if n%1000000000000 != 0 {
			result += " " + FromInt(n%1000000000000)
		}
	}

	return strings.TrimSpace(result)
}

func FromFloat(f float64) string {
	if math.IsNaN(f) || math.IsInf(f, 0) {
		return ""
	}
	if f < 0 {
		return "minus " + FromFloat(-f)
	}

	intPart := int64(math.Floor(f))
	fracPart := f - float64(intPart)

	res := FromInt(intPart)

	if fracPart > 0.000001 {
		res += " koma"
		str := fmt.Sprintf("%.6f", fracPart)
		str = strings.TrimPrefix(str, "0.")
		str = strings.TrimRight(str, "0")
		for _, ch := range str {
			digit, _ := strconv.Atoi(string(ch))
			if digit == 0 {
				res += " nol"
			} else {
				res += " " + units[digit]
			}
		}
	}
	return res
}

func ToRupiah(amount float64) string {
	if amount < 0 {
		return "minus " + ToRupiah(-amount)
	}
	intPart := int64(math.Floor(amount))
	senPart := int64(math.Round((amount - float64(intPart)) * 100))

	res := FromInt(intPart) + " rupiah"
	if senPart > 0 {
		res += " " + FromInt(senPart) + " sen"
	}
	return res
}

func Format(s string, caseStyle CasingStyle) string {
	switch caseStyle {
	case CaseUpper:
		return strings.ToUpper(s)
	case CaseTitle:
		words := strings.Fields(s)
		for i, w := range words {
			if len(w) > 0 {
				words[i] = strings.ToUpper(w[:1]) + strings.ToLower(w[1:])
			}
		}
		return strings.Join(words, " ")
	case CaseLower:
		fallthrough
	default:
		return strings.ToLower(s)
	}
}

func normalizeNumber(s string) (string, bool) {
	s = strings.TrimPrefix(s, "Rp")
	s = strings.TrimPrefix(s, "rp")
	s = strings.TrimPrefix(s, "RP")
	s = strings.TrimPrefix(s, "Rp.")
	s = strings.TrimPrefix(s, "rp.")
	s = strings.TrimPrefix(s, "RP.")
	s = strings.TrimSpace(s)

	if strings.Contains(s, ",") {
		s = strings.ReplaceAll(s, ".", "")
		s = strings.ReplaceAll(s, ",", ".")
		return s, true
	}
	if strings.Contains(s, ".") {
		parts := strings.Split(s, ".")
		if len(parts) > 2 {
			s = strings.ReplaceAll(s, ".", "")
			return s, false
		}
		if len(parts) == 2 && len(parts[1]) == 3 {
			s = strings.ReplaceAll(s, ".", "")
			return s, false
		}
		return s, true
	}
	return s, false
}

func processInput(rawInput string, opts ConvertOptions) string {
	rawInput = strings.TrimSpace(rawInput)
	if rawInput == "" {
		return ""
	}

	var converted string
	cleaned, isFloat := normalizeNumber(rawInput)

	if opts.Currency {
		val, err := strconv.ParseFloat(cleaned, 64)
		if err != nil {
			return ""
		}
		converted = ToRupiah(val)
	} else {
		if !isFloat {
			val, err := strconv.ParseInt(cleaned, 10, 64)
			if err == nil {
				converted = FromInt(val)
			} else {
				fval, ferr := strconv.ParseFloat(cleaned, 64)
				if ferr != nil {
					return ""
				}
				converted = FromFloat(fval)
			}
		} else {
			fval, err := strconv.ParseFloat(cleaned, 64)
			if err != nil {
				return ""
			}
			converted = FromFloat(fval)
		}
	}

	return Format(converted, opts.CaseStyle)
}

func main() {
	var isCurrency bool
	var casing string

	flag.BoolVar(&isCurrency, "rupiah", false, "format output as rupiah currency")
	flag.BoolVar(&isCurrency, "r", false, "format output as rupiah currency (shorthand)")
	flag.StringVar(&casing, "case", "lower", "output casing: lower, upper, title")
	flag.StringVar(&casing, "c", "lower", "output casing: lower, upper, title (shorthand)")
	flag.Parse()

	opts := ConvertOptions{
		Currency:  isCurrency,
		CaseStyle: CasingStyle(casing),
	}

	// Workaround for flag package limitation: flag package does not automatically process piped stdin when no positional args are provided
	args := flag.Args()
	if len(args) > 0 {
		for _, arg := range args {
			out := processInput(arg, opts)
			if out != "" {
				fmt.Println(out)
			}
		}
		return
	}

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			out := processInput(line, opts)
			if out != "" {
				fmt.Println(out)
			}
		}
	} else {
		flag.Usage()
	}
}
