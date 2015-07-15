// Bijectiv - A library for URL shortener encoding/decoding in Go
// Copyright 2015 Leonardo Eloy
// https://github.com/leonardoeloy/bijectiv
// Algorithm based on http://stackoverflow.com/questions/742013/how-to-code-a-url-shortener

package bijectiv

import (
	"bytes"
	"strings"
)

const (
	default_alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

type Bijectiv struct {
	Alphabet string
	Base int
}

func NewAlphabet(alphabet string) *Bijectiv {
	return &Bijectiv{Alphabet: alphabet, Base: len(alphabet)}
}

func New() *Bijectiv {
	return NewAlphabet(default_alphabet)
}

func (b *Bijectiv) Encode(number int) string {
	if number == 0 {
		return b.Alphabet[:0]
	}
	var buffer bytes.Buffer

	for number > 0 {
		base := number % b.Base
		buffer.WriteString(b.Alphabet[base:base+1])

		number /= b.Base
	}

	return reverse(buffer.String())
}

func (b *Bijectiv) Decode(value string) int {
	var number int

	for i := 0; i < len(value); i++ {
		number = number * b.Base + strings.Index(b.Alphabet, value[i:i+1])
	}

	return number
}

// Taken from https://github.com/golang/example/blob/master/stringutil/reverse.go
// Copyright 2014 Google Inc.
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
