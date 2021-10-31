package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ShopCache struct {
	gorm.Model
	Keyword   string
	ID        string
	Name      string
	LogoImage string
	Urls      string
}

func SearchShopCache(keyword string) []ShopCache {
	// db開く
	db, err := gorm.Open(sqlite.Open("shop.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var shops []ShopCache
	db.Find(&shops)
	return shops
}

// func CreateShopCache(keyword string, shop []parse.Shop) {
// 	// db開く
// 	db, err := gorm.Open(sqlite.Open("shop.db"), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	shopcache := ShopCache{Keyword: keyword, ID: shop.ID, Name: shop.Name, LogoImage: shop.LogoImage, Urls: shop.Urls}
// 	db.Create(&shopcache)
// }
