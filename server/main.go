package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"net/http"
	"github.com/influxdata/influxdb/client/v2"
	"log"
)

const (
	Addr               = "http://localhost:8086"
	Username           = "root"
	Password           = "root"
	Database           = "status"
	Precision          = "s"
	status_field       = "status"
	stattus_code_field = "status_code"
)

func main() {

	log.Println("Starting server")

	server := echo.New()
	server.Get("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello, World!")
	})
	server.Run(standard.New(":1323"))
}


func read(table string){
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
		Username: "root",
		Password: "root",
	})

	defer c.Close()

	if err != nil {
		log.Fatalln("Cannot connect to influxdb")
	}


}

func queryDB(clnt client.Client, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: Database,
	}
	if response, err := clnt.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}
