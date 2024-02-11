package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rmrcunha/Oportunity.git/schemas"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "eror listing openings")
	}

	sendSuccess(ctx, "list-openings", openings)
}
