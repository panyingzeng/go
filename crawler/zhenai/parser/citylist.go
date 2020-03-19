package parser

import (
	"learngo/crawler/engine"
	"regexp"
)

const cityListRe  =`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParserCityList(contents []byte) engine.ParserResult{
	re := regexp.MustCompile(cityListRe)
	matches :=re.FindAllSubmatch(contents,-1)

	result := engine.ParserResult{}
	limit :=10
	for _, m := range matches{
		result.Items = append(result.Items,"City " + string(m[2]))
		result.Request = append(result.Request,engine.Request{
			Url:string(m[1]),
			ParserFunc:ParserCity,
		})
	//	fmt.Printf("city:%s, url:%s  ", m[2],m[1])
	//	fmt.Println()
	limit--
	if limit==0{
		break
	}
	}
	//fmt.Printf("matches numbers %v\n",len(matches))
	return  result
}