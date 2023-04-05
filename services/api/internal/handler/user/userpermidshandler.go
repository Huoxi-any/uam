package user

import (
	"net/http"

	"uam/services/api/internal/logic/user"
	"uam/services/api/internal/svc"
	"uam/services/api/internal/types"
	"uam/tools/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserPermIdsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserPermIdsReq
		if err := httpx.Parse(r, &req); err != nil {
			response.FailByArgsErr(w)
			return
		}

		l := user.NewUserPermIdsLogic(r.Context(), svcCtx)
		resp, err := l.UserPermIds(&req)
		if err != nil {
			response.FailBySvcErr(w, err.Error())
		} else {
			response.OkWithData(w, resp.List)
		}
	}
}
