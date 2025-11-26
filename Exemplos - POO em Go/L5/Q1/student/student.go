package student

import "fmt"

type Student struct {
	Name    string
	RNumber uint8
	Adress  Adress
	PNotes  [4]float32
}

type Adress struct {
	CEP    string
	Number uint8
}

func New(name string, rNumber uint8, adress Adress, pNotes [4]float32) *Student {
	return &Student{
		name,
		rNumber,
		adress,
		pNotes,
	}
}

func (s *Student) Average() float32 {
	var (
		sum float32
	)

	for _, note := range s.PNotes {
		sum += note
	}

	return sum / 4
}

func (s *Student) ToJSON() string {
	return fmt.Sprintf(`{
  "name": "%s",
  "rNumber": %d,
  "adress": {
    "cep": "%s",
    "number": %d,
  },
  "pNotes": [
    %.1f,
    %.1f,
    %.1f,
    %.1f
  ]
}`, s.Name, s.RNumber, s.Adress.CEP, s.Adress.Number, s.PNotes[0], s.PNotes[1], s.PNotes[2], s.PNotes[3])
}

func (a *Adress) ToJSON() string {
	return fmt.Sprintf(`
"adress": {
  "cep": "%s",
  "number": %d
}`, a.CEP, a.Number)
}
