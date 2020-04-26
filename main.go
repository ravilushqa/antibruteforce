package main

import (
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq"
	"gitlab.com/ravilushqa/antibruteforce/cmd"
)

func main() {
	cmd.Execute()
}
