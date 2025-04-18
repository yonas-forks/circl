// Code generated by fiat.go using fiat-crypto v0.1.4.
//
// Autogenerated: './FiatCrypto_v.0.1.4-issue1672' word-by-word-montgomery --output 'fp64/fiatMont.go' --lang Go --package-name fp64 --doc-prepend-header 'Code generated by fiat.go using fiat-crypto v0.1.4.' --package-case lowerCamelCase --public-function-case lowerCamelCase --public-type-case lowerCamelCase --doc-newline-before-package-declaration --no-primitives --widen-carry --no-field-element-typedefs --relax-primitive-carry-to-bitwidth 64 Fp 64 0xffffffff00000001 add sub mul square
//
// curve description: Fp
//
// machine_wordsize = 64 (from "64")
//
// requested operations: add, sub, mul, square
//
// m = 0xffffffff00000001 (from "0xffffffff00000001")
//
//
//
// NOTE: In addition to the bounds specified above each function, all
//
//   functions synthesized for this Montgomery arithmetic require the
//
//   input to be strictly less than the prime modulus (m), and also
//
//   require the input to be in the unique saturated representation.
//
//   All functions also ensure that these two properties are true of
//
//   return values.
//
//
//
// Computed values:
//
//   eval z = z[0]
//
//   bytes_eval z = z[0] + (z[1] << 8) + (z[2] << 16) + (z[3] << 24) + (z[4] << 32) + (z[5] << 40) + (z[6] << 48) + (z[7] << 56)
//
//   twos_complement_eval z = let x1 := z[0] in
//
//                            if x1 & (2^64-1) < 2^63 then x1 & (2^64-1) else (x1 & (2^64-1)) - 2^64

package fp64

import "math/bits"

// The function fiatFpAdd adds two field elements in the Montgomery domain.
//
// Preconditions:
//   0 ≤ eval arg1 < m
//   0 ≤ eval arg2 < m
// Postconditions:
//   eval (from_montgomery out1) mod m = (eval (from_montgomery arg1) + eval (from_montgomery arg2)) mod m
//   0 ≤ eval out1 < m
//
// Input Bounds:
//   arg1: [[0x0 ~> 0xffffffffffffffff]]
//   arg2: [[0x0 ~> 0xffffffffffffffff]]
// Output Bounds:
//   out1: [[0x0 ~> 0xffffffffffffffff]]
func fiatFpAdd(out1 *Fp, arg1 *Fp, arg2 *Fp) {
	var x1 uint64
	var x2 uint64
	x1, x2 = bits.Add64(arg1[0], arg2[0], uint64(0x0))
	var x3 uint64
	var x4 uint64
	x3, x4 = bits.Sub64(x1, 0xffffffff00000001, uint64(uint64(0x0)))
	var x6 uint64
	_, x6 = bits.Sub64(x2, uint64(0x0), uint64(x4))
	var x7 uint64
	fiatFpCmovznzU64(&x7, x6, x3, x1)
	out1[0] = x7
}

// The function fiatFpSub subtracts two field elements in the Montgomery domain.
//
// Preconditions:
//   0 ≤ eval arg1 < m
//   0 ≤ eval arg2 < m
// Postconditions:
//   eval (from_montgomery out1) mod m = (eval (from_montgomery arg1) - eval (from_montgomery arg2)) mod m
//   0 ≤ eval out1 < m
//
// Input Bounds:
//   arg1: [[0x0 ~> 0xffffffffffffffff]]
//   arg2: [[0x0 ~> 0xffffffffffffffff]]
// Output Bounds:
//   out1: [[0x0 ~> 0xffffffffffffffff]]
func fiatFpSub(out1 *Fp, arg1 *Fp, arg2 *Fp) {
	x1 := arg2[0]
	var x2 uint64
	var x3 uint64
	x2, x3 = bits.Sub64(arg1[0], x1, uint64(0x0))
	var x4 uint64
	fiatFpCmovznzU64(&x4, x3, uint64(0x0), 0xffffffffffffffff)
	var x5 uint64
	x5, _ = bits.Add64(x2, (x4 & 0xffffffff00000001), uint64(0x0))
	out1[0] = x5
}

// The function fiatFpMul multiplies two field elements in the Montgomery domain.
//
// Preconditions:
//   0 ≤ eval arg1 < m
//   0 ≤ eval arg2 < m
// Postconditions:
//   eval (from_montgomery out1) mod m = (eval (from_montgomery arg1) * eval (from_montgomery arg2)) mod m
//   0 ≤ eval out1 < m
//
// Input Bounds:
//   arg1: [[0x0 ~> 0xffffffffffffffff]]
//   arg2: [[0x0 ~> 0xffffffffffffffff]]
// Output Bounds:
//   out1: [[0x0 ~> 0xffffffffffffffff]]
func fiatFpMul(out1 *Fp, arg1 *Fp, arg2 *Fp) {
	x1 := arg1[0]
	var x2 uint64
	var x3 uint64
	x3, x2 = bits.Mul64(x1, arg2[0])
	var x4 uint64
	_, x4 = bits.Mul64(x2, 0xfffffffeffffffff)
	var x6 uint64
	var x7 uint64
	x7, x6 = bits.Mul64(x4, 0xffffffff00000001)
	var x9 uint64
	_, x9 = bits.Add64(x2, x6, uint64(0x0))
	var x10 uint64
	var x11 uint64
	x10, x11 = bits.Add64(x3, x7, uint64(x9))
	var x12 uint64
	var x13 uint64
	x12, x13 = bits.Sub64(x10, 0xffffffff00000001, uint64(uint64(0x0)))
	var x15 uint64
	_, x15 = bits.Sub64(x11, uint64(0x0), uint64(x13))
	var x16 uint64
	fiatFpCmovznzU64(&x16, x15, x12, x10)
	out1[0] = x16
}

// The function fiatFpSquare squares a field element in the Montgomery domain.
//
// Preconditions:
//   0 ≤ eval arg1 < m
// Postconditions:
//   eval (from_montgomery out1) mod m = (eval (from_montgomery arg1) * eval (from_montgomery arg1)) mod m
//   0 ≤ eval out1 < m
//
// Input Bounds:
//   arg1: [[0x0 ~> 0xffffffffffffffff]]
// Output Bounds:
//   out1: [[0x0 ~> 0xffffffffffffffff]]
func fiatFpSquare(out1 *Fp, arg1 *Fp) {
	x1 := arg1[0]
	var x2 uint64
	var x3 uint64
	x3, x2 = bits.Mul64(x1, arg1[0])
	var x4 uint64
	_, x4 = bits.Mul64(x2, 0xfffffffeffffffff)
	var x6 uint64
	var x7 uint64
	x7, x6 = bits.Mul64(x4, 0xffffffff00000001)
	var x9 uint64
	_, x9 = bits.Add64(x2, x6, uint64(0x0))
	var x10 uint64
	var x11 uint64
	x10, x11 = bits.Add64(x3, x7, uint64(x9))
	var x12 uint64
	var x13 uint64
	x12, x13 = bits.Sub64(x10, 0xffffffff00000001, uint64(uint64(0x0)))
	var x15 uint64
	_, x15 = bits.Sub64(x11, uint64(0x0), uint64(x13))
	var x16 uint64
	fiatFpCmovznzU64(&x16, x15, x12, x10)
	out1[0] = x16
}
