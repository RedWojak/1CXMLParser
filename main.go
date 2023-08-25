package main

import (
	"IMAXMLParser/redis"
	"IMAXMLParser/xmlParser"
	"flag"
	"log"
	"os"
	"time"
)

// Define the flag
var help = flag.Bool("help", false, "Show help")
var xmlPath = "./" +
	""
var redisIP = "localhost"
var redisPassword = ""
var redisPort = "6379"

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func main() {
	// Bind the flag
	flag.StringVar(&xmlPath, "path", xmlPath, "Path to XML Documents folder")
	flag.StringVar(&redisIP, "redisAddress", redisIP, "Redis IP flag")
	flag.StringVar(&redisPort, "redisPort", "6379", "Redis Port Flag")
	flag.StringVar(&redisPassword, "redisPassword", redisPassword, "Redis Password")

	// Parse the flag
	flag.Parse()

	// Usage Demo
	if *help {
		flag.Usage()
		os.Exit(0)
	}

	redis.Redis = redis.RedisNewClient(redisIP, redisPort, redisPassword)

	tStart := time.Now()
	xmlParser.ParseXML(xmlPath + "test.xml")
	timeTrack(tStart, "Parser")

	time.Sleep(10 * time.Second)

}
