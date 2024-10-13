package numeric

// Implements go-redis encoding interface.
func (n Numeric) MarshalBinary() ([]byte, error) {
	return []byte(n.String()), nil
}

// Implements go-redis decoding interface.
func (n *Numeric) UnmarshalBinary(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	str := string(data)
	num, err := NewWithError(str)
	if err != nil {
		return err
	}
	*n = num

	return nil
}

// Implements go-redis encoding interface.
func (n NullNumeric) MarshalBinary() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return []byte(n.Numeric.String()), nil
}

// Implements go-redis decoding interface.
func (n *NullNumeric) UnmarshalBinary(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	num, err := NewNullWithError(string(data))
	if err != nil {
		return err
	}
	*n = num

	return nil
}
