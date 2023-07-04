package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"l0_wb/internal/utils/logger"
)

type errResp struct {
	Error string `json:"error"`
}

const errPrefix = "HTTP_ERROR"

func writeErrorAndLog(w http.ResponseWriter, errMsg string, intErr error) {
	logger.Log().Errorw(fmt.Sprintf("%s %s", errPrefix, errMsg), "internal_error", intErr)

	body := errResp{Error: errMsg}
	bytes, err := json.Marshal(body)
	if err != nil {
		logger.Log().Errorw(fmt.Sprintf("%s can't marshall body to json", errPrefix), "body", body)
		return
	}

	writeRespAndLogIfCant(w, bytes)
}

func writeRespAndLogIfCant(w http.ResponseWriter, bytes []byte) {
	bytes = append(bytes, '\n')
	_, err := w.Write(bytes)
	if err != nil {
		logger.Log().Errorw(fmt.Sprintf("%s can't write to response writer", errPrefix), "error", err)
	}
}
