//No nosso conteúdo, temos um exemplo de validação de CPF. Porém, esse se encontra incompleto. Complete o exemplo adicionando o
//algoritimo de verificação dos digitos verificadores e faça com que o cpf formatado (e.g 111.111.110-30) seja aceito também por meio de um construtor.

package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type CPF string

func main() {
	cpf := CPF("11111111030")
	if cpf.EhValido() {
		fmt.Println("CPF válido")
	} else {
		fmt.Println("CPF inválido")
	}
}

var (
	cpfFirstDigitTable  = []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	cpfSecondDigitTable = []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
)

const (
	// CPFFormatPattern is the pattern for string replacement
	// with Regex
	CPFFormatPattern string = `([\d]{3})([\d]{3})([\d]{3})([\d]{2})`
)

// NewCPF is a helper function to convert and clean a new CPF
// from a string
func NewCPF(s string) CPF {
	return CPF(Clean(s))
}

// IsValid returns if CPF is a valid CPF document
func (c *CPF) EhValido() bool {
	return ValidateCPF(string(*c))
}

// String returns a formatted CPF document
// 000.000.000-00
func (c *CPF) String() string {

	str := string(*c)

	if !c.EhValido() {
		return str
	}

	expr, err := regexp.Compile(CPFFormatPattern)

	if err != nil {
		return str
	}

	return expr.ReplaceAllString(str, "$1.$2.$3-$4")
}

// ValidateCPF validates a CPF document
// You should use without punctuation
func ValidateCPF(cpf string) bool {
	if len(cpf) != 11 {
		return false
	}

	firstPart := cpf[0:9]
	sum := sumDigit(firstPart, cpfFirstDigitTable)

	r1 := sum % 11
	d1 := 0

	if r1 >= 2 {
		d1 = 11 - r1
	}

	secondPart := firstPart + strconv.Itoa(d1)

	dsum := sumDigit(secondPart, cpfSecondDigitTable)

	r2 := dsum % 11
	d2 := 0

	if r2 >= 2 {
		d2 = 11 - r2
	}
	finalPart := fmt.Sprintf("%s%d%d", firstPart, d1, d2)
	return finalPart == cpf
}

func sumDigit(s string, table []int) int {

	if len(s) != len(table) {
		return 0
	}

	sum := 0

	for i, v := range table {
		c := string(s[i])
		d, err := strconv.Atoi(c)
		if err == nil {
			sum += v * d
		}
	}

	return sum
}

// Clean can be used for cleaning formatted documents
func Clean(s string) string {
	s = strings.Replace(s, ".", "", -1)
	s = strings.Replace(s, "-", "", -1)
	s = strings.Replace(s, "/", "", -1)
	return s
}
