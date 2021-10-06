package info

import "gitlab.com/notula/go-tools/dbms"

func GetDatabaseProperties() dbms.Properties {
	return dbms.Properties{
		Username:     "postgres",
		Password:     "docker",
		Name:         "db_notula",
		Host:         "10.226.174.200",
		Port:         "5433",
		MaxIdleConns: 10,
	}
}
