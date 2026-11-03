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

