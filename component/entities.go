package component

import (
	"fmt"
)

type Record struct {
	ID string
}

// For the easy printing of the record
func (r Record) String() string {
	return fmt.Sprintf("Record (%v)", r.ID)
}
