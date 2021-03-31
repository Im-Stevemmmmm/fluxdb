package database_test

import (
	"reflect"
	"testing"

	"github.com/Im-Stevemmmmm/fluxdb/database"
)

func TestGetSet(t *testing.T) {
	db := database.NewDB(nil)

	data := map[string]interface{}{
		"Username": "Stevemmmmm",
		"Email":    "example@domain.com",
		"Password": "1234",
	}

	if err := db.Set("key", data); err != nil {
		t.Fatalf("error while setting value: %s", err)
	}

	res, err := db.Get("key")
	if err != nil {
		t.Fatalf("error while getting value: %s", err)
	}

	if !reflect.DeepEqual(res, data) {
		t.Fatalf("get value is not the expected value for the key; got %+v expected %+v", res, data)
	}
}
