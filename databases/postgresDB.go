package databases

import (
	"fmt"
	"log"
	"sync"

	"github.com/natersland/case-codex-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type posgresDB struct {
	*gorm.DB
}

var (
	postgresDBInstance *posgresDB
	once               sync.Once
)

func NewPostgresDB(conf *config.Database) Database {
	once.Do(func() {
		// Connect to postgres
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s search_path=%s anon=%s url=%s jwtSecret=%s",
			conf.Host,
			conf.User,
			conf.Password,
			conf.DBName,
			conf.Port,
			conf.SSLMode,
			conf.Schema,
			conf.Anon,
			conf.Url,
			conf.JwtSecret,
		)

		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		log.Printf("Connected to postgres database %s", conf.DBName)

		postgresDBInstance = &posgresDB{conn}
	},
	)

	return postgresDBInstance
}

func (db *posgresDB) Connect() *gorm.DB {
	return postgresDBInstance.DB
}
