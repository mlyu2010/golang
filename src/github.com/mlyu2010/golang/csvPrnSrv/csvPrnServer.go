//Server to render content of CSV, PRN files as HTML

package csvPrnSrv


import "fmt"
import "net/http"

//HTTP server wrapper
type Server struct {
    	//Port of the application, for example 8080
	port string 
	//Relative URL of the application in the path, for example 'app1' provided with full URL like 'http://localhost:8080/app1    
	url string  
	//Title of the HTTP server
	title string 
	//Relative directory where CSV/PRN files are located
	datadir string
	//Instance of the HTTP server	
	srv *http.Server
}


//HTTP server constructor
func NewServer( port, url, title, datadir string ) (*Server) { 
    	s := &Server{port, url, title,datadir,&http.Server{Addr: ":"+port, Handler: nil } }
	mux := http.NewServeMux()

    	//Closure for directory content handler
	mux.HandleFunc( "/" + url, 
        func ( rw http.ResponseWriter, req *http.Request ) {
            fmt.Fprintf( rw, "<h1>Content located in '" + datadir + "' folder is served by '" + title + "' from URL:%s </h1>\n", url )
	    ListIndex( rw, datadir, url )
        })

	//Closure for file content handler
	mux.HandleFunc( "/" + url + "/" + DATASTR + "/", 
        func ( rw http.ResponseWriter, req *http.Request ) {
            	pathSize := len( url ) + len(DATASTR) + 3
		fileName := req.URL.Path[pathSize:]
		fileTypeName := GetFileTypeName(fileName)
		fmt.Fprintf( rw, "<h1>" + fileTypeName + " data file '%s' located in '%s' folder</h1>", fileName, datadir )
		ListFile( rw, fileName, datadir )
        })

    	s.srv.Handler = mux
    	return s
}

//Getter for Server Port
func (s *Server) Port() (string) {
    return s.port;
}

//Getter for Server Url 
func (s *Server) Url() (string) {
    return s.url;
}

//Getter for Server Title
func (s *Server) Title() (string) {
    return s.title;
}

//Getter for Server Datadir
func (s *Server) Datadir() (string) {
    return s.datadir;
}

//Launch Server and handle HTTP requests
func (s *Server) Serve() {
    s.srv.ListenAndServe()
}




