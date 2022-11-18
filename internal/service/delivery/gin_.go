package delivery

import (
	"context"

	"github.com/vldcreation/simple-auth-golang/internal/feature"
)

type GinObject struct {
	features
}

type features struct {
	feature.SetupUser
}

func NewGinHandler(ctx context.Context, f features) {
	obj := &GinObject{f}
	obj.InitRoutes(ctx)
}
