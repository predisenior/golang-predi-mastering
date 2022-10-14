package main

import (
	"project-2/config"
	"project-2/routes"
)



func main(){
config.KonekDB()
e:=routes.New()
e.Start(":3000")
   // Start server
// e.Logger.Fatal(e.Start(":3000"))
}

