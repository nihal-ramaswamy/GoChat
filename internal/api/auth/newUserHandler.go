package auth_api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nihal-ramaswamy/GoVid/internal/constants"
	"github.com/nihal-ramaswamy/GoVid/internal/db"
	"github.com/nihal-ramaswamy/GoVid/internal/dto"
	"go.uber.org/zap"
)

type NewUserHandler struct {
	db  *sql.DB
	log *zap.Logger
}

func (*NewUserHandler) Pattern() string {
	return "/register"
}

func (n *NewUserHandler) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := dto.NewUser()

		if err := c.ShouldBindJSON(&user); err != nil {
			c.Error(err)
			n.log.Info("Responding with error", zap.Error(err))
			c.AbortWithStatus(http.StatusBadRequest)
		}

		if db.DoesEmailExist(n.db, user.Email) {
			c.JSON(http.StatusBadRequest,
				gin.H{"error": fmt.Sprintf("User with email %s already exists", user.Email)})

		}

		id := db.RegisterNewUser(n.db, user, n.log)

		c.JSON(http.StatusAccepted, gin.H{"id": id})

	}
}

func NewNewUserHandler(db *sql.DB, log *zap.Logger) *NewUserHandler {
	return &NewUserHandler{
		db:  db,
		log: log,
	}
}

func (*NewUserHandler) RequestMethod() string {
	return constants.POST
}
