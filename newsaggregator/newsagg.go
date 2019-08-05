package main

import ("fmt"
"net/http" 
"io/ioutil"
"encoding/json"
"strings"
)

type NewsItem struct{
 Author string
 Title string
 Description string 
 PublishedAt string
 Url string
 UrlAt string

}

func main(){
res,_:=http.Get("https://newsapi.org/v1/articles?source=cnn&apiKey=d5af3160c2cb45dcb4fea791e0c24f9f");
bytes,_:=ioutil.ReadAll(res.Body);
var s string =strings.Split(string(bytes), "[")[1];
var p string =strings.Split(s,"]")[0];
var z string="["+p[0:]+"]";

var xs []NewsItem;
 json.Unmarshal([]byte(z),&xs);


fmt.Printf("%+v",xs[0].Author);

res.Body.Close();

}