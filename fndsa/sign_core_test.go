package fndsa

import (
	"bytes"
	sha3 "golang.org/x/crypto/sha3"
	"testing"
)

func TestSignCore(t *testing.T) {
	logn := uint(9)
	n := 1 << logn
	sh := sha3.NewShake256()
	sh.Write(kat_vrfy_key)
	var hvk [64]byte
	sh.Read(hvk[:])
	sig := make([]byte, SignatureSize(logn))
	err := sign_core(logn,
		kat_f, kat_g, kat_F, kat_G,
		hvk[:], DOMAIN_NONE, 0, []byte("test"), []byte("seed"),
		sig, make([]int16, n), make([]uint16, n), make([]f64, n*9))
	if err != nil {
		t.Fatalf("failure, err = %v", err)
	}
	if !bytes.Equal(sig, kat_sig) {
		t.Fatalf("ERR: wrong signature value: %v", sig)
	}
}

var kat_f = []int8{
	9, 1, -2, 3, 4, 0, -2, 8, 0, -4, 5, 0, 4, -2, 4, 3, -7, 0, 2, 3, 2,
	-2, 3, 3, -1, 3, 1, 0, 4, -1, -4, 8, 1, 1, -3, -1, 5, -3, -2, -4, 3,
	-2, 2, 0, 0, -6, 1, -1, 0, -3, 4, 3, 1, -3, 5, -7, 10, -3, 2, 3, 4,
	-6, -10, 2, 3, -3, 0, 1, 1, -1, -1, 4, 5, -4, -1, 6, 4, -4, -2, 1,
	0, -2, 3, 0, -4, 10, -2, -10, 1, 1, 6, 4, -4, -1, -1, 3, 0, 0, -4,
	-4, -1, -3, -3, -1, -2, -2, 0, -1, 4, 2, 1, -2, 0, 3, -2, 2, -5, -5,
	-7, 1, -4, 12, -2, -1, -3, 2, -4, 3, -7, -8, 5, 1, -9, -3, 1, 10, 0,
	1, 1, 5, -4, 0, 2, -3, 4, 2, -4, 0, 8, -3, 0, 6, 2, 2, -4, 5, -8,
	-1, 3, 6, 2, -6, -6, -1, 3, 0, 3, 2, -3, 0, 3, -4, 6, 1, -4, 0, 4,
	0, 4, -10, -3, 4, -5, 2, 1, -5, 0, 3, 5, 1, -1, -3, -4, 1, 4, 0, -5,
	3, -2, 4, 1, 2, -5, -8, -6, -2, 8, 6, -1, -3, -5, 8, -1, 5, 9, -1,
	1, -6, -6, 6, 1, 4, -2, -4, -1, -2, -4, -1, -5, 5, -4, -6, 2, 3, -8,
	0, 6, 1, -9, 7, -6, 5, 0, 5, 0, -5, 8, 0, 2, -3, 3, -5, 13, 5, 10,
	-4, 10, 0, 1, -2, 1, 0, -5, 1, 6, 8, 0, 0, 4, 8, -2, -7, -2, 1, 2,
	6, 7, 7, -2, 2, -5, -6, -3, 0, 4, 0, -2, 4, 2, 4, 5, -7, -2, -5, 5,
	-2, 0, 5, 0, -2, 1, 6, 1, 0, 4, 3, -2, 1, 2, 3, 0, 2, -2, 3, 4, -1,
	5, -1, 3, -7, -5, -7, 1, 2, 2, 2, 4, -1, 11, -2, 4, -2, 0, 3, 2, -2,
	3, -1, -5, -2, 0, -2, 7, 1, 6, -6, -4, -7, -1, 9, 2, -1, 2, 5, 0,
	-3, -4, -3, -2, -5, 0, 2, 3, -2, -4, -4, -1, 1, -3, 0, -1, -2, 3,
	-1, -6, 0, 0, -5, 2, 0, -1, 4, -6, 3, -1, -5, -1, -1, -5, 5, 3, 3,
	2, 10, 5, -3, -1, 3, 3, -8, 2, -3, -2, 3, 4, 4, -2, 9, -4, -3, 2, 6,
	0, -3, 0, 1, 3, 3, -5, -4, 1, -3, -2, 6, 1, -4, 3, 1, 6, -1, 3, -4,
	8, 3, 1, -8, -3, 5, 1, 0, 0, -7, 4, -11, -3, -2, 1, 0, 2, -2, 1, -1,
	-4, 0, 2, -2, 6, -6, 7, 3, -1, -4, -1, -3, -5, -6, -1, -3, 1, 5, 5,
	6, -5, 4, 5, -2, 2, 3, 5, -4, -2, 0, -4, 8, -8, 0, 4, -1, -3, 5, 2,
	6, 5, 3, 8, -1, -9, 1, -1, 9, -1, -4, 2, 5, 1, 5, -1, -1, -5, 1, 4, 3,
}

var kat_g = []int8{
	-5, -2, -4, 1, -5, 2, -4, -6, 6, -1, -1, 0, 0, -4, 3, -1, -6, 2, 3,
	-2, 0, -4, 5, -2, 5, 2, -2, -8, 4, -1, 11, -4, 2, -4, -1, 6, 9, 0,
	-7, 4, -3, 0, 0, 1, 11, 3, 3, 1, 0, -2, -1, -1, 6, -5, 7, 6, 9, -3,
	-5, 0, -6, 1, -4, -2, 0, 2, -5, 2, -2, 0, 1, 1, 0, 1, -6, 2, 0, 4,
	-6, 3, 5, 2, 1, -6, -1, 2, 4, 2, -1, -1, -7, 2, -1, 0, -6, -3, 5,
	-1, 0, -4, 0, 1, 4, 4, 0, -1, -5, -1, -4, -8, 2, 2, 2, 0, -3, 0, -1,
	1, 8, 2, 0, 6, -3, -5, -5, 3, -2, 7, 3, -3, -1, 2, 5, 8, -1, 6, 0,
	-2, 3, 2, -4, 5, 8, 1, -1, 0, -6, 0, -5, 4, -6, 0, -3, 2, 3, -3, -3,
	3, 6, -4, 5, 2, 3, 0, -2, 0, -4, 0, 0, -4, -7, -5, 0, 2, 2, 1, -2,
	-1, 4, 2, 7, 0, -3, -2, -6, 7, 1, -1, -4, -7, -5, 2, 2, 0, -3, 3, 7,
	-3, 5, -2, -10, -2, 0, 4, -2, 1, 2, -5, 6, -2, 2, 5, -3, 4, -2, 9,
	-5, -4, 1, 3, 2, 3, -3, 4, -2, -2, -5, 1, 4, 2, 1, -1, 1, -2, -3,
	-2, 8, 0, -5, -4, 5, -1, -4, -4, 0, 3, -3, -1, 0, 1, -1, 1, 6, 2, 1,
	5, 0, 8, -2, 2, 6, 1, 2, -3, 4, -6, 1, -5, -2, 1, 6, 1, 0, -1, -1,
	1, -1, -3, -1, 8, -3, 7, 0, 2, -3, 1, 1, 0, 4, 0, 6, -3, 0, -7, 1,
	7, -1, 2, 6, 2, 2, -8, -3, -3, 3, -2, 1, -5, -5, 0, 4, -4, 4, 3, -7,
	-2, -4, 2, -1, 8, 2, -2, -6, -4, 4, 1, -5, -8, -6, 2, -1, 3, 2, 1,
	0, 1, 2, 3, -1, 2, -6, 5, 10, 0, 1, 1, -3, -1, 0, 2, -12, 3, -4, 2,
	7, -6, -2, -5, 8, 0, -6, 5, 2, -6, 4, -5, -12, 0, 2, 1, 0, -1, -1,
	-4, -3, -3, -2, -1, -4, 0, 2, -5, -2, -3, -4, -4, 4, 6, 3, 4, 6, 6,
	-6, 2, 4, -3, 2, 3, 4, 1, 2, 1, -4, -2, 0, -6, 4, -4, -4, 7, 0, 6,
	-4, 2, 5, 1, -1, 5, 1, -2, 0, -2, -1, 3, 4, 8, 0, 0, 1, -5, 2, 0, 5,
	-2, -3, -6, 6, -6, 3, 5, 3, -1, 4, -3, 3, 2, -3, 2, 2, -3, 1, 3, 4,
	3, -2, -4, 0, 6, -2, -1, -3, 1, -10, -3, 0, 3, -2, -2, 1, -2, -5,
	-1, -4, 12, 2, 4, 3, -4, -3, 5, 0, -2, 3, -4, -7, 2, -2, -8, 0, -4,
	-2, 0, -2, 0, -1, -6, 5, -2, -1, 1, 4, 4, -3, -5, 3, 2, 0, 4, 2, 4,
	-1, 9,
}

var kat_F = []int8{
	40, 26, -25, 24, 27, -30, -10, 20, 62, -46, 8, 5, -34, 13, 4, 9, 32,
	-2, 37, -23, 12, 9, -15, -42, -24, 31, 18, 18, -18, -50, -34, 35,
	19, 0, 11, 0, 50, 6, -38, 27, 19, -17, 16, -24, -1, -51, 41, 14,
	-13, 4, -53, 15, -6, 19, 11, 37, 41, -31, -38, -11, -25, -8, 2, 23,
	11, -4, -24, 5, -24, 41, 10, 9, 8, -23, 15, -18, -7, -40, 6, -21,
	-12, 4, 20, 18, -26, 6, 26, 16, -28, -8, 31, -20, 21, -72, -1, 12,
	24, -20, -36, -4, 28, -7, 30, -15, -17, -18, 39, -13, -3, 5, 19, 35,
	-9, 3, 14, -68, -33, 18, 24, 50, -40, 4, -8, -13, 69, 3, 6, 3, -20,
	-23, -12, 10, 9, -11, 2, 6, -33, -16, -49, 35, -5, 17, 28, -49, -41,
	20, -54, -50, 2, 13, 23, -40, -5, 41, -53, 24, 16, 14, 17, -28, -47,
	-22, 35, -5, 19, -9, -1, 25, -31, 16, 12, 67, -20, -27, -45, 20, -5,
	44, -31, 1, 38, -8, -14, 21, 24, -7, -7, 24, 52, 52, -25, 6, 13, 3,
	18, -40, -1, -15, -28, 5, 17, -2, 16, 3, -20, -22, -7, 49, 21, -8,
	-1, -25, 32, 3, -62, 12, -4, 6, -7, -41, 41, -21, -44, -26, 33, -16,
	-45, -19, -16, -20, 33, 13, 4, 56, -31, -17, 7, 22, -16, -50, 4,
	-43, -21, 0, 3, 11, 33, 4, -16, 36, -9, -39, 0, -25, 4, 65, -9, 43,
	10, 41, 3, 30, -46, -24, 5, 42, 9, 19, 0, -6, -28, 14, 5, -4, -25,
	2, -7, 1, -15, 12, -16, -8, -10, 31, -8, 23, -11, 16, -11, 4, -47,
	-1, -30, 3, 22, -11, 14, -18, -10, 21, 29, 13, -14, -21, -19, 24,
	-7, -65, 39, -10, 2, 9, 2, -13, -26, -38, -6, 26, -8, 3, -6, 22,
	-27, 24, -28, -5, -17, 19, -27, 30, 61, 55, 4, 22, 14, -41, -21, 26,
	14, -35, 25, 21, 6, -13, 16, 2, 39, 18, 40, -8, -38, 9, 8, 22, 12,
	22, 8, -60, 44, 57, -19, -62, -15, 19, -31, -8, -11, 0, -6, -25, 46,
	18, -19, 9, -13, 17, 21, -29, 11, -36, 37, 20, 14, -33, -30, -12,
	49, -10, 16, -37, -16, 0, -20, -31, 63, 29, 35, -18, 16, 42, -7, 11,
	-1, -26, -43, 1, -12, 79, 31, 20, -1, 49, 12, -48, -40, -7, 18, 12,
	-18, -20, -21, 49, 45, -36, 39, 24, -5, -40, 5, 14, 64, 0, -41, -4,
	32, 26, -30, -26, -42, 53, -34, 6, 25, -1, -20, -18, 31, 15, -27,
	10, -32, 18, 21, 7, -39, -26, 57, -21, -7, -48, 24, 58, -19, -3, 10,
	32, 23, -15, -58, 49, 20, 20, 43, -51, 4, -5, -42, 44, 15, 34, -12,
	-6, 29, -41, -21, -11, -27, 37, 32, 10, -16, -34, 26, -30, 17, 7,
	20, 19, -19, 21, 11, -13, 12, -40, -19, 4, 22, -41, -14, 3, 39, 9,
}

var kat_G = []int8{
	34, -5, 41, 26, -7, 4, -23, 18, -10, -25, -5, -24, 8, -55, 29, 22,
	11, 25, -16, 21, -3, -29, -31, 13, 17, -18, 9, 25, 23, -33, -52,
	-21, -18, 27, -23, 1, 65, 0, -3, 0, -47, -3, -73, -20, -4, -5, 24,
	-26, 54, -1, 29, 4, -13, 30, -14, 15, -31, 2, 13, -11, 21, 37, -19,
	-16, -39, -48, -15, -5, 24, 2, -14, 0, 41, -8, 28, -52, 50, 31, 23,
	-32, -12, -15, -35, 7, 4, -20, -8, -41, 20, -5, -3, -19, -10, 14,
	-17, -11, 13, 16, -43, 35, -26, 6, 2, 12, 3, -74, -1, -53, 20, 43,
	28, 21, -12, -4, -23, 30, 26, 46, 8, -22, 6, 26, -53, 3, -2, 37, -9,
	4, 3, 5, -5, 7, 34, -21, -49, -6, 46, 3, 50, 11, -77, -21, 1, -32,
	-5, -39, 77, 38, -11, -15, -35, -30, 12, 14, 26, -4, -52, -22, 39,
	-47, 9, 8, -22, -17, -6, 4, 35, 27, 28, -13, 40, 25, -71, 21, 5, 5,
	0, 11, -9, 18, -17, -4, 25, 31, -2, -24, -2, 32, -13, 7, -17, 20,
	-3, -33, -21, 17, 15, -17, 21, 7, 28, -45, 31, 1, 20, -27, 11, 24,
	-36, 34, 29, 24, -5, 3, 0, -56, 17, 27, -42, -43, 11, 19, -12, -15,
	20, -8, -5, -25, -6, 17, -5, 34, -22, -48, -29, 19, -27, 18, 19,
	-28, 17, 5, -26, -17, -41, 2, -10, 22, 5, -9, -39, -13, -3, -30, 21,
	-13, 24, 8, -28, 33, 21, -39, -67, 31, 27, 45, -6, 11, -3, -6, -30,
	-7, 3, 0, 24, -6, 12, 24, 1, 13, 31, 21, 0, -40, 15, -1, -32, 13,
	-2, -23, 26, 3, 52, 18, 23, 3, -13, 25, -42, 19, 29, -20, 29, 4,
	-28, -15, -13, -10, 13, 27, 30, -77, 1, 61, 32, 10, -42, -25, -77,
	-3, 7, 8, 24, -11, -4, -28, -30, -3, -10, 30, -3, 65, -11, -57, 32,
	-37, 25, -34, 44, 15, -57, -14, 22, -4, -33, 9, 29, -14, 18, -18,
	46, -29, -4, -15, -41, 14, 2, 1, 4, -33, -12, 14, 6, -34, -35, 19,
	41, 35, 27, -6, -26, -25, 25, 5, 30, 5, 3, 30, 7, -7, 19, 8, 10, 0,
	32, 11, 7, 43, -3, -10, 13, 8, 17, -1, -10, -37, -11, 8, 40, -57,
	-27, 25, -26, 14, -15, -41, 45, -11, -3, -2, -28, -6, -19, -16, -15,
	12, 58, -16, -29, 21, -15, -17, -18, 5, -44, 11, -1, -12, 34, -68,
	2, 7, 62, 29, -38, 18, 16, -6, -10, -14, 23, 46, 50, 47, 60, -49, 0,
	-8, 5, 0, -62, -2, -39, 25, -16, 22, 2, 6, 1, 29, 11, 15, -19, 5, 8,
	18, -5, -14, 27, 38, -9, -5, -27, -42, -26, 53, 49, -9, 1, -19, -21,
	-13, -16, -5, 37, 4, 6, -21, -15, -39, -1, 42, 11, 8, -6, -33, -7,
	3, -5, -42, 21, -10, 27, -39, -64, 7, 22, 29, -7, 23, -26, -27,
}

var kat_sig = []byte{
	0x39, 0x59, 0x0D, 0x3E, 0xC0, 0x05, 0x4C, 0x1A, 0x2C, 0x45, 0xEC, 0xEE,
	0xB0, 0x4D, 0x3D, 0x07, 0xEC, 0xB3, 0x1D, 0xD3, 0x67, 0x0F, 0xEB, 0x7C,
	0x47, 0x3E, 0xAB, 0x4A, 0x74, 0xB9, 0x34, 0x8A, 0xF2, 0xB3, 0xF4, 0x5B,
	0x6C, 0xF2, 0xFA, 0xC5, 0x3F, 0x80, 0x17, 0xC3, 0xA2, 0x45, 0xDA, 0xA2,
	0x4F, 0x6B, 0xE4, 0x32, 0x27, 0x36, 0x3B, 0xAE, 0x20, 0xDB, 0xCD, 0x5E,
	0x19, 0x30, 0x79, 0xA9, 0xF8, 0x24, 0x5C, 0xFA, 0xB4, 0xCC, 0x32, 0x56,
	0x61, 0x21, 0xDB, 0x6F, 0xF9, 0x96, 0x9B, 0x92, 0x64, 0x42, 0x3F, 0x9F,
	0xD6, 0xD8, 0xDD, 0x37, 0x45, 0x8C, 0x19, 0x63, 0x7D, 0x99, 0xBD, 0x5D,
	0x5A, 0x0F, 0xF1, 0xA6, 0x8D, 0xC6, 0xD2, 0x47, 0x3E, 0x81, 0x9F, 0xB4,
	0xA6, 0x98, 0xC9, 0x5E, 0x45, 0xA5, 0x28, 0x0F, 0x02, 0xBD, 0xF0, 0xED,
	0x08, 0x59, 0x96, 0xEB, 0x49, 0xA3, 0xEE, 0x2C, 0x77, 0x2F, 0x43, 0x89,
	0xEA, 0xF9, 0x08, 0x9F, 0x71, 0xE3, 0x81, 0x2D, 0x5D, 0x9E, 0x16, 0xC6,
	0xF5, 0xB3, 0x94, 0x1A, 0x06, 0xB7, 0x79, 0x73, 0x22, 0x7D, 0x3E, 0xDF,
	0xF6, 0x3C, 0xD3, 0x79, 0x54, 0x03, 0x13, 0x82, 0x87, 0x52, 0x8F, 0xA5,
	0x4F, 0x51, 0xB8, 0xB5, 0x5C, 0x55, 0xA8, 0x5C, 0xF0, 0xDB, 0x8C, 0xC4,
	0xFA, 0x0F, 0x08, 0x61, 0x96, 0x0C, 0x83, 0xDD, 0x48, 0x61, 0xDD, 0x0B,
	0x4F, 0x79, 0xCB, 0xC3, 0x9D, 0x38, 0xAB, 0xAD, 0x7D, 0xD8, 0xF4, 0x31,
	0xCB, 0x36, 0xDB, 0x96, 0xC2, 0x59, 0xD8, 0x56, 0x10, 0xC2, 0xD0, 0xFB,
	0xFC, 0x9E, 0x02, 0x54, 0xF1, 0xE6, 0x17, 0x26, 0xAD, 0x89, 0x93, 0xF5,
	0x11, 0xAC, 0x8F, 0xCA, 0xD1, 0x4A, 0x4E, 0x10, 0xBE, 0x7A, 0x38, 0xD0,
	0x4A, 0x26, 0xDB, 0x98, 0x31, 0x25, 0x78, 0x07, 0xB7, 0x82, 0x68, 0xA1,
	0xDD, 0x7F, 0x4F, 0xA9, 0x2B, 0xCF, 0x7E, 0xEC, 0xB0, 0x19, 0xF4, 0x24,
	0xDC, 0xC0, 0x57, 0x1F, 0xC1, 0x29, 0x35, 0xC4, 0xC6, 0xBF, 0xE1, 0x77,
	0xEE, 0x39, 0x00, 0x63, 0x60, 0x28, 0x19, 0xBA, 0x56, 0x74, 0x1A, 0x66,
	0x5B, 0x4F, 0xDF, 0xEE, 0xFD, 0x9D, 0x98, 0x0C, 0x15, 0xB8, 0x56, 0x90,
	0xD3, 0xD7, 0x1C, 0x83, 0xEB, 0x5C, 0x2C, 0x5F, 0x6F, 0x5E, 0xAF, 0x9B,
	0x2F, 0xAA, 0x44, 0xC1, 0xA1, 0x5C, 0x17, 0xEB, 0x69, 0xD8, 0x4D, 0x10,
	0x65, 0x69, 0x47, 0x83, 0x3E, 0x32, 0xDE, 0x64, 0x9A, 0x6F, 0xCB, 0x9B,
	0x26, 0x6F, 0xDE, 0xAC, 0xCC, 0x42, 0x51, 0xAC, 0x1B, 0xAE, 0x85, 0x7B,
	0x7D, 0x08, 0x89, 0x6D, 0xC6, 0x38, 0xD5, 0x2C, 0x4F, 0xF7, 0x0E, 0xFB,
	0x62, 0x19, 0xEA, 0x77, 0x82, 0x0A, 0xBC, 0x4F, 0x20, 0x24, 0x71, 0x2E,
	0xB6, 0xFF, 0x22, 0x90, 0x03, 0x34, 0x95, 0x90, 0x38, 0x92, 0x21, 0xFC,
	0xA1, 0x22, 0x95, 0xB4, 0xA1, 0x00, 0xC2, 0xEE, 0x36, 0xBC, 0xAD, 0xD4,
	0x2B, 0xD1, 0xD3, 0x28, 0x49, 0x5A, 0xE1, 0x92, 0x7C, 0x14, 0x6C, 0xFC,
	0x12, 0xCB, 0xF2, 0xFE, 0x58, 0xF4, 0xE8, 0x4A, 0x97, 0x1B, 0x53, 0x22,
	0x9E, 0xC7, 0x8A, 0x6F, 0xDF, 0xA0, 0xD5, 0x2C, 0xCC, 0x32, 0x16, 0x8B,
	0xC8, 0x23, 0x12, 0xD7, 0x07, 0x94, 0x8E, 0x91, 0xE6, 0x77, 0x8E, 0x9D,
	0xFA, 0x56, 0xCC, 0x79, 0x47, 0x67, 0xD9, 0x24, 0xE2, 0x03, 0xBA, 0x9B,
	0x7B, 0x3F, 0xF6, 0x0D, 0x4B, 0x36, 0x83, 0xDE, 0xE1, 0xC2, 0xCD, 0xA7,
	0xF1, 0x6A, 0x2C, 0xAC, 0x6D, 0x77, 0x4C, 0x9B, 0x55, 0x7B, 0x58, 0xD7,
	0x7E, 0x30, 0xA5, 0x64, 0x99, 0x16, 0xDE, 0xC8, 0x87, 0x6F, 0xDE, 0x28,
	0x12, 0xB9, 0x47, 0xAB, 0xDD, 0x99, 0xF8, 0xAA, 0xC4, 0xC7, 0xAC, 0x18,
	0xF8, 0xB5, 0xCA, 0x6F, 0xE2, 0x6F, 0x10, 0x06, 0x7D, 0x3A, 0x11, 0xB5,
	0x8C, 0x37, 0xF9, 0x34, 0x4B, 0x87, 0x00, 0xF7, 0x47, 0x1A, 0xE4, 0x82,
	0x52, 0xF3, 0x48, 0x5B, 0x5B, 0xDC, 0xF6, 0x65, 0x54, 0x4C, 0x0D, 0xBF,
	0x71, 0x83, 0xEF, 0xE8, 0x7A, 0x7C, 0x39, 0x8A, 0x21, 0x58, 0xDE, 0xE8,
	0x66, 0xB9, 0x1D, 0x9E, 0x2E, 0x55, 0x83, 0xC1, 0xA2, 0x10, 0x02, 0x4D,
	0xE2, 0xC2, 0xAE, 0xFE, 0x5E, 0x4C, 0x2D, 0x7D, 0x8A, 0x0A, 0x8D, 0x06,
	0x3B, 0x80, 0x19, 0xC5, 0x0F, 0x6F, 0x85, 0x57, 0xE7, 0x26, 0x78, 0xE9,
	0xC4, 0x9A, 0xBC, 0xB0, 0x86, 0xF8, 0x94, 0x97, 0x6A, 0x6A, 0xDA, 0x77,
	0xD2, 0x26, 0xF5, 0x1B, 0x95, 0x96, 0xA4, 0xDF, 0xF0, 0x8D, 0x4E, 0x85,
	0xD7, 0x7B, 0x4D, 0x32, 0xB7, 0x4C, 0x5B, 0xBB, 0xB1, 0xD0, 0x68, 0x9B,
	0xDE, 0x43, 0x6F, 0x08, 0xBB, 0xD8, 0x2D, 0xCB, 0x33, 0x44, 0xE5, 0xD0,
	0xF9, 0x9C, 0xAF, 0xD6, 0x9D, 0x5F, 0x81, 0x90, 0x13, 0x0B, 0xDA, 0x99,
	0xE0, 0x99, 0x78, 0xF4, 0x1D, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
}

var kat_vrfy_key = []byte{
	0x09, 0x54, 0x61, 0xAF, 0xA4, 0xD8, 0x9E, 0x69, 0x45, 0x66, 0x94, 0x99,
	0x26, 0x62, 0xAF, 0x9A, 0x6D, 0xED, 0x0B, 0x63, 0x2F, 0xCE, 0x3D, 0x0D,
	0x29, 0x02, 0xA4, 0x0C, 0xB0, 0x7F, 0x28, 0x30, 0x70, 0xD9, 0x19, 0x9A,
	0xB6, 0x36, 0xB8, 0x78, 0x4C, 0x95, 0x8B, 0x72, 0x68, 0xCA, 0x06, 0x3C,
	0x5D, 0x53, 0x3C, 0x2C, 0xEB, 0x10, 0x64, 0x4A, 0xD1, 0x44, 0xEA, 0x50,
	0x57, 0x1F, 0xDF, 0x98, 0x37, 0x6E, 0xFB, 0xCA, 0xD0, 0x64, 0x2A, 0x83,
	0xD4, 0x10, 0x25, 0xB7, 0x27, 0xFC, 0xB2, 0xDD, 0xDF, 0xA1, 0x70, 0xA3,
	0xF9, 0x0A, 0xF0, 0x37, 0x98, 0x46, 0x5A, 0x1C, 0x8F, 0x6A, 0x13, 0x6B,
	0xFD, 0x86, 0x0C, 0x2F, 0xAA, 0x74, 0x92, 0x19, 0xAA, 0xF3, 0x8D, 0x8E,
	0x0C, 0x49, 0xA9, 0x8B, 0x78, 0x90, 0xE1, 0xF4, 0xA9, 0x73, 0x8E, 0x52,
	0x61, 0xC0, 0xB3, 0x4B, 0xD0, 0xAF, 0x2D, 0x97, 0xD0, 0xDF, 0xD7, 0x3F,
	0x5E, 0x9D, 0x5A, 0xF5, 0x1D, 0xC5, 0x2D, 0x57, 0x95, 0x17, 0xF9, 0x51,
	0x28, 0x85, 0x24, 0x83, 0x34, 0xD1, 0x44, 0x31, 0x2C, 0x63, 0xC2, 0x67,
	0xA8, 0x0A, 0x13, 0x1E, 0x96, 0x5D, 0x55, 0x52, 0xAC, 0x72, 0x9F, 0x2D,
	0x59, 0xA9, 0xC1, 0x7C, 0xB5, 0xA6, 0x6A, 0x71, 0x1B, 0x48, 0x64, 0x87,
	0x1D, 0x04, 0x85, 0x25, 0x4C, 0x0C, 0xA9, 0x04, 0xDB, 0x7C, 0x7A, 0x94,
	0x1F, 0x60, 0x12, 0x1F, 0x49, 0x10, 0x7C, 0x58, 0xF8, 0x19, 0xA1, 0x0B,
	0xAB, 0x0E, 0x21, 0x7B, 0x0F, 0x89, 0x0E, 0x40, 0x25, 0xC5, 0x72, 0x90,
	0x55, 0xA6, 0x02, 0xA0, 0x10, 0x48, 0x03, 0x83, 0x52, 0x92, 0x50, 0x54,
	0x9B, 0x1E, 0xC1, 0x9B, 0x46, 0xA4, 0x35, 0xE5, 0x07, 0x03, 0x82, 0x3D,
	0x79, 0xBC, 0xEA, 0x62, 0x65, 0xB7, 0x93, 0x00, 0x26, 0x40, 0x9B, 0xCF,
	0xC8, 0x00, 0x34, 0xDD, 0xDB, 0x26, 0xCA, 0x22, 0x90, 0xB5, 0x9B, 0xA2,
	0x63, 0x97, 0x2F, 0x90, 0xD5, 0xEA, 0x15, 0x0C, 0xE2, 0xD6, 0x76, 0xFD,
	0x13, 0x20, 0xEA, 0x81, 0xFF, 0x36, 0xC5, 0xBB, 0x59, 0xF2, 0xEC, 0xCB,
	0xB2, 0xA5, 0x8F, 0x4B, 0x33, 0xA9, 0xE3, 0x97, 0x28, 0x21, 0x5A, 0xA5,
	0x22, 0x41, 0x43, 0xE0, 0xE9, 0xA6, 0xCE, 0x65, 0x0D, 0xA8, 0xD5, 0x29,
	0x35, 0xE4, 0xC9, 0x89, 0x16, 0x25, 0x1F, 0x79, 0x15, 0x26, 0x1C, 0xAC,
	0xEA, 0xB4, 0x53, 0x9E, 0x5E, 0xA7, 0x08, 0x4E, 0xDC, 0x00, 0xC7, 0x2E,
	0x15, 0x96, 0x35, 0x63, 0x87, 0x09, 0x17, 0x6F, 0x9C, 0x85, 0x9C, 0x56,
	0xCC, 0x50, 0x36, 0x70, 0xBA, 0xFA, 0x6A, 0x9D, 0x54, 0xCA, 0x96, 0x01,
	0x10, 0xD3, 0x02, 0x17, 0x03, 0x70, 0x95, 0x05, 0x60, 0xB7, 0x01, 0x3B,
	0xB9, 0x6A, 0x04, 0x00, 0x60, 0x9B, 0x9D, 0x1E, 0x11, 0xA9, 0x65, 0xFD,
	0x88, 0x35, 0x30, 0xA4, 0xC7, 0x45, 0xC7, 0x41, 0x57, 0x90, 0x59, 0x03,
	0xC9, 0xA0, 0xE0, 0xE2, 0x3B, 0xF9, 0x15, 0xD4, 0xAE, 0x85, 0xC7, 0x33,
	0xB6, 0xF0, 0xA8, 0x65, 0x46, 0x67, 0x0B, 0xFA, 0x08, 0xA6, 0x2A, 0xE1,
	0x10, 0x04, 0x04, 0x90, 0x52, 0x7A, 0x84, 0x10, 0x58, 0x11, 0x3B, 0x95,
	0x28, 0x20, 0x3C, 0x90, 0x9D, 0xE3, 0x61, 0xF1, 0xA1, 0xC6, 0x43, 0x94,
	0xDA, 0xA7, 0x7F, 0x0D, 0xD8, 0x93, 0x00, 0x1E, 0x53, 0xE0, 0x81, 0x7C,
	0x62, 0xED, 0x63, 0x2A, 0x6A, 0x4B, 0x56, 0x07, 0xA2, 0xA2, 0xE5, 0x24,
	0x02, 0x9C, 0xB9, 0x55, 0x25, 0x01, 0x19, 0x16, 0xA2, 0x09, 0xD4, 0x4D,
	0x7B, 0xB3, 0x26, 0x7E, 0x02, 0x8A, 0x19, 0xA3, 0x01, 0x61, 0xDB, 0x13,
	0xB6, 0x75, 0xD6, 0x27, 0x43, 0x1E, 0x79, 0x51, 0x2D, 0xE3, 0x29, 0xEC,
	0x4B, 0x13, 0x7A, 0xCD, 0x80, 0xDF, 0x20, 0xA4, 0x27, 0xD6, 0x67, 0x77,
	0x53, 0x53, 0x0A, 0x93, 0x38, 0x57, 0xD2, 0xBF, 0x87, 0x7A, 0x4C, 0x65,
	0x01, 0xD9, 0x76, 0x87, 0x16, 0x0B, 0x31, 0xE4, 0xD0, 0x95, 0x0D, 0x45,
	0xAA, 0x1E, 0x1E, 0xA0, 0x06, 0xDD, 0xEB, 0x8D, 0x12, 0x51, 0xC7, 0xBB,
	0x4E, 0x90, 0x12, 0xE2, 0x24, 0xF2, 0x8D, 0x43, 0x87, 0x00, 0xC6, 0xA0,
	0x80, 0x57, 0xDA, 0x1E, 0xAA, 0x2A, 0x4C, 0x9A, 0xCB, 0xD0, 0x19, 0x26,
	0xF4, 0x1B, 0x6B, 0xBC, 0x46, 0x5B, 0x72, 0x9E, 0x23, 0x18, 0x78, 0xEF,
	0xF5, 0x44, 0x3E, 0xE1, 0xD9, 0xD7, 0x91, 0x95, 0x07, 0x8A, 0x30, 0xF6,
	0x0A, 0xA4, 0x4A, 0x77, 0x58, 0x46, 0xA8, 0x68, 0xDC, 0x9C, 0x6C, 0x92,
	0x3A, 0xE0, 0x15, 0xE8, 0xF1, 0xBE, 0x94, 0x7B, 0xA7, 0x2B, 0x15, 0x2B,
	0x48, 0xF4, 0x65, 0x93, 0xF2, 0xED, 0x94, 0x5F, 0xF0, 0x1D, 0x79, 0xB6,
	0x4C, 0xB9, 0x54, 0xBD, 0x8A, 0xA8, 0x7D, 0x56, 0x8F, 0x1A, 0x72, 0x1F,
	0x1A, 0x85, 0x2B, 0xDD, 0x04, 0x41, 0x53, 0x1B, 0x83, 0x0E, 0x96, 0xB7,
	0x55, 0x7F, 0x17, 0x46, 0x1E, 0x28, 0x1E, 0xE6, 0xC3, 0xB6, 0x68, 0x6F,
	0x25, 0x56, 0xD0, 0x0C, 0x76, 0xC0, 0x2B, 0xE1, 0x95, 0x26, 0x66, 0xD8,
	0x5A, 0x23, 0x71, 0x2C, 0xE9, 0xC7, 0xA6, 0x56, 0x57, 0x71, 0x7D, 0x80,
	0x4F, 0xE1, 0x4F, 0xEE, 0xD2, 0x7D, 0x4C, 0xCD, 0xA7, 0x81, 0x6A, 0xB1,
	0x9D, 0xA2, 0x37, 0xC6, 0xD2, 0xE9, 0xE7, 0x80, 0x0D, 0x86, 0xEA, 0x0E,
	0x9B, 0xEE, 0x8C, 0x2C, 0x72, 0xE1, 0x0C, 0xE6, 0x09, 0x8C, 0x76, 0xD4,
	0x19, 0x69, 0xA8, 0x91, 0x3B, 0xB9, 0x73, 0x41, 0x45, 0x5B, 0xAB, 0x59,
	0xC6, 0x83, 0xD1, 0xA0, 0x97, 0xBC, 0x00, 0x42, 0x76, 0xE2, 0x7A, 0x68,
	0xFB, 0x77, 0x71, 0x3B, 0x6B, 0x63, 0xEC, 0x62, 0xAD, 0xAA, 0x3E, 0xD5,
	0xFD, 0xCB, 0x70, 0x53, 0x79, 0xB7, 0x46, 0x90, 0x4C, 0x5F, 0x17, 0x9A,
	0x3A, 0xC8, 0x04, 0xE5, 0xC0, 0xAF, 0x58, 0x6A, 0x22, 0xEE, 0x05, 0x74,
	0xAB, 0x65, 0xD2, 0xC0, 0x92, 0x64, 0x48, 0x47, 0x7C, 0xB5, 0x33, 0xBA,
	0xCC, 0x97, 0x80, 0x80, 0x68, 0x96, 0xF8, 0xD5, 0x9C, 0x95, 0xDD, 0x5F,
	0xC6, 0x3D, 0x19, 0x6D, 0x8E, 0xAC, 0x3D, 0xA0, 0x7B, 0x58, 0x67, 0xB3,
	0x99, 0xBE, 0xA0, 0xED, 0x08, 0xEB, 0xB4, 0xBE, 0x73, 0x94, 0x09, 0x56,
	0x12, 0x39, 0x0A, 0x03, 0x06, 0xF1, 0x08, 0x9D, 0x0B, 0x5C, 0x1F, 0x20,
	0xDF, 0x41, 0x95, 0x99, 0x24, 0x0C, 0xE8, 0x18, 0x6F, 0x47, 0x2E, 0x9D,
	0x98, 0x71, 0xC7, 0xDB, 0xBA, 0x77, 0xAD, 0x82, 0x38, 0xF7, 0xAB, 0xF7,
	0xA5, 0x2C, 0x36, 0xE1, 0x82, 0x5E, 0x93, 0x75, 0x12, 0x97, 0x84, 0xEB,
	0xCC, 0x8D, 0x59, 0x54, 0x8E, 0x9A, 0xC0, 0x5B, 0xB3,
}
