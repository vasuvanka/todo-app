package middlewares

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/vasuvanka/todo-app/backend/models"
	"github.com/vasuvanka/todo-app/backend/shared"
)

// JwtValidatorWrap - jwt validator middleware
func JwtValidatorWrap(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		token, err := shared.GetTokenFromHeader(r)
		if err != nil {
			shared.SendError(w, models.Response{
				Status:  http.StatusUnauthorized,
				Message: err.Error(),
			})
			return
		}
		claims, err := shared.ValidateJWT(token)
		if err != nil {
			shared.SendError(w, models.Response{
				Status:  http.StatusUnauthorized,
				Message: err.Error(),
			})
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, shared.KeyClaims, claims)
		r = r.WithContext(ctx)
		h(w, r, p)
	}
}
