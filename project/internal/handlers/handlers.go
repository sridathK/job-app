package handlers

import (
	"project/internal/auth"
	"project/internal/middlewear"
	"project/internal/services"

	"github.com/gin-gonic/gin"
)

func Api(a *auth.Auth, s services.Service) *gin.Engine {
	r := gin.New()
	h := handler{a: a, us: s}
	m, _ := middlewear.NewMiddleWear(a)
	r.Use(m.Log(), gin.Recovery())
	r.POST("/signup", h.userSignin)
	r.POST("/login", h.userLoginin)
	r.POST("/createCompany", h.companyCreation)
	r.GET("/getAllCompany", h.getAllCompany)
	r.GET("/getCompany/:id", h.getCompany)
	return r
}
