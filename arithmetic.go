package numeric

/*
#cgo LDFLAGS: -lmpfr
#include <mpfr.h>
*/
import "C"

// region Public

// Add a number and return the result. This will not modify the original number.
func (n Numeric) Add(x any) Numeric {
	result := New(0)

	switch x := x.(type) {
	case Numeric:
		if !x.init {
			return n
		}

		C.mpfr_add(&result.val[0], &n.val[0], &x.val[0], C.MPFR_RNDN)

	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		_x := New(x)
		C.mpfr_add(&result.val[0], &n.val[0], &_x.val[0], C.MPFR_RNDN)
	}

	return result
}

// Subtract a number and return the result. This will not modify the original number.
func (n Numeric) Subtract(x any) Numeric {
	result := New(0)

	switch x := x.(type) {
	case Numeric:
		if !x.init {
			return n
		}

		C.mpfr_sub(&result.val[0], &n.val[0], &x.val[0], C.MPFR_RNDN)

	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		_x := New(x)
		C.mpfr_sub(&result.val[0], &n.val[0], &_x.val[0], C.MPFR_RNDN)
	}

	return result
}

// Multiply a number and return the result. This will not modify the original number.
func (n Numeric) Multiply(x any) Numeric {
	result := New(0)

	switch x := x.(type) {
	case Numeric:
		if !x.init {
			return n
		}

		C.mpfr_mul(&result.val[0], &n.val[0], &x.val[0], C.MPFR_RNDN)

	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		_x := New(x)
		C.mpfr_mul(&result.val[0], &n.val[0], &_x.val[0], C.MPFR_RNDN)
	}

	return result
}

// Divide a number and return the result. This will not modify the original number.
func (n Numeric) Divide(x any) Numeric {
	result := New(0)

	switch x := x.(type) {
	case Numeric:
		if !x.init {
			return n
		}

		C.mpfr_div(&result.val[0], &n.val[0], &x.val[0], C.MPFR_RNDN)

	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		_x := New(x)
		C.mpfr_div(&result.val[0], &n.val[0], &_x.val[0], C.MPFR_RNDN)
	}

	return result
}

// Exponent the current number to the power of `x` and return the result. This will not modify the original number.
func (n Numeric) Pow(power any) Numeric {
	result := New(0)

	switch x := power.(type) {
	case Numeric:
		if !x.init {
			return n
		}

		C.mpfr_pow_ui(&result.val[0], &n.val[0], C.ulong(x.Uint()), C.MPFR_RNDN)

	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		_x := New(x)
		C.mpfr_pow_ui(&result.val[0], &n.val[0], C.ulong(_x.Uint()), C.MPFR_RNDN)
	}

	return result
}

// Makes the number negative. This will modify the original number.
func (n Numeric) Neg() Numeric {
	C.mpfr_neg(&n.val[0], &n.val[0], C.MPFR_RNDN)
	return n
}

// Makes the number positive. This will modify the original number.
func (n Numeric) Abs() Numeric {
	C.mpfr_abs(&n.val[0], &n.val[0], C.MPFR_RNDN)
	return n
}

// endregion
