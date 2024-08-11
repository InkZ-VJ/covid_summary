package main

import (
	"context"
	"fmt"

	"covid/config"
	"covid/internal/adapters/covidadt"
)

func main() {
	config.Init()
	ctx := context.Background()
	ca := covidadt.New()
	out, err := ca.GetCovidStat(ctx)
	fmt.Println(out, err)
}
