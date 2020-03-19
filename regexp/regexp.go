package main

import (
	"fmt"
	"regexp"
)

func main() {
	const text  = `<div class="m-btn" data-v-bff6f798>45岁</div>
					<div class="m-btn" data-v-bff6f798>天蝎座(10.23-11.21)</div>`

	 IncomeRe := regexp.MustCompile(`<div class="m-btn" data-v-bff6f798>（*^<）</div>`)
	 match :=IncomeRe.FindAllString(text,-1)
	 for _,data:=range match{
	 	fmt.Println(data[0])
	 	fmt.Println(data[1])
	 }

 }
