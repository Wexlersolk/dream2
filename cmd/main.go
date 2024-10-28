package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Wexlersolk/dream2/handlers/env"
)

func main() {

	operation := flag.String("operation", string(None), "Specify the operation: Right, Left")
	flag.Parse()

	cfg := config{
		readFilePath:  env.GetString("READ_FILE_PATH", "/home/wexlersolk/work/dream2/file/Week.wiki"),
		writeFilePath: env.GetString("WRITE_FILE_PATH", "/home/wexlersolk/work/dream2/file/dream.json"),
		testFile:      env.GetString("WRITE_FILE_PATH", "/home/wexlersolk/work/dream2/file/test.json"),
		testWriteFile: env.GetString("WRITE_FILE_PATH", "/home/wexlersolk/work/dream2/file/Additional.wiki"),
		graphFilePath: env.GetString("GRAPH_FILE_PATH", "/home/wexlersolk/work/dream2/file/graph.png"),
		daysToDisplay: env.GetInt("DAYS_TO_DISPLAY", 100),
		Operation:     Operation(*operation),
	}

	fmt.Printf("rfile: %s\n", cfg.readFilePath)
	fmt.Printf("wfile: %s\n", cfg.writeFilePath)
	fmt.Printf("Operation: %s\n", cfg.Operation)

	err := run(cfg)
	if err != nil {
		log.Fatal(err)
	}

}
