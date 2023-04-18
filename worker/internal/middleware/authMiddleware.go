package middleware

import (
	"bytes"
	"io"
	"k2edge/auth"
	"k2edge/worker/internal/config"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type AuthMiddleware struct {
	auther auth.Auther
}

func NewAuthMiddleware(c config.Config, etcd *clientv3.Client) *AuthMiddleware {
	return &AuthMiddleware{
		auther: auth.NewEtcdAuther(c.Name, etcd),
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("K2EDGE-AUTH-TOKEN")
		logx.Debugf("token=%s", token)
		if len(token) != 0 && !m.auther.CheckToken(token) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if len(token) == 0 {
			secret := r.Header.Get("K2EDGE-NODE-SECRET")
			if len(secret) == 0 {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			logx.Debugf("auth get K2EDGE-NODE-SECRET, secret=%s", secret)
			token = m.auther.GenerateToken(secret)
			w.Header().Add("K2EDGE-AUTH-TOKEN", token)
		}
		logx.Debugf("auth get K2EDGE-AUTH-TOKEN, token=%s", token)
		body := r.Body
		bodyBytes, err := io.ReadAll(body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		r.Body = io.NopCloser(
			bytes.NewBuffer(
				m.auther.Decode(token, bodyBytes),
			),
		)
		next(
			auth.NewEncodeResponseWriter(token, m.auther, w),
			r,
		)
	}
}
