# numeric

Ultra-accurate and reliable big number library for Go. Powered with C GNU MPFR library.

# Requirements

1. [Install GMP](https://gmplib.org/manual/Installing-GMP) (GNU Multi Precision Arithmetic Library):

```bash
sudo apt-get install libgmp3-dev
```

2. [Install GNU MPFR](https://www.mpfr.org/mpfr-current/mpfr.html#Installing-MPFR) (GNU Multiple Precision Floating-Point Reliable Library):

```bash
sudo apt-get install libmpfr-dev
```

# Examples

```go
package main

import (
	"fmt"

	"github.com/acailuv/numeric"
)

func main() {
	a1 := numeric.New(10001.123)
	a2 := numeric.New(100001.123123)
	a3 := a1.Add(a2)
	fmt.Println(a1, "+", a2, "=", a3)

	s1 := numeric.New(3)
	s2 := numeric.New(4)
	s3 := s1.Subtract(s2)
	fmt.Println(s1, "-", s2, "=", s3)

	m1 := numeric.New(5)
	m2 := numeric.New(6)
	m3 := m1.Multiply(m2)
	fmt.Println(m1, "*", m2, "=", m3)

	d1 := numeric.New(19)
	d2 := numeric.New(6)
	d3 := d1.Divide(d2).Truncate(5)
	fmt.Println(d1, "/", d2, "=", d3.StringDecimalPlaces(3))

	x1 := numeric.New(1)
	x2 := x1.Divide(3)
	fmt.Println(x1, "/", 3, "=", x2)

	gt1 := numeric.New(1)
	gt2 := numeric.New(2)
	fmt.Println(gt1, ">", gt2, "=", gt1.GreaterThan(gt2))

	lt1 := numeric.New(1)
	lt2 := numeric.New(2)
	fmt.Println(lt1, "<", lt2, "=", lt1.LessThan(lt2))

	gte1 := numeric.New(1)
	gte2 := numeric.New(2)
	fmt.Println(gte1, ">=", gte2, "=", gte1.GreaterThanOrEqual(gte2))

	lte1 := numeric.New(1)
	lte2 := numeric.New(2)
	fmt.Println(lte1, "<=", lte2, "=", lte1.LessThanOrEqual(lte2))

	eq1 := numeric.New(1)
	eq2 := numeric.New(1)
	fmt.Println(eq1, "==", eq2, "=", eq1.Equal(eq2))

	taxRate := numeric.New("0.0011")
	feeRate := numeric.New("0.0008")
	buyerBrokerage := numeric.New("0.0285")
	sumTaxFeeRate := taxRate.Add(feeRate)
	taxProportion := taxRate.Divide(sumTaxFeeRate)
	tax := buyerBrokerage.Multiply(taxProportion)

	fmt.Println("Tax:", tax) // 0.0165000000 -- Other libraries would return something like 0.016499999...
}
```
