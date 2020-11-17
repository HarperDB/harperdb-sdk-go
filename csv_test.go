package harperdb

import (
	"fmt"
	"os"
	"testing"
)

func TestCSVURLLoad(t *testing.T) {
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

	fd, err := os.Open("testdata/companies100.csv")
	if err != nil {
		t.Fatal("unable to open test data file")
	}
	defer fd.Close()

	jobID, err := c.CSVDataLoad(schema, table, false, fd)
	t.Log(fmt.Sprintf("Job ID: %s", jobID))
	if err != nil {
		t.Fatal(err)
	}

	for {
		job, err := c.GetJob(jobID)
		if err != nil {
			t.Fatal(err)
			break
		}
		if job.Status == JobStatusCompleted {
			break
		} else {
			t.Log(fmt.Sprintf("Job status was: %s", job.Status))
		}
		wait()
	}

	// The table should now have 100 rows
	meta, err := c.DescribeTable(schema, table)
	if err != nil {
		t.Fatal(err)
	}

	if meta.RecordCount != 100 {
		t.Log(meta)
		t.Fatal(fmt.Errorf("expected table to have 100 rows after csv load"))
	}

	if err := c.DropSchema(schema); err != nil {
		t.Fatal(err)
	}
}
