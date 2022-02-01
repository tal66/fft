package main

import (
	"errors"
	"fmt"
	"math"
	"math/cmplx"
)

func multiply(P []complex128, Q []complex128) []complex128 {

	l := len(P) + len(Q)
	var newP = make([]complex128, l)
	var newQ = make([]complex128, l)

	copy(newP, P)
	copy(newQ, Q)

	A := FFT(newP)
	B := FFT(newQ)

	var R = make([]complex128, len(A))
	for i := 0; i < len(R); i++ {
		R[i] = A[i] * B[i]
	}

	RInv := inverseFFT(R)
	for i := 0; i < len(R); i++ {
		RInv[i] = complex(math.Round(real(RInv[i]*100))/100,
			math.Round(imag(RInv[i]*100)/100))
	}
	return RInv
}

func FFT(P []complex128) []complex128 {
	var n int = len(P)
	if n == 1 {
		return P
	}

	if (n & (n - 1)) != 0 {
		pow, err := nextPowerOf2(n)
		if err != nil {
			fmt.Println(err)
		}

		n = pow
		var temp []complex128 = make([]complex128, n)
		copy(temp, P)
		P = temp
	}

	var O []complex128 = make([]complex128, n/2)
	var E []complex128 = make([]complex128, n/2)

	for i := 0; i < n/2; i++ {
		E[i] = P[2*i]
		O[i] = P[2*i+1]
	}

	var FFT_E []complex128 = FFT(E)
	var FFT_O []complex128 = FFT(O)

	var x complex128 = complex(2*math.Pi/float64(n), 0)
	var nRoot complex128 = cmplx.Cos(x) + complex(0, 1)*cmplx.Sin(x)
	var c complex128 = 1
	var R = make([]complex128, n)

	for i := 0; i < n/2; i++ {
		t := c * FFT_O[i]
		R[i] = FFT_E[i] + t
		R[i+n/2] = FFT_E[i] - t
		c *= nRoot
	}

	return R
}

func inverseFFT(P []complex128) []complex128 {
	var F []complex128 = FFT(P)
	var R = make([]complex128, len(F))

	nint := len(F)
	n := complex(float64(len(F)), 0)

	R[0] = cmplx.Conj(F[0]) / n
	for i := 1; i < nint; i++ {
		R[i] = cmplx.Conj(F[nint-i]) / n
	}

	return R
}

func nextPowerOf2(n int) (int, error) {
	switch {
	case n > 128 || n < 0:
		return n, errors.New("argument error")
	case n > 64:
		n = 128
	case n > 32:
		n = 64
	case n > 16:
		n = 32
	case n > 8:
		n = 16
	case n > 4:
		n = 8
	case n > 2:
		n = 4
	}
	return n, nil
}
