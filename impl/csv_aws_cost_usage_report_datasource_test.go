package impl

import (
	"testing"
	"time"
)

func TestNewAwsCostUsageReportDatapoint(t *testing.T) {
	_, err := NewAwsCostUsageReportDatapoint(
		time.Now(),
		"Amazon EC2",
		"123456789012",
		"BoxUsage:t2.micro",
		10.0,
		15.0,
		18.0,
		5.0,
		13.0,
	)
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}

}
func TestNewCsvAwsCostUsageReportCostDatasource(t *testing.T) {
	csvDpSource, err := NewCsvAwsCostUsageReportCostDatasource("../data/aws_cur_report_sample.csv")
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
	if csvDpSource.filepath != "../data/aws_cur_report_sample.csv" {
		t.Errorf("Expected filepath ../data/aws_cur_report_sample.csv but got %s", csvDpSource.filepath)
	}

}

func TestFetch(t *testing.T) {
	// Good Path
	csvDpSource, _ := NewCsvAwsCostUsageReportCostDatasource("../data/aws_cur_report_sample.csv")
	results, err := csvDpSource.Fetch(time.Now(), time.Now())
	if err != nil {
		t.Errorf("An unexpected err occured %s", err)

	}

	if len(results) != 10 {
		t.Errorf("Expected 10 items got %d", len(results))
	}

	// Test the first datapoint
	first_dp := results[0]
	if first_dp.date != time.Date(2024, 9, int(time.January), 0, 0, 0, 0, time.UTC) {
		t.Errorf("Unexpected Date %s", first_dp.date)
	}
	if first_dp.account != "123456789012" {
		t.Errorf("Expected account 123456789012 got %s", first_dp.account)
	}

	if first_dp.service != "Amazon EC2" {
		t.Errorf("Expected service Amazon EC2 got %s", first_dp.service)
	}

	if first_dp.total_cost != 13.0 {
		t.Errorf("Expected total cost 13.0 got %f", first_dp.total_cost)
	}

	// Bad Path
	badCsvDpSource, _ := NewCsvAwsCostUsageReportCostDatasource("../non-existent/path")
	_, err = badCsvDpSource.Fetch(time.Now(), time.Now())
	if err == nil {
		t.Error("Expected an error to be thrown")

	}

	// Bad Path
	badCsvDpSource, _ = NewCsvAwsCostUsageReportCostDatasource("../data/malformatted.csv")
	_, err = badCsvDpSource.Fetch(time.Now(), time.Now())
	if err == nil {
		t.Error("Expected an error to be thrown")

	}

}
