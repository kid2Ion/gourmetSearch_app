package main

import (
	"fmt"
	"gourmetSearch_app/parse"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

type GourmetSearcher struct {
	apikey string
}

func (g *GourmetSearcher) SearchGourmet(ctx echo.Context) error {
	keyWord := ctx.QueryParam("keyword")

	xml, err := g.callApi(keyWord)
	if err != nil {
		return err
	}
	// shop, err := parse.Parse(xml)
	shop, err := parse.MyParse(xml)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, shop)
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
