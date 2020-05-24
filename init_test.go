package blog

func init() {
	SetDatastore(DatastoreConfig{Type: DatastoreTypeMySQLMock})
}
