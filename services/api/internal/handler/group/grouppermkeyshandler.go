package group

import (
	"net/http"

	"uam/services/api/internal/logic/group"
	"uam/services/api/internal/svc"
	"uam/services/api/internal/types"
	"uam/tools/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GroupPermKeysHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupPermKeysReq
		if err := httpx.Parse(r, &req); err != nil {
			response.FailByArgsErr(w)
			return
		}

		l := group.NewGroupPermKeysLogic(r.Context(), svcCtx)
		resp, err := l.GroupPermKeys(&req)
		if err != nil {
			response.FailBySvcErr(w, err.Error())
		} else {
			response.OkWithData(w, resp.List)
		}
	}
}
