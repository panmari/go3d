package mat4

import (
	"testing"
	"github.com/ungerik/go3d/mat3"
	"github.com/ungerik/go3d/vec3"
	"github.com/ungerik/go3d/vec4"
)

const EPSILON = 0.0001

func TestDeterminant(t *testing.T) {
	detId := Ident.Determinant()
	if detId != 1 {
		t.Errorf("Wrong determinant for id: %f", detId)
	}
	
	detTwo := Ident
	detTwo[0][0] = 2
	if det := detTwo.Determinant(); det != 2 {
		t.Errorf("Wrong determinant: %f", det)
	}
	
	scale2 := Ident.Scale(2)
	if det := scale2.Determinant(); det != 2*2*2*1 {
		t.Errorf("Wrong determinant: %f", det)
	}
	
	row1changed, _ := Parse("3   0   0   0   2   2   0   0   1   0   2   0   2   0   0   1")
	if det := row1changed.Determinant(); det != 12 {
		t.Errorf("Wrong determinant: %f", det)
	}
	
	row12changed, _ := Parse("3     1     0     0     2     5     0     0     1     6     2     0     2   100     0     1")
	if det := row12changed.Determinant(); det != 26 {
		t.Errorf("Wrong determinant: %f", det)
	}
	
	row123changed,_ := Parse("3.00000     1.00000     0.50000     0.00000     2.00000     5.00000     2.00000     0.00000     1.00000     6.00000     7.00000     0.00000     2.00000   100.00000	 1.00000     1.00000")
	if det := row123changed.Determinant3x3(); det != 60.500 {
		t.Errorf("Wrong determinant for 3x3 matrix: %f", det)
	}
	if det := row123changed.Determinant(); det != 60.500 {
		t.Errorf("Wrong determinant: %f", det)
	}
	randomMatrix, err := Parse("0.43685   0.81673   0.63721   0.23421   0.16600   0.40608   0.53479   0.43210   0.37328   0.36436   0.56356   0.66830   0.32475   0.14294   0.42137   0.98046")
	randomMatrix.Transpose() //transpose for easy comparability with octave output
	if err != nil {
		t.Errorf("Could not parse random matrix: %v", err)
	}
	if det := randomMatrix.Determinant3x3(); det - 0.043437 > EPSILON {
		t.Errorf("Wrong determinant for random sub 3x3 matrix: %f", det)
	}
	
	if det := randomMatrix.Determinant(); det - 0.012208 > EPSILON {
		t.Errorf("Wrong determinant for random matrix: %f", det)
	}
}

func TestInvert(t *testing.T) {
	row123changed,_ := Parse("3.00000     1.00000     0.50000     0.00000     2.00000     5.00000     2.00000     0.00000     1.00000     6.00000     7.00000     0.00000     2.00000   100.00000	 1.00000     1.00000")
	blocked_expected := mat3.T{vec3.T{5, 2, 0}, vec3.T{6, 7, 0}, vec3.T{100, 1, 1}}
	if blocked := row123changed.maskedBlock(0,0); *blocked != blocked_expected {
		t.Errorf("Did not block 0,0 correctly: %#v", blocked)
	}
	
	adj_expected := T{vec4.T{23, -4, -0.5, -0}, vec4.T{-12, 20.5, -5, 0}, vec4.T{7, -17, 13, -0}, vec4.T{1147, -2025, 488, 60.5}}	
	adj := row123changed
	adj.Adjugate()
	if adj != adj_expected {
		t.Errorf("Adjugate not computed correctly: %#v", adj)
	}
	
	inv_expected :=  T{vec4.T{0.38016528, -0.0661157, -0.008264462, -0}, vec4.T{-0.19834709, 0.33884296, -0.08264463, 0}, vec4.T{0.11570247, -0.28099173, 0.21487603, -0}, vec4.T{18.958677, -33.471073, 8.066115, 0.99999994}}
	inv := row123changed
	inv.Invert()
	if inv != inv_expected {
		t.Errorf("Inverse not computed correctly: %#v", inv)
	}
}

