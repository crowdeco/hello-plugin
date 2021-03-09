package main

import (
	"encoding/json"
	"net/http"

	"github.com/crowdeco/bima/v2/configs"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var BimaPluginName = "Hello"

var BimaPlugin = New()

type hello struct {
}

func New() hello {
	return hello{}
}

func (h hello) GetRoutes() []configs.Route {
	return []configs.Route{
		&endPoint{},
	}
}

func (h hello) GetLoggers() []logrus.Hook {
	return nil
}

func (h hello) GetMiddlewares() []configs.Middleware {
	return nil
}

func (h hello) GetListeners() []configs.Listener {
	return nil
}

func (h hello) GetServers() []configs.Server {
	return nil
}

func (h hello) GetUpgrades() []configs.Upgrade {
	return nil
}

func (h hello) GetVersion() string {
	return "v1.0"
}

type endPoint struct {
}

func (x endPoint) Path() string {
	return "/hello"
}

func (x endPoint) Method() string {
	return http.MethodGet
}

func (x endPoint) SetClient(client *grpc.ClientConn) {
}

func (x endPoint) Middlewares() []configs.Middleware {
	return nil
}

func (x endPoint) Handle(w http.ResponseWriter, r *http.Request, params map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello from Plugin",
	})
}

func main() {}
