package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"net/http"
)

type ProductRequest struct {
	Id      *int  `json:"id"`
	Price   *int  `json:"price"`
	Deleted *bool `json:"deleted"`
}

type Product struct {
	Id      *int  `json:"id"`
	Price   *int  `json:"price"`
	Deleted *bool `json:"deleted"`
}

type ProductRecord struct {
	Id      int
	Price   int
	Deleted bool
}

func main() {
	r := gin.Default()
	r.POST("/", func(gc *gin.Context) {
		req := &ProductRequest{}
		if err := bind(gc, req); err != nil {
			gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		dp := Product{
			Id:      req.Id,
			Price:   req.Price,
			Deleted: req.Deleted,
		}
		update(&dp)
		gc.JSON(200, req)
	})
	r.Run(":3000")
}

func bind(gc *gin.Context, v interface{}) error {
	if err := gc.Bind(v); err != nil {
		return err
	}
	return nil
}

func update(product *Product) {

	record := toRecord(product)
	db, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3316)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.Model(&record).Updates(&record)
}

func toRecord(d *Product) *ProductRecord {
	return &ProductRecord{
		Id:      *d.Id,
		Price:   *d.Price,
		Deleted: *d.Deleted,
	}
}

func (ProductRecord) TableName() string {
	return "products"
}
