package dogController

import (
	"github.com/go-chi/chi/v5"
	"github.com/jedzeins/go-chi-webserver/domains"
	"github.com/jedzeins/go-chi-webserver/services/dogService"

	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GetDogs(w http.ResponseWriter, r *http.Request) {
	dogs := dogService.GetDogs()

	res, err := json.Marshal(dogs)
	if err != nil {
		fmt.Printf("error in marshall: %s", err)
	}

	w.Write(res)
}

func GetOneDog(w http.ResponseWriter, r *http.Request) {

	index, errStrconv := strconv.Atoi(chi.URLParam(r, "index"))
	if errStrconv != nil {
		apiError := domains.DogError{
			StatusCode: http.StatusBadRequest,
			Message:    "The index needs to be an integer",
		}

		res, _ := json.Marshal(apiError)
		w.WriteHeader(400)
		w.Write(res)
		return
	}

	dog, err := dogService.GetOneDog(index)
	if err != nil {
		res, _ := json.Marshal(err)
		w.WriteHeader(400)
		w.Write(res)
	} else {
		res, err := json.Marshal(dog)
		if err != nil {
			fmt.Printf("error in marshall: %s", err)
		}
		w.WriteHeader(200)
		w.Write(res)
	}

}
