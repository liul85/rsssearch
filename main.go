package main

import (
	"os"
	"log"
	"github.com/liul85/goinaction/search/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}