package db

import (
	"github.com/hiroki-kondo-git/gourmetSearch_app/parse"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ShopCache struct {
	gorm.Model
	Keyword   string
	ShopID    string
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
	// keywordをプライマリーキーにして、dbから探す
	db.Find(&shops, "Keyword = ?", keyword)
	return shops
}

func CreateShopCache(keyword string, shops []parse.Shop) {
	// db開く
	db, err := gorm.Open(sqlite.Open("shop.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	for _, shop := range shops {
		shopCache := ShopCache{}
		shopCache.Keyword = keyword
		shopCache.ShopID = shop.ID
		shopCache.Name = shop.Name
		shopCache.LogoImage = shop.LogoImage
		shopCache.Urls = shop.Urls.Pc
		db.Create(&shopCache)
	}
}
