package main

import (
	"crypto/sha256"
)

func Encript(token1 string, token2 string) (string, error) {
  combined := token1 + token2
  if len(combined) == 0 {
    return Encript("Allen", "Turing")
  }

  hash := sha256.Sum256([]byte(combined))
  return string(hash[:]), nil
}

