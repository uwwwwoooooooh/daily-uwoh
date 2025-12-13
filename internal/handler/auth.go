package handler

import (
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	// service service.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Register(c *gin.Context) {
	// TODO
}

func (h *AuthHandler) Login(c *gin.Context) {
	// TODO
}
