package main

import (
	"context"
	"fmt"

	"covid/config"
	"covid/internal/adapters/covidadt"
	"covid/internal/core/service/covidsvc"
)

func main() {
	config.Init()
	ctx := context.Background()
	ca := covidadt.New()
	csvc := covidsvc.New(ca)
	out, err := csvc.GetSummary(ctx)
	fmt.Println(out, err)
}
