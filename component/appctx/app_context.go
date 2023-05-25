package appctx

import "gorm.io/gorm"

// đồ chơi nội bộ
type AppContext interface {
	GetMainDBConnection() *gorm.DB
	SecretKey() string
}
type appCtx struct {
	db        *gorm.DB
	secretKey string
}

func NewAppContext(db *gorm.DB, secretKey string) *appCtx {
	return &appCtx{db: db, secretKey: secretKey}
}
func (ctx *appCtx) GetMainDBConnection() *gorm.DB { return ctx.db }
func (ctx *appCtx) SecretKey() string             { return ctx.secretKey }
