package app

import (
	"context"

	"github.com/vldcreation/simple-auth-golang/internal/constants"
	"github.com/vldcreation/simple-auth-golang/internal/feature"
	setup_user "github.com/vldcreation/simple-auth-golang/internal/feature/account_creation"
	"github.com/vldcreation/simple-auth-golang/internal/service/delivery"
)

// Not complete yet, currently we are declaring service depends on feature level (using new function)
// We supposed to register all service in init function(soon)
func (ox *App) initService(ctx context.Context) error {
	featureSetupUser := setup_user.New(
		setup_user.Configuration{},
		setup_user.Dependency{
			Postgresql: constants.DB,
		},
	)

	delivery.NewGinHandler(ctx,
		struct {
			feature.SetupUser
		}{
			featureSetupUser,
		})

	return nil
}
