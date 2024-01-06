// @Author huzejun 2024/1/1 20:35:00
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"logrus2024/logs"
	"net/http"
	"time"
)

var log *logs.Log

func init() {
	conf := logs.LogConf{
		Level:       logrus.InfoLevel,
		AdapterName: "fileRotate",
	}
	log = logs.InitLog(conf)
}

func main() {
	r := gin.New()
	r.Use(myLogger)
	//r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})
	r.Run(":8080")
}

func myLogger(ctx *gin.Context) {
	start := time.Now()
	path := ctx.Request.URL.Path
	raw := ctx.Request.URL.RawQuery
	ctx.Next()
	mp := make(map[string]interface{})
	mp["path"] = path
	mp["client_ip"] = ctx.ClientIP()
	mp["method"] = ctx.Request.Method
	mp["status_code"] = ctx.Writer.Status()
	mp["size"] = ctx.Writer.Size()
	if raw != "" {
		mp["path"] = path + "?" + raw
	}
	mp["latency"] = time.Since(start)
	log.WithFields(mp).Info()
}

func main1() {
	defer func() {
		log.Flush()
	}()

	log.Info()
	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(logrus.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")

	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	contextLogger := log.WithFields(logrus.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
	//return log

	log.Flush()

}
