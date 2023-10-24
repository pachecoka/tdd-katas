package fizzbuzz

import "strconv"

const FIZZ = "Fizz"
const Buzz = "Buzz"
const Whizz = "Whizz"
const Bang = "Bang"

type FizzbuzzUsecaseI interface {
	ExecuteFizzbuzzUsecase(number int) string
}

type FizzbuzzUsecase struct {
}

func (f *FizzbuzzUsecase) ExecuteFizzbuzzUsecase(number int) string {
	var response = ""
	if number%3 == 0 {
		response += FIZZ
	}
	if number%5 == 0 {
		response += Buzz
	}
	if number%7 == 0 {
		response += Whizz
	}
	if number%11 == 0 {
		response += Bang
	}
	if response != "" {
		return response
	}
	return strconv.Itoa(number)
}
