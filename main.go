package main

import (
	"log"
	"net"
)

// Application is the main handler for chat server
type Application struct {
	// DBs            database.Database
	// AggregatorHndl aggregator.AggregatorInterface
	//KafkaHndl      mq.KafkaInterface
}

var App Application

func NewApplicationHandler() {
	log.Println("Initializing Chat server app")
	// error handling done in caller routine, app exits if error occurs
	// App.DBs.PostgresDBAnalyticsHndl, _ = database.NewAnalyticsDBHandler()
	// App.DBs.EsDBAnalyticsHndl, _ = database.NewElasticDBHandler()
	// App.AggregatorHndl, _ = aggregator.NewAggregatorHandler(App.DBs)
}

func main() {
	//init listener
	listen, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalf("Could not listen @ %s :: %v", listen.Addr(), err)
	}
}