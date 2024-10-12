package numeric

import "fmt"

// UnmarshalJSON implements the json.Unmarshaler interface.
func (n *Numeric) UnmarshalJSON(bytes []byte) error {
	if string(bytes) == "null" {
		return nil
	}

	str, err := unquote(bytes)
	if err != nil {
		return err
	}

	num, err := NewWithError(str)
	if err != nil {
		return fmt.Errorf("numeric: Error decoding string '%s': %s", str, err)
	}
	*n = num

	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (n Numeric) MarshalJSON() ([]byte, error) {
	str := "\"" + n.String() + "\""
	return []byte(str), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (n *NullNumeric) UnmarshalJSON(bytes []byte) error {
	if string(bytes) == "null" {
		return nil
	}

	str, err := unquote(bytes)
	if err != nil {
		return err
	}

	num, err := NewNullWithError(str)
	if err != nil {
		return fmt.Errorf("numeric: Error decoding string '%s': %s", str, err)
	}
	*n = num

	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (n NullNumeric) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	str := "\"" + n.Numeric.String() + "\""
	return []byte(str), nil
}
