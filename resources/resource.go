package resources

type Resource struct {
	ResourceName string
	Status       int
	Fail         int
	ResourceType string
	Handler      func()
	Critical     bool
}

func (r *Resource) setResource(name string) {
	r.resourceName = name
}

func (r *Resource) getResource() string {
	return r.resourceName
}

func (r *Resource) setResourceType(name string) {
	r.resourceName = name
}

func (r *Resource) getResourceType() string {
	return r.resourceType
}

func (r *Resource) setFail(fail int) {
	r.fail = fail
}

func (r *Resource) getFail() int {
	return r.fail
}

func (r *Resource) setStatus(status int) {
	r.status = status
}

func (r *Resource) getStatus() int {
	return r.status
}

func (r *Resource) setCritical(critical bool) {
	r.critical = critical
}

func (r *Resource) getCritical() bool {
	return r.critical
}

func (r *Resource) setHandler(handler func()) {
	r.handler = handler
}

func (r *Resource) getHandler() func() {
	return r.handler
}
