package rooms_api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nihal-ramaswamy/GoVid/internal/constants"
	"github.com/nihal-ramaswamy/GoVid/internal/db"
	"github.com/nihal-ramaswamy/GoVid/internal/dto"
	"github.com/nihal-ramaswamy/GoVid/internal/utils"
	"go.uber.org/zap"
)

type CreateRoomHandler struct {
	db          *sql.DB
	log         *zap.Logger
	middlewares []gin.HandlerFunc
}

func NewCreateRoomHandler(db *sql.DB, log *zap.Logger) *CreateRoomHandler {
	return &CreateRoomHandler{
		db:          db,
		log:         log,
		middlewares: []gin.HandlerFunc{},
	}
}

func (*CreateRoomHandler) Pattern() string {
	return "/create"
}

func (*CreateRoomHandler) RequestMethod() string {
	return constants.POST
}

func (c *CreateRoomHandler) Middlewares() []gin.HandlerFunc {
	return c.middlewares
}

func (c *CreateRoomHandler) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email := ctx.GetString("email")

		code := utils.NewCode(9)
		conference := dto.NewConference(code, email)
		err := db.CreateNewMeeting(c.db, conference)

		if nil != err {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"code": code,
		})
	}
}
