package harperdb

import (
	"testing"
)

func TestCreateTable(t *testing.T) {
	schema := randomID()
	table := randomID()

	if err := c.CreateSchema(schema); err != nil {
		t.Fatal(err)
	}

	wait()

	if err := c.CreateTable(schema, table, "id"); err != nil {
		t.Log(err)
		t.FailNow()
	}

	wait()

	if err := c.DropSchema(schema); err != nil {
		t.Fatal(err)
	}
}

func TestCreateDuplicateTable(t *testing.T) {
	schema := randomID()
	table := randomID()

	if err := c.CreateSchema(schema); err != nil {
		t.Fatal(err)
	}

	wait()

	if err := c.CreateTable(schema, table, "id"); err != nil {
		t.Log(err)
		t.FailNow()
	}

	wait()

	err := c.CreateTable(schema, table, "id")
	if e, ok := err.(*AlreadyExistsError); !ok {
		t.Fatal(e.err)
	}

	wait()

	if err := c.DropSchema(schema); err != nil {
		t.Fatal(err)
	}
}

func TestCreateTableInUnknownSchema(t *testing.T) {
	schema := randomID()
	table := randomID()

	err := c.CreateTable(schema, table, "id")
	if err == nil {
		t.Fatal("should have raised DoesNotExistError")
	}
	if e, ok := err.(*DoesNotExistsError); !ok {
		t.Fatal(e)
	}

}
