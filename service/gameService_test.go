package service

import (
	"github.com/jboulet/fizzbuzz-go/dto"
	"reflect"
	"testing"
)

const (
	INT1  = 3
	INT2  = 5
	LIMIT = 15
	FIZZ  = "fizz"
	BUZZ  = "buzz"
)

func TestFizzBuzz(t *testing.T) {
	gameParameter := dto.GameParameter{Int1: INT1, Int2: INT2, Limit: LIMIT, Str1: FIZZ, Str2: BUZZ}

	result := make([]string, 0, gameParameter.Limit)
	expected := []string{"1", "2", FIZZ, "4", BUZZ, FIZZ, "7", "8", FIZZ, BUZZ, "11", FIZZ, "13", "14", FIZZ + BUZZ}

	for value := range FizzBuzz(&gameParameter) {
		result = append(result, value)
	}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("error %v", result)
	}
}

func BenchmarkXxx(b *testing.B) {

	limit := b.N

	gameParameter := dto.GameParameter{Int1: INT1, Int2: INT2, Limit: limit, Str1: FIZZ, Str2: BUZZ}
	for value := range FizzBuzz(&gameParameter) {
		b.Log(value)
	}
}
