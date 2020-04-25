package app

func (receiver *server) InitRoutes() {
	receiver.router.GET("/api/answers", receiver.handleAnswersList())
	receiver.router.DELETE("/api/answers/{id}", receiver.handleRemoveAnswers())
	receiver.router.POST("/api/answers",receiver.handleAddAnswersAndQuestions())
	receiver.router.GET("/api/answers/search",receiver.search())
	// Category
	receiver.router.GET("/api/categories",receiver.handleCategoryList())
	receiver.router.POST("/api/categories",receiver.handleAddCategory())
	receiver.router.DELETE("/api/categories/{id}",receiver.handleRemovedCategory())
}
