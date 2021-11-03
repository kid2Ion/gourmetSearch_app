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
	// クエリから検索キーワード取得
	keyWord := ctx.QueryParam("keyword")

	//sqliteにキャッシュ保持していないか検索
	shops := db.SearchShopCache(keyWord)
	if len(shops) != 0 {
		// cacheにあればjsonをフロントに返す
		return ctx.JSON(http.StatusOK, shops)
	} else {
		// cacheになければapiから取得
		xml, err := g.callApi(keyWord)
		if err != nil {
			return err
		}
		// パース処理
		shops, err := parse.MyParse(xml)
		if err != nil {
			return err
		}
		//新たにcacheに書き込む
		db.CreateShopCache(keyWord, shops)
		// jsonをフロントに返す
		return ctx.JSON(http.StatusOK, shops)
	}
}

func (g *GourmetSearcher) callApi(keyWord string) ([]byte, error) {
	// api叩くURI作成
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
	// コマンドライン第二引数にapikey
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s Apikey", os.Args[0])
	}
	apiKey := os.Args[1]
	// db開く
	shopdb, err := gorm.Open(sqlite.Open("shop.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	shopdb.AutoMigrate(&db.ShopCache{})
	searcher := &GourmetSearcher{apikey: apiKey}
	e := echo.New()
	e.Static("/", "vue/dist/")
	e.GET("/keyword", searcher.SearchGourmet)
	e.Logger.Fatal(e.Start(":8080"))
}
