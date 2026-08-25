package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func celsiusParaFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func fahrenheitParaCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func celsiusParaKelvin(c float64) float64 {
	return c + 273.15
}

func kelvinParaCelsius(k float64) float64 {
	return k - 273.15
}

func lerFloat(leitor *bufio.Reader, mensagem string) float64 {
	for {
		fmt.Print(mensagem)
		entrada, _ := leitor.ReadString('\n')
		entrada = strings.TrimSpace(entrada)
		entrada = strings.Replace(entrada, ",", ".", 1)

		valor, err := strconv.ParseFloat(entrada, 64)
		if err != nil {
			fmt.Println("Valor inválido, tente novamente.")
			continue
		}
		return valor
	}
}

func main() {
	leitor := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== Conversor de Temperatura ===")
		fmt.Println("1 - Celsius para Fahrenheit")
		fmt.Println("2 - Fahrenheit para Celsius")
		fmt.Println("3 - Celsius para Kelvin")
		fmt.Println("4 - Kelvin para Celsius")
		fmt.Println("0 - Sair")
		fmt.Print("Escolha uma opção: ")

		opcao, _ := leitor.ReadString('\n')
		opcao = strings.TrimSpace(opcao)

		switch opcao {
		case "1":
			c := lerFloat(leitor, "Digite a temperatura em Celsius: ")
			fmt.Printf("%.2f°C = %.2f°F\n", c, celsiusParaFahrenheit(c))
		case "2":
			f := lerFloat(leitor, "Digite a temperatura em Fahrenheit: ")
			fmt.Printf("%.2f°F = %.2f°C\n", f, fahrenheitParaCelsius(f))
		case "3":
			c := lerFloat(leitor, "Digite a temperatura em Celsius: ")
			fmt.Printf("%.2f°C = %.2fK\n", c, celsiusParaKelvin(c))
		case "4":
			k := lerFloat(leitor, "Digite a temperatura em Kelvin: ")
			fmt.Printf("%.2fK = %.2f°C\n", k, kelvinParaCelsius(k))
		case "0":
			fmt.Println("Encerrando o programa...")
			return
		default:
			fmt.Println("Opção inválida, tente novamente.")
		}
	}
}
