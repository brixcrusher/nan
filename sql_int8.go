package nan

import "database/sql/driver"

func (n *NullInt8) Scan(value interface{}) error {
	if value == nil {
		n.Int8, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	return convertAssign(&n.Int8, value)
}

func (n NullInt8) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return int64(n.Int8), nil
}
