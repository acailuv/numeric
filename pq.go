package numeric

import "database/sql/driver"

// Scan implements the sql.Scanner interface for database deserialization.
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

// Value implements the driver.Valuer interface for database serialization.
func (n Numeric) Value() (driver.Value, error) {
	return n.String(), nil
}
