package blog

import (
	"testing"
)

func TestDatastoreInvalidType(t *testing.T) {

	_, err := NewMySQL(DatastoreConfig{
		Type: DatastoreTypeMySQLMock,
	})

	if err != nil {
		t.Errorf("Unexpected error %s", err.Error())
	}

	// store.Create

}
