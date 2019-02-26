package ginsession

import (
	"context"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/go-session/session"
)

type (
	// ErrorHandleFunc error handling function
	ErrorHandleFunc func(*gin.Context, error)
	// Config defines the config for Session middleware
	Config struct {
		// error handling when starting the session
		ErrorHandleFunc ErrorHandleFunc
		// keys stored in the context
		StoreKey string
	}
)

var (
	once            sync.Once
	internalManager *session.Manager
	storeKey        string

	// DefaultConfig is the default Recover middleware config.
	DefaultConfig = Config{
		ErrorHandleFunc: func(ctx *gin.Context, err error) {
			ctx.AbortWithError(500, err)
		},
		StoreKey: "github.com/go-session/gin-session",
	}
)

func manager(opt ...session.Option) *session.Manager {
	once.Do(func() {
		internalManager = session.NewManager(opt...)
	})
	return internalManager
}

// New create a session middleware
func New(opt ...session.Option) gin.HandlerFunc {
	return NewWithConfig(DefaultConfig, opt...)
}

// NewWithConfig create a session middleware
func NewWithConfig(config Config, opt ...session.Option) gin.HandlerFunc {
	if config.ErrorHandleFunc == nil {
		config.ErrorHandleFunc = DefaultConfig.ErrorHandleFunc
	}

	storeKey = config.StoreKey
	if storeKey == "" {
		storeKey = DefaultConfig.StoreKey
	}

	return func(ctx *gin.Context) {
		store, err := manager(opt...).Start(context.Background(), ctx.Writer, ctx.Request)
		if err != nil {
			config.ErrorHandleFunc(ctx, err)
			return
		}
		ctx.Set(storeKey, store)
		ctx.Next()
	}
}

// FromContext Get session storage from context
func FromContext(ctx *gin.Context) session.Store {
	return ctx.MustGet(storeKey).(session.Store)
}

// Destroy a session
func Destroy(ctx *gin.Context) error {
	return manager().Destroy(context.Background(), ctx.Writer, ctx.Request)
}

// Refresh a session and return to session storage
func Refresh(ctx *gin.Context) (session.Store, error) {
	return manager().Refresh(context.Background(), ctx.Writer, ctx.Request)
}
