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

