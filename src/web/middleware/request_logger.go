package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

func RequestLogger(r *gin.Engine)  {
	log := logrus.New()
	r.Use(ginlogrus.Logger(log))
}
