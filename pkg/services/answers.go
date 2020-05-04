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

//----------------------------Category-------------------------
func (service *AnswersSvc) CategoryList() (list []models.СategorySelect, err error) {
	list = make([]models.СategorySelect, 0)
	conn, err := service.pool.Acquire(context.Background())
	if err != nil {
		return nil, errors2.QueryErrors("can't execute pool: ", err)
	}
	defer conn.Release()
	rows, err := conn.Query(context.Background(), "select c.id,c.name,a.questions,a.answers from categories c inner join answers a on c.id = a.id_category where c.removed =false and a.removed= false;")
	if err != nil {
		return nil, errors2.QueryErrors("can't query: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.СategorySelect{}
		err := rows.Scan(&item.Id, &item.Name, &item.Question, &item.Answer)
		if err != nil {
			return nil, errors2.QueryErrors("can't scan: ", err)
		}
		list = append(list, item)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (service *AnswersSvc) AddCategory(model models.Сategory) (err error) {
	save, err := service.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't acuire: %d", err)
		return errors2.QueryErrors("can't execute pool: ", err)
	}
	defer save.Release()
	_, err = save.Exec(context.Background(), "INSERT INTO categories (name) VALUES ($1)", model.Name)
	if err != nil {
		log.Printf("can't exec: %d", err)
		return errors2.QueryErrors("can't save:  ", err)
	}

	return nil
}

func (service *AnswersSvc) RemovedCategoryById(id int) (err error)  {
	removed, err := service.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't acuire: %d", err)
		return errors2.QueryErrors("can't execute pool: ", err)
	}
	defer removed.Release()
	_, err = removed.Exec(context.Background(), "UPDATE categories SET removed = true WHERE id = $1", id)
	if err != nil {
		return errors2.QueryErrors("can't remove : ", err)
	}
	return nil

}

func (service *AnswersSvc) UpdateCategory(c models.Сategory) (err error) {
	update, err := service.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't acuire: %d", err)
		return errors2.QueryErrors("can't execute pool: ", err)
	}
	defer update.Release()
	_, err = update.Exec(context.Background(), "Update categories SET name = $1 WHERE id = $2 AND removed = false", c.Name,c.Id)
	if err != nil {
		return errors2.QueryErrors("can't remove : ", err)
	}
	return nil
}

// ----------------------------Answers-------------------------------------
func (service *AnswersSvc) QuestionsAndAnswersList() (list []models.QuestionsAndAnswersSelect, err error) {
	list = make([]models.QuestionsAndAnswersSelect, 0)
	conn, err := service.pool.Acquire(context.Background())
	if err != nil {
		return nil, errors2.QueryErrors("can't execute pool: ", err)
	}
	defer conn.Release()
	rows, err := conn.Query(context.Background(), "select a.id, a.questions, a.answers, c.name from answers a inner join categories c on a.id_category = c.id where a.removed = false;")
	if err != nil {
		return nil, errors2.QueryErrors("can't query: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.QuestionsAndAnswersSelect{}
		err := rows.Scan(&item.Id, &item.Question, &item.Answer, &item.СategoryName)
		if err != nil {
			return nil, errors2.QueryErrors("can't scan: ", err)
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
		log.Printf("can't acuire: %d", err)
		return errors2.QueryErrors("can't execute pool: ", err)
	}
	defer save.Release()
	_, err = save.Exec(context.Background(), "INSERT INTO answers (questions,answers,id_category) VALUES ($1,$2,$3)", model.Question, model.Answer, model.IdCategory)
	if err != nil {
		log.Printf("can't exec: %d", err)
		return errors2.QueryErrors("can't save:  ", err)
	}

	return nil
}

func (service *AnswersSvc) RemoveById(id int) (err error) {
	remove, err := service.pool.Acquire(context.Background())
	if err != nil {
		return errors2.QueryErrors("can't execute pool: ", err)
	}
	defer remove.Release()
	_, err = remove.Exec(context.Background(), "UPDATE answers SET removed = true WHERE id = $1", id)
	if err != nil {
		return errors2.QueryErrors("can't remove : ", err)
	}
	return nil
}

func (service *AnswersSvc) UpdateAnswerAndQuestion(model models.QuestionsAndAnswers2) (err error) {
	update, err := service.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't acuire: %d", err)
		return errors2.QueryErrors("can't execute pool: ", err)
	}
	defer  update.Release()
	_, err = update.Exec(context.Background(), "UPDATE answers SET questions = $1 , answers = $2 WHERE id = $3 and removed = false", model.Question, model.Answer, model.Id)
	if err != nil {
		return errors2.QueryErrors("can't remove : ", err)
	}
	return nil
}

//---------------------DB Tables----------------------

func (service *AnswersSvc) Start() {
	conn, err := service.pool.Acquire(context.Background())
	if err != nil {
		panic(errors.New("can't create database"))
	}
	defer conn.Release()
	_, err = conn.Exec(context.Background(), `
CREATE TABLE if not exists categories (
  id BIGSERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  removed BOOLEAN DEFAULT FALSE
);
`)

	if err != nil {
		panic(errors.New("can't create database"))
	}
	_, err = conn.Exec(context.Background(), `
CREATE TABLE if not exists answers (
  id BIGSERIAL PRIMARY KEY,
  questions    TEXT NOT NULL,
  answers    TEXT NOT NULL ,
  id_category BIGSERIAL references categories(id),
  removed BOOLEAN DEFAULT FALSE
);
`)
	if err != nil {
		panic(errors.New("can't create database"))
	}

}

//------------------------Search----------------------------------------

func (service AnswersSvc) Search(lists string) (list []models.QuestionsAndAnswers, err error) {
	conn, err := service.pool.Acquire(context.Background())
	if err != nil {
		return nil, errors2.QueryErrors("can't execute pool: ", err)
	}
	defer conn.Release()
	rows, err := conn.Query(context.Background(), "SELECT id, questions, answers, id_category FROM answers WHERE questions  LIKE $1 AND removed = FALSE", "%"+lists+"%")
	if err != nil {
		return nil, errors2.QueryErrors("can't query: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		item := models.QuestionsAndAnswers{}
		err := rows.Scan(&item.Id, &item.Question, &item.Answer, &item.IdCategory)
		if err != nil {
			return nil, errors2.QueryErrors("can't scan: ", err)
		}
		list = append(list, item)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return list, nil
}
