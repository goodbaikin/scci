package compute

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const asciiZero = 48

var errToLarge error = fmt.Errorf("Error: The number of significant digits seem too large")

func exitTooLarge() {
	fmt.Println(errToLarge)
	os.Exit(1)
}

func formatError(value float64, n int) (string, int) {
	var formatted string
	var exp int

	str := strconv.FormatFloat(value, 'f', -1, 64)

	var firstDigitIndex int
	for i, c := range str {
		if (c != '0') && (c != '.') {
			firstDigitIndex = i
			break
		}
	}

	var lastDigitIndex int
	if firstDigitIndex+n > len(str) {
		exitTooLarge()
	}
	if strings.Contains(str[firstDigitIndex:firstDigitIndex+n], ".") {
		lastDigitIndex = firstDigitIndex + (n - 1) + 1
	} else {
		lastDigitIndex = firstDigitIndex + (n - 1)
	}
	if lastDigitIndex >= len(str) {
		exitTooLarge()
	}

	digitIndexToRound := lastDigitIndex + 1
	if digitIndexToRound >= len(str) {
		exitTooLarge()
	}
	if str[digitIndexToRound] == '.' {
		digitIndexToRound += 1
	}

	dotIndex := strings.Index(str, ".")
	if dotIndex > digitIndexToRound {
		exp = dotIndex - digitIndexToRound
	} else {
		exp = 0
	}

	digitToRound := int(str[digitIndexToRound]) - asciiZero

	formattedValue, _ := strconv.ParseFloat(str[:digitIndexToRound], 64)
	if digitToRound >= 5 {
		if dotIndex > lastDigitIndex {
			formattedValue = formattedValue + 1
		} else {
			formattedValue = formattedValue + math.Pow10(dotIndex-lastDigitIndex)
		}
	}

	formatted = strconv.FormatFloat(formattedValue, 'f', -1, 64)
	if len(formatted) < digitIndexToRound {
		formatted = strconv.FormatFloat(formattedValue, 'f', lastDigitIndex-dotIndex, 64)
	} else {
		formatted = formatted[:(lastDigitIndex + 1)]
	}

	return formatted, exp
}

func getDigitIndexToRound(formattedErr string, exp int) int {
	dotIndex := strings.Index(formattedErr, ".")

	if dotIndex == -1 {
		if exp > 0 {
			return exp
		}
		return exp - 1
	}

	return dotIndex - len(formattedErr)
}

func formatAvg(value float64, digitIndexToRound int) string {
	var formatted string

	str := strconv.FormatFloat(value, 'f', -1, 64)

	dotIndex := strings.Index(str, ".")
	if dotIndex == -1 {
		dotIndex = len(str)
	}
	indexToRound := dotIndex + (-1)*digitIndexToRound
	if indexToRound >= len(str) {
		exitTooLarge()
	}
	digitToRound := int(str[indexToRound]) - asciiZero

	lastDigitIndex := indexToRound - 1
	if str[lastDigitIndex] == '.' {
		lastDigitIndex -= 1
	}

	formattedValue, _ := strconv.ParseFloat(str[:indexToRound], 64)
	if digitToRound >= 5 {
		if dotIndex > lastDigitIndex {
			formattedValue = formattedValue + 1
		} else {
			formattedValue = formattedValue + math.Pow10(dotIndex-lastDigitIndex)
		}
	}

	formatted = strconv.FormatFloat(formattedValue, 'f', -1, 64)
	if len(formatted) < indexToRound {
		formatted = strconv.FormatFloat(formattedValue, 'f', lastDigitIndex-dotIndex, 64)
	} else {
		formatted = formatted[:(lastDigitIndex + 1)]
	}

	return formatted
}

func Format(avg, err float64, n int) string {
	var formattedAvg, formattedErr string
	var exp int

	if n <= 0 {
		formattedAvg = strconv.FormatFloat(avg, 'f', -1, 64)
		formattedErr = strconv.FormatFloat(err, 'f', -1, 64)
		exp = 0
	} else {
		formattedErr, exp = formatError(err, n)
		digitIndexToRound := getDigitIndexToRound(formattedErr, exp)
		formattedAvg = formatAvg(avg, digitIndexToRound)
	}

	if exp == 0 {
		return fmt.Sprintf("%s±%s", formattedAvg, formattedErr)
	} else if exp == 1 {
		return fmt.Sprintf("(%s±%s)×10", formattedAvg, formattedErr)
	}

	return fmt.Sprintf("(%s±%s)×10^{%d}", formattedAvg, formattedErr, exp)
}
