package terbilang

import (
	"math"
	"strconv"
	"strings"
)

type CurrencyFormatter struct {
	CurrencyName string
	SubunitName  string
}

func DefaultRupiahFormatter() CurrencyFormatter {
	return CurrencyFormatter{
		CurrencyName: "rupiah",
		SubunitName:  "sen",
	}
}

func FormatCurrency(amount float64, fmtConfig CurrencyFormatter) string {
	if math.IsNaN(amount) || math.IsInf(amount, 0) {
		return ""
	}
	if amount < 0 {
		return "minus " + FormatCurrency(-amount, fmtConfig)
	}

	intPart := int64(math.Floor(amount))
	senPart := int64(math.Round((amount - float64(intPart)) * 100))

	if senPart >= 100 {
		intPart++
		senPart -= 100
	}

	result := FromInt(intPart)
	if fmtConfig.CurrencyName != "" {
		result += " " + fmtConfig.CurrencyName
	}

	if senPart > 0 {
		result += " " + FromInt(senPart)
		if fmtConfig.SubunitName != "" {
			result += " " + fmtConfig.SubunitName
		}
	}

	return strings.TrimSpace(result)
}

func ParseRupiahString(input string) (float64, error) {
	cleaned := strings.TrimSpace(input)
	cleaned = strings.TrimPrefix(cleaned, "Rp")
	cleaned = strings.TrimPrefix(cleaned, "rp")
	cleaned = strings.TrimPrefix(cleaned, "RP")
	cleaned = strings.TrimSpace(cleaned)
	cleaned = strings.ReplaceAll(cleaned, ".", "")
	cleaned = strings.ReplaceAll(cleaned, ",", ".")
	return strconv.ParseFloat(cleaned, 64)
}
