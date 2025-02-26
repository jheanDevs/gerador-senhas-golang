package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var lenSenha int

	//GERAR SENHA PELO INPUT DO USUARIO

	fmt.Print("Digite o tamanho da senha: ")
	fmt.Scanln(&lenSenha)
	fmt.Println("\n")
	password := gerenatePassword(lenSenha)

	fmt.Println("============================")
	fmt.Println("Senha gerada: ", password)
	fmt.Println("============================")
}

func gerenatePassword(length int) string {
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	numbers := "0123456789"
	specialChars := "!@#$%^&*()+?><:{}[]"

	//JUNTAR CARACTERES
	allChars := upperCase + lowerCase + numbers + specialChars

	//GERAR CARACTERES OBRIGATÓRIOS
	mandatoryChars := []byte{
		upperCase[rand.Intn(len(upperCase))],
		numbers[rand.Intn(len(numbers))],
	}

	//GERAR CARACTERES ALEATÓRIOS
	password := make([]byte, length-len(mandatoryChars))

	for i := range password {
		password[i] = allChars[rand.Intn(len(allChars))]
	}

	//JUNTAR TUDO
	password = append(password, mandatoryChars...)

	//EMBARALHAR
	rand.Shuffle(len(password), func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})

	return string(password)
}

// gerar senha aleatória
// https://pkg.go.dev/github.com/sethvargo/go-password/password