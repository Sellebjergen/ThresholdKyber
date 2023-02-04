package TKyber

import (
	"reflect"
	"testing"
)

// ================= Add tests =================
func TestAdd(t *testing.T) {
	lhs := &Polynomial{Coeffs: []int{1, 2}}
	rhs := &Polynomial{Coeffs: []int{0}}

	res := add(lhs, rhs)

	if !reflect.DeepEqual(res.Coeffs, []int{1, 2}) {
		t.Errorf("Add failed!")
	}
}

func TestAdd2(t *testing.T) {
	lhs := &Polynomial{Coeffs: []int{17, 42}}
	rhs := &Polynomial{Coeffs: []int{73, 100}}

	res := add(lhs, rhs)

	if !reflect.DeepEqual(res.Coeffs, []int{90, 142}) {
		t.Errorf("Add failed!")
	}
}

func TestAdd3(t *testing.T) {
	lhs := &Polynomial{Coeffs: []int{0, 2, 3, 6, 3, 2}}
	rhs := &Polynomial{Coeffs: []int{2, 4, 0, 5}}

	res := add(lhs, rhs)

	if !reflect.DeepEqual(res.Coeffs, []int{2, 6, 3, 11, 3, 2}) {
		t.Errorf("Add failed!")
	}
}

// ================= Mult tests =================
func TestMult1(t *testing.T) {
	lhs := &Polynomial{Coeffs: []int{3, 5}}
	rhs := &Polynomial{Coeffs: []int{2, 7, 2}}

	res := mult(lhs, rhs)

	if !reflect.DeepEqual(res.Coeffs, []int{6, 31, 41, 10}) {
		t.Errorf("mult failed!")
	}
}

func TestMult2(t *testing.T) {
	lhs := &Polynomial{Coeffs: []int{3, 2, 0, 5}}
	rhs := &Polynomial{Coeffs: []int{0, 1, 8, 2}}

	res := mult(lhs, rhs)

	if !reflect.DeepEqual(res.Coeffs, []int{0, 3, 26, 22, 9, 40, 10}) {
		t.Errorf("mult failed!")
	}
}

func TestMult3(t *testing.T) {
	lhs := &Polynomial{Coeffs: []int{1}}
	rhs := &Polynomial{Coeffs: []int{1, 6, 2, 1}}

	res := mult(lhs, rhs)

	if !reflect.DeepEqual(res.Coeffs, []int{1, 6, 2, 1}) {
		t.Errorf("mult failed!")
	}
}

func TestMult4(t *testing.T) {
	lhs := &Polynomial{Coeffs: []int{17, 3, 1, 0}}
	rhs := &Polynomial{Coeffs: []int{0}}

	res := mult(lhs, rhs)

	if !reflect.DeepEqual(res.Coeffs, []int{0}) {
		t.Errorf("mult failed!")
	}
}

// ================= Mult w/ constant tests =================
func TestMultConst(t *testing.T) {
	lhs := &Polynomial{Coeffs: []int{42, 10, 30}}
	rhs := int(7)

	res := mult_const(lhs, rhs)

	if !reflect.DeepEqual(res.Coeffs, []int{42 * rhs, 10 * rhs, 30 * rhs}) {
		t.Errorf("Add failed!")
	}
}

// ================= Convert kyber poly tests =================
func TestConvertPoly(t *testing.T) {
	initial := &Polynomial{Coeffs: []int{42, 10, 30}}

	kPoly := initial.toKyberPoly()
	result := fromKyberPoly(kPoly)

	if !reflect.DeepEqual(initial, result) {
		t.Errorf("Polynomials not identical!")
	}
}
