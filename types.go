package harperdb

import (
	"time"
)

type Attribute string

type AttributeList interface{}

func FromStringSlice(ss []string) AttributeList {
	return AttributeList(ss)
}

// Record type gives you convenient access to HarperDB's meta data fields
// which are automatically added to every record. Currently these are
// "__createdtime__" and "__updatedtime__".
//
// Meta data fields are read-only.
// Overwriting .CreatedTime and .UpdatedTime will have no effect.
type Record struct {
	CreatedTime harperTimestamp `json:"__createdtime__"`
	UpdatedTime harperTimestamp `json:"__updatedtime__"`
}

type harperTimestamp int64

func (t harperTimestamp) ToTime() time.Time {
	return time.Unix(int64(t)/1000, int64(t)&1000)
}

var AllAttributes = FromStringSlice([]string{"*"})
