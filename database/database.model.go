package database

type Database interface {
	Set(key string, user User) (User, error)
	Get(key string) (User, error)
}

// Factory looks up according to the databaseName the database implementation
func Factory(databaseName string) (Database, error) {
	switch databaseName {
	case "redis":
		return createRedisDatabase()
	default:
		return nil, &NotImplementedDatabaseError{databaseName}
	}
}
