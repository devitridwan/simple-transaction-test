package info

import "fmt"

func PrintHeader() {
	fmt.Println(
		`
      ######     ###     ######      ###    ########     ###    
      ##    ##   ## ##   ##    ##    ## ##   ##     ##   ## ##   
      ##        ##   ##  ##         ##   ##  ##     ##  ##   ##  
       ######  ##     ## ##   #### ##     ## ########  ##     ## 
            ## ######### ##    ##  ######### ##   ##   ######### 
      ##    ## ##     ## ##    ##  ##     ## ##    ##  ##     ## 
       ######  ##     ##  ######   ##     ## ##     ## ##     ## 

`)
	PrintInfo()
}