# Gerador de Senhas em Golang

## 📌 Sobre o projeto
Este é um simples gerador de senhas aleatórias escrito em **Golang**. O programa permite ao usuário especificar o tamanho da senha e escolher quais tipos de caracteres incluir na senha.

## 🚀 Como funciona
1. O usuário informa o tamanho da senha desejada.
2. O programa solicita quais tipos de caracteres incluir:
   - ✅ Letras maiúsculas
   - ✅ Letras minúsculas
   - ✅ Números
   - ✅ Caracteres especiais
3. A senha gerada é exibida no terminal.

## 🛠 Tecnologias utilizadas
- **Linguagem:** Go (Golang)
- **Pacotes padrão:** fmt, math/rand, time, strings

## ▶ Como executar o programa
1. Certifique-se de ter o **Go** instalado em sua máquina.
2. Clone o repositório:
   ```sh
   git clone https://github.com/jheanDevs/gerador-senhas-golang.git
   ```
3. Acesse a pasta do projeto:
   ```sh
   cd gerador-senhas-golang
   ```
4. Compile e execute o programa:
   ```sh
   go run main.go
   ```
5. Insira o tamanho da senha desejada e selecione os tipos de caracteres (S/N).
6. A senha gerada será exibida no terminal.

## 📝 Exemplo de uso
```
Digite o tamanho da senha: 10
Incluir letras maiúsculas? (S/N): S
Incluir letras minúsculas? (S/N): S
Incluir números? (S/N): N
Incluir caracteres especiais? (S/N): S

============================
Senha gerada: Kx!PaW@yVq
============================
```

## 🔧 Melhorias futuras
- Melhorar a aleatoriedade usando **crypto/rand** em vez de **math/rand**.
- Criar uma interface gráfica para o gerador.
- Implementar a possibilidade de salvar senhas geradas em um arquivo.

## 📚 Referências
- [Pacote rand](https://pkg.go.dev/math/rand)
- [Gerador de senhas seguro](https://pkg.go.dev/crypto/rand)

## 📜 Licença
Este projeto é de código aberto e está licenciado sob a licença **MIT**.

