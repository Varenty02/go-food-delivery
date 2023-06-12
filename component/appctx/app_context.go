package appctx

import (
	"g05-fooddelivery/pubsub"
	"gorm.io/gorm"
)

// đồ chơi nội bộ
type AppContext interface {
	GetMainDBConnection() *gorm.DB
	SecretKey() string
	GetPubSub() pubsub.Pubsub
}
type appCtx struct {
	db        *gorm.DB
	secretKey string
	ps        pubsub.Pubsub
}

func NewAppContext(db *gorm.DB, secretKey string, ps pubsub.Pubsub) *appCtx {
	return &appCtx{db: db, secretKey: secretKey, ps: ps}
}
func (ctx *appCtx) GetMainDBConnection() *gorm.DB { return ctx.db }
func (ctx *appCtx) SecretKey() string             { return ctx.secretKey }
func (ctx *appCtx) GetPubSub() pubsub.Pubsub      { return ctx.ps }
