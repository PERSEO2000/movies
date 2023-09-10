package main

import (
  "fmt"
  "strconv"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "github.com/gin-gonic/gin"
)

type Movies struct {
  Id int `json: "id"`
  Name string  `json: "nam>"`
  Img_url string  `json: "img"`
  Url string `json: "url"`
  Description string `json: "description"`
}


var database *sql.DB
func main() {
   var err error
  database,err = sql.Open("sqlite3","./Movies.db")
  
  if err != nil {
    fmt.Println("No se puede habrir la base de datos")
    return
  }
  
  defer database.Close()

  
  app := gin.Default()
  
  app.GET("/",func(gn *gin.Context) {
    mv := get_all()
    fmt.Println(mv)
    gn.JSON(200,gin.H{"movies": mv})
  })
  
  app.GET("/:id",func(gn *gin.Context) {
    ide := gn.Param("id")
    nm,_ := strconv.Atoi(ide)
    mv := getById(nm)
    gn.JSON(200,gin.H{"movie" : mv})
  })
  
  
  fmt.Println("escuchanfo....")
  app.Run(":3000")
}


func get_all() []Movies {
  
  var mv []Movies

  getall := `SELECT  id,img_url,name,url,description FROM movies`
  
  
  filas,err := database.Query(getall)
  
  if err != nil {
    fmt.Println("Error")
  }
  defer filas.Close()
  
  for filas.Next() {
    var id int
    var img_url string
    var url string
    var name string
    var description string
    err = filas.Scan(&id,&img_url,&url,&name,&description)
    movies := Movies{id,name,img_url,url,description}
    mv = append(mv,movies)
  }
  return mv
}

func getById(ide int) Movies {
  findId := `SELECT id,name,img_url,url,description FROM movies WHERE id = ?`
  row := database.QueryRow(findId,ide)
  var id int
  var img_url string
  var url string
  var name string
  var description string
  err := row.Scan(&id,&name,&img_url,&url,&description)
  if err != nil {
    
  }
  
  mv := Movies{id,name,img_url,url,description}
  
  return mv

}
