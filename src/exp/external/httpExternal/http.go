package httpExternal

import (
	"encoding/json"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/exp/external/common"
	"github.com/KevinZonda/GoX/pkg/panicx"
	"io"
	"net/http"
	"net/url"
)

type Library struct {
	EndPoint string
	hc       http.Client
}

func NewLibrary(ep string) *Library {
	return &Library{
		EndPoint: ep,
	}
}

func (l *Library) AvailablePackage() []string {
	//TODO implement me
	panic("implement me")
}

func (l *Library) PackageInfo() map[string]common.PackageInfoElement {
	//TODO implement me
	panic("implement me")
}

func (l *Library) InvokeFunc(fxName string, args ...interface{}) common.InvokeResult {
	ep, _ := url.JoinPath(l.EndPoint, fxName)
	req, _ := http.NewRequest(http.MethodPost, ep, FuncCallRequest{
		FuncName: fxName,
		Args:     args,
	}.IO())
	rsp, err := l.hc.Do(req)
	panicx.NotNilErr(err)
	defer rsp.Body.Close()
	body, err := io.ReadAll(rsp.Body)
	panicx.NotNilErr(err)
	var rspT common.InvokeResult
	err = json.Unmarshal(body, &rspT)
	panicx.NotNilErr(err)

	return rspT
}

var _ common.Server = (*Library)(nil)
