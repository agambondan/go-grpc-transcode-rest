package repo

import (
	"fmt"
	"github.com/agambondan/web-go-blog-grpc-rest/app/config"
	"github.com/agambondan/web-go-blog-grpc-rest/app/migrations"
	"github.com/morkid/gocache"
	cache_redis "github.com/morkid/gocache-redis/v8"
	"github.com/morkid/paginate"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"strconv"
	"time"
)

var PG *paginate.Pagination

type Repositories struct {
	//Role     RoleRepository
	//Category CategoryRepository
	//Article  ArticleRepository
	//Firebase FirebaseRepository
	User UserRepository
	db   *gorm.DB
}

func NewRepositories() (*Repositories, error) {
	logLevel := logger.Info

	switch os.Getenv("ENVIRONMENT") {
	case "development":
		logLevel = logger.Error
	case "staging":
		logLevel = logger.Error
	case "production":
		logLevel = logger.Silent
	}
	gormConfig := gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   os.Getenv("DB_TABLE_PREFIX"),
			SingularTable: true,
		},
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", config.Config.DBHost, config.Config.DBPort, config.Config.DBUser, config.Config.DBPassword, config.Config.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gormConfig)
	if err != nil {
		dsn = os.Getenv("URI")
		db, err = gorm.Open(postgres.Open(dsn), &gormConfig)
		if err != nil {
			fmt.Printf("Cannot connect to %s database url %s", config.Config.DBDriver, dsn)
			log.Println("\nThis is the error:", err)
			panic(err)
		}
	}

	var cache *gocache.AdapterInterface
	cacheSeconds := 1
	cacheSeconds, _ = strconv.Atoi(os.Getenv("CACHE_TTL_SECONDS"))

	if nil != REDIS && cacheSeconds > 0 {
		cache = cache_redis.NewRedisCache(cache_redis.RedisCacheConfig{
			Client:    REDIS,
			ExpiresIn: time.Duration(cacheSeconds) * time.Second,
		})
	}

	PG = paginate.New(&paginate.Config{
		CacheAdapter:         cache,
		FieldSelectorEnabled: true,
	})

	fmt.Printf("We are connected to the %s database with url %s\n", config.Config.DBDriver, dsn)
	return &Repositories{
		User: NewUserRepository(db),
		db:   db,
	}, nil
}

// Close closes the  database connection
func (s *Repositories) Close() error {
	db, _ := s.db.DB()
	return db.Close()
}

// Migrations to migrate any struct to table/model in database
func (s *Repositories) Migrations() error {
	err := s.db.AutoMigrate(migrations.ModelMigrations...)
	if err != nil {
		return err
	}
	err = s.db.Migrator().DropTable("schema_migration")
	if err != nil {
		return err
	}
	return nil
}

// Seeder to migrate any struct with data to insert to table
func (s *Repositories) Seeder() error {
	for i := range migrations.DataSeeds {
		tx := s.db.Begin()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

		if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(migrations.DataSeeds[i]).Error; nil != err {
			tx.Rollback()
		}

		if err := tx.Commit().Error; nil != err {
			tx.Rollback()
		}
	}
	return nil
}

// AddForeignKey to
//func (s *Repositories) AddForeignKey() error {
//	var err error
//	return err
//}
