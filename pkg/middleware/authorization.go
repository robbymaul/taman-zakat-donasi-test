package middleware

import (
	pkgjwt "donasitamanzakattest/pkg/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *Auth) Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		value, exists := ctx.Get(BearerToken)
		if !exists {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		response := value.(*pkgjwt.JwtResponse)

		ctx.Set(Session, response)
		ctx.Next()
	}
}
