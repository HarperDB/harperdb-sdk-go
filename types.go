package harperdb

type Attribute string

type AttributeList interface{}

func FromStringSlice(ss []string) AttributeList {
	return AttributeList(ss)
}

type Record interface{}

var AllAttributes = FromStringSlice([]string{"*"})
