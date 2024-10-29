package main

import (
	"context"
	"os"

	"github.com/happycrud/icons/material"
)

func main() {
	material.AddTask("40", "red", "m-1").Render(context.Background(), os.Stdout)
}
