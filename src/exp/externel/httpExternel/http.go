package httpExternel

import (
	"encoding/json"
	"git.cs.bham.ac.uk/projects-2023-24/xxs166/src/exp/externel"
	"github.com/KevinZonda/GoX/pkg/panicx"
	"io"
	"net/http"
	"net/url"
)

type Library struct {
	EndPoint string
	hc       http.Client
}

func (l *Library) AvailablePackage() []string {
	//TODO implement me
	panic("implement me")
}

func (l *Library) PackageInfo() map[string]externel.PackageInfoElement {
	//TODO implement me
	panic("implement me")
}

func (l *Library) InvokeFunc(fxName string, args ...interface{}) externel.InvokeResult {
	ep, _ := url.JoinPath(l.EndPoint, fxName)
	req, _ := http.NewRequest(http.MethodPost, ep, FuncCallRequest{
		FuncName: fxName,
		Args:     args,
	}.IO())
	rsp, err := l.hc.Do(req)
	panicx.PanicIfNotNil(err, err)
	defer rsp.Body.Close()
	body, err := io.ReadAll(rsp.Body)
	panicx.PanicIfNotNil(err, err)
	var rspT externel.InvokeResult
	err = json.Unmarshal(body, &rspT)
	panicx.PanicIfNotNil(err, err)

	return rspT
}

var _ externel.Server = (*Library)(nil)
