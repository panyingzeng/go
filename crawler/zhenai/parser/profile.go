package parser

import (
	"learngo/crawler/engine"
	"learngo/crawler/mode"
	"regexp"
	"strconv"
)

 var ageRe =regexp.MustCompile(`<div class=[^>]*>(\d+)岁</div>`)

 //var MarriageRe =regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`)

 var HeightRe = regexp.MustCompile(`<div class=[^>]*>(\d+)cm</div>`)

 var WeightRe = regexp.MustCompile(`<div class=[^>]*>(\d+)kg</div>`)

 var IncomeRe = regexp.MustCompile(`<div class=[^>]*>月收入:([^<]+)</div>`)

 //var EducationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`)

// var OccupationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`)

 var HukouRe = regexp.MustCompile(`<div class=[^>]*>籍贯:([^<]+)</div>`)

 var XinzuoRe = regexp.MustCompile(`<div class=[^>]*>([^(]+\(\d{2}.\d{2}-\d{2}.\d{2}\))</div>`)

// var HouseRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`)

// var CarRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`)

func ParseProfile(contents []byte, name string) engine.ParserResult{
	profile := mode.Profile{}

	profile.Name=name

	//整数部分

	//age, err :=strconv.Atoi(extractString(contents,ageRe))
	age1 :=extractString(contents,ageRe)
	age2, err:= strconv.Atoi(age1)
	if err!=nil{
		profile.Age = age2
	}

	height, err :=strconv.Atoi(extractString(contents,HeightRe))
	if err!=nil{
		profile.Height = height
	}
	weight, err :=strconv.Atoi(extractString(contents,WeightRe))
	if err!=nil{
		profile.Weight = weight
	}


	//字符串部分
	profile.Hukou = extractString(contents,HukouRe)
	profile.Xinzuo = extractString(contents,XinzuoRe)
//	profile.Marriage = extractString(contents,MarriageRe)
//	profile.Education = extractString(contents,EducationRe)
//	profile.Occupation = extractString(contents,OccupationRe)
	profile.Income = extractString(contents,IncomeRe)
//	profile.Car = extractString(contents,CarRe)
//	profile.House = extractString(contents,HouseRe)




	result :=engine.ParserResult{
		Items:[]interface{}{profile},
	}
	return result
}

//字符串公共处理部分
func extractString(contents []byte, re *regexp.Regexp) string{
	match :=re.FindSubmatch(contents)

	if len(match)>=2{
		return string(match[1])
	}else {
		return ""
	}
}
