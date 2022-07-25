package mysql

import (
	"fmt"
	"log"
	"socialmediabackendproject/config"
	commentData "socialmediabackendproject/feature/comments/data"
	postImagesData "socialmediabackendproject/feature/post_images/data"
	"socialmediabackendproject/feature/posts/data"
	usersData "socialmediabackendproject/feature/users/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cfg.Username, cfg.Password, cfg.Address, cfg.Port, cfg.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	return db
}

func MigrateData(db *gorm.DB) {
	db.AutoMigrate(usersData.User{}, data.Post{}, postImagesData.Post_Image{}, commentData.Comment{})
}
