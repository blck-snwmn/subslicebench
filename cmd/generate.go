package main

import (
	"fmt"
	"os"

	"github.com/blck-snwmn/subslicebench/cmd/schema/fbs"
	flatbuffers "github.com/google/flatbuffers/go"
)

func main() {
	const elem = 10000
	positions := make([]*fbs.UserPositionT, 0, elem)
	for i := 0; i < elem; i++ {
		positions = append(positions, &fbs.UserPositionT{
			Name: "name",
			Pos: &fbs.PositionT{
				X: 1,
				Y: 2,
				Z: 3,
			},
		})
	}

	b := fbs.UserPositionsT{
		Poss: positions,
	}
	builder := flatbuffers.NewBuilder(200)
	builder.FinishSizePrefixed(b.Pack(builder))
	buf := builder.FinishedBytes()

	file, err := os.OpenFile("./testdata/data", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = file.Write(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully wrote to file")
}
