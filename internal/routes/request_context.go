package routes

import (
	"context"
	"net/http"

	"github.com/KnoblauchPilze/go-server/pkg/rest"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type requestData struct {
	id       uuid.UUID
	response rest.Response
}

var requestDataKey stringDataKeyType = "requestData"

func newRequestData() requestData {
	var out requestData

	out.id, _ = uuid.NewUUID()
	out.response = rest.NewSuccessResponse(out.id)

	return out
}

func (rd requestData) failWithErrorAndCode(err error, code int, w http.ResponseWriter) {
	rd.response.WithCode(code)
	rd.response.WithDetails(err)
	rd.response.Write(w)
}

func (rd requestData) writeDetails(details interface{}, w http.ResponseWriter) {
	rd.response.WithDetails(details)
	rd.response.Write(w)
}

func requestCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rd := newRequestData()

		ctx := context.WithValue(r.Context(), requestDataKey, rd)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getRequestDataFromContextOrFail(w http.ResponseWriter, r *http.Request) (requestData, bool) {
	ctx := r.Context()
	reqData, ok := ctx.Value(requestDataKey).(requestData)
	if !ok {
		logrus.Errorf("Failed to get request data")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	return reqData, ok
}
