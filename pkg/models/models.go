package models

type  QuestionsAndAnswers struct {
	Id int64 `json:"id"`
	Question string `json:"question"`
	Answer  string   `json:"answer"`
	IdCategory int64 `json:"id_category"`
}
type  QuestionsAndAnswersSelect struct {
	Id int64 `json:"id"`
	Question string `json:"question"`
	Answer  string   `json:"answer"`
	СategoryName string `json:"сategory_name"`
}

type QuestionsAndAnswers2 struct {
	Id int64 `json:"id"`
	Question string `json:"question"`
	Answer  string   `json:"answer"`
}

type  Сategory struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
}


type  СategorySelect struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Question string `json:"question"`
	Answer  string   `json:"answer"`
}


