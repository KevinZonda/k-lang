package common

type Server interface {
	AvailablePackage() []string
	PackageInfo() map[string]PackageInfoElement
	InvokeFunc(method string, args ...interface{}) InvokeResult
}
