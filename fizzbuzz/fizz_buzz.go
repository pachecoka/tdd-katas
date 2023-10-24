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
	if number%3 == 0 {
		return FIZZ
	}
	if number%5 == 0 {
		return Buzz
	}
	if number%7 == 0 {
		return Whizz
	}
	if number%11 == 0 {
		return Bang
	}
	return strconv.Itoa(number)
}
