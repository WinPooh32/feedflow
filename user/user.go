package user

import (
	"log"

	"github.com/WinPooh32/feedflow/model"

	"github.com/jinzhu/gorm"

	"github.com/WinPooh32/feedflow/database"
	"github.com/WinPooh32/feedflow/user/previlegies"
	"github.com/WinPooh32/feedflow/user/session"
	"github.com/gin-gonic/gin"
)

type User struct {
	sess *session.Session
	db   *gorm.DB

	record *model.SigninRequest
}

func New(ctx *gin.Context) *User {
	db, ok := database.FromContext(ctx)
	if !ok {
		log.Println("There is no database in current context")
	}

	return &User{
		sess: session.New(ctx),
		db:   db,
	}
}

func logErr(err error) {
	log.Println("<User>:", err)
}

//Find record in database
func (u *User) Find() bool {
	u.record = &model.SigninRequest{
		Base: model.Base{ID: uint64(u.sess.GetUserID())},
	}

	if err := u.db.First(u.record).Error; err != nil {
		logErr(err)
		return false
	}

	return true
}

//Upgrade user session role
func (u *User) SessionUpgrade(record *model.SigninRequest, role previlegies.Role) {
	u.sess.SetUserID(int64(record.ID))
	u.sess.SetUserRole(role)
}

//Hit increments page hits counter
func (u *User) SessionHit() {
	u.sess.SetHits(u.sess.GetHits() + 1)
}

func (u *User) SessionGetHits() int64 {
	return u.sess.GetHits()
}

func (u *User) SessionGetID() int64 {
	return u.sess.GetUserID()
}

func (u *User) SessionSave() {
	if err := u.sess.Commit(); err != nil {
		logErr(err)
	}
}
