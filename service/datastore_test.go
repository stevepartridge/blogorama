package service

import (
	blog "github.com/stevepartridge/blogorama"
)

func setupDatastore() error {

	return blog.SetDatastore(blog.DatastoreConfig{
		Type: blog.DatastoreTypeMySQLMock,
	})

}
