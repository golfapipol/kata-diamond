package diamond

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Diamond_of_A_should_be_A(t *testing.T) {
	expectResult := "A"
	actualResult := diamond("A")
	assert.Equal(t, expectResult, actualResult)
}

func Test_Diamond_of_B_should_be_ABBA(t *testing.T) {
	expectResult := " A \nB B\n A "
	actualResult := diamond("B")
	assert.Equal(t, expectResult, actualResult)
}

func Test_Diamond_of_C_should_be_ABBCCBBA(t *testing.T) {
	expectResult := "  A  \n B B \nC   C\n B B \n  A  "
	actualResult := diamond("C")
	assert.Equal(t, expectResult, actualResult)
}

func Test_Diamond_of_D_should_be_ABBCCDDCCBBA(t *testing.T) {
	expectResult := "   A   \n  B B  \n C   C \nD     D\n C   C \n  B B  \n   A   "
	actualResult := diamond("D")
	assert.Equal(t, expectResult, actualResult)
}
