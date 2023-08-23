package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
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
    result = base64.URLEncoding.EncodeToString(hash[:length])
  } else {
    result = base64.StdEncoding.EncodeToString(hash[:length])
  }
  return fmt.Sprintf("%x", result), nil
}

