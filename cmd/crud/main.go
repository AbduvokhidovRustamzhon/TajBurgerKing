package main

// package
// import
// var + type
// method + function



import (
	"context"
	"crud/cmd/crud/app"
	"crud/pkg/crud/services"
	"crud/pkg/crud/services/burgers"
	"flag"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net"
	"net/http"
	"path/filepath"
)

const (
	EHost = "HOST"
	EPort = "PORT"
	EDsn  = "DATABASE_URL"
)
var (
	host = flag.String("host", "", "Server host")
	port = flag.String("port", "", "Server port")
	dsn  = flag.String("dsn", "", "Postgres DSN")
)

func main() {
	flag.Parse()
	host, ok := services.FlagOrEnv(*host, EHost)
	if !ok {
		log.Panic("can't get port")
	}

	port, ok := services.FlagOrEnv(*port, EPort)
	if !ok {
		log.Panic("can't get port")
	}

	log.Println("set address to connect")
	addr := net.JoinHostPort(host, port)
	log.Printf("address to connect: %s", addr)

	dsn, ok := services.FlagOrEnv(*dsn, EDsn)
	if !ok {
		log.Panic("can't get dsn")
	}

	log.Printf("try start server on: %s, dbUrl: %s", addr, dsn)
	start(addr, dsn)
	log.Printf("server success on: %s, dbUrl: %s", addr, dsn)
}

func start(addr string, dsn string) {
	router := app.NewExactMux()
	// Context: <-
	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		panic(err)
	}
	burgersSvc := burgers.NewBurgersSvc(pool)
	server := app.NewServer(
		router,
		pool,
		burgersSvc, // DI + Containers
		filepath.Join("web", "templates"),
		filepath.Join("web", "assets"),
	)
	server.InitRoutes()

	// server'ы должны работать "вечно"
	panic(http.ListenAndServe(addr, server)) // поднимает сервер на определённом адресе и обрабатывает запросы
}
