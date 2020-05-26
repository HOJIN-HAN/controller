package controller // import "github.com/HOJIN-HAN/controller"

import (
	"log"

	"github.com/HOJIN-HAN/models"
	"github.com/revel/revel"
)

// Board ..
type Board struct {
	App
}

// Index ...
func (c Board) Index() revel.Result {
	results, err := c.Txn.Select(models.Board{}, `select * from tbl_user`)
	if err != nil {
		panic(err)
	}

	var articles []*models.Board
	for _, r := range results {
		b := r.(*models.Board)
		articles = append(articles, b)
	}
	log.Println(articles)
	return c.Render(articles)
}

// Index ...
func (c Board) Read(Id int) revel.Result {
	return c.Render(Id)
}
