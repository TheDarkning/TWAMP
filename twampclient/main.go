package main

import (
	"flag"
	"fmt"
	"log"

	twamp "github.com/tcaine/twamp"
)

func main() {
	twampserver := flag.String("server", "localhost:2000", "twampserver address:port")
	numpings := flag.Int("nping", 10, "number of pings to send")
	flag.Parse()

	c := twamp.NewClient()

	connection, err := c.Connect(*twampserver)
	if err != nil {
		log.Fatal(err)
	}

	session, err := connection.CreateSession(
		twamp.TwampSessionConfig{
			Port:    6666,
			Timeout: 1,
			Padding: 42,
			TOS:     twamp.EF,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	test, err := session.CreateTest()
	if err != nil {
		log.Fatal(err)
	}

	results := test.RunX(*numpings)

	//for _, result := range results.Results {
	//	fmt.Println(result.GetRTT())
	//}
	fmt.Printf("Media: %v, StDev %v, Persi: %v\n", results.Stat.Avg, results.Stat.StdDev, results.Stat.Loss)

	session.Stop()
	connection.Close()

}
