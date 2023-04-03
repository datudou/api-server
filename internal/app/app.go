package app

import (
	"github.com/retail-ai-test/internal/app/initialize"
)

func Run() {
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()
	initialize.InitServer()
}
