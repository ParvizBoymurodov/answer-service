package app

func (receiver *server) InitRoutes() {
	receiver.router.GET("/api/answers", receiver.handleAnswersList())
	receiver.router.DELETE("/api/answers/{id}", receiver.handleRemoveAnswers())
	receiver.router.POST("/api/answers",receiver.handleAddAnswersAndQuestions())
}
