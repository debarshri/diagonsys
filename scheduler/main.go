package main

import (
	"github.com/influxdata/influxdb/client/v2"
	"github.com/jasonlvhit/gocron"
	"log"
	"net/http"
	"time"
)

const (
	Address            = "http://localhost:8086"
	Username           = "root"
	Password           = "root"
	Database           = "status"
	Precision          = "s"
	status_field       = "status"
	stattus_code_field = "status_code"
)

func main() {

	log.Println("Starting scheduler")

	gocron.Every(10).Minutes().Do(task)
	<-gocron.Start()
}

func task() {

	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     Address,
		Username: Username,
		Password: Password,
	})

	defer c.Close()

	is_error(err)

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  Database,
		Precision: Precision,
	})

	save(bp)
	c.Write(bp)

	log.Println("Point saved")
}

func save(bp client.BatchPoints) {

	//Should come from database
	status, status_code := get("http://demo.recyclingcourse.eu")
	tags := map[string]string{"site": "http://demo.recyclingcourse.eu"}

	fields := map[string]interface{}{
		status_field:       status,
		stattus_code_field: status_code,
	}

	pt, err := client.NewPoint("site_status", tags, fields, time.Now())
	bp.AddPoint(pt)

	is_error(err)
}

func is_error(err error) {
	if err != nil {
		log.Fatalln("Cannot connect to influxdb")
	}
}

func get(url string) (bool, int) {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln("Cannot connect to get the site " + url)
	}

	c := resp.StatusCode

	if c == 200 {
		return true, c
	}

	return false, c
}
