package producer

import (
	"context"
	"fmt"
	"github.com/sajjad1993/todo/pkg/errs"
	"github.com/sajjad1993/todo/pkg/meesage_broker/command_utils"
)

func getKeyString(ctx context.Context, key string) (string, error) {
	hashValue := ctx.Value(key)
	hash := hashValue.(string)
	if hash == "" {
		return hash, errs.NewNoSuchKeyError(fmt.Sprintf("%s is not found in userWriter", command_utils.RequestHashKey))
	}
	return hash, nil
}
