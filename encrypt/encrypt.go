/*!!DO NOT UNDER ANY CIRCUMSTANCES CHANGE ANY OF THE CODE IN THIS FILE
UNLESS YOU ARE ABSOLUTELY SURE OF WHAT YOU ARE DOING. CHANGING THIS CODE
CAN BREAK THE THE CONSISTANCY OF THE ENCRYPTION AND WILL
RENDER THE PROGRAM USELESS.
IF YOU MUST REFACTOR OR MODIFY THIS CODE, DO SO IN A SEPERATE FILE AND
UNIT-TEST IT VERY THOROUGHLY BEFORE USING IT IN PRODUCTION.
YOUR MODIFIED CODE SHOULD RETURN THE EXACT SAME RESULTS AS THE ORIGINAL.
IF THIS ALGORITHM IS CHANGED TURING PASS WILL NO LONGER BE ABLE TO
RETURN CONSISTANT RESULTS. THIS WILL RENDER THE PROGRAM USELESS.
!!*/

package encrypt

import (
	"crypto/sha256"
	"encoding/base64"
	"math"
	"strings"
)

func Encrypt(token1, token2 string, symbols bool, length int) (string, error) {
	var result string
	combined := token1 + token2
	if len(combined) == 0 {
		result, err := Encrypt("Allen", "Turing", symbols, length)
		if err != nil {
			return "", err
		}
		return result, nil
	}

	hash := sha256.Sum256([]byte(combined))
	if length > len(hash) {
		length = len(hash)
	}

	if symbols {
		result = base64.URLEncoding.EncodeToString(hash[:])
	} else {
		result = base64.StdEncoding.EncodeToString(hash[:])
	}
	return format(result, symbols, length), nil
}

func format(str string, symbols bool, length int) string {
	result := strings.TrimRight(str, "=")
	if !symbols {
		result = strings.ReplaceAll(result, "+", "")
		result = strings.ReplaceAll(result, "/", "")
		result = strings.ReplaceAll(result, "=", "")
		result = strings.ReplaceAll(result, "*", "")
		result = strings.ReplaceAll(result, "#", "")
		result = strings.ReplaceAll(result, "@", "")
		result = strings.ReplaceAll(result, "$", "")
		result = strings.ReplaceAll(result, "!", "")
		result = strings.ReplaceAll(result, "-", "")
		result = strings.ReplaceAll(result, "_", "")
	}
	if symbols && !hasSymbol(result) {
		position := int(math.Pow(float64(length), 1.0/3.0)) * 2
		result = result[:position] + "!" + result[position:]
	}
	if !hasCapital(result) {
		start := int(math.Sqrt(float64(length)))
		for i := start; i < len(result); i++ {
			if result[i] >= 'a' && result[i] <= 'z' {
				result = result[:i] + strings.ToUpper(result[i:i+1]) + result[i+1:]
				break
			}
		}
	}
	return result[:length]
}

func hasCapital(str string) bool {
	for _, char := range str {
		if char >= 'A' && char <= 'Z' {
			return true
		}
	}
	return false
}

func hasSymbol(str string) bool {
	for _, char := range str {
		if char == '+' || char == '/' || char == '=' || char == '*' || char == '#' || char == '$' || char == '@' {
			return true
		}
	}
	return false
}
