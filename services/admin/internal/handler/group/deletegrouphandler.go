package group

import (
	"net/http"

	"uam/services/admin/internal/logic/group"
	"uam/services/admin/internal/svc"
	"uam/services/admin/internal/types"
	"uam/tools/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteGroupHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteGroupReq
		if err := httpx.Parse(r, &req); err != nil {
			response.FailByArgsErr(w)
			return
		}

		l := group.NewDeleteGroupLogic(r.Context(), svcCtx)
		_, err := l.DeleteGroup(&req)
		if err != nil {
			response.FailBySvcErr(w, err.Error())
		} else {
			response.Ok(w)
		}
	}
}
