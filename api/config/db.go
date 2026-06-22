package config

import (
	"context"
	"fmt"
	"os"
	"github.com/jackc/pgx/v5"
)


func Connection()(*pgx.Conn , error){
	connStr := os.Getenv("DATABASE_URL")
	conn , err := pgx.Connect(context.Background() , connStr)
	if err != nil{
		return nil, fmt.Errorf("erro connecting to database: %v" , err)
	}

	return conn , nil
}