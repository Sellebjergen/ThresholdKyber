package gladius

import (
	"fmt"
	"strconv"
)

func bytesToGladiusMsg(params *GladiusParams, b []byte) []int64 {
	msg := make([]int64, params.N)
	for i, oneByte := range b {
		byte_as_string := fmt.Sprintf("%08b", oneByte)
		for j, char := range reverse(byte_as_string) {
			bit, _ := strconv.Atoi(string(char))
			msg[i*8+j] = int64(bit)

		}
	}
	return msg
}

// NOTE: Only supports 32 bytes, since this is length of key
func gladiusMsgToBytes(msg []int64) []byte {
	k := make([]byte, 32)
	for i := range k {
		for j := 0; j < 8; j++ {
			k[i] += byte(int(msg[i*8+j]) << j)
		}

	}
	return k
}

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}
