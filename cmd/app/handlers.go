package app

import (
	"answers/pkg/models"
	"github.com/ParvizBoymurodov/mux/pkg/mux"
	"github.com/ParvizBoymurodov/rest/pkg"
	"log"
	"net/http"
	"strconv"
)

func (receiver server) handleAnswersList() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		list, err := receiver.answersSvc.QuestionsAndAnswersList()
		if err != nil {
			log.Print(err)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		err = rest.WriteJSONBody(writer, &list)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Print(err)
			return
		}

	}
}

func (receiver server) handleRemoveAnswers() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		idContext, ok := mux.FromContext(request.Context(), "id")
		if !ok {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(idContext)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		err = receiver.answersSvc.RemoveById(id)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Print(err)
			return
		}
	}
}

func (receiver server) handleAddAnswersAndQuestions() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		get := request.Header.Get("Content-Type")
		if get != "application/json" {
			log.Println("can't")
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		answer := models.QuestionsAndAnswers{}
		err := rest.ReadJSONBody(request, &answer)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			log.Printf("can'r read json: %d", err)
			return
		}
		err = receiver.answersSvc.Save(&answer)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Print(err)
			return
		}

		err = rest.WriteJSONBody(writer, &answer)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}

}

func (receiver server) handleUpdateAnswerAndQuestion() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		get := request.Header.Get("Content-Type")
		if get != "application/json" {
			log.Println("can't")
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		ansQue := models.QuestionsAndAnswers2{}
		err := rest.ReadJSONBody(request, &ansQue)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			log.Printf("can'r read json: %d", err)
			return
		}
		err = receiver.answersSvc.UpdateAnswerAndQuestion(ansQue)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Print(err)
			return
		}

		err = rest.WriteJSONBody(writer, &ansQue)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

	}

}

// --------------------------------------------Category---------------------------------------------------------

func (receiver server) handleCategoryList() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		list, err := receiver.answersSvc.CategoryList()
		if err != nil {
			log.Print(err)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		err = rest.WriteJSONBody(writer, &list)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Print(err)
			return
		}

	}
}

func (receiver server) handleAddCategory() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		get := request.Header.Get("Content-Type")
		if get != "application/json" {
			log.Println("can't")
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		category := models.Сategory{}
		err := rest.ReadJSONBody(request, &category)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			log.Printf("can'r read json: %d", err)
			return
		}
		err = receiver.answersSvc.AddCategory(category)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Print(err)
			return
		}

		err = rest.WriteJSONBody(writer, &category)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

func (receiver server) handleRemovedCategory() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		idCategory, ok := mux.FromContext(request.Context(), "id")
		if !ok {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(idCategory)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		err = receiver.answersSvc.RemovedCategoryById(id)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Print(err)
			return
		}

	}
}

func (receiver server) handleUpdateCategory() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		get := request.Header.Get("Content-Type")
		if get != "application/json" {
			log.Println("can't")
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		category := models.Сategory{}
		err := rest.ReadJSONBody(request, &category)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			log.Printf("can'r read json: %d", err)
			return
		}
		err = receiver.answersSvc.UpdateCategory(category)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Print(err)
			return
		}

		err = rest.WriteJSONBody(writer, &category)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

	}

}

//------------------------------------Search----------------------------------

func (receiver server) search() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		val := request.FormValue("q")
		search, err := receiver.answersSvc.Search(val)
		if err != nil {
			log.Print(err)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		err = rest.WriteJSONBody(writer, search)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Print(err)
			return
		}
	}
}
