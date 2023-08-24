package main

import (
	"crypto/sha256"
	"encoding/base64"
	"log"
	"math"
	"strings"
	"unicode"
)

func Encrypt(token1, token2 string, simbols bool, length int) (string, error) {
  var result string
  combined := token1 + token2
  if len(combined) == 0 {
    result, err := Encrypt("Allen", "Turing" , simbols, length)
    if err != nil {
      return "", err
    }
    return result, nil
  }


  hash := sha256.Sum256([]byte(combined))
  if length > len(hash) {
    length = len(hash)
  }

  if simbols {
    result = ensureContainsCapital(base64.URLEncoding.EncodeToString(hash[:length]))
  } else {
    result = ensureContainsCapital(base64.StdEncoding.EncodeToString(hash[:length]))
  }
  return result, nil
}

func ensureContainsCapital(str string) string {
  var result string
  //use the square root of the length of str to determine which character to capitalize
  //fancy type casting magic to get the square root of the length of str as an int
  start := int(math.Sqrt(float64(len(str))))
  
  for i := start - 1; i < len(str); i += start {
    letter := str[i]
    if unicode.IsLetter(rune(letter)) {
      result = str[:i] + strings.ToUpper(string(letter)) + str[i+1:]
      break
    }
  }
  log.Println(result)
  return result
}

