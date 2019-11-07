package service

import (
	"fmt"
	"github.com/jboulet/fizzbuzz-go/dto"
)

func FizzBuzz(gameParameter *dto.GameParameter) *[]string {
	out := make(chan string, gameParameter.Limit)
	result := make([]string, 0, gameParameter.Limit)

	go run(gameParameter, out)

	for value := range out {
		result = append(result, value)
	}

	return &result
}

func run(gameParameter *dto.GameParameter, out chan string) {
	for i := 1; i <= gameParameter.Limit; i++ {
		result := ""
		if i%gameParameter.Int1 == 0 {
			result += gameParameter.Str1
		}
		if i%gameParameter.Int2 == 0 {
			result += gameParameter.Str2
		}
		if result == "" {
			result = fmt.Sprintf("%v", i)
		}
		out <- result
	}
	close(out)
}
