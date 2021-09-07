package main

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
	"github.com/ravilushqa/antibruteforce/cmd"
)

func main() {
	cmd.Execute()
}
