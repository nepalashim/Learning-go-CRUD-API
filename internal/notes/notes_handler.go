package notes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo *Repo
}

func NewHandler(repo *Repo) *Handler {
	return &Handler{
		repo: repo,
	}
}

// hit /post--> tslks to repo to create note in db-->talk to database.
func (h *Handler) CreateNote(c *gin.Context) { //gin.context(everything about single http reuest and response) is a wrapper around http request and response writer. It provides us with methods to read request data and write response data.
	var req CreateNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil { //shouldbindjson- parse JSON into request Struct and validate the struct based on the binding you have passed.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdNote, err := h.repo.CreateNote(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdNote)
}
