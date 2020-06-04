package function

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	//"xorm.io/xorm"
	//"github.com/go-redis/redis"
)

type RuntimeConfig struct {
	//define your config params
}

type Function struct {
	RuntimeConfig
	*gin.Engine

	//interal orm or redis

	//orm *xorm.Engine
	//rds redis.UniversalClient
}

func NewFunctionApp(cfg *RuntimeConfig) (http.Handler, error) {
	var app Function
	app.RuntimeConfig = *cfg

	r := gin.Default()

	//implement your router
	r.GET("/hello", app.hello)

	app.Engine = r

	//TODO: init app
	return &app, fmt.Errorf("you should implement")
}

func (self *Function) hello(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}
