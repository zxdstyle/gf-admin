package request

import "github.com/gogf/gf/v2/text/gstr"

type (
	HasWith struct {
		With  string `json:"with"`
		_with []string
	}
)

func (req *HasWith) GetWith() []string {
	if len(req.With) == 0 {
		return []string{}
	}
	if len(req._with) == 0 {
		with := gstr.Split(req.With, ",")
		for i, _ := range with {
			with[i] = gstr.UcFirst(with[i])
		}
		req._with = with
	}
	return req._with
}
