package impl

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
	"time"
)

type CsvAwsCostUsageReportDpSource struct {
	filepath string
}

type AwsCostUsageReportDatapoint struct {
	date                 time.Time
	service              string
	account              string
	usage_type           string
	usage_qty            float64
	blended_cost         float64
	unblended_cost       float64
	savings_plan_applied float64
	total_cost           float64
}

func NewAwsCostUsageReportDatapoint(
	date time.Time,
	service string,
	account string,
	usage_type string,
	usage_qty float64,
	blended_cost float64,
	unblended_cost float64,
	savings_plan_applied float64,
	total_cost float64,
) (*AwsCostUsageReportDatapoint, error) {
	return &AwsCostUsageReportDatapoint{
		date:                 date,
		service:              service,
		account:              account,
		usage_type:           usage_type,
		usage_qty:            usage_qty,
		blended_cost:         blended_cost,
		unblended_cost:       unblended_cost,
		savings_plan_applied: savings_plan_applied,
		total_cost:           total_cost,
	}, nil
}

func NewCsvAwsCostUsageReportDpSource(filepath string) (*CsvAwsCostUsageReportDpSource, error) {
	return &CsvAwsCostUsageReportDpSource{
		filepath: filepath,
	}, nil
}

func (c *CsvAwsCostUsageReportDpSource) Fetch(start_date time.Time, end_date time.Time) ([]AwsCostUsageReportDatapoint, error) {

	var datapoints []AwsCostUsageReportDatapoint

	rows, err := readRows(c.filepath)
	if err != nil {
		return nil, err
	}
	for _, r := range rows {
		d := convToAwsCostUsageReportDatapoint(r)
		datapoints = append(datapoints, *d)
	}

	return datapoints, nil

}

// readRows will read all rows of a CSV file
// and map them to header-value pairs
// eg: example.CSV
// 			User, Score
// 			Alice, 28.0
// 			Bob, 12.0

// will return
//
//	[
//		{"User":"Alice", "Score":28.0},
//		{"User":"Bob", "Score":12.0},
//	]
func readRows(filepath string) ([]map[string]string, error) {
	var rows []map[string]string

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err

	}
	defer file.Close()

	reader := csv.NewReader(file)
	headers, _ := reader.Read()

	for {
		rowMap := make(map[string]string)

		row, err := reader.Read()
		if err != nil {
			// Break the loop on EOF
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}

		for i, header := range headers {
			key := strings.ToLower(header)
			rowMap[key] = row[i]

		}
		rows = append(rows, rowMap)
	}

	return rows, nil

}

func convToAwsCostUsageReportDatapoint(row_d map[string]string) *AwsCostUsageReportDatapoint {
	date, _ := time.Parse("2006-01-02", row_d["date"])
	service_name := row_d["service name"]
	account := row_d["linked account"]
	usage_type := row_d["usage type"]
	usage_qty_float, _ := strconv.ParseFloat(row_d["usage quanitity"], 64)
	blended_cost_float, _ := strconv.ParseFloat(row_d["blended cost"], 64)
	unblended_cost_float, _ := strconv.ParseFloat(row_d["unblended_cost"], 64)
	savings_plan_applied_float, _ := strconv.ParseFloat(row_d["savings plan applied"], 64)
	total_cost_float, _ := strconv.ParseFloat(row_d["total cost"], 64)

	d, _ := NewAwsCostUsageReportDatapoint(
		date,
		service_name,
		account,
		usage_type,
		usage_qty_float,
		blended_cost_float,
		unblended_cost_float,
		savings_plan_applied_float,
		total_cost_float,
	)
	return d
}
