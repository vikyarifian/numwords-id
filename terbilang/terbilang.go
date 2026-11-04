package terbilang

import (
	"fmt"
	"math"
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

