package middleware

import (
	"be-dse/datastruct"
	"encoding/json"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Header.Get("X-API-Key") == "RAHASIA" {
		// ok
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		// error
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		// webResponse := web.WebResponse{
		// 	Code:   http.StatusUnauthorized,
		// 	Status: "UNAUTHORIZED",
		// }
		var response datastruct.ResponseGetAll
		response.Status = http.StatusUnauthorized
		response.Message = "UNAUTHORIZED"

		json.NewEncoder(writer).Encode(response)
		// helper.WriteToResponseBody(writer, webResponse)
	}
}
