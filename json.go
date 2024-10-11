package numeric

import "fmt"

// UnmarshalJSON implements the json.Unmarshaler interface.
func (n *Numeric) UnmarshalJSON(bytes []byte) error {
	if string(bytes) == "null" {
		return nil
	}

	// If the amount is quoted, strip the quotes
	if len(bytes) > 2 && bytes[0] == '"' && bytes[len(bytes)-1] == '"' {
		bytes = bytes[1 : len(bytes)-1]
	}
	str := string(bytes)

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
