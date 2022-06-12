package tax

import "math"

type Taxpayer struct {
	Name   string
	Income float64
	Year   int
	Status string
	States []string // not using in this article. states also impose taxes.
}

type Bracket struct {
	Status    string
	Threshold []float64
}

type Taxes interface {
	fedTax() float64
}

func Calc(t Taxes) float64 {
	return t.fedTax()
}

func (tp Taxpayer) fedTax() float64 {

	rates := make(map[int][]float64)
	rates[2022] = []float64{.1, .12, .22, .24, .32, .35, .37}
	rates[2021] = []float64{.1, .12, .22, .24, .32, .35, .37}
	rates[2020] = []float64{.1, .12, .22, .24, .32, .35, .37}

	br := make(map[int][]Bracket)
	br[2022] = []Bracket{
		{"IND", []float64{0, 10275, 41775, 89075, 170050, 215950, 539900}},
		{"MFS", []float64{0, 10275, 41775, 89075, 170050, 215950, 323925}},
		{"MFJ", []float64{0, 20550, 83550, 178150, 340100, 431900, 647850}},
		{"HOH", []float64{0, 14650, 55900, 89050, 170050, 215950, 539900}},
	}
	br[2021] = []Bracket{
		{"IND", []float64{0, 9950, 40525, 86375, 164925, 209425, 523600}},
		{"MFS", []float64{0, 9950, 40525, 86375, 164925, 209425, 314150}},
		{"MFJ", []float64{0, 19900, 81050, 172750, 329850, 418850, 628300}},
		{"HOH", []float64{0, 14200, 54200, 86350, 164900, 209400, 314150}},
	}
	br[2020] = []Bracket{
		{"IND", []float64{0, 9875, 40125, 85525, 163300, 207350, 518400}},
		{"MFS", []float64{0, 9875, 40125, 85525, 163300, 207350, 311025}},
		{"MFJ", []float64{0, 19750, 80250, 171050, 326600, 414700, 622050}},
		{"HOH", []float64{0, 14100, 53700, 85500, 163300, 207350, 518400}},
	}

	brackets := make(map[int][]Bracket)
	brackets[2022] = []Bracket{
		{"IND", []float64{0, 10275, 41175, 89075, 170050, 215950, 539900}},
		{"MFS", []float64{0, 10275, 41175, 89075, 170050, 215950, 323925}},
		{"MFJ", []float64{0, 20550, 83550, 178150, 340100, 431900, 647850}},
		{"HOH", []float64{0, 14650, 55900, 89050, 170050, 215950, 539900}},
	}
	brackets[2021] = []Bracket{
		{"IND", []float64{0, 9950, 40525, 86375, 164925, 209425, 523600}},
		{"MFS", []float64{0, 9950, 40525, 86375, 164925, 209425, 314150}},
		{"MFJ", []float64{0, 19900, 81050, 172750, 329850, 418850, 628300}},
		{"HOH", []float64{0, 14200, 54200, 86350, 164900, 209400, 314150}},
	}
	brackets[2020] = []Bracket{
		{"IND", []float64{0, 9875, 40125, 85525, 163300, 207350, 518400}},
		{"MFS", []float64{0, 9875, 40125, 85525, 163300, 207350, 311025}},
		{"MFJ", []float64{0, 19750, 80250, 171050, 326600, 414700, 622050}},
		{"HOH", []float64{0, 14100, 53700, 85500, 163300, 207350, 518400}},
	}

	var myBrackets []float64

	for _, v := range br[tp.Year] {
		if tp.Status == v.Status {
			myBrackets = append(myBrackets, v.Threshold...)
		}
	}

	var tax float64

	for i := 0; i < len(myBrackets)-1; i++ {
		if tp.Income > myBrackets[i] {
			tax += (math.Min(tp.Income, myBrackets[i+1]) - myBrackets[i]) * rates[tp.Year][i]
		}
	}

	if tp.Income > myBrackets[len(myBrackets)-1] {
		tax += (tp.Income - myBrackets[len(myBrackets)-1]) * rates[tp.Year][len(rates[tp.Year])-1]
	}

	return tax

}
