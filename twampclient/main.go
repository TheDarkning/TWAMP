package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	twamp "github.com/tcaine/twamp"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	twampserver := flag.String("server", "localhost:2000", "twampserver address:port")
	numpings := flag.Int("nping", 10, "number of pings to send")
	flag.Parse()

	c := twamp.NewClient()
	port := 6666 + rand.Intn(1000)
	//fmt.Println(port)
	connection, err := c.Connect(*twampserver)
	if err != nil {
		log.Fatal("Connect error: ", err.Error())
	}

	session, err := connection.CreateSession(
		twamp.TwampSessionConfig{
			Port:    port,
			Timeout: 1,
			Padding: 42,
			TOS:     twamp.EF,
		},
	)
	if err != nil {
		log.Fatal("Session creation impossible: ", err.Error())
	}

	test, err := session.CreateTest()
	if err != nil {
		log.Fatal("Creating test impossible: ", err.Error())
	}

	results := test.RunX(*numpings)

	//for _, result := range results.Results {
	//	fmt.Println(result.GetRTT())
	//}
	fmt.Printf("Media: %v, StDev %v, %%Persi: %.2f\n", results.Stat.Avg, results.Stat.StdDev, results.Stat.Loss)

	session.Stop()
	connection.Close()

}
