package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"l0_wb/internal/app"

	"github.com/gorilla/mux"
)

type modelHandler struct {
	app app.Service
}

func NewModelHandler(app app.Service) modelHandler {
	return modelHandler{app: app}
}

func (h modelHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	id, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("no variable %q in url", "id")
		writeErrorAndLog(w, msg, fmt.Errorf(msg))
		return
	}

	model, err := h.app.Get(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// TODO: обработку ошибок в этом месте бы...
		writeErrorAndLog(w, "can't get model or model with this id not exists :(", err)
		return
	}

	bytes, err := json.Marshal(model)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeErrorAndLog(w, "can't marshal model in json :(", err)
		return
	}

	writeRespAndLogIfCant(w, bytes)
}
