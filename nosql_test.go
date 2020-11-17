package harperdb

import (
	"fmt"
	"testing"
	"time"
)

type aRecord struct {
	Record
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
	resp, err := c.Insert(schema, table, []Record{})
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.InsertedHashes) != 0 {
		t.Fatal(fmt.Errorf("expected zero inserted hashes"))
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

	resp, err := c.Insert(schema, table, []interface{}{
		record,
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.InsertedHashes) != 1 {
		t.Fatal(fmt.Errorf("expected 1 inserted hashes"))
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

	data := []interface{}{}

	for i := 0; i < 1000; i++ {
		data = append(data, createTestRecord())
	}

	// TODO Count records in table to verify

	resp, err := c.Insert(schema, table, data)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.InsertedHashes) != len(data) {
		t.Fatal(fmt.Errorf("expected %d inserted hashes", len(data)))
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

	if _, err := c.Insert(schema, table, []interface{}{
		record,
	}); err != nil {
		t.Fatal(err)
	}

	record.Name = "Harper, the wolf"

	resp, err := c.Update(schema, table, []interface{}{
		record,
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.UpdatedHashes) != 1 {
		t.Fatal(fmt.Errorf("expected 1 updated hash"))
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
	resp, err := c.Update(schema, table, []Record{})
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.UpdatedHashes) != 0 {
		t.Fatal(fmt.Errorf("expected 0 updated hash"))
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

	if _, err := c.Insert(schema, table, []interface{}{
		record,
	}); err != nil {
		t.Fatal(err)
	}

	hashes := []string{record.ID}

	resp, err := c.Delete(schema, table, hashes)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.DeletedHashes) != 1 {
		t.Fatal(fmt.Errorf("expected 1 deleted hash"))
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

	if _, err := c.Insert(schema, table, []interface{}{
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

	if _, err := c.Insert(schema, table, []interface{}{
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
	if !found[0].CreatedTime.ToTime().Before(time.Now()) {
		t.Fatal(fmt.Errorf("record timestamp is too recent"))
	}

	if err := c.DropSchema(schema); err != nil {
		t.Fatal(err)
	}

}
