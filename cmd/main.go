package cmd

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/tusmasoma/clean-architecture-campfinder/driver"
)

func main() {
	// .envファイルから環境変数を読み込む
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	/* ===== サーバの起動 ===== */
	driver.Run()
}
