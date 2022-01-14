package dm

import (
	"net/http"

	"github.com/go-things/things/src/webapi/internal/logic/backgrand/dm"
	"github.com/go-things/things/src/webapi/internal/svc"
	"github.com/go-things/things/src/webapi/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func GetProductTemplateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetProductTemplateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := dm.NewGetProductTemplateLogic(r.Context(), svcCtx)
		resp, err := l.GetProductTemplate(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}