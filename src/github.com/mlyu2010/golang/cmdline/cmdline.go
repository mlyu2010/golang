package main
import "fmt"
import "os"
func main() {
    fmt.Printf( "Number of args is %d\n", len( os.Args ) )
    for idx,val := range( os.Args ) {
        fmt.Printf( "arg %d is %s\n", idx, val )
    }
}
