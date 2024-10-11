package numeric

/*
#cgo LDFLAGS: -lmpfr
#include <mpfr.h>
#include <stdlib.h>

extern void _panic(char* msg);

static char* _str(mpfr_t num, unsigned long decimal_digits) {
	// Determine the size needed for the buffer
	int size_needed = mpfr_snprintf(NULL, 0, "%.*Rf", decimal_digits, num);
	if (size_needed < 0) {
		_panic("numeric: Error determining buffer size");
	}

	// Allocate memory for the string
	char* str = malloc(size_needed + 1);  // +1 for the null terminator
	if (str == NULL) {
		_panic("numeric: Unable to allocate memory");
	}

	// Format the mpfr_t value as a string into the allocated buffer
	mpfr_snprintf(str, size_needed + 1, "%.*Rf", decimal_digits, num);

	return str;
}
*/
import "C"

// region Public

// Returns numeric as a string.
// The default number of decimal places is 10.
// You can modify the default number of decimal places by using SetStringDecimalPlaces function.
func (n Numeric) String() string {
	if !n.init {
		n = New(0)
	}

	out := C._str(&n.val[0], C.ulong(StringDecimalPlaces))

	return C.GoString(out)
}

// Returns numeric as a string with a specified number of decimal places.
func (n Numeric) StringDecimalPlaces(dp uint64) string {
	if !n.init {
		n = New(0)
	}

	out := C._str(&n.val[0], C.ulong(dp))

	return C.GoString(out)
}

// Returns numeric as int
func (n Numeric) Int() int {
	return int(n.getInt())
}

// Returns numeric as int8
func (n Numeric) Int8() int8 {
	return int8(n.getInt())
}

// Returns numeric as int16
func (n Numeric) Int16() int16 {
	return int16(n.getInt())
}

// Returns numeric as int32
func (n Numeric) Int32() int32 {
	return int32(n.getInt())
}

// Returns numeric as int64
func (n Numeric) Int64() int64 {
	return n.getInt()
}

// Returns numeric as uint
func (n Numeric) Uint() uint {
	return uint(n.getUInt())
}

// Returns numeric as uint8
func (n Numeric) Uint8() uint8 {
	return uint8(n.getUInt())
}

// Returns numeric as uint16
func (n Numeric) Uint16() uint16 {
	return uint16(n.getUInt())
}

// Returns numeric as uint32
func (n Numeric) Uint32() uint32 {
	return uint32(n.getUInt())
}

// Returns numeric as uint64
func (n Numeric) Uint64() uint64 {
	return n.getUInt()
}

// Returns numeric as float32
func (n Numeric) Float32() float32 {
	return float32(n.getFloat())
}

// Returns numeric as float64
func (n Numeric) Float64() float64 {
	return n.getFloat()
}

// endregion

// region Private
func (n Numeric) getInt() int64 {
	if !n.init {
		n = New(0)
	}

	return int64(C.mpfr_get_si(&n.val[0], C.MPFR_RNDN))
}

func (n Numeric) getUInt() uint64 {
	if !n.init {
		n = New(0)
	}

	return uint64(C.mpfr_get_ui(&n.val[0], C.MPFR_RNDN))
}

func (n Numeric) getFloat() float64 {
	if !n.init {
		n = New(0)
	}

	return float64(C.mpfr_get_d(&n.val[0], C.MPFR_RNDN))
}

//export _panic
func _panic(msg *C.char) {
	panic(C.GoString(msg))
}

// endregion
