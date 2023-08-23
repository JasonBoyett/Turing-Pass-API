package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func Encrypt(token1, token2 string) (string, error) {
  combined := token1 + token2
  if len(combined) == 0 {
    result, err := Encrypt("Allen", "Turing")
    if err != nil {
      return "", err
    }
    return result, nil
  }

  hash := sha256.Sum256([]byte(combined))
  base64 := base64.StdEncoding.EncodeToString(hash[:])
  return fmt.Sprintf("%x", base64), nil
}

