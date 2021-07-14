package main

import (
    "github.com/TonyFeng628/http-server/internal/core"
)




func main(){
	core := core.NewCore()
	core.Init()
	core.Serve()
}
  
