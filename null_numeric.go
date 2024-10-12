package numeric

import "fmt"

type NullNumeric struct {
	Valid   bool
	Numeric Numeric
}

// region Public

// Creates a new null numeric value.
// The type of x has to be int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64 or string.
func NewNull(x any) NullNumeric {
	switch x := x.(type) {
	case NullNumeric:
		return x
	case Numeric:
		return NullNumeric{true, x}
	case int:
		return NullNumeric{true, newInt(int64(x))}
	case int8:
		return NullNumeric{true, newInt(int64(x))}
	case int16:
		return NullNumeric{true, newInt(int64(x))}
	case int32:
		return NullNumeric{true, newInt(int64(x))}
	case int64:
		return NullNumeric{true, newInt(x)}
	case uint:
		return NullNumeric{true, newUint(uint64(x))}
	case uint8:
		return NullNumeric{true, newUint(uint64(x))}
	case uint16:
		return NullNumeric{true, newUint(uint64(x))}
	case uint32:
		return NullNumeric{true, newUint(uint64(x))}
	case uint64:
		return NullNumeric{true, newUint(x)}
	case float32:
		return NullNumeric{true, newFloat(float64(x))}
	case float64:
		return NullNumeric{true, newFloat(x)}
	case string:
		return NullNumeric{true, newString(x)}
	default:
		panic(fmt.Sprintf("numeric: Invalid type. Type has to be int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64 or string. Got: %T", x))
	}
}

// Creates a new null numeric value, with error handling.
// The type of x has to be int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64 or string.
func NewNullWithError(x any) (NullNumeric, error) {
	switch x := x.(type) {
	case NullNumeric:
		return x, nil
	case Numeric:
		return NullNumeric{true, x}, nil
	case int:
		return newNullIntWithError(int64(x))
	case int8:
		return newNullIntWithError(int64(x))
	case int16:
		return newNullIntWithError(int64(x))
	case int32:
		return newNullIntWithError(int64(x))
	case int64:
		return newNullIntWithError(x)
	case uint:
		return newNullUintWithError(uint64(x))
	case uint8:
		return newNullUintWithError(uint64(x))
	case uint16:
		return newNullUintWithError(uint64(x))
	case uint32:
		return newNullUintWithError(uint64(x))
	case uint64:
		return newNullUintWithError(x)
	case float32:
		return newNullFloatWithError(float64(x))
	case float64:
		return newNullFloatWithError(x)
	case string:
		return newNullStringWithError(x)
	default:
		return NullNumeric{}, fmt.Errorf("numeric: Invalid type. Type has to be int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64 or string. Got: %T", x)
	}
}

// endregion

// region Private

func newNullIntWithError(x int64) (NullNumeric, error) {
	num, err := newIntWithError(x)
	if err != nil {
		return NullNumeric{}, err
	}

	return NullNumeric{true, num}, nil
}

func newNullUintWithError(x uint64) (NullNumeric, error) {
	num, err := newUintWithError(x)
	if err != nil {
		return NullNumeric{}, err
	}

	return NullNumeric{true, num}, nil
}

func newNullFloatWithError(x float64) (NullNumeric, error) {
	num, err := newFloatWithError(x)
	if err != nil {
		return NullNumeric{}, err
	}

	return NullNumeric{true, num}, nil
}

func newNullStringWithError(x string) (NullNumeric, error) {
	num, err := newStringWithError(x)
	if err != nil {
		return NullNumeric{}, err
	}

	return NullNumeric{true, num}, nil
}

// endregion
