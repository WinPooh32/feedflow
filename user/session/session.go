package session

import (
	"log"

	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	ginsessionstore "github.com/go-session/session"
)

//Field is string enum type for session keys
type Field string

//Session storage fields
const (
	FieldUserID   Field = "user_id"
	FieldRoleID   Field = "user_role"
	FieldPageHits Field = "page_hits"
)

//Session -
type Session struct {
	store ginsessionstore.Store
}

func logFieldError(f Field) {
	log.Println("<Session>: can not get field from storage:", f)
}

//New - Creates new session wrapper
func New(ctx *gin.Context) *Session {
	return &Session{
		store: ginsession.FromContext(ctx),
	}
}

//Commit changes to storage
func (s *Session) Commit() error {
	return s.store.Save()
}

//get value from session storage
func (s *Session) get(key Field) (interface{}, bool) {
	v, ok := s.store.Get(string(key))
	return v, ok
}

func (s *Session) getFloat64(key Field) (float64, bool) {
	v, ok := s.store.Get(string(key))
	if value, isValid := v.(float64); ok && isValid {
		return value, isValid
	}
	return 0, false
}

func (s *Session) getFloat32(key Field) (float32, bool) {
	v, ok := s.getFloat64(key)
	if ok {
		return float32(v), ok
	}
	return 0, false
}

func (s *Session) getInt64(key Field) (int64, bool) {
	v, ok := s.getFloat64(key)
	if ok {
		return int64(v), ok
	}
	return 0, false
}

func (s *Session) getInt32(key Field) (int32, bool) {
	v, ok := s.getFloat64(key)
	if ok {
		return int32(v), ok
	}
	return 0, false
}

func (s *Session) getString(key Field) (string, bool) {
	v, ok := s.store.Get(string(key))
	if value, isValid := v.(string); ok && isValid {
		return value, isValid
	}
	return "", false
}

//Set value to session storage as float64 number
func (s *Session) setNumber(key Field, value float64) {
	s.store.Set(string(key), value)
}

//Set value to session storage as string
func (s *Session) setString(key Field, value string) {
	s.store.Set(string(key), value)
}
