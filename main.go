package main

import (
	"fmt"
	"os"

	"github.com/gogo/protobuf/proto"
	"github.com/knights-analytics/onnx-gomlx/onnx/ir"
)

func main() {
	// read the onnx model
	b, err := os.ReadFile("models/sentence-transformers_all-MiniLM-L6-v2/onnx/model.onnx")
	if err != nil {
		panic(err)
	}

	// decode into protobuf
	pbModel := &ir.ModelProto{}
	err = proto.Unmarshal(b, pbModel)
	if err != nil {
		panic(err)
	}

	// look at the graph
	graph := pbModel.Graph

	count := 0

	for _, n := range graph.Node {
		fmt.Println("name is", n.GetName())
		fmt.Println("optype is ", n.GetOpType())
		fmt.Println("input is ", n.GetInput())
		fmt.Println("attrs are", n.GetAttribute())
		count++
		if count >= 10 {
			break
		}
	}
}
