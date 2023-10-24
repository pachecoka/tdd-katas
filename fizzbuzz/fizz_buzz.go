package fizzbuzz

import "strconv"

type FizzbuzzUsecaseI interface {
	ExecuteFizzbuzzUsecase(number int) string
}

type FizzbuzzUsecase struct {
}

func (f *FizzbuzzUsecase) ExecuteFizzbuzzUsecase(number int) string {
	return strconv.Itoa(number)
}
