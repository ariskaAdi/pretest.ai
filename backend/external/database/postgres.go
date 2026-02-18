package database

import (
	"ariskaAdi-pretest-ai/internal/config"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectPostgres(cfg config.DBConfig) (db *sqlx.DB, err error ){
	dsn := fmt.Sprintf("host=%s port=%s  user=%s  dbname=%s password=%s sslmode=disable ",
	cfg.Host,
	cfg.Port,
	cfg.User,
	cfg.Name,
	cfg.Password,
)
	db, err = sqlx.Open("postgres", dsn)
	if err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		return
	}

	db.SetMaxIdleConns(int(cfg.ConnectionPool.MaxIdle))
	db.SetMaxOpenConns(int(cfg.ConnectionPool.MaxOpen))
	db.SetConnMaxIdleTime(time.Duration(cfg.ConnectionPool.MaxIdleTime) * time.Minute) 
	db.SetConnMaxLifetime(time.Duration(cfg.ConnectionPool.MaxLifetime) * time.Minute)

	

	return
}

// migrate init
// $ migrate create -ext sql -dir db/migrations init

// MIGRATION DEV
// migrate -database "postgres://postgres:mysecretpassword@localhost:5432/ewallet?sslmode=disable" -path db/migrations up

// VERIFY MIGRATION
// migrate -path db/migration -database "postgres://postgres:mysecretpassword@localhost:5432/ewallet?sslmode=disable" version

// CHECK PG DOCKER
// docker exec -it container-postgres psql -U postgres -d ewallet
// \dt


// <!-- $ docker run -d \

// > -p 5432:5432 \
// > --name todo-postgres \
// > -e POSTGRES_USER=postgres \
// > -e POSTGRES_PASSWORD=secret \
// > -e POSTGRES_DB=pretest \
// > -v $(pwd)/backend/external/database:/var/lib/postgresql/data \
// > postgres -->