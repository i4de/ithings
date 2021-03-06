package middleware

import (
	"context"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/apisvr/internal/types"
	"github.com/i-Things/things/src/usersvr/user"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type CheckTokenMiddleware struct {
	UserRpc user.User
}

func NewCheckTokenMiddleware(UserRpc user.User) *CheckTokenMiddleware {
	return &CheckTokenMiddleware{UserRpc: UserRpc}
}

func (m *CheckTokenMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		strIP, _ := utils.GetIP(r)
		r.Header.Set(types.USER_IP, strIP)
		strToken := r.Header.Get(types.USER_TOKEN)
		if strToken == "" {
			logx.WithContext(r.Context()).Errorf("%s|CheckToken|ip=%s|not find token",
				utils.FuncName(), strIP)
			http.Error(w, errors.TokenMalformed.Error(), http.StatusUnauthorized)
			return
		}
		resp, err := m.UserRpc.CheckToken(r.Context(), &user.CheckTokenReq{
			Ip:    strIP,
			Token: strToken,
		})
		if err != nil {
			er := errors.Fmt(err)
			logx.WithContext(r.Context()).Errorf("%s|CheckToken|ip=%s|token=%s|return=%s",
				utils.FuncName(), strIP, strToken, err)
			http.Error(w, er.Error(), http.StatusUnauthorized)
			return
		}
		if resp.Token != "" {
			w.Header().Set(types.USER_SET_TOKEN, resp.Token)
		}
		logx.WithContext(r.Context()).Infof("CheckToken|ip=%s|uid=%s|token=%s|newToken=%s",
			strIP, resp.Uid, strToken, resp.Token)
		ctx := context.WithValue(r.Context(), types.USER_UID, &types.UserCtx{
			Uid: resp.Uid,
		})
		next(w, r.WithContext(ctx))
	}
}
