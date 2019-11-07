package service

import (
	"fmt"
	"github.com/jboulet/fizzbuzz-go/dto"
)

func FizzBuzz(gameParameter dto.GameParamater) []string {
	result := make([]string, gameParameter.Limit)
	i := 0
	for out := range run(gameParameter) {
		result[i] = out
		i++
	}
	return result
}

func run(gameParameter dto.GameParamater) <-chan string {

	out := make(chan string, gameParameter.Limit)

	go func() {
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
	}()

	return out
}
