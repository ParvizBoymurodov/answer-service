package models

type  QuestionsAndAnswers struct {
	Id int64 `json:"id"`
	Question string `json:"question"`
	Answer  string   `json:"answer"`
}

//type  QuestionsAndAnswers2 struct {
//	Id int64 `json:"id"`
//	Question string `json:"question"`
//	Answer  string   `json:"answer"`
//}
