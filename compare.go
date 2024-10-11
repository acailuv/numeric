package numeric

/*
#cgo LDFLAGS: -lmpfr
#include <mpfr.h>
*/
import "C"

// region Public

// GreaterThan returns true if the number is greater than `x`.
func (n Numeric) GreaterThan(x any) bool {
	switch x := x.(type) {
	case Numeric:
		if !x.init {
			return false
		}

		return C.mpfr_greater_p(&n.val[0], &x.val[0]) != 0
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		_x := New(x)
		return C.mpfr_greater_p(&n.val[0], &_x.val[0]) != 0
	}

	return false
}

// GreaterThanOrEqual returns true if the number is greater than or equal to `x`.
func (n Numeric) GreaterThanOrEqual(x any) bool {
	switch x := x.(type) {
	case Numeric:
		if !x.init {
			return false
		}

		return C.mpfr_greaterequal_p(&n.val[0], &x.val[0]) != 0
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		_x := New(x)
		return C.mpfr_greaterequal_p(&n.val[0], &_x.val[0]) != 0
	}

	return false
}

// LessThan returns true if the number is less than `x`.
func (n Numeric) LessThan(x any) bool {
	switch x := x.(type) {
	case Numeric:
		if !x.init {
			return false
		}

		return C.mpfr_less_p(&n.val[0], &x.val[0]) != 0
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		_x := New(x)
		return C.mpfr_less_p(&n.val[0], &_x.val[0]) != 0
	}

	return false
}

// LessThanOrEqual returns true if the number is less than or equal to `x`.
func (n Numeric) LessThanOrEqual(x any) bool {
	switch x := x.(type) {
	case Numeric:
		if !x.init {
			return false
		}

		return C.mpfr_lessequal_p(&n.val[0], &x.val[0]) != 0
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		_x := New(x)
		return C.mpfr_lessequal_p(&n.val[0], &_x.val[0]) != 0
	}

	return false
}

// Equal returns true if the number is equal to `x`.
func (n Numeric) Equal(x any) bool {
	switch x := x.(type) {
	case Numeric:
		if !x.init {
			return false
		}

		return C.mpfr_equal_p(&n.val[0], &x.val[0]) != 0
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		_x := New(x)
		return C.mpfr_equal_p(&n.val[0], &_x.val[0]) != 0
	}

	return false
}

// endregion
