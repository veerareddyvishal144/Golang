package main

import ("fmt"
"net/http" 
"io/ioutil"
"encoding/json"

)
type News struct{
	SortBy string
	Status string
	Source string
	Articles []NewsItem
}
type NewsItem struct{
 Author string
 Title string
 Description string 
 PublishedAt string
 Url string
 UrlAt string

}

func main(){
res,_:=http.Get("https://newsapi.org/v1/articles?source=cnn&apiKey=****");
bytes,_:=ioutil.ReadAll(res.Body);


var xs News;
 json.Unmarshal(bytes,&xs);


fmt.Printf("%+v",xs);

res.Body.Close();

}