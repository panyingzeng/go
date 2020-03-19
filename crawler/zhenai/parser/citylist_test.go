package parser

import (
	"io/ioutil"
	"testing"
)

func TestParserCityList(t *testing.T) {
	contents, err :=ioutil.ReadFile("citylist_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParserCityList(contents)

	const ResultSize = 470
	expectedUrls :=[]string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCitys :=[]string{
		"City 阿坝",
		"City 阿克苏",
		"City 阿拉善盟",
	}

	if len(result.Request) != ResultSize {
		t.Errorf("result should have %d request, but have %d",ResultSize,len(result.Request))
	}

	for i, url := range expectedUrls {
		if result.Request[i].Url !=url {
			t.Errorf("expected url #%d: %s;but was %s",i,url,result.Request[i].Url)
		}
		
	}
	
	if len(result.Items) != ResultSize {
		t.Errorf("result should have %d request, but have %d",ResultSize,len(result.Items))
	}

	for i, city := range expectedCitys {
		if result.Items[i] !=city {
			t.Errorf("expected city #%d: %s;but was %s",i,city,result.Items[i])
		}

	}
	/*//verify result*/
	//fmt.Printf("%s\n",contents)
}
