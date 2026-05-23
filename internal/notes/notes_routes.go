package notes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func RegisterRoutes(r *gin.Engine, db *mongo.Database) {
	//create repo and Handler once at starter
	repo := NewRepo(db)
	handler := NewHandler(repo)
	// /notes .post"" --> POST --> create note,  but if /notes and inside it: .post"/add"--> "/notes/add"--> POST --> create note
	notesGroup := r.Group("/notes")
	notesGroup.POST("", handler.CreateNote)
}