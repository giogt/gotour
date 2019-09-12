package interfaces

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_vertexShow_whenVertexIsValued_mustReturnItsStringRepresentation(t *testing.T) {
	var s Show = &Vertex{1, 2}
	res := s.Show()
	assert.Equal(t, "[1, 2]", res)
}

func Test_vertexShow_whenVertexIsNull_mustReturnTheStringNull(t *testing.T) {
	var s Show
	var v *Vertex
	s = v

	res := s.Show()
	assert.Equal(t, "<nil>", res)
}

func Test_personShow_whenPersonIsValued_mustReturnItsStringRepresentation(t *testing.T) {
	var s Show = &Person{Name: "John", Surname: "Doe", Age: 35}
	res := s.Show()
	assert.Equal(t, "[name: John, surname: Doe, age: 35]", res)
}

func Test_personShow_whenPersonIsNull_mustReturnTheStringNull(t *testing.T) {
	var s Show
	var v *Person
	s = v

	res := s.Show()
	assert.Equal(t, "<nil>", res)
}

func Test_describe_forDifferentParameterTypes_mustReturnAStringDescribingValueAndTypePlusTypeDependentSuffix(t *testing.T) {
	assert.Equal(t, "(1, uint) [unsigned integer]", describe(uint(1)))
	assert.Equal(t, "({1 2}, interfaces.Vertex) [2D vertex]", describe(Vertex{1, 2}))
	assert.Equal(t, "({John Doe 35}, interfaces.Person) [a person]", describe(Person{Name: "John", Surname: "Doe", Age: 35}))
}

func Test_showToVertex_whenSuccessful_mustReturnVertexValueAndTrue(t *testing.T) {
	var s Show
	v := Vertex{1, 2}
	s = &v

	res, ok := ShowToVertex(s)
	assert.Equal(t, Vertex{1, 2}, *res)
	assert.Equal(t, true, ok)
}

func Test_showToVertex_whenNotSuccessful_mustReturnVertexZeroValueAndFalse(t *testing.T) {
	var s Show
	p := Person{Name: "John", Surname: "Doe", Age: 35}
	s = &p

	res, ok := ShowToVertex(s)
	assert.Nil(t, res)
	assert.Equal(t, false, ok)
}

func Test_showToVertexUnsafe_whenSuccessful_mustReturnVertexValue(t *testing.T) {
	var s Show
	v := Vertex{1, 2}
	s = &v

	res := ShowToVertexUnsafe(s)
	assert.Equal(t, Vertex{1, 2}, *res)
}

func Test_showToVertexUnsafe_whenUnsuccessful_mustPanic(t *testing.T) {
	var s Show
	p := Person{Name: "John", Surname: "Doe", Age: 35}
	s = &p

	assert.Panics(t, func() { ShowToVertexUnsafe(s) })
}
