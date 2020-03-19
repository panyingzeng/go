package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)


//fetch
// 输入URL 输出byte，error
func Fetch(url string) ([]byte, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		return nil, fmt.Errorf("wrong error %d ",resp.StatusCode)
	}
	//解决gbk编码问题
	bodyReader := bufio.NewReader(resp.Body)
	e :=determineEncodeing(bodyReader)
	utf8Reader :=transform.NewReader(bodyReader,e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

//自动识别编码
func determineEncodeing(r *bufio.Reader) encoding.Encoding{
	bytes, err :=r.Peek(1024)
	if err !=nil{
		log.Printf("fetcher error %d",err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes,"")
	return e
}