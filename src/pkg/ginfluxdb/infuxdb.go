package ginfluxdb

import (
	"fmt"
	"time"
	"encoding/json"
	_ "github.com/influxdata/influxdb1-client"
	client "github.com/influxdata/influxdb1-client/v2"
)

const (
	Addr  = "http://10.6.124.21:8086"
)

type Tags map[string]string
type Fields map[string]interface{}

func Connect() (client.Client, error) {
	cli, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     Addr,
	})
	if err != nil {
		return cli, err
	}
	return cli, nil
}

func Query(cli client.Client, sql, database string) (res []client.Result, err error) {
	q := client.NewQuery(sql, database, "")
	response, err := cli.Query(q)
	if err != nil {
		return res, err
	}
	if response.Error() != nil {
		return res, response.Error()
	}
	res = response.Results
	return res, nil
}

func InsertOne(cli client.Client, database string, measurement string, tags Tags, fields Fields) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Precision: "s",
		Database: database,
	})
	if err != nil {
		return err
	}
	pt, err := client.NewPoint(
		measurement,
		tags,
		fields,
		time.Now(),
	)
	if err != nil {
		return err
	}
	bp.AddPoint(pt)
	if err := cli.Write(bp); err != nil {
		return err
	}
	return nil
}

func InsertExample() string {
	cli, err := Connect()
	if err != nil {
		return err.Error()
	}
	defer cli.Close()

	tags := Tags{
		"device":		"computer",
		"method":		"earn",
		"product":		"CPU",
	}
	fields := Fields{
		"billed":		float64(108),
		"licenses":		6,
	}
	err = InsertOne(cli, "dubbo", "payment", tags, fields)
	fmt.Println("insert err:", err)
	if err != nil {
		return err.Error()
	}
	return "success"
}

func QueryExample() string {
	c, err := Connect()
	if err != nil {
		fmt.Println("error:", err)
	}
	defer c.Close()

	sql := "select device, billed from payment"
	res, err := Query(c, sql, "dubbo")
	series := res[0].Series
	if len(series) == 0 {
		return ""
	}
	for i, row := range series[0].Values {
		t, err := time.Parse(time.RFC3339, row[0].(string))
		if err != nil {
			fmt.Println("time parse error:", err)
		}
		device := row[1].(string)
		billed := row[2].(json.Number)
		fmt.Println(i, t, billed, device)
	}
	return "good"
}