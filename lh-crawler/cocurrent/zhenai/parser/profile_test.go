package parser

import (
	"io/ioutil"
	"lh-example/lh-crawler/cocurrent/model"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile(
		"profile_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "限时拥抱万能萌妹")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 "+
			"element; but was %v", result.Items)
	}

	actual := result.Items[0]

	expected := model.Profile{
		Age:        16,
		Height:     116,
		Weight:     148,
		Income:     "5001-8000元",
		Gender:     "女",
		Name:       "限时拥抱万能萌妹",
		Xinzuo:     "巨蟹座",
		Occupation: "测试工程师",
		Marriage:   "离异",
		House:      "无房",
		Hokou:      "西安市",
		Education:  "硕士",
		Car:        "有车",
	}

	if actual != expected {
		t.Errorf("expected %v; but was %v",
			expected, actual)
	}
}
