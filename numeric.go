package numeric

/*
#cgo LDFLAGS: -lmpfr
#include <mpfr.h>
*/
import "C"
import (
	"errors"
	"fmt"
	"regexp"
)

// region Global Variables
var (
	PrecisionBits       uint64 = 53 // Default MPFR precision
	StringDecimalPlaces uint64 = 10 // Default decimal places for string conversion
)

// endregion

type Numeric struct {
	init bool
	val  C.mpfr_t
}

// region Public

// Creates a new numeric value.
// The type of x has to be int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64 or string.
func New(x any) Numeric {
	switch x := x.(type) {
	case Numeric:
		return x
	case int:
		return newInt(int64(x))
	case int8:
		return newInt(int64(x))
	case int16:
		return newInt(int64(x))
	case int32:
		return newInt(int64(x))
	case int64:
		return newInt(x)
	case uint:
		return newUint(uint64(x))
	case uint8:
		return newUint(uint64(x))
	case uint16:
		return newUint(uint64(x))
	case uint32:
		return newUint(uint64(x))
	case uint64:
		return newUint(x)
	case float32:
		return newFloat(float64(x))
	case float64:
		return newFloat(x)
	case string:
		return newString(x)
	default:
		panic(fmt.Sprintf("numeric: Invalid type. Type has to be int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64 or string. Got: %T", x))
	}
}

// Creates a new numeric value, with error handling.
// The type of x has to be int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64 or string.
func NewWithError(x any) (Numeric, error) {
	switch x := x.(type) {
	case Numeric:
		return x, nil
	case int:
		return newIntWithError(int64(x))
	case int8:
		return newIntWithError(int64(x))
	case int16:
		return newIntWithError(int64(x))
	case int32:
		return newIntWithError(int64(x))
	case int64:
		return newIntWithError(x)
	case uint:
		return newUintWithError(uint64(x))
	case uint8:
		return newUintWithError(uint64(x))
	case uint16:
		return newUintWithError(uint64(x))
	case uint32:
		return newUintWithError(uint64(x))
	case uint64:
		return newUintWithError(x)
	case float32:
		return newFloatWithError(float64(x))
	case float64:
		return newFloatWithError(x)
	case string:
		return newStringWithError(x)
	default:
		return Numeric{}, fmt.Errorf("numeric: Invalid type. Type has to be int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64 or string. Got: %T", x)
	}
}

// Sets the default precision bits when doing arithmetic operations.
// The upper limit is "virtually" unlimited. However, more precision bits will make your app use more RAM.
// Cranking this number up to a large number will also make arithmetic operations slower. So bear this in mind.
// Default value is 53.
func SetPrecisionBits(bits uint64) {
	C.mpfr_set_default_prec(C.long(bits))
	PrecisionBits = bits
}

// Sets the decimal places shown when .String() is called.
// The upper limit is "virtually" unlimited. However, more decimal places will make your number inaccurate.
// Example: 1.23 will be represented as 1.22999999... if you set a high decimal places, with not enough precision bits.
// You can overcome this by setting a higher precision bits, but bear in mind the consequences.
// Default value is 10.
func SetStringDecimalPlaces(dp uint64) {
	StringDecimalPlaces = dp
}

// Clears the memory for the numeric value.
func (n Numeric) Destroy() {
	if n.init {
		C.mpfr_clear(&n.val[0])
		n.init = false
	}
}

// endregion

// region Private
func newInt(x int64) Numeric {
	return newString(fmt.Sprintf("%d", x))
}

func newUint(x uint64) Numeric {
	return newString(fmt.Sprintf("%d", x))
}

func newFloat(x float64) Numeric {
	return newString(fmt.Sprintf("%f", x))
}

func newString(x string) Numeric {
	// Validate numeric string
	if !regexp.MustCompile(`^\d+(\.\d+)?$`).MatchString(x) {
		panic("numeric: Invalid string. String has to be numerical")
	}

	num := Numeric{}
	num.init = true

	if ok := C.mpfr_init_set_str(&num.val[0], C.CString(x), C.int(10), C.MPFR_RNDN); ok != 0 {
		panic("numeric: Failed to initialize mpfr_t")
	}

	return num
}

func newIntWithError(x int64) (Numeric, error) {
	return newStringWithError(fmt.Sprintf("%d", x))
}

func newUintWithError(x uint64) (Numeric, error) {
	return newStringWithError(fmt.Sprintf("%d", x))
}

func newFloatWithError(x float64) (Numeric, error) {
	return newStringWithError(fmt.Sprintf("%f", x))
}

func newStringWithError(x string) (Numeric, error) {
	// Validate numeric string
	if !regexp.MustCompile(`^\d+(\.\d+)?$`).MatchString(x) {
		return Numeric{}, errors.New("numeric: Invalid string. String has to be numerical")
	}

	num := Numeric{}
	num.init = true

	if ok := C.mpfr_init_set_str(&num.val[0], C.CString(x), C.int(10), C.MPFR_RNDN); ok != 0 {
		return Numeric{}, errors.New("numeric: Failed to initialize mpfr_t")
	}

	return num, nil
}

// endregion
