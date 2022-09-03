package tests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"net/http/httptest"
	"snappfood_challenge/db"
	"snappfood_challenge/forms"
	"snappfood_challenge/internal"
	"strings"
	"testing"
)

func Init() {

	if err := gotenv.Load("../.env"); err != nil {
		log.Print("No .env file found")
	}

	viper.AddConfigPath("../config")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	db.Init()
	db.CreateOrderTable()

	db.InitRedis()
}

func TestPublishOrder(t *testing.T)  {
	order := forms.AddOrder{
		Id:    100,
		Price: 3000,
		Title: "test",
	}

	Init()

	orderService := internal.OrderService{}
	err := orderService.PublishOrder(order)

	assert.Equal(t, nil, err)
}

func handler(c *gin.Context) {
	var info forms.AddOrder
	if err := c.ShouldBindJSON(&info); err != nil {
		log.Panic(err)
	}
	fmt.Println(info)
	c.Writer.Write([]byte(`{"status": 200}`))
}

func TestAddOrder(t *testing.T)  {
	payload := strings.NewReader(`{
    "id": 200,
    "price": 2000,
	"title": "test"}`)

	path := "/api/v1/order"
	router := gin.Default()
	router.POST(path, handler)
	req, _ := http.NewRequest("POST", path, payload)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, http.StatusOK, w.Code)
}
