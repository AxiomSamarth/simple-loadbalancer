package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/AxiomSamarth/loadbalancers/loadbalancer"
	"github.com/AxiomSamarth/loadbalancers/server"
)

func main() {
	lb := loadbalancer.NewLoadbalancer("8080")

	firstServerAddress := "localhost:8000"
	firstServer := server.NewSimpleServer(firstServerAddress)
	err := lb.RegisterServer(firstServer)
	if err != nil {
		log.Fatalf("error registering firstServer to Loadbalancer: %s", err)
	}

	secondServerAddress := "localhost:8001"
	secondServer := server.NewSimpleServer(secondServerAddress)
	err = lb.RegisterServer(secondServer)
	if err != nil {
		log.Fatal("error registering secondServer to Loadbalancer")
	}

	thirdServerAddress := "localhost:8002"
	thirdServer := server.NewSimpleServer(thirdServerAddress)
	err = lb.RegisterServer(thirdServer)
	if err != nil {
		log.Fatal("error registering thirdServer to Loadbalancer")
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = firstServer.Start()
		if err != nil {
			log.Printf("error starting firstServer: %s", err.Error())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = secondServer.Start()
		if err != nil {
			log.Printf("error starting secondServer: %s", err.Error())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = thirdServer.Start()
		if err != nil {
			log.Printf("error starting thirdServer: %s", err.Error())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		handleRedirect := func(w http.ResponseWriter, r *http.Request) {
			lb.ServeProxy(w, r)
		}
		http.HandleFunc("/", handleRedirect)
		fmt.Printf("\nserving requests at localhost:%s", lb.GetPort())
		if err := http.ListenAndServe(":"+lb.GetPort(), nil); err != nil {
			log.Printf("error starting the server: %s", err)
		}
	}()

	wg.Wait()
}
