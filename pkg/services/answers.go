package services

import (
	errors2 "answers/pkg/errors"
	"answers/pkg/models"
	"context"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type AnswersSvc struct {
	pool *pgxpool.Pool
}

func NewAnswerSvc(pool *pgxpool.Pool) *AnswersSvc {
	if pool == nil {
		panic(errors.New("pool can't be nil"))
	}
	return &AnswersSvc{pool: pool}
}

func (service *AnswersSvc) QuestionsAndAnswersList() (list []models.QuestionsAndAnswers, err error) {
	list = make([]models.QuestionsAndAnswers, 0)
	conn, err := service.pool.Acquire(context.Background())
	if err != nil 	{
		return nil,errors2.QueryErrors("can't execute pool: ",err)
	}
	defer conn.Release()
	rows, err := conn.Query(context.Background(), "SELECT id, questions, answers FROM answers WHERE removed = FALSE")
	if err != nil {
		return nil, errors2.QueryErrors("can't query: ",err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.QuestionsAndAnswers{}
		err := rows.Scan(&item.Id, &item.Question, &item.Answer)
		if err != nil {
			return nil, errors2.QueryErrors("can't scan: ",err)
		}
		list = append(list, item)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (service *AnswersSvc) Save(model *models.QuestionsAndAnswers) (err error) {
	save, err := service.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't acuire: %d",err)
		return errors2.QueryErrors("can't execute pool: ",err)
	}
	defer save.Release()
	_, err = save.Exec(context.Background(), "INSERT INTO answers (questions,answers) VALUES ($1,$2)", model.Question, model.Answer)
	if err != nil {
		log.Printf("can't exec: %d",err)
		return errors2.QueryErrors("can't save:  ",err)
	}

	return nil
}

func (service *AnswersSvc) RemoveById(id int) (err error) {
	remove, err := service.pool.Acquire(context.Background())
	if err != nil {
		return errors2.QueryErrors("can't execute pool: ",err)
	}
	defer remove.Release()
	_, err = remove.Exec(context.Background(), "UPDATE answers SET removed = true WHERE id = $1", id)
	if err != nil {
		return errors2.QueryErrors("can't remove : ",err)
	}
	return nil
}

func (service *AnswersSvc) Start() {
	conn, err := service.pool.Acquire(context.Background())
	if err != nil {
		panic(errors.New("can't create database"))
	}
	defer conn.Release()
	_, err = conn.Exec(context.Background(), `
CREATE TABLE if not exists answers (
  id BIGSERIAL PRIMARY KEY,
  questions    TEXT NOT NULL,
  answers    TEXT NOT NULL ,
  removed BOOLEAN DEFAULT FALSE
);
`)
	if err != nil {
		panic(errors.New("can't create database"))
	}
}