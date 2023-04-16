package large_mod

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"ThresholdKyber.com/m/ddec"
	kyberk2so "ThresholdKyber.com/m/kyber-k2so"
	owcpa "ThresholdKyber.com/m/owcpa_TKyber"
)

func TestEuclid(t *testing.T) {
	d, x, y := euclidsExtendedAlgorithm(7, 10)

	if d != 1 {
		t.Errorf("d incorrect!")
	}

	if x != 3 {
		t.Errorf("x incorrect!")
	}

	if y != -2 {
		t.Errorf("y incorrect!")
	}
}

// From http://www-math.ucdenver.edu/~wcherowi/courses/m5410/exeucalg.html
func TestEuclid2(t *testing.T) {
	d, x, y := euclidsExtendedAlgorithm(81, 57)

	if d != 3 {
		t.Errorf("d incorrect!")
	}

	if x != -7 {
		t.Errorf("x incorrect!")
	}

	if y != 10 {
		t.Errorf("y incorrect!")
	}
}

func testWritePartialTestVectorToFile(t *testing.T) {
	t_param := 2
	n_param := 3

	msg := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	params := owcpa.NewParameterSet("TKyber-Test")
	pk, sk_shares := owcpa.Setup(params, n_param, t_param)
	ct := owcpa.Enc(params, msg, pk)

	d_is := make([][]kyberk2so.Poly, n_param)
	for i := 0; i < t_param+1; i++ {
		d_is[i] = owcpa.PartDec(params, sk_shares[i], ct, i)
	}

	y := owcpa.Combine(params, ct, d_is, n_param, t_param)

	file_poly, err := os.Create("C:/Users/Kasper/Desktop/Speciale/ThresholdKyber/large_mod/test_vectors_crt/vector1_q3329") // creating...
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	defer file_poly.Close()

	ddec.WritePoly(y, file_poly)

	file_expected, err := os.Create("C:/Users/Kasper/Desktop/Speciale/ThresholdKyber/large_mod/test_vectors_crt/vector1_expected") // creating...
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	defer file_expected.Close()

	ddec.WriteExpectedKey(msg, file_expected)
}

func readTestVectorFromFile(name string) kyberk2so.Poly {
	content, err := ioutil.ReadFile("C:/Users/Kasper/Desktop/Speciale/ThresholdKyber/large_mod/test_vectors_crt/" + name)
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)
	numbers_as_str := strings.Fields(text)

	poly := kyberk2so.Poly{}

	for i := 0; i < kyberk2so.ParamsN; i++ {
		value, _ := strconv.ParseInt(numbers_as_str[i], 10, 16)
		poly[i] = int16(value)
	}

	return poly
}

func testMerge(t *testing.T) {
	poly1 := readTestVectorFromFile("vector1_q3329")
	poly2 := readTestVectorFromFile("vector1_q3313")

	q1 := 3329
	q2 := 3313

	res := Merge(poly1, poly2, q1, q2)

	fmt.Println(res)

	//t.Errorf("AAAAAAAA")
}
