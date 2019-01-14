package models

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func FindAMovie(str string,movie_Id int64)MovieInfo{
	var movie MovieInfo
	movie.Movie_id = movie_Id
	movie.Movie_name = FindMovieName(str)
	movie.Movie_director = FindMovieDirector(str)
	movie.Movie_writer = FindMovieWriter(str)
	movie.Movie_main_character = FindMovieCharacter(str)
	movie.Movie_type = FindMovieType(str)
	movie.Movie_country = FindMovieCountry(str)
	movie.Movie_on_time = FindMovieOnTime(str)
	movie.Movie_picture = FindMoviePicture(str)
	movie.Movie_span = FindMovieSpan(str)
	movie.Movie_grade = FindMovieGrade(str)
	return movie
}

//匹配name
func FindMovieName(str string)string{
	if str == "" {
		return ""
	}
	reg := regexp.MustCompile(`<span\s*property=.*>(.*)</span>`)
	result := reg.FindAllStringSubmatch(str,-1)

	if len(result) == 0 {
		return ""
	}
	name := result[0][1]

	return string(name)
}

//匹配导演
func FindMovieDirector(str string)string{
	if str == "" {
		return ""
	}
	reg := regexp.MustCompile(`<a\s*href="/celebrity/[0-9]{7}/"\s*rel="v:directedBy">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(str,-1)

	director := ""
	for _,v := range result{
		director += v[1]+"/"
	}
	//fmt.Println(director)

	//去掉末尾的“/”
	return strings.Trim(director,"/")
}

//匹配编剧
func FindMovieWriter(str string)string{
	if str == "" {
		return ""
	}
	reg := regexp.MustCompile(`<a\s*href="/celebrity/[0-9]{7}/">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(str,-1)//-1表示匹配的个数不限制

	writers := ""

	for _,v := range result{
		writers += v[1] + "/"
	}
	//fmt.Println(writers)
	return strings.Trim(writers,"/")
}

//匹配主演
func FindMovieCharacter(str string) string{
	if str == ""{
		return ""
	}
	reg := regexp.MustCompile(`<a\s*href="/celebrity/[0-9]{7}/"\s*rel="v:starring">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(str,-1)
	Character := ""
	//fmt.Println(result)
	for _,v := range result{
		Character += v[1] + "/"
	}
	return strings.Trim(Character,"/")
}

//匹配type
func FindMovieType(str string) string{
	if str == ""{
		return ""
	}
	reg := regexp.MustCompile(`<span\s*property="v:genre">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(str,-1)
	MovieType := ""
	//fmt.Println(result)
	for _,v := range result{
		MovieType += v[1] + "/"
	}
	return strings.Trim(MovieType,"/")
}

//匹配country
func FindMovieCountry(str string) string{
	if str == ""{
		return ""
	}
	reg := regexp.MustCompile(`<span\s*class="pl">制片国家/地区:</span>\s*(.*?)<br/>`)
	result := reg.FindAllStringSubmatch(str,-1)
	MovieCountry := ""
	//fmt.Println(result)
	for _,v := range result{
		MovieCountry += v[1]
	}
	return MovieCountry
}

//匹配ontime
func FindMovieOnTime(str string) string{
	if str == ""{
		return ""
	}
	reg := regexp.MustCompile(`<span\s*property="v:initialReleaseDate"\s*content=".*">(.*?)\(.*\)</span>`)
	result := reg.FindAllStringSubmatch(str,-1)

	//fmt.Println(result)
	if len(result) == 0 {
		return ""
	}
	MovieOnTime := result[0][1]
	//fmt.Println(MovieOnTime)
	return MovieOnTime
}

//匹配ontime
func FindMoviePicture(str string) string{
	if str == ""{
		return ""
	}
	reg := regexp.MustCompile(`<img\s*src="(.*?)"\s*title="点击看更多海报" alt=".*"\s*rel="v:image"\s*/>`)
	result := reg.FindAllStringSubmatch(str,-1)

	//fmt.Println(result)
	if len(result) == 0{
		return ""
	}
	MoviePicture := result[0][1]
	return MoviePicture
}

//匹配language
func FindMovieLanguage(str string) string{
	if str == ""{
		return ""
	}
	reg := regexp.MustCompile(`<span\s*class="pl">语言:</span>\s*(.*?)<br/>`)
	result := reg.FindAllStringSubmatch(str,-1)
	if len(result) == 0 {
		return ""
	}
	MovieLanguage := ""
	//fmt.Println(result)
	MovieLanguage = result[0][1]
	return MovieLanguage
}

//匹配时长
func FindMovieSpan(str string) string{
	if str == ""{
		return ""
	}
	reg := regexp.MustCompile(`<span property="v:runtime"\s*content="\d*">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(str,-1)
	if len(result) == 0 {
		return ""
	}
	MovieSpan := ""
	//fmt.Println(result)
	MovieSpan = result[0][1]
	return MovieSpan
}

//匹配评分
func FindMovieGrade(str string) string{
	if str == ""{
		return ""
	}
	reg := regexp.MustCompile(`<strong\s*class="ll\s*rating_num"\s*property="v:average">(.*?)</strong>`)
	result := reg.FindAllStringSubmatch(str,-1)
	MovieGrade := ""
	fmt.Println(result[0][1])
	MovieGrade = result[0][1]
	return MovieGrade
}

//find movie id
func FindMovieId(sUrl string)int64{
	//fmt.Println(sUrl)
	strings.Trim(sUrl,"/")
	strs := strings.Split(sUrl,"/")
	id,err := strconv.ParseInt(strs[4],10,64)
	if err != nil {
		fmt.Println("tag")
		return 0
	}
	//fmt.Println(id)
	return id
}

//find movie urls
func FindMovieURLs(str string)[]string{

	reg := regexp.MustCompile(`href="(https://movie.douban.com/subject/.*?/)`)
	result := reg.FindAllStringSubmatch(str,-1)
	//fmt.Println(result)

	var strs []string
	for _,value := range result{
		strs = append(strs,value[1])
	}
	//fmt.Println(len(strs))
	//fmt.Println(strs)
	return strs
}