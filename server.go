package simpleWeb

import "net/http"

type Option func(server *HTTPServer)

type HandlerFunc func()

type Server interface {
	http.Handler
	// Start 启动服务器
	Start(addr string) error
	// 注册路由
	addRoute(method string, path string, handler HandlerFunc)
}

type HTTPServer struct {
}

func NewHttpServer(opts ...Option) *HTTPServer {
	s := &HTTPServer{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (h *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("simple! simple! simple!"))
}

func (h *HTTPServer) Start(addr string) error {
	return http.ListenAndServe(addr, h)
}

func (h *HTTPServer) addRoute(method string, path string, handler HandlerFunc) {
	//TODO implement me
	panic("implement me")
}
