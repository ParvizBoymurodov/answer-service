package main

import (
	"answers/cmd/app"
	"answers/pkg/services"
	"context"
	"flag"
	"github.com/ParvizBoymurodov/mux/pkg/mux"
	_ "github.com/ParvizBoymurodov/mux/pkg/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"net"
	"net/http"
	"os"
)

var (
	host = flag.String("host", "", "Server host")
	port = flag.String("port", "", "Server port")
	dsn  = flag.String("dsn", "", "Postgres DSN")
)

const (
	envHost = "HOST"
	envPort = "PORT"
	envDSN  = "DATABASE_URL"
)

func fromFLagOrEnv(flag *string, envName string) (server string, ok bool) {
	if *flag != "" {
		return *flag, true
	}
	return os.LookupEnv(envName)
}

func main() {
	flag.Parse()
	hostf, ok := fromFLagOrEnv(host, envHost)
	if !ok {
		hostf = *host
	}
	portf, ok := fromFLagOrEnv(port, envPort)
	if !ok {
		portf = *port
	}
	dsnf, ok := fromFLagOrEnv(dsn, envDSN)
	if !ok {
		dsnf = *dsn
	}

	addr := net.JoinHostPort(hostf, portf)
	start(addr, dsnf)
}

func start(addr string, dsn string) {

	router := mux.NewExactMux()
	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		panic(err)
	}
	answersSvc := services.NewAnswerSvc(pool)
	server := app.NewServer(
		pool,
		router,
		answersSvc,
	)
	server.InitRoutes()
	answersSvc.Start()
	panic(http.ListenAndServe(addr, server))
}
