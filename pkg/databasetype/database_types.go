package databasetype

// DatabaseType defines the type of the database
type DatabaseType int

const (
	// Sqlite 0, Sqlite
	Sqlite DatabaseType = iota
	// Postgres 1, postgres sql
	Postgres
	// MySQL 2, mysql
	MySQL
)

// ToInt converts database type to integer
func (databaseType DatabaseType) ToInt() int {
	switch databaseType {
	case Sqlite:
		return 0
	case Postgres:
		return 1
	case MySQL:
		return 2
	default:
		return 0
	}
}

// ToDatabaseType converts int to DatabaseType
func ToDatabaseType(data int) DatabaseType {
	switch data {
	case 0:
		return Sqlite
	case 1:
		return Postgres
	case 2:
		return MySQL
	default:
		return Sqlite
	}
}
