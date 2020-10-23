package services

import (
	"log"
	"webproject/rsystem/models"
)

type RJenkinsServer struct {
	Client *models.RJenkins
}

var DefaultRJenkinsServer *RJenkinsServer

func NewDefaultJenkinsServer() *RJenkinsServer {
	return &RJenkinsServer{
		Client: models.DefaultJenkins,
	}
}

func init() {
	DefaultRJenkinsServer = NewDefaultJenkinsServer()
	log.Println("Jenkins server 初始化成功")
}
