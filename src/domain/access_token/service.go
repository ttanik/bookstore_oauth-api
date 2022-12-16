package access_token

import (
	"github.com/ttanik/bookstore_oauth-api/utils/errors"
)

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}
