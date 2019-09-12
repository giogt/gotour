package methods

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVertex_Abs(t *testing.T) {
	vertex := Vertex{3.0, 4.0}
	abs := vertex.Abs()
	assert.Equal(t, 5.0, abs)
}

func TestVertex_Scale(t *testing.T) {
	vertex := Vertex{3.0, 4.0}
	vertex.Scale(10.0)
	abs := vertex.Abs()
	assert.Equal(t, 50.0, abs)
}
