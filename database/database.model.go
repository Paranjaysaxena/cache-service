package database

type Database interface {
	Set(key string, value []byte) ([]byte, error)
	Get(key string) ([]byte, error)
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
