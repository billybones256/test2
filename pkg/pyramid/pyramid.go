package pyramid

import (
	"context"
	"strings"
	"task2/pkg/api"
)

//GRPCServer ...
type GRPCServer struct{}

//Square ...
func (server *GRPCServer) BuildPyramid(ctx context.Context, req *api.PyramidRequest) (*api.PyramidResponse, error) {
	ret := Pyramid(int(req.GetH()))
	return &api.PyramidResponse{Answer: ret}, nil
}

func Pyramid(n int) string {
	var result string
	for i := 0; i < n; i++ {
		result += strings.Repeat(" ", n-(i+1))
		result += strings.Repeat("@", (i*2)+1)
		result += "\n"
	}
	return strings.TrimRight(result, "\n")
}
