package client_server

import (
	"encoding/json"
	"github.com/google/uuid"
	"mime/multipart"
	"net/http"
	"pafaul/overcomplicated-image-compressor/server_errors"
	"strconv"
)

type (
	RequestType string

	ClientRequest struct {
		ClientId      string          `json:"clientId"`
		RequestType   RequestType     `json:"requestType"`
		RequestParams json.RawMessage `json:"requestParams"`
	}

	CompressRequest struct {
		WidthMultiplier  float64 `json:"widthMultiplier"`
		HeightMultiplier float64 `json:"heightMultiplier"`
		TargetWidth      int32   `json:"targetWidth"`
		TargetHeight     int32   `json:"targetHeight"`
	}
)

var (
	_ = RequestType("compress")
)

var maxMemory = int64(1024 * 1024 * 64)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/file-upload.html")
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(maxMemory)
	if err != nil {
		server_errors.ReplyWithError(w, server_errors.NewServerError(err.Error()))
		return
	}

	clientIdStr := r.FormValue("clientId")
	if clientIdStr == "" {
		server_errors.ReplyWithError(w, server_errors.NewBadRequest("Client id cannot be empty"))
		return
	}

	_, err = strconv.ParseUint(clientIdStr, 10, 64)
	if err != nil {
		server_errors.ReplyWithError(w, server_errors.NewBadRequest(err.Error()))
		return
	}

	f, header, err := r.FormFile("file")
	filename := getUploadedFilename(header)
	err = saveFile(f, filename)
	if err != nil {
		server_errors.ReplyWithError(w, server_errors.NewServerError(err.Error()))
		return
	}
}

func getUploadedFilename(fh *multipart.FileHeader) string {
	return uuid.NewString() + "." + getFileExt(fh.Filename)
}

var _ = http.HandlerFunc(requestHandler)
var _ = http.HandlerFunc(indexHandler)
