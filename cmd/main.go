package main

import (
	"fmt"
	"log"

	"github.com/Wexlersolk/dream2/handlers/env"
)

func main() {
	cfg := config{
		readFilePath:  env.GetString("READ_FILE_PATH", "/home/wexlersolk/work/dream2/file/Week.wiki"),
		writeFilePath: env.GetString("WRITE_FILE_PATH", "/home/wexlersolk/work/dream2/file/dream.json"),
	}

	fmt.Printf("rfile: %s\n", cfg.readFilePath)
	fmt.Printf("wfile: %s\n", cfg.writeFilePath)

	err := run(cfg)
	if err != nil {
		log.Fatal(err)
	}

}
