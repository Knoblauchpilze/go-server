package routes

import (
	"net/http"
	"strings"

	"github.com/KnoblauchPilze/go-server/pkg/errors"
	"github.com/KnoblauchPilze/go-server/pkg/rest"
	"github.com/sirupsen/logrus"
)

var authenticationScheme = "bearer"

func authenticationCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// https://stackoverflow.com/questions/33265812/best-http-authorization-header-type-for-jwt
		// https://reqbin.com/req/5k564bhv/get-request-bearer-token-authorization-header-example
		// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Authorization
		reqData, ok := getRequestDataFromContextOrFail(w, r)
		if !ok {
			return
		}

		authData, err := rest.GetSingleHeaderFromHttpRequest(r, "Authorization")
		if err != nil {
			reqData.failWithErrorAndCode(err, http.StatusBadRequest, w)
			return
		}

		token, err := checkAndConsumeScheme(authData)
		if err != nil {
			reqData.failWithErrorAndCode(err, http.StatusBadRequest, w)
			return
		}

		logrus.Warnf("Should check token: %v", token)

		next.ServeHTTP(w, r)
	})
}

func checkAndConsumeScheme(authData string) (string, error) {
	tokens := strings.Split(authData, " ")
	if len(tokens) != 2 {
		return "", errors.New("invalid authentication header")
	}

	if strings.ToLower(tokens[0]) != authenticationScheme {
		return "", errors.New("invalid authentication scheme")
	}

	return tokens[1], nil
}
