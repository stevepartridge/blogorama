package blog

import (
	"github.com/stevepartridge/blogorama/pkg/log"
)

const (
	// DatastoreTypeMaria MariaDB type
	DatastoreTypeMaria = "maria"

	// DatastoreTypeMySQL MySQL type
	DatastoreTypeMySQL = "mysql"

	// DatastoreTypeMaria MySQL/MariaDB mock type
	DatastoreTypeMySQLMock = "mysql_mock"
)

var (
	Store Datastore

	defaultLimit = 0
)

// Datastore interface
type Datastore interface {
	Migrations() error

	// Users
	CreateUser(user *User) error
	UpdateUser(user *User) error
	GetUserByID(id int) (User, error)
	GetUserByAPIKey(apiKey string) (User, error)
	GetAPIKeyForUserID(userID int) (string, error)
	GetUserIDByAPIKey(apiKey string) (int, error)
	DeleteUser(id, deletedBy int) error
	GetUsers(opts ...int) ([]User, ResultsInfo, error)
	GetUsersByIDs(ids ...int) ([]User, error)

	// Posts
	CreatePost(post *Post) error
	UpdatePost(post *Post) error
	GetPostByID(id int) (Post, error)
	DeletePost(id, deletedBy int) error
	GetPosts(opts ...int) ([]Post, ResultsInfo, error)
	GetPostsByIDs(ids ...int) ([]Post, error)
	GetPostsByCreatedByID(createdBy int, opts ...int) ([]Post, ResultsInfo, error)
}

// DatastoreConfig holds the specific connection info and type
type DatastoreConfig struct {
	Type       string
	Host       string
	Port       string
	User       string
	Pass       string
	Name       string
	SSLEnabled bool
	Migrations bool
}

// ResultsInfo is leveraged for pagination type info from a query
type ResultsInfo struct {
	Limit  int
	Offset int
	Total  int
}

// SetDatastore sets or overrides the current Store implementation
func SetDatastore(config DatastoreConfig) error {

	var err error

	switch config.Type {
	case DatastoreTypeMaria, DatastoreTypeMySQL, DatastoreTypeMySQLMock:
		Store, err = NewMySQL(config)
		if log.IfError(err) {
			return err
		}
	default:
		if config.Type == "" {
			config.Type = "unset"
		}
		return ErrReplacer(ErrDatastoreUnsupportedOrMissingType, config.Type)
	}

	return nil
}
