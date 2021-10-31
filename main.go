package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/hiroki-kondo-git/gourmetSearch_app/db"

	"github.com/hiroki-kondo-git/gourmetSearch_app/parse"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/labstack/echo"
)

type GourmetSearcher struct {
	apikey string
}

func (g *GourmetSearcher) SearchGourmet(ctx echo.Context) error {
	keyWord := ctx.QueryParam("keyword")

	//sqlite検索
	shops := db.SearchShopCache(keyWord)
	if len(shops) != 0 {
		return ctx.JSON(http.StatusOK, shops)
	} else {
		// apiから取得
		xml, err := g.callApi(keyWord)
		if err != nil {
			return err
		}
		shop, err := parse.MyParse(xml)
		if err != nil {
			return err
		}
		//cacheに書き込む
		// db.CreateShopCache(keyWord, shop)
		return ctx.JSON(http.StatusOK, shop)
	}
}

func (g *GourmetSearcher) callApi(keyWord string) ([]byte, error) {
	url := fmt.Sprintf("https://webservice.recruit.co.jp/hotpepper/gourmet/v1/?key=%s&keyword=%s&count=20", g.apikey, keyWord)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s Apikey", os.Args[0])
	}
	apiKey := os.Args[1]
	// Migrate the schema
	shopdb, err := gorm.Open(sqlite.Open("shop.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	shopdb.AutoMigrate(&db.ShopCache{})
	searcher := &GourmetSearcher{apikey: apiKey}
	e := echo.New()
	e.Static("/", "vue/dist/")
	e.GET("/keyword", searcher.SearchGourmet)
	e.Logger.Fatal(e.Start(":8080"))
}

func GourmetSerach(APIkey, small_area string) error {
	url := fmt.Sprintf("https://webservice.recruit.co.jp/hotpepper/gourmet/v1/?key=%s&small_area=%s", APIkey, small_area)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}
