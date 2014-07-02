package mat4

import (
	"testing"
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

