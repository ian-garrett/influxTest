package main

import (
	"github.com/bugsnag/bugsnag-go"
	"github.com/influxdata/influxdb/client/v2"
	//"sync"
	"fmt"
	//"net/http"
	"log"
	"time"
	"math/rand"
	"sync"
)

const (
	MyDB = "authtest"
	username = ""
	password = ""
)


func main() {
	// Initialize Bugsnag with your API key
	bugsnag.Configure(bugsnag.Configuration{
		APIKey: "f2fa85c805629a398253ad33ac72bdcc",
		ReleaseStage: "sandbox",
	})

	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: "todd",
		Password: "123456",
	})
	if err != nil {
		log.Fatal(err)
	}

	writePoints(c)

	// testing goroutine bug testing
	//runProcesses()

	fmt.Println("PROGRAM COMPLETE")
}

func writePoints(clnt client.Client) {
	sampleSize := 5000

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "us",
	})
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < sampleSize; i++ {
		tags := map[string]string{
			"success": "true",
		}

		fields := map[string]interface{}{
			"to":    string(rand.Int())+"@gmail.com",
			"from":   string(rand.Int())+"@gmail.com",
			"subject": "Tup",
		}

		pt, err := client.NewPoint(
			"email_infoTEST",
			tags,
			fields,
			time.Now(),
		)
		if err != nil {
			log.Fatal(err)
		}
		bp.AddPoint(pt)
	}

	if err := clnt.Write(bp); err != nil {
		log.Fatal(err)
	}
}

//TEST FUNCTIONS FOR BUGSNAG
func runProcesses() {
	//bugsnag.Configure(bugsnag.Configuration{
	//	APIKey: "f2fa85c805629a398253ad33ac72bdcc",
	//	ReleaseStage: "sandbox",
	//})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		// AutoNotify captures any panics, repanicking after error reports
		// are sent
		defer bugsnag.AutoNotify()

		var object struct{}
		crash(object)
	}()
	go func() {
		defer wg.Done()
		// AutoNotify captures any panics, repanicking after error reports
		// are sent
		defer bugsnag.AutoNotify()

		var object struct{}
		crash2(object)
	}()

	wg.Wait()
}

//func runProcesses2() {
//	//bugsnag.Configure(bugsnag.Configuration{
//	//	APIKey: "f2fa85c805629a398253ad33ac72bdcc",
//	//	ReleaseStage: "sandbox",
//	//})
//	defer bugsnag.AutoNotify()
//	var wg sync.WaitGroup
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		// AutoNotify captures any panics, repanicking after error reports
//		// are sent
//		defer bugsnag.AutoNotify()
//
//		var object struct{}
//		crash2(object)
//	}()
//
//	wg.Wait()
//}

func crash(a interface{}) string {
	return a.(string)
}

func crash2(a interface{}) int {
	return a.(int)
}