package parser

import (
	"testing"

	"io/ioutil"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile(
		"citylist_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)

	const RESULT_SIZE = 470
	if len(result.Requests) != RESULT_SIZE {
		t.Errorf("需要 %d 得到 %d", RESULT_SIZE, len(result.Requests))
	}

}
