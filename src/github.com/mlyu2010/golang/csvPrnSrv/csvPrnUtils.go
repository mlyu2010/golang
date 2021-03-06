//Various utils for csvPrnServer

package  csvPrnSrv

import "regexp"
import "fmt"
import "os"
import "io"
import "bufio"
import "encoding/csv"
import "net/http"

// data part of URL
const DATASTR = "data"

//get a list of eligible file names for a given directory
func ContentList( dirname string )  ([]string, error) {

	dirf,err := os.Open( dirname )
    if err != nil {
        return nil, err
    }
	defer dirf.Close()

	dir,err := dirf.Readdir(-1)
    if err != nil {
        return nil, err
    }

	var files []string
	for _,val := range( dir ) {
		if  (! val.IsDir()) && IsFileEligible( val.Name() ) {
			files = append( files, val.Name() )
		}
	}

	return files, nil
}

// whether a file is eligible to be listed
func IsFileEligible( fname string ) bool {
	return IsFileCsv(fname) || IsFilePrn(fname)
}

// is it CSV File
func IsFileCsv( fname string ) bool {
	match,_ := regexp.MatchString( "\\.[Cc][Ss][Vv]$",  fname )
	return match
}

// is it PRN File
func IsFilePrn( fname string ) bool {
	match,_ := regexp.MatchString( "\\.[Pp][Rr][Nn]$",  fname )
	return match
}

// get file type name CSV or PRN
func GetFileTypeName( fname string ) string {
	if IsFileCsv(fname) {
		return "CSV"
        }
	if IsFilePrn(fname) {
		return "PRN"
        }
	return "Generic"
}

// list index of files as links 
func ListIndex( rw http.ResponseWriter, datadir string, url string ) {
	dir,err := ContentList( datadir )
	if err != nil {
		fmt.Fprintf( rw, "Error reading directory %s", datadir )
	} else {
		fmt.Fprintf( rw, "<ul>\n")
		for _,file := range( dir ) {
			fmt.Fprintf( rw, "<li>%s</li>\n", makeLink( file, url ) )
		}
		fmt.Fprintf( rw, "</ul>\n" )
	}
}

// render a specified file as HTML table
func ListFile(  rw http.ResponseWriter, fname string, datadir string  ) {
	inp,err := os.Open( datadir + "/" + fname )
	
	defer inp.Close()

	if err != nil {
		fmt.Fprintf( rw, "Open file %s error ", fname )
	} else {
		
		table,err := FileContentToHTML( inp, fname )
		if  err != nil {		
			fmt.Fprintf( rw, "Error reading file %s", fname)
		} else {
			fmt.Fprintf( rw, "%s", table )
		}
		
	}
}

// render a link from filename
func makeLink( fname string, url string ) (string) {
	link := "<a href=\""
	link += "/" + url + "/" + DATASTR + "/" + fname + "\">"
	link += fname + "</a>"
	return link
}


// convert data file content into HTML table
func FileContentToHTML( r io.Reader, fname string ) (string,error) {
	//var html string     
	if IsFileCsv(fname) {	    
   		html,err := csvFileContent2HTML(r)
		if err != nil {
			return "",err
		}
    		return "<table border=\"1\">\n" + html + "\n</table>\n", nil
        }
	html,err := prnFileContent2HTML(r)
	if err != nil {
		return "",err
	}
    	return "<table border=\"1\">\n" + html + "\n</table>\n", nil
}

// convert string into table data tag
func makeHTMLData( val string ) string {
    return "<td>" + val + "</td>\n"
}

// convert CSV record to table row
func csvRow2HTML( rec []string ) string {
    var row string
    row += "<tr>"
    for _,val := range( rec ) {
        row += makeHTMLData( val )
    }
    row += "</tr>\n"
    return row
}

// convert PRN record to table row
func prnRow2HTML( rec string ) string {
    var row string
    row += "<tr>"
    row += makeHTMLData( rec )
    row += "</tr>\n"
    return row
}

//transform CSV file content into HTML
//Field separator is comma
//Hence able to create a good formating with TD for each field
func csvFileContent2HTML( r io.Reader ) (string,error) {
	var html string     
		    
    	reader := csv.NewReader( r )
    	
    	for {
    	    record,err := reader.Read()
    	    if err == io.EOF {
    	        break
    	    } else if err != nil {
    	        return "", err
    	    }
    	    html += csvRow2HTML( record )
    	}

    	return "<table border=\"1\">\n" + html + "\n</table>\n", nil
}

//transform PRN file content into HTML
//No info for PRN what's the field separator
//Hence output as a single TD only
func prnFileContent2HTML( r io.Reader ) (string,error) {
	var html string
	reader := bufio.NewReader( r )
	for  {
		record,err := reader.ReadString( '\n' )
		if err == io.EOF {
    	        	break
    	    	} else if err != nil {
    	        	return "", err
    	    	}
		html += prnRow2HTML(record)
	}
	
	return "<table border=\"1\">\n" + html + "\n</table>\n", nil
}
