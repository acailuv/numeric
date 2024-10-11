package numeric

/*
#cgo LDFLAGS: -lmpfr
#include <mpfr.h>
*/
import "C"

// region Public

// Ceil the number to the specified decimal places.
func (n Numeric) Ceil(dp int) Numeric {
	if !n.init {
		n = New(0)
	}

	scale := New(10).Pow(dp)
	C.mpfr_mul_ui(&n.val[0], &n.val[0], C.ulong(scale.Uint()), C.MPFR_RNDN)
	C.mpfr_ceil(&n.val[0], &n.val[0])
	C.mpfr_div_ui(&n.val[0], &n.val[0], C.ulong(scale.Uint()), C.MPFR_RNDN)

	return n
}

// Floor the number to the specified decimal places.
func (n Numeric) Floor(dp int) Numeric {
	if !n.init {
		n = New(0)
	}

	scale := New(10).Pow(dp)
	C.mpfr_mul_ui(&n.val[0], &n.val[0], C.ulong(scale.Uint()), C.MPFR_RNDN)
	C.mpfr_floor(&n.val[0], &n.val[0])
	C.mpfr_div_ui(&n.val[0], &n.val[0], C.ulong(scale.Uint()), C.MPFR_RNDN)

	return n
}

// Truncate the number to the specified decimal places.
func (n Numeric) Truncate(dp int) Numeric {
	if !n.init {
		n = New(0)
	}

	scale := New(10).Pow(dp)
	C.mpfr_mul_ui(&n.val[0], &n.val[0], C.ulong(scale.Uint()), C.MPFR_RNDN)
	C.mpfr_trunc(&n.val[0], &n.val[0])
	C.mpfr_div_ui(&n.val[0], &n.val[0], C.ulong(scale.Uint()), C.MPFR_RNDN)

	return n
}

// endregion
