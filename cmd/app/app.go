package app

import (
	"answers/pkg/services"
	"github.com/ParvizBoymurodov/mux/pkg/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"net/http"
)

type server struct {
	pool          *pgxpool.Pool
	router        *mux.ExactMux
	answersSvc    *services.AnswersSvc
}

func NewServer(pool *pgxpool.Pool, router *mux.ExactMux, answersSvc *services.AnswersSvc) *server {
	if router == nil {
		panic(errors.New("router can't be nil"))
	}
	if pool == nil {
		panic(errors.New("pool can't be nil"))
	}
	if answersSvc == nil {
		panic(errors.New("burgersSvc can't be nil"))
	}
	return &server{
		pool: pool,
		router: router,
		answersSvc: answersSvc,
	}

}



func (receiver *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	receiver.router.ServeHTTP(writer, request)
}
