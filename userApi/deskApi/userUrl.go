package deskApi

type userUrlConfig struct {
	userUrl string
}

func NewUserUrlFunc() *userUrlConfig {
	return &userUrlConfig{
		userUrl: "https://sit-webapi.mggametransit.com",
	}
}
