package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func HelloGET(c *gin.Context) {
    c.String(http.StatusOK, "hello chika")
}