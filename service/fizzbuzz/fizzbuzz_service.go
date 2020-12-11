package fizzbuzz

import (
	"fiz-buzz-server/model"
	"strconv"
)

type FizzBuzzService interface {
	ComputeFizzBuzz(param model.FizzbuzzParam) model.FizzBuzzRresponse
}

type fizzBuzzService struct{}

func NewFizzBuzzService() *fizzBuzzService {
	return new(fizzBuzzService)
}

func (f fizzBuzzService) ComputeFizzBuzz(param model.FizzbuzzParam) model.FizzBuzzRresponse {
	sentenceSeq := []string{}
	for i := 1; i <= param.Limit; i++ {
		isMultipleOf1, isMultipleOf2 := isMultipleOf(i, param.FstMultiple, param.SecMultiple)
		if isMultipleOf1 && isMultipleOf2 {
			sentenceSeq = append(sentenceSeq, param.Label1+param.Label2)
		} else if isMultipleOf1 {
			sentenceSeq = append(sentenceSeq, param.Label1)
		} else if isMultipleOf2 {
			sentenceSeq = append(sentenceSeq, param.Label2)
		} else {
			sentenceSeq = append(sentenceSeq, strconv.Itoa(i))
		}

	}

	return model.FizzBuzzRresponse{Sequences: sentenceSeq}
}

func isMultipleOf(numberToTest int, multiple1 int, multiple2 int) (bool, bool) {

	isMultipleOf1 := numberToTest%multiple1 == 0
	isMultipleOf2 := numberToTest%multiple2 == 0

	return isMultipleOf1, isMultipleOf2
}
