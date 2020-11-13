package harperdb

import (
	"testing"
)

func TestCreateSchema(t *testing.T) {
	schema := randomID()
	if err := c.CreateSchema(schema); err != nil {
		t.Fatal(err)
	}

	wait()

	if err := c.DropSchema(schema); err != nil {
		t.Fatal(err)
	}
}

func TestDuplicateSchema(t *testing.T) {
	schema := randomID()
	err := c.CreateSchema(schema)
	if err != nil {
		t.Fatal(err)
	}

	wait()

	err = c.CreateSchema(schema)
	if err == nil {
		t.Fatal("should have raised AlreadyExistsError")
	}
	if e, ok := err.(*AlreadyExistsError); !ok {
		t.Fatal(e)
	}

	if err := c.DropSchema(schema); err != nil {
		t.Fatal(err)
	}
}

func TestDropSchema(t *testing.T) {
	schema := randomID()
	err := c.CreateSchema(schema)
	if err != nil {
		t.Fatal(err)
	}

	wait()

	err = c.DropSchema(schema)
	if err != nil {
		t.Fatal(err)
	}

	wait()

	err = c.DropSchema(schema)
	if err == nil {
		t.Fatal("should have raised DoesNotExist error")
	}
	if e, ok := err.(*DoesNotExistsError); !ok {
		t.Fatal(e)
	}
}
