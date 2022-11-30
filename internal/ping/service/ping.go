package service

type PingService struct {
}

func (p PingService) Ping() string {
	return "hello world"
}

func NewPingService() PingService {
	return PingService{}
}
