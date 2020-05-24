// +build !unit_test

package service

import (
	"github.com/ory/viper"

	blog "github.com/stevepartridge/blogorama"
)

func setupDatastore() error {

	return blog.SetDatastore(blog.DatastoreConfig{
		Type:       blog.DatastoreTypeMaria,
		Host:       viper.GetString(blog.EnvBlogDbHost),
		Port:       viper.GetString(blog.EnvBlogDbPort),
		User:       viper.GetString(blog.EnvBlogDbUser),
		Pass:       viper.GetString(blog.EnvBlogDbPass),
		Name:       viper.GetString(blog.EnvBlogDbName),
		Migrations: viper.GetBool(blog.EnvBlogDbMigrations),
	})

}
