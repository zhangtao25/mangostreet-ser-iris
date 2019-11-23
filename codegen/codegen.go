package main

import (
	"github.com/mlogclub/simple"

	"mangostreet-ser-iris/model"
)

func main() {
	simple.Generate("./", "github.com/mlogclub/bbs-go", simple.GetGenerateStruct(&model.ThirdAccount{}))
}
