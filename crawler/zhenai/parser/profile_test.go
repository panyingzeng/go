package parser

import (
	"io/ioutil"
	"learngo/crawler/mode"
	"testing"
)

func TestParserProfile(t *testing.T) {
	contents, err :=ioutil.ReadFile("profile_test_data.html")

	if err != nil {
		panic(err)
	}
	result :=ParseProfile(contents,"手心的余温慢慢蔓")

	if len(result.Items) !=1 {
		t.Errorf("Item should contain 1"+
			"element; but was %v", result.Items)
	}

	profile :=result.Items[0].(mode.Profile)

	expected := mode.Profile{
		Name:"手心的余温慢慢蔓",
		Age:26,
		Height:156,
		Weight:40,
		Hukou:"四川阿坝",
		Xinzuo:"天蝎座(10.23-11.21)",
	//	Marriage:"未婚",
	//	Education:"大学本科",
	//	Occupation:"政府机构",
		Income:"5-8千",
	//	Car:"未买车",
	//	House:"和家人同住",

	}

	if profile !=expected {
		t.Errorf("expected %v; but was %v",expected,profile)
	}

}
