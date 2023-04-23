package modularHTTPGo

import (
	. "ModularHTTPGo/constants"
	. "ModularHTTPGo/utils"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func CreateConnection() (*sql.DB, error) {
	host := GetEnvString("DB_HOST", DefaultPostgresHost)
	port := GetEnvInt("DB_PORT", DefaultPostgresPort)
	user := GetEnvString("DB_USER", DefaultPostgresUser)
	pass := GetEnvString("DB_PASSWORD", DefaultPostgresPassword)
	dbName := GetEnvString("DB_NAME", DefaultPostgresDbname)

	psqlInfo := fmt.Sprintf(`
	host=%s 
	port=%d 
	user=%s 
	password=%s 
	dbname=%s 
	sslmode=disable`,
		host, port, user, pass, dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to ping the database: %v", err)
	}

	log.Println("Successfully Connected to the database")

	return db, nil
}
