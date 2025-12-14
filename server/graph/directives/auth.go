package directives

import (
	"context"

	"github.com/brightsidedeveloper/go-native-template/internal/auth"

	"github.com/99designs/gqlgen/graphql"
)

func AuthDirective(ctx context.Context, obj any, next graphql.Resolver) (res any, err error) {
	_, err = auth.GetUserFromContext(ctx)
	if err != nil {
		return nil, auth.Unauthorized(ctx)
	}
	return next(ctx)
}
