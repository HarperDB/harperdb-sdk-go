package harperdb

import (
	"fmt"
	"testing"
)

func TestSQL(t *testing.T) {
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

	var rows []aRecord
	if err := c.SQLSelect(&rows, "select * from `%s`.`%s`", schema, table); err != nil {
		t.Fatal(err)
	}
	if found := len(rows); !(found > 0 && found <= 1) {
		t.Fatal(fmt.Errorf("expected one row from select"))
	}
	if rows[0].Name != "Harper, the dog" {
		t.Fatal(fmt.Errorf("returned row has different data"))
	}

	newName := "Harper, the wolf"
	resp, err := c.SQLExec("update `%s`.`%s` set name = '%s' where id = '%s'", schema, table, newName, rows[0].ID)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.UpdatedHashes) != 1 {
		t.Log(fmt.Errorf("expected one updated row"))
		t.Fail()
	}

	// TODO Did it really update the name?

	resp, err = c.SQLExec("DELETE FROM `%s`.`%s` where id = '%s'", schema, table, rows[0].ID)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.DeletedHashes) != 1 {
		t.Fatal(fmt.Errorf("expected 1 deleted hash"))
	}

	// TODO Fetch record and check if delete was done

	if err := c.DropSchema(schema); err != nil {
		t.Fatal(err)
	}
}
