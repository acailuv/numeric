package numeric

import (
	"database/sql/driver"
	"fmt"
)

// region Numeric

// Scan implements the sql.Scanner interface.
func (n *Numeric) Scan(value any) error {
	switch v := value.(type) {
	case float32, float64, int64, uint64:
		*n = New(v)
	default:
		// default is trying to interpret value stored as string
		str, err := unquote(v)
		if err != nil {
			return err
		}

		*n, err = NewWithError(str)
		if err != nil {
			return err
		}
	}

	return nil
}

// Value implements the driver.Valuer interface.
func (n Numeric) Value() (driver.Value, error) {
	return n.String(), nil
}

// Array form of Numeric, used for scanning and storing arrays of Numeric in postgresql.
type NumericArray []Numeric

// Scan implements the sql.Scanner interface.
func (a *NumericArray) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		return a.scanBytes(src)
	case string:
		return a.scanBytes([]byte(src))
	case nil:
		*a = nil
		return nil
	}

	return fmt.Errorf("numeric: cannot convert %T to StringArray", src)
}

func (a *NumericArray) scanBytes(src []byte) error {
	elems, err := scanLinearArray(src, []byte{','}, "StringArray")
	if err != nil {
		return err
	}
	if *a != nil && len(elems) == 0 {
		*a = (*a)[:0]
	} else {
		b := make(NumericArray, len(elems))
		for i, v := range elems {
			num, err := NewWithError(string(v))
			if err != nil {
				return fmt.Errorf("numeric: parsing array element index %d: cannot convert to numeric", i)
			}

			b[i] = num
		}
		*a = b
	}
	return nil
}

// Value implements the driver.Valuer interface.
func (a NumericArray) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}

	if n := len(a); n > 0 {
		// There will be at least two curly brackets, 2*N bytes of quotes,
		// and N-1 bytes of delimiters.
		b := make([]byte, 1, 1+3*n)
		b[0] = '{'

		b = appendArrayQuotedBytes(b, []byte(a[0].String()))
		for i := 1; i < n; i++ {
			b = append(b, ',')
			b = appendArrayQuotedBytes(b, []byte(a[i].String()))
		}

		return string(append(b, '}')), nil
	}

	return "{}", nil
}

// endregion

// region NullNumeric

// Scan implements the sql.Scanner interface.
func (n *NullNumeric) Scan(value any) error {
	if value == nil {
		*n = NullNumeric{}
		return nil
	}

	switch v := value.(type) {
	case float32, float64, int64, uint64:
		*n = NewNull(v)
	default:
		// default is trying to interpret value stored as string
		str, err := unquote(v)
		if err != nil {
			return err
		}

		*n, err = NewNullWithError(str)
		if err != nil {
			return err
		}
	}

	return nil
}

// Value implements the driver.Valuer interface.
func (n NullNumeric) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}

	return n.Numeric.String(), nil
}

// endregion
