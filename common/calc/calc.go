package calc

//Grpc请求返回数据结构
type GrpcResponse struct {
	ErrNo  int64       `json:"errNo"`
	ErrStr string      `json:"errStr"`
	Data   interface{} `json:"data"`
}