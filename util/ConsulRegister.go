package util

import (
	consulApi "github.com/hashicorp/consul/api"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strconv"
)

func RegisterToConsul(e *echo.Echo, serverName string, currentIP string, currentPort int) {
	config := consulApi.DefaultConfig()
	client, err := consulApi.NewClient(config)

	if err != nil {
		log.Fatal("consul client error : ", err)
	}

	registration := new(consulApi.AgentServiceRegistration)
	registration.Name = serverName
	registration.ID = serverName + ": " + currentIP + ":" + strconv.Itoa(currentPort)
	registration.Address = currentIP
	registration.Port = currentPort
	registration.Check = &consulApi.AgentServiceCheck{
		HTTP: "http://" + currentIP + ":" + strconv.Itoa(currentPort) + "/check",
		Timeout: "3s",
		Interval: "5s",
		DeregisterCriticalServiceAfter: "10s", //check失败后10秒删除本服务
	}

	err = client.Agent().ServiceRegister(registration)

	if err != nil {
		log.Fatal("register server error : ", err)
	}

	e.GET("/check", consulHealthCheck)
}

func consulHealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "Health Check Successful")
}
