package permission

import (
	"net/http"

	"uam/services/api/internal/logic/permission"
	"uam/services/api/internal/svc"
	"uam/services/api/internal/types"
	"uam/tools/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeletePermHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeletePermReq
		if err := httpx.Parse(r, &req); err != nil {
			response.FailByArgsErr(w)
			return
		}

		l := permission.NewDeletePermLogic(r.Context(), svcCtx)
		_, err := l.DeletePerm(&req)
		if err != nil {
			response.FailBySvcErr(w, err.Error())
		} else {
			response.Ok(w)
		}
	}
}
