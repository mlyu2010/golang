//Run 3 different HTTP Servers serving content from  diffent Port and Url
//Port, Url and folder name with CSV, PRN files are configurable 

package  main

import "fmt"
import "github.com/mlyu2010/golang/csvPrnSrv"

func main() {

	server1 := csvPrnSrv.NewServer( "8080", "app1", "Server1", "csvprndata1" )
	fmt.Printf( "Server Name: %s Port: %s URL: %s Data Dir: %s\n", server1.Title(), server1.Port(), server1.Url(), server1.Datadir() )
	//New Thread1	
	go server1.Serve()

	server2 := csvPrnSrv.NewServer( "8085", "app2", "Server2", "csvprndata2" )
	fmt.Printf( "Server Name: %s Port: %s URL: %s Data Dir: %s\n", server2.Title(), server2.Port(), server2.Url(), server2.Datadir() )
	//New Thread2	
	go server2.Serve()

	server3 := csvPrnSrv.NewServer( "8090", "app3", "Server3", "csvprndata3" )
	fmt.Printf( "Server Name: %s Port: %s URL: %s Data Dir: %s\n", server3.Title(), server3.Port(), server3.Url(), server3.Datadir() )
	//Main Process	
	server3.Serve()
}

