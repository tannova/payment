package main

import (
	_ "github.com/go-sql-driver/mysql"
	"gitlab.com/greyhole/night-kit/pkg/config"

	"gitlab.com/mcuc/monorepo/backend/natasha/internal/server"
)

func main() {
	flags := config.ParseFlags()
	server.Run(flags)
}
