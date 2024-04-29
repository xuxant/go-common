package errors

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/rs/zerolog/log"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func GQLErrPresenter(ctx context.Context, e error) *gqlerror.Error {
	err := graphql.DefaultErrorPresenter(ctx, e)
	var er *Error

	if errors.As(e, &er) {
		err.Message = er.Message
		err.Extensions = map[string]interface{}{
			"code": er.Code,
		}
	} else {
		err.Message = Internal.Message
		err.Extensions = map[string]interface{}{
			"code": Internal.Code,
		}
	}
	log.Err(err.Unwrap()).Msgf("error executing graphql query: %s", err.Message)
	return err
}
