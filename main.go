package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)


func generatePassword(length int, useUpper, useLower, useNumbers, useSpecial bool) string {
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	numbers := "0123456789"
	specialChars := "!@#$%^&*()+?><:{}[]"
	
	//JUNTAR CARACTERES
	allChars := ""
		if useUpper {
			allChars += upperCase
		} 
		if useLower {
			allChars += lowerCase
		}
		if useNumbers {
			allChars += numbers
		}
		if useSpecial {
			allChars += specialChars
		}

		if allChars == "" {
			return "Erro: Nenhum tipo de caractere selecionado"
		}

		//INICIAR O GERADOR DE NÚMEROS ALEATÓRIOS
		rand.Seed(time.Now().UnixNano())

		//CRIAR A SENHA ALEATORIA
		password := make([]byte, length)
		for i := range password {
			password[i] = allChars[rand.Intn(len(allChars))]
		}

		return string(password)
}

func getUserChoice(prompt string) bool {
	var choice  string
	for {
		fmt.Print(prompt, " (s/n): ")
		fmt.Scanln(&choice)
		choice = strings.ToLower(strings.TrimSpace(choice))

		if choice == "s" {
			return true
		} else if choice == "n" {
			return false
		} else {
		fmt.Println("Resposta inválida, digite 's' para sim ou 'n' para não")
		}
	}
}

func main() {

	var lenSenha int

	//GERAR SENHA PELO INPUT DO USUARIO
	fmt.Print("Digite o tamanho da senha: ")
	fmt.Scanln(&lenSenha)
	fmt.Println()

	useUpper := getUserChoice("Incluir letras maiúsculas?")
	useLower := getUserChoice("Incluir letras minúsculas?")
	useNumbers := getUserChoice("Incluir números?")
	useSpecial := getUserChoice("Incluir caracteres especiais?")
	fmt.Println()

	//GERANDO A SENHA COM BASE NAS PREFERENCIAS DO USUARIO
	password := generatePassword(lenSenha, useUpper, useLower, useNumbers, useSpecial)

	//EXIBIR SENHA GERADA
	fmt.Println("============================")
	fmt.Println("Senha Gerada: ", password)
	fmt.Println("============================")
}