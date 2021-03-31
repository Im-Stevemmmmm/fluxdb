package database_test

import (
	"testing"

	"github.com/Im-Stevemmmmm/fluxdb/database"
)

func BenchmarkReadWrite(b *testing.B) {
	db := database.NewDB(nil)

	data := map[string]string{
		"FirstName": "John",
		"LastName":  "Smith",
	}

	if err := db.Set("key", data); err != nil {
		b.Fatal(err)
	}

	_, _ = db.Get("key")
}
