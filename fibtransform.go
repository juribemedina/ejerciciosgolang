package main

import (
	"math/big"
)

func fibtransform(n int) *big.Int {
	fibresul := fibitertrans(big.NewInt(1), big.NewInt(0), big.NewInt(0), big.NewInt(1), Abs(n))

	if n < 0 && n%2 == 0 {
		return fibresul.Mul(fibresul, big.NewInt(-1))
	}
	return fibresul
}

func fibitertrans(a, b, p, q *big.Int, counter int) *big.Int {
	if counter == 0 {
		return b
	}

	if counter%2 == 0 {
		qq, dosp := big.NewInt(0), big.NewInt(0)
		qq, dosp = qq.Mul(q, q), dosp.Mul(big.NewInt(2), p)
		p = p.Add(p.Mul(p, p), qq)
		q = q.Add(q.Mul(q, dosp), qq)
		return fibitertrans(a, b, p, q, counter/2)
	}

	aq, ap := big.NewInt(0), big.NewInt(0)
	aq, ap = aq.Mul(a, q), ap.Mul(a, p)
	bp, bq := big.NewInt(0), big.NewInt(0)
	bp, bq = bp.Mul(b, p), bq.Mul(b, q)
	a = a.Add(ap, a.Add(bq, aq))
	b = b.Add(bp, aq)

	return fibitertrans(a, b, p, q, counter-1)
}
