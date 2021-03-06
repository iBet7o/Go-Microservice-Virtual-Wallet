package api

import (
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"

	"github.com/Marlos-Rodriguez/go-postgres-wallet-back/auth/internal/cache"
	"github.com/Marlos-Rodriguez/go-postgres-wallet-back/auth/internal/storage"
)

//Start Start a new User server API
func Start() {
	var DB *gorm.DB

	DB = storage.ConnectDB()
	defer DB.Close()

	var RDB *redis.Client

	RDB = cache.NewRedisClient()
	defer RDB.Close()

	app := routes(DB, RDB)
	createServer(app)
}
