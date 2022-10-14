package allergies

const eggs = 0b00000001
const peanuts = 0b00000010
const shellfish = 0b00000100
const strawberries = 0b00001000
const tomatoes = 0b00010000
const chocolate = 0b00100000
const pollen = 0b01000000
const cats = 0b10000000

var score = map[uint]string{
	eggs: "eggs", peanuts: "peanuts",
	shellfish: "shellfish", strawberries: "strawberries",
	tomatoes: "tomatoes", chocolate: "chocolate",
	pollen: "pollen", cats: "cats"}

func Allergies(allergies uint) []string {
	var result []string
	for _, i := range []uint{eggs, peanuts, shellfish, strawberries, tomatoes, chocolate, pollen, cats} {
		if allergies&i == i {
			result = append(result, score[i])
		}
	}
	return result
}

func AllergicTo(allergies uint, allergen string) bool {
	for _, b := range Allergies(allergies) {
		if b == allergen {
			return true
		}
	}
	return false
}
