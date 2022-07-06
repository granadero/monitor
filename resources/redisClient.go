package resources

type redisClient struct {
	Resource
}

func newRedisClient() iResource {
	return &redisClient{
		Resource: Resource{
			ResourceName: "AK47 gun",
			Fail:         4,
			Status:       200,
			Critical:     true,
			ResourceType: "serviceUrl",
		},
	}
}
