package main

import (
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq"
	"gitlab.com/otus_golang/antibruteforce/cmd"
)

func main() {
	cmd.Execute()
}
