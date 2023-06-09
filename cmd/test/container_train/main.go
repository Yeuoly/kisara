package main

/*
	Date: 2023/04/01
	Author: Yeuoly

	this package is used to test create containers in serval networks
	using birdge/host/overlay mode
	and test weather the container can connect to the internet
*/

import (
	"strings"
	"time"

	"github.com/Yeuoly/kisara/src/client"
	docker "github.com/Yeuoly/kisara/src/routine/docker"
)

func main() {
	go client.Main()
	time.Sleep(time.Second * 10)
	testTrain()

	select {}
}

func testTrain() {
	client := docker.NewDocker()
	defer client.Stop()
	_, err := client.CreateNetwork("172.127.0.0/16", "irina-train", false, "bridge")
	if err != nil && !strings.ContainsAny(err.Error(), "already exists") {
		panic(err)
	}

	// launch container
	_, err = client.LaunchTargetMachine("ctfhub/base_web_nginx_mysql_php_56", "80/tcp", "irina-train", 9, "train")
	if err != nil {
		panic(err)
	}

	// stop container
	// err = client.StopContainer(container.Id)
	// if err != nil {
	// 	panic(err)
	// }
}
