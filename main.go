package main

import (
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq"
	"github.com/ravilushqa/antibruteforce/cmd"
)

func main() {
	cmd.Execute()
}
