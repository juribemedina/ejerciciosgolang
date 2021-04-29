package main

import (
	"math/big"
)

func fibmatrix(n int) *big.Int {
	f := [4]*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(1), big.NewInt(1)}
	fibresul := mmultiter(f, Abs(n))[1]

	if n < 0 && n%2 == 0 {
		return fibresul.Mul(fibresul, big.NewInt(-1))
	}
	return fibresul
}

func mmultiter(fm [4]*big.Int, counter int) [4]*big.Int {
	if counter == 0 {
		return [4]*big.Int{big.NewInt(1), big.NewInt(0), big.NewInt(0), big.NewInt(1)}
	}

	if counter%2 == 0 {
		fn := mmultiter(fm, counter/2)
		return mmult(fn, fn)
	}

	fn := mmultiter(fm, counter-1)
	return mmult(fm, fn)
}

func mmult(fm, fn [4]*big.Int) [4]*big.Int {
	a, b, c, d := fm[0], fm[1], fm[2], fm[3]
	e, f, g, h := fn[0], fn[1], fn[2], fn[3]

	var resul [4]*big.Int

	x, y, z := big.NewInt(0), big.NewInt(0), big.NewInt(0)
	resul[0] = x.Add(y.Mul(a, e), z.Mul(b, g))

	m, n, p := big.NewInt(0), big.NewInt(0), big.NewInt(0)
	resul[1] = m.Add(n.Mul(a, f), p.Mul(b, h))

	r, s, t := big.NewInt(0), big.NewInt(0), big.NewInt(0)
	resul[2] = r.Add(s.Mul(c, e), t.Mul(d, g))

	u, v, w := big.NewInt(0), big.NewInt(0), big.NewInt(0)
	resul[3] = u.Add(v.Mul(c, f), w.Mul(d, h))

	return resul
}
