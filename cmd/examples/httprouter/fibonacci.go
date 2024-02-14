package main

import (
	"math/big"
	"sync"
)

type Fibonacci struct {
	index *big.Int
}

var (
	fib     = Fibonacci{index: big.NewInt(0)}
	fibLock sync.RWMutex
)

func matrixFib(n *big.Int) *big.Int {
	if n.Cmp(big.NewInt(0)) == -1 {
		return big.NewInt(0)
	}
	if n.Cmp(big.NewInt(2)) == -1 {
		return n
	}
	base := [][2]*big.Int{
		{big.NewInt(1), big.NewInt(1)},
		{big.NewInt(1), big.NewInt(0)},
	}
	result := matrixPower(base, n)
	return result[0][1]
}

func matrixPower(M [][2]*big.Int, n *big.Int) [][2]*big.Int {
	result := [][2]*big.Int{{big.NewInt(1), big.NewInt(0)}, {big.NewInt(0), big.NewInt(1)}}
	for n.Cmp(big.NewInt(0)) == 1 {
		if new(big.Int).And(n, big.NewInt(1)).Cmp(big.NewInt(0)) != 0 {
			result = matrixMultiply(result, M)
		}
		M = matrixMultiply(M, M)
		n.Rsh(n, 1)
	}
	return result
}

func matrixMultiply(a, b [][2]*big.Int) [][2]*big.Int {
	x := new(big.Int).Add(new(big.Int).Mul(a[0][0], b[0][0]), new(big.Int).Mul(a[0][1], b[1][0]))
	y := new(big.Int).Add(new(big.Int).Mul(a[0][0], b[0][1]), new(big.Int).Mul(a[0][1], b[1][1]))
	z := new(big.Int).Add(new(big.Int).Mul(a[1][0], b[0][0]), new(big.Int).Mul(a[1][1], b[1][0]))
	w := new(big.Int).Add(new(big.Int).Mul(a[1][0], b[0][1]), new(big.Int).Mul(a[1][1], b[1][1]))
	return [][2]*big.Int{{x, y}, {z, w}}
}
