package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	"github.com/stevenkitter/tools/database"

	"os"
)

type Controller struct {
	g     *gin.Engine
	d     *gorm.DB
	cache *redis.Client
}

// NewController init gin
func NewController() *Controller {
	r := gin.Default()
	user := "tools"
	da, err := database.ConnectMysqlDB(
		user, os.Getenv("MYSQL_PWD"),
		os.Getenv("MYSQL_HOST"), user)
	if err != nil {
		panic(err)
	}
	rd := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ct := &Controller{
		g:     r,
		d:     da,
		cache: rd,
	}
	ct.Route()
	return ct
}

// Run start server
func (ct *Controller) Run(addr string) error {
	appId := "1Yv4vFyYK27tOdh1CTtL25ObC19"
	key := fmt.Sprintf("app:secret:%s", appId)
	ct.cache.Set(key, "be010be1127d3547ec72d57561e61b8a81a03209", 0)
	return ct.g.Run(addr)
}
