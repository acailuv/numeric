package numeric

import "database/sql/driver"

// Scan implements the sql.Scanner interface for database deserialization.
func (n *Numeric) Scan(value any) error {
	num, err := NewWithError(value)
	if err != nil {
		return err
	}

	*n = num

	return nil
}

// Value implements the driver.Valuer interface for database serialization.
func (n Numeric) Value() (driver.Value, error) {
	return n.String(), nil
}
