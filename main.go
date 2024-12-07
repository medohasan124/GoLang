package main

import (
    "fmt" 
    
)

 type Playable interface{
    Play() string
 }

 type Song struct{
    Title string
    Artist string
 }

 type Movie struct{
    Title string
    Dircetor string
 }

 func (s Song) Play() string{
    return "play " + s.Title + "by " + s.Artist
 }

 func (m Movie) Play() string{
    return "now play " + m.Title + "direction By" + m.Dircetor
 }

 func PrintPlay(p Playable){
    fmt.Println(p.Play())
 }

 func main(){
    songs := Song{Title: "mafish haga" , Artist:"shreen"}

    movie := Movie{Title: "Wight ibrahim" , Dircetor:"Action"}

    PrintPlay(songs)
    PrintPlay(movie)
   
 }
