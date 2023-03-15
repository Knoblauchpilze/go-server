package routes

import (
	"context"
	"net/http"

	"github.com/KnoblauchPilze/go-server/pkg/rest"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var requestIdDataKey stringDataKeyType = "requestId"

func requestCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.NewUUID()
		if err != nil {
			logrus.Errorf("Failed to generate uuid (%v)", err)
			code := http.StatusInternalServerError
			http.Error(w, http.StatusText(code), code)
			return
		}

		ctx := context.WithValue(r.Context(), requestIdDataKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func buildServerResponseFromHttpRequest(r *http.Request) rest.Response {
	var err error

	ctx := r.Context()
	id, ok := ctx.Value(requestIdDataKey).(uuid.UUID)
	if !ok {
		id, err = uuid.NewUUID()
		if err != nil {
			logrus.Errorf("Failed to generate request ID (%v)", err)
		}
	}

	return rest.NewSuccessResponse(id)
}
