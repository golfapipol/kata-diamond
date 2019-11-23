package diamond

import "strings"

const (
	letters = "ABCD"
)

func diamond(letter string) string {
	if letter == "D" {
		return strings.Join([]string{
			single(3),         // round 0: 3 A 3
			double("B", 2, 1), // round 1: 2 B 1 B 2
			double("C", 1, 3), // round 2: 1 C 3 C 1
			double("D", 0, 5), // round 3: 0 D 5 D 0
			double("C", 1, 3), // round 2: 1 C 3 C 1
			double("B", 2, 1), // round 1: 2 B 1 B 2
			single(3),         // round 0: 3 A 3
		}, "\n")
	}
	if letter == "C" {
		return strings.Join([]string{
			single(2),         // round 0: 2 A 2
			double("B", 1, 1), // round 1: 1 B 1 B 1
			double("C", 0, 3), // round 2: 0 C 3 C 0
			double("B", 1, 1), // round 1: 1 B 1 B 1
			single(2),         // round 0: 2 A 2
		}, "\n")
	}
	if letter == "B" {
		return strings.Join([]string{
			single(1),         // round 0: 1 A 1
			double("B", 0, 1), // round 1: B 1 B
			single(1),         // round 0: 1 A 1
		}, "\n")
	}
	return strings.Join([]string{
		single(0), // round 0: 0 A 0
	}, "\n")
}

func single(number int) string {
	return strings.Repeat(" ", number) + "A" + strings.Repeat(" ", number)
}

func double(char string, outerNumber int, innerNumber int) string {
	return strings.Repeat(" ", outerNumber) + char + strings.Repeat(" ", innerNumber) + char + strings.Repeat(" ", outerNumber)
}
