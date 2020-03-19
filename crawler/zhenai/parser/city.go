package parser

import (
	"learngo/crawler/engine"
	"regexp"
)

const cityRe  = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*">([^<]+)</a>`
const userGenderRe  = `<span class="grayL">性别：</span>([^<]+)</td>`
func ParserCity(contents []byte)engine.ParserResult{
	re :=regexp.MustCompile(cityRe)
	matches :=re.FindAllSubmatch(contents,-1)

	result :=engine.ParserResult{}
	for _, m :=range matches{
		name :=string(m[2])

		result.Items = append(result.Items,"User "+ name)
		result.Request = append(result.Request, engine.Request{
			Url:string(m[1]),
			ParserFunc: func(contents []byte) engine.ParserResult {
				return ParseProfile(contents,name)
			},})

		/*
		r :=regexp.MustCompile(userGenderRe)
		match :=r.FindSubmatch(contents)

		if match!=nil{
			result.Gender = append(result.Gender,"Gender "+string(match[1]))
		}*/
	}


	return  result
}
