package resources

type iResource interface {
	setResource(name string)
	setStatus(code int)
	setFail(counter int)
	getResource() string
	getStatus() int
	getFail() int
	setResourceType(rType string)
	setCritical(critical bool)
	getResourceType() string
	getCritical() bool
	setHandler(handler func())
	getHandler() func()
}
