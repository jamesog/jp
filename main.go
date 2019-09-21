package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type JWT struct {
	header, payload []byte
}

func indent(b []byte) string {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	if err != nil {
		return ""
	}
	return out.String()
}

func (jwt JWT) Header() string {
	return indent(jwt.header)
}

func (jwt JWT) Payload() string {
	return indent(jwt.payload)
}

func parseJWT(jwt string) (JWT, error) {
	parts := strings.SplitN(jwt, ".", 3)
	if len(parts) != 3 {
		return JWT{}, errors.New("invalid JWT")
	}
	header, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return JWT{}, err
	}
	payload, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return JWT{}, err
	}
	return JWT{
		header:  header,
		payload: payload,
	}, nil
}

func tokenFromStdin() string {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("couldn't read stdin: %v", err)
	}
	return strings.TrimSpace(string(b))
}

func main() {
	var token string
	// Read the token from either the first argument or stdin
	switch {
	case len(os.Args) >= 2:
		token = os.Args[1]
	default:
		token = tokenFromStdin()
	}

	jwt, err := parseJWT(token)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(">> Header")
	fmt.Println(jwt.Header())
	fmt.Println(">> Payload")
	fmt.Println(jwt.Payload())
}
