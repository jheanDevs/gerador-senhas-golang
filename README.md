# gerador-senhas-golang
Gerador de Senhas em Golang

# Gerador de Senhas em Golang

Este é um simples gerador de senhas aleatórias escrito em Golang. O programa permite ao usuário especificar o tamanho da senha e gera uma senha contendo letras maiúsculas, minúsculas, números e caracteres especiais.

## Como funciona

1. O usuário informa o tamanho da senha desejada.
2. O programa gera uma senha aleatória garantindo que pelo menos um número e uma letra maiúscula estejam presentes.
3. A senha gerada é exibida no terminal.

## Tecnologias utilizadas
- Linguagem: Go (Golang)
- Pacotes padrão: `fmt`, `math/rand`

## Como executar o programa

1. Certifique-se de ter o Go instalado em sua máquina.
2. Salve o código em um arquivo chamado `main.go`.
3. No terminal, navegue até o diretório onde o arquivo está salvo e execute o seguinte comando:
   ```sh
   go run main.go
   ```
4. Insira o tamanho da senha desejada e pressione `Enter`.
5. A senha gerada será exibida no terminal.

## Exemplo de uso
```
Digite o tamanho da senha: 10

============================
Senha gerada: A9f@bC*3!
============================
```

## Melhorias futuras
- Adicionar a opção de personalizar quais tipos de caracteres incluir na senha.
- Melhorar a aleatoriedade usando `crypto/rand` em vez de `math/rand`.
- Criar uma interface gráfica para o gerador.

## Referências
- [Pacote rand](https://pkg.go.dev/math/rand)
- [Gerador de senhas seguro](https://pkg.go.dev/github.com/sethvargo/go-password/password)

