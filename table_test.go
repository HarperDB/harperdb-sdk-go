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
	if e, ok := err.(*OperationError); ok && e.IsAlreadyExistsError() {
		return
	} else {
		t.Log(e)
		t.Fatalf("should have raised AlreadyExistsError")
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
	if e, ok := err.(*OperationError); ok && e.IsDoesNotExistError() {
		return
	} else {
		t.Log(e)
		t.Fatalf("should have raised DoesNotExistError")
	}
}

func TestDescribeTable(t *testing.T) {
	schema := randomID()
	table := randomID()

	if err := c.CreateSchema(schema); err != nil {
		t.Fatal(err)
	}

	wait()

	err := c.CreateTable(schema, table, "id")
	if err != nil {
		t.Fatal(err)
	}

	meta, err := c.DescribeTable(schema, table)
	if err != nil {
		t.Fatal(err)
	}

	if !(meta.Schema == schema &&
		meta.Name == table &&
		meta.HashAttribute == "id" &&
		meta.RecordCount == 0) {
		t.Log(meta)
		t.Fail()
	}
}
