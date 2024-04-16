package server

import (
	"bytes"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"io"
	v1 "nft-market-backend/api/moralis/v1"
	"nft-market-backend/internal/conf"
)

var ERROR_UNAUTH = errors.New(403, "error signature", "error signature")

func signatureMiddleware(conf *conf.Moralis) middleware.Middleware {

	return selector.Server(
		func(handler middleware.Handler) middleware.Handler {
			return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
				if tr, ok := transport.FromServerContext(ctx); ok {
					token := tr.RequestHeader().Get("x-signature")
					fmt.Println("signature", token)
					fmt.Println("signature", conf.Secret)
					if token == "" {
						err = ERROR_UNAUTH
						return
					}
					if ht, ok := tr.(*http.Transport); ok {
						request := ht.Request()
						if request.Method == "POST" {
							data, _ := io.ReadAll(request.Body)
							request.Body = io.NopCloser(bytes.NewBuffer(data))
							if len(data) > 0 {
								fmt.Println("signature", string(data))
							}
						}
					}
				}
				reply, err = handler(ctx, req)
				return
			}
		},
	).Match(NewSignatureListMatcher()).Build()
}

// NewWhiteListMatcher 创建签名验证接口名单
func NewSignatureListMatcher() selector.MatchFunc {
	signatureList := make(map[string]struct{})
	//签名验证接口
	signatureList[v1.OperationStreamReceive] = struct{}{}

	return func(ctx context.Context, operation string) bool {
		if _, ok := signatureList[operation]; ok {
			return true
		}
		return false
	}
}
