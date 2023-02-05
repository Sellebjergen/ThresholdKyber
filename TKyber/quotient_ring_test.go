package TKyber

import (
	"reflect"
	"testing"
)

// ================= Add tests =================
func TestAddQuotRing(t *testing.T) {
	fx := &Polynomial{Coeffs: []int{1, 0, 1}}
	quot_ring := new(quotRing)
	quot_ring.q = 32
	quot_ring.mod = fx

	lhs := &Polynomial{Coeffs: []int{3, 6, 4, 2, 1}}
	rhs := &Polynomial{Coeffs: []int{-17, 38, -12, 1}}

	res := quot_ring.add(lhs, rhs)
	res = quot_ring.reduce_coefficients(res)

	if !reflect.DeepEqual(res.Coeffs, []int{27, 9}) {
		t.Errorf("Add failed!")
	}
}

// ================= Mult tests =================
func TestMultQuotRing(t *testing.T) {
	fx := &Polynomial{Coeffs: []int{1, 0, 1}}
	quot_ring := new(quotRing)
	quot_ring.q = 32
	quot_ring.mod = fx

	lhs := &Polynomial{Coeffs: []int{3, 5, 0, 8}}
	rhs := &Polynomial{Coeffs: []int{1, 1, 5}}

	res := quot_ring.mult(lhs, rhs)
	res = quot_ring.reduce_coefficients(res)

	if !reflect.DeepEqual(res.Coeffs, []int{23, 15}) {
		t.Errorf("Mult failed!")
	}
}

// ================= Mult w/ constant tests =================
func TestMultConstQuotRing(t *testing.T) {
	fx := &Polynomial{Coeffs: []int{1, 0, 1}}
	quot_ring := new(quotRing)
	quot_ring.q = 32
	quot_ring.mod = fx

	lhs := &Polynomial{Coeffs: []int{3, 17, 2, -3, 6}}
	rhs := 3

	res := quot_ring.mult_const(lhs, rhs)
	res = quot_ring.reduce_coefficients(res)

	if !reflect.DeepEqual(res.Coeffs, []int{21, 28}) {
		t.Errorf("Mult const failed!")
	}
}

// ================= Reduce tests =================
func TestReduce(t *testing.T) {
	fx := &Polynomial{Coeffs: []int{1, 3}}
	quot_ring := new(quotRing)
	quot_ring.q = 32
	quot_ring.mod = fx
	to_reduce := &Polynomial{Coeffs: []int{5, 7, 3}}

	if !reflect.DeepEqual(quot_ring.reduce(to_reduce).Coeffs, []int{3}) {
		t.Errorf("Reduce failed!")
	}
}

func TestReduceNegativeNumb(t *testing.T) {
	fx := &Polynomial{Coeffs: []int{1, 0, 1}}
	quot_ring := new(quotRing)
	quot_ring.q = 32
	quot_ring.mod = fx
	to_reduce := &Polynomial{Coeffs: []int{-17, 38, -12, 1}}
	res := quot_ring.reduce(to_reduce)
	res = quot_ring.reduce_coefficients(res)
	if !reflect.DeepEqual(res.Coeffs, []int{27, 5}) {
		t.Errorf("Reduce failed for negative numbers!")
	}
}

func TestReduceComplex(t *testing.T) {
	fx := &Polynomial{Coeffs: []int{1, 0, 0, 0, 0, 1}}
	quot_ring := new(quotRing)
	quot_ring.q = 32
	quot_ring.mod = fx
	to_reduce := &Polynomial{Coeffs: []int{13, 2, 5, -1}}
	res := quot_ring.reduce(to_reduce)
	res = quot_ring.reduce_coefficients(res)
	if !reflect.DeepEqual(res.Coeffs, []int{13, 2, 5, 31}) {
		t.Errorf("Reduce failed for complex case!")
	}
}
