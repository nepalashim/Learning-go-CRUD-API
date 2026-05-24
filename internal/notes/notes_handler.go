package notes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Handler struct {
	repo *Repo
}

func NewHandler(repo *Repo) *Handler {
	return &Handler{
		repo: repo,
	}
}

// hit /post--> talks to repo to create note in db-->talk to database.
func (h *Handler) CreateNote(c *gin.Context) { //gin.context(everything about single http reuest and response) is a wrapper around http request and response writer. It provides us with methods to read request data and write response data.
	var req CreateNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil { //shouldbindjson- parse JSON into request Struct and validate the struct based on the binding you have passed.
		c.JSON(http.StatusBadRequest, gin.H{"error; Invalid JSON": err.Error()})
		return
	}
	now := time.Now()
	note := Note{
		ID:        bson.NewObjectID(),
		Title:     req.Title,
		Content:   req.Content,
		Pinned:    req.Pinned,
		CreatedAt: now,
		UpdatedAt: now,
	}

	created, err := h.repo.CreateNote(c.Request.Context(), note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}

func (h *Handler) ListNotes(c *gin.Context) {
	notes, err := h.repo.ListNotes(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error: Failed to Fetch all notes": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"notes": notes})
}

func (h *Handler) GetNoteByID(c *gin.Context) {
	idstr := c.Param("id")
	//convert 24 character hex string to bson object id
	noteID, err := bson.ObjectIDFromHex(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}
	note, err := h.repo.GetByID(c.Request.Context(), noteID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}
	c.JSON(http.StatusOK, note)
}

func (h *Handler) UpdateNoteByID(c *gin.Context) {
	idstr := c.Param("id")
	objID, err := bson.ObjectIDFromHex(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}
	var req UpdateNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	updated, err := h.repo.UpdateNote(c.Request.Context(), objID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (h *Handler) DeleteNoteByID(c *gin.Context) {
	idstr := c.Param("id")
	objID, err := bson.ObjectIDFromHex(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}
	deleted, err := h.repo.DeleteNote(c.Request.Context(), objID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Note deleted successfully"})
}