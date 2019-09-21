package weather

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"fmt"
	"encoding/json"
)
type Index struct {
	Des   string `json:"des"`
	Tipt  string `json:"tipt"`
	Title string `json:"title"`
	Zs    string `json:"zs"`
}
type WeatherData struct {
	Date            string `json:"date"`
	DayPictureURL   string `json:"dayPictureUrl"`
	NightPictureURL string `json:"nightPictureUrl"`
	Weather         string `json:"weather"`
	Wind            string `json:"wind"`
	Temperature     string `json:"temperature"`
}
type Results struct {
	CurrentCity string 			`json:"currentCity"`
	Pm25        string 			`json:"pm25"`
	//Index       []Index 		`json:"index"`
	WeatherData []WeatherData 	`json:"weather_data"`
}
type results struct {
	//Error   int    		`json:"error"`
	//Status  string 		`json:"status"`
	//Date    string 		`json:"date"`
	Results []Results 	`json:"results"`
}

// 接口来自网络http://api.jirengu.com/
func GetWeather(city string) (results, error) {
	//请求地址
	originURL :="http://api.jirengu.com/getWeather.php"

	//初始化参数
	param:=url.Values{}

	//配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
	param.Set("cityname",city) //要查询的城市，如：温州、上海、北京

	//发送请求
	data,err:=Get(originURL,param)

	var netReturn results
	json.Unmarshal(data,&netReturn)

	if err!=nil{
		fmt.Errorf("请求失败,错误信息:\r\n%v",err)
		return netReturn, err
	}
	return netReturn, err
}
// get 网络请求
func Get(apiURL string,params url.Values)(rs[]byte ,err error){
	var Url *url.URL
	Url,err=url.Parse(apiURL)
	if err!=nil{
		fmt.Printf("解析url错误:\r\n%v",err)
		return nil,err
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery=params.Encode()
	resp,err:=http.Get(Url.String())
	if err!=nil{
		fmt.Println("err:",err)
		return nil,err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// post 网络请求 ,params 是url.Values类型
func Post(apiURL string, params url.Values)(rs[]byte,err error){
	resp,err:=http.PostForm(apiURL, params)
	if err!=nil{
		return nil ,err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}