package main

import (
	"encoding/base64"
	"regexp"
	"os"
	"log"
)

const (
	Base64       string = "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$"
)

func isBase64(str string) bool {
	return regexp.MustCompile(Base64).MatchString(str)
}

func encode(str string) string {
	if isBase64(str) {
		return str
	}

	return base64.StdEncoding.EncodeToString([]byte(str))
}

func decode(str string) string {
	if !isBase64(str) {
		return str
	}

	data, err := base64.StdEncoding.DecodeString(str)

	if err != nil {
		log.Fatalf("Exception occurred during decoding of %s", str)
		os.Exit(1)
	}

	return string(data[:])
}
