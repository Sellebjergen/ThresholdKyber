package TKyber

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) { // Fails likely due to mod being regular and not euclid
	fx := &Polynomial{Coeffs: []int32{1, 0, 1}}
	quot_ring := new(polyRing)
	quot_ring.q = 32
	quot_ring.mod = fx

	lhs := &Polynomial{Coeffs: []int32{3, 6, 4, 2, 1}}
	rhs := &Polynomial{Coeffs: []int32{-17, 38, -12, 1}}

	res := quot_ring.add(lhs, rhs)

	fmt.Println(res)
	if !reflect.DeepEqual(res.Coeffs, []int32{27, 9}) {
		t.Errorf("Add failed!")
	}
}

func TestMul(t *testing.T) {
	fx := &Polynomial{Coeffs: []int32{1, 0, 1}}
	quot_ring := new(polyRing)
	quot_ring.q = 32
	quot_ring.mod = fx

	lhs := &Polynomial{Coeffs: []int32{3, 5, 0, 8}}
	rhs := &Polynomial{Coeffs: []int32{1, 1, 5}}

	res := quot_ring.mult(lhs, rhs)

	fmt.Println(res)
	if !reflect.DeepEqual(res.Coeffs, []int32{23, 15}) {
		t.Errorf("Sub failed!")
	}
}

func TestReduce(t *testing.T) {
	fx := &Polynomial{Coeffs: []int32{1, 3}}
	quot_ring := new(polyRing)
	quot_ring.q = 32
	quot_ring.mod = fx
	to_reduce := &Polynomial{Coeffs: []int32{5, 7, 3}}
	fmt.Println(quot_ring.reduce(*to_reduce))
	if !reflect.DeepEqual(quot_ring.reduce(*to_reduce).Coeffs, []int32{3}) {
		t.Errorf("Reduce failed!")
	}

	fx_2 := &Polynomial{Coeffs: []int32{1, 0, 1}}
	quot_ring_2 := new(polyRing)
	quot_ring_2.q = 32
	quot_ring_2.mod = fx_2
	to_reduce_2 := &Polynomial{Coeffs: []int32{-17, 38, -12, 1}}
	if !reflect.DeepEqual(quot_ring_2.reduce(*to_reduce_2).Coeffs, []int32{27, 5}) {
		t.Errorf("Reduce failed!")
	}

}
