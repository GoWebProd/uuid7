package uuid7

import (
	"database/sql/driver"
	"fmt"
)

func (uuid *UUID) Scan(src interface{}) error {
	switch src := src.(type) {
	case nil:
		return nil

	case string:
		if src == "" {
			return nil
		}

		u, err := Parse(src)
		if err != nil {
			return fmt.Errorf("scan uuid error: %w", err)
		}

		*uuid = u
	case []byte:
		switch len(src) {
		case 0:
			return nil
		case 16:
			copy((*uuid)[:], src)
		case 36:
			return uuid.Scan(string(src))
		default:
			return fmt.Errorf("unable to scan type []byte with length %d into UUID", len(src))
		}
	default:
		return fmt.Errorf("unable to scan type %T into UUID", src)
	}

	return nil
}

func (uuid UUID) Value() (driver.Value, error) {
	return uuid.String(), nil
}
