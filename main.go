package main

import (
	"fmt"
	"goredis/repositories"
	"goredis/services"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDatabase()
	redisClient := initRedis()
	_ = redisClient

	productRepo := repositories.NewProductRepositoryDB(db)
	//productRepo := repositories.NewProductRepositoryRedis(db, redisClient)
	productService := services.NewCatalogService(productRepo)

	products, err := productService.GetProducts()
	if err != nil {
		fmt.Println("Error GetProducts:: ", err)
		return
	}

	fmt.Println(products)
}

func initDatabase() *gorm.DB {
	dial := mysql.Open("root:password@1234@tcp(localhost:3307)/testdb")
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
