package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"obliviate/app"
	"obliviate/config"
)

type SaveRequest struct {
	Message           []byte `json:"message"`
	TransmissionNonce string `json:"nonce"`
	Hash              string `json:"hash"`
	PublicKey         string `json:"publicKey"`
}

type ReadRequest struct {
	Hash      string `json:"hash"`
	PublicKey string `json:"publicKey"`
}

type ReadResponse struct {
	Message string `json:"message"`
}

type templateData struct {
	PublicKey string
}

const (
	jsonErrMsg = "input json error"
	emptyBody  = "empty body post, no json expected"
)

func ProcessTemplate(config *config.Configuration, publicKey string) http.HandlerFunc {
	data := templateData{PublicKey: publicKey}

	var t *template.Template
	if config.ProdEnv {
		t = template.Must(template.New("template.html").ParseFiles("template.html"))
	}
	return func(w http.ResponseWriter, r *http.Request) {

		logrus.Trace("ProcessTemplate Handler")

		if !config.ProdEnv {
			t, _ = template.New("template.html").ParseFiles("template.html")
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		t.Execute(w, data)
	}
}

func Save(app *app.App) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		logrus.Debug("Save Handler")

		defer r.Body.Close()
		if r.Body == nil {
			finishRequestWithErr(w, emptyBody, http.StatusBadRequest)
			return
		}

		data := SaveRequest{}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			finishRequestWithErr(w, jsonErrMsg, http.StatusBadRequest)
			return
		}
		if len(data.Message) == 0 {
			finishRequestWithErr(w, "Message is empty", http.StatusBadRequest)
			return
		}
		if len(data.TransmissionNonce) == 0 {
			finishRequestWithErr(w, "TransmissionNonce is empty", http.StatusBadRequest)
			return
		}
		if len(data.Hash) == 0 {
			finishRequestWithErr(w, "Hash is empty", http.StatusBadRequest)
			return
		}

		nonce, err := base64.StdEncoding.DecodeString(data.TransmissionNonce)
		if err != nil {
			finishRequestWithErr(w, "TransmissionNonce cannot be decoded", http.StatusBadRequest)
			return
		}
		if len(nonce) != 24 {
			finishRequestWithErr(w, "TransmissionNonce length is wrong !=24", http.StatusBadRequest)
			return
		}

		publicKey, err := base64.StdEncoding.DecodeString(data.PublicKey)
		if err != nil {
			finishRequestWithErr(w, "PublicKey cannot be decoded", http.StatusBadRequest)
			return
		}
		if len(publicKey) != 32 {
			finishRequestWithErr(w, "PublicKey length is wrong !=24", http.StatusBadRequest)
			return
		}

		err = app.ProcessSave(r.Context(), data.Message, nonce, data.Hash, publicKey)
		if err != nil {
			finishRequestWithErr(w, fmt.Sprintf("Cannot process input message, err: %v", err), http.StatusBadRequest)
			return
		}

		setStatusAndHeader(w, http.StatusOK)
		w.Write([]byte("[]"))
	}
}

func Read(app *app.App) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		logrus.Trace("Read handler")

		defer r.Body.Close()
		if r.Body == nil {
			finishRequestWithErr(w, emptyBody, http.StatusBadRequest)
			return
		}

		data := ReadRequest{}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			finishRequestWithErr(w, jsonErrMsg, http.StatusBadRequest)
			return
		}
		if len(data.Hash) == 0 {
			finishRequestWithErr(w, "Hash not found", http.StatusBadRequest)
			return
		}
		if len(data.PublicKey) == 0 {
			finishRequestWithErr(w, "PublicKey not found", http.StatusBadRequest)
			return
		}

		publicKey, err := base64.StdEncoding.DecodeString(data.PublicKey)
		if err != nil {
			finishRequestWithErr(w, "PublicKey cannot be decoded", http.StatusBadRequest)
			return
		}
		if len(publicKey) != 32 {
			finishRequestWithErr(w, "PublicKey length is wrong !=32", http.StatusBadRequest)
			return
		}

		encrypted, err := app.ProcessRead(r.Context(), data.Hash, publicKey)
		if err != nil {
			finishRequestWithErr(w, fmt.Sprintf("Cannot process read message, err: %v", err), http.StatusBadRequest)
			return
		}
		if encrypted == nil {
			// not found
			finishRequestWithWarn(w, "Message not found", http.StatusBadRequest)
			return
		}

		msgBase64 := base64.StdEncoding.EncodeToString(encrypted)
		message := ReadResponse{Message: msgBase64}

		setStatusAndHeader(w, http.StatusOK)
		w.Write([]byte(jsonStruct(message)))

	}
}

func Expired(config *config.Configuration) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		logrus.Trace("Expired handler")

		if err := config.Db.DeleteBeforeNow(r.Context()); err != nil {
			logrus.Errorf("Delete expired error: %v", err)
		} else {
			logrus.Info("Delete expired done")
		}

		setStatusAndHeader(w, http.StatusOK)
		w.Write([]byte("[]"))
	}
}
