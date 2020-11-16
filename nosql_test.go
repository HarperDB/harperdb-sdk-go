package harperdb

import (
	"fmt"
	"testing"
)

type aRecord struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func createTestRecord() aRecord {
	return aRecord{
		ID:   randomID(),
		Name: "This is an arbitrary string",
	}
}

func TestInsertEmptySlice(t *testing.T) {
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

	// inserting an empty slice should not return an error
	if err := c.Insert(schema, table, []Record{}); err != nil {
		t.Fatal(err)
	}

	if err := c.DropSchema(schema); err != nil {
		t.Fatal(err)
	}
}

func TestInsertOne(t *testing.T) {
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

	record := createTestRecord()

	if err := c.Insert(schema, table, []Record{
		record,
	}); err != nil {
		t.Fatal(err)
	}

	// TODO Fetch and check

	if err := c.DropSchema(schema); err != nil {
		t.Fatal(err)
	}
}

func TestInsertThousand(t *testing.T) {
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

	data := []Record{}

	for i := 0; i < 1000; i++ {
		data = append(data, createTestRecord())
	}

	// TODO Count records in table to verify

	if err := c.Insert(schema, table, data); err != nil {
		t.Fatal(err)
	}

	if err := c.DropSchema(schema); err != nil {
		t.Fatal(err)
	}
}

func TestUpdateOne(t *testing.T) {
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

	record := aRecord{
		ID:   randomID(),
		Name: "Harper, the dog",
	}

	if err := c.Insert(schema, table, []Record{
		record,
	}); err != nil {
		t.Fatal(err)
	}

	record.Name = "Harper, the wolf"

	if err := c.Update(schema, table, []Record{
		record,
	}); err != nil {
		t.Fatal(err)
	}

	// TODO Fetch record and check if update was done

	if err := c.DropSchema(schema); err != nil {
		t.Fatal(err)
	}
}

func TestUpdateEmptySlice(t *testing.T) {
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

	// inserting an empty slice should not return an error
	if err := c.Update(schema, table, []Record{}); err != nil {
		t.Fatal(err)
	}

	if err := c.DropSchema(schema); err != nil {
		t.Fatal(err)
	}
}

func TestDeleteOne(t *testing.T) {
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

	record := createTestRecord()

	if err := c.Insert(schema, table, []Record{
		record,
	}); err != nil {
		t.Fatal(err)
	}

	hashes := []string{record.ID}

	if err := c.Delete(schema, table, hashes); err != nil {
		t.Fatal(err)
	}

	if err := c.DropSchema(schema); err != nil {
		t.Fatal(err)
	}
}

func TestSearchByHash(t *testing.T) {
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

	record := createTestRecord()

	if err := c.Insert(schema, table, []Record{
		record,
	}); err != nil {
		t.Fatal(err)
	}

	wait()

	lookupList := []string{record.ID}

	found := []aRecord{}
	err := c.SearchByHash(schema, table, &found, lookupList, AllAttributes)
	if err != nil {
		t.Fatal(err)
	}
	if num := len(found); num != 1 {
		t.Fatal(fmt.Errorf("wanted 1, got %d", num))
	}
	if !(found[0].ID == record.ID && found[0].Name == record.Name) {
		t.Fatal(fmt.Errorf("record data is not the same"))
	}

	if err := c.DropSchema(schema); err != nil {
		t.Fatal(err)
	}

}

func TestSearchByValue(t *testing.T) {
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

	record := createTestRecord()

	if err := c.Insert(schema, table, []Record{
		record,
	}); err != nil {
		t.Fatal(err)
	}

	wait()

	found := []aRecord{}
	err := c.SearchByValue(schema, table, &found, "name", record.Name, AllAttributes)
	if err != nil {
		t.Fatal(err)
	}
	if num := len(found); num != 1 {
		t.Fatal(fmt.Errorf("wanted 1, got %d", num))
	}
	if !(found[0].ID == record.ID && found[0].Name == record.Name) {
		t.Fatal(fmt.Errorf("record data is not the same"))
	}

	if err := c.DropSchema(schema); err != nil {
		t.Fatal(err)
	}

}
