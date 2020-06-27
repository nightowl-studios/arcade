// package nullable is a package that provides a wrapper to the sql.Int64 etc
// to ensure that any JSON it receives can be null
// it also ensures future proofing just in case we need to write to a
// sql database
package nullable

import (
	"database/sql"
	"encoding/json"
)

type NullFloat64 struct {
	sql.NullFloat64
}

type NullString struct {
	sql.NullString
}

type NullBool struct {
	sql.NullBool
}

type NullInt64 struct {
	sql.NullInt64
}

func (n NullString) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *string
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		n.Valid = true
		n.String = *x
	} else {
		n.Valid = false
		n.String = "."
	}
	return nil
}
