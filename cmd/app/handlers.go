package app

import (
	"answers/pkg/models"
	"github.com/ParvizBoymurodov/mux/pkg/mux"
	"github.com/ParvizBoymurodov/rest/pkg"
	"log"
	"net/http"
	"strconv"
)

func (receiver server) handleAnswersList() func( http.ResponseWriter,*http.Request ) {
	return func(writer http.ResponseWriter, request *http.Request) {
		list, err := receiver.answersSvc.QuestionsAndAnswersList()
		if err != nil {
			log.Print(err)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		err =rest.WriteJSONBody(writer, &list)
		return
		http.Error(writer,http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Print(err)
	}
}

func (receiver server) handleRemoveAnswers() func( http.ResponseWriter,*http.Request ) {
  return func(writer http.ResponseWriter, request *http.Request) {
	  idContext, ok := mux.FromContext(request.Context(), "id")
	  if !ok {
		  http.Error(writer,http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		  return
	  }
	  id, err := strconv.Atoi(idContext)
	  if err != nil {
		  http.Error(writer,http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		  return
	  }
	  err= receiver.answersSvc.RemoveById(id)
	  if err != nil {
		  http.Error(writer,http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		  log.Print(err)
		  return
	  }
  }
}

func (receiver server) handleAddAnswersAndQuestions() func( http.ResponseWriter,*http.Request) {
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
			log.Printf("can'r read json: %d",err)
			return
		}
		err = receiver.answersSvc.Save(&answer)
		if err != nil {
			http.Error(writer,http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Print(err)
			return
		}


		err = rest.WriteJSONBody(writer, &answer)
		if err != nil {
			http.Error(writer,http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}

}



