package main

import (
	"api-sample/database"
	"api-sample/server"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	tracer.Start(
		tracer.WithEnv("dev"),
		tracer.WithService("api-sample-totvs"),
		tracer.WithServiceVersion("v1"),
	)
	defer tracer.Stop()
	database.StartDB()
	server := server.NewServer()
	server.Run()
}
