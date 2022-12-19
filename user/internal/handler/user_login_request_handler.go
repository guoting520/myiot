package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"myiot/user/internal/logic"
	"myiot/user/internal/svc"
	"myiot/user/internal/types"
)

func UserLoginRequestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserLoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserLoginRequestLogic(r.Context(), svcCtx)
		resp, err := l.UserLoginRequest(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
