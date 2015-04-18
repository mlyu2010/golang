//Unit tests for the package 'csvPrnSrv'

package csvPrnSrv

import "testing"
import "os"

// assertion
func check( test bool, msg string, t *testing.T ) {
	if ! test  {
		t.Errorf( msg )
	}
}

// test valid file extension detected correctly
func Test_HasCSVExt( t *testing.T ) {
	check( IsFileEligible( "test1.csv" ), "IsFileEligible (1)", t )
	check( IsFileEligible( "test2.CSV" ), "IsFileEligible (2)", t )
	check( ! IsFileEligible( "test3.crv" ), "IsFileEligible (3)", t )
	check( ! IsFileEligible( "fest4.csvx" ), "IsFileEligible (4)", t )
	check( IsFileEligible( "test5.prn" ), "IsFileEligible (5)", t )
	check( IsFileEligible( "test6.PRN" ), "IsFileEligible (6)", t )
}

// test conversion of CSV file to HTML
func Test_csv2HTML( t *testing.T ) {
	csvfile :=  "../../../../../csvprndata1/Workbook2.csv"
	afile,err := os.Open( csvfile )
	if err != nil {
		t.Errorf( "can't open csv test file: " + csvfile )
		return
	}
	defer afile.Close()
	html,err := csvFileContent2HTML( afile )
	if  err != nil {
		t.Errorf( "error parsing csv file: " + csvfile)
		return
	}
	check( html !=  "", "csvFileContent2HTML (1)", t )
}

// test conversion of PRN file to HTML
func Test_prn2HTML( t *testing.T ) {
	prnfile :=  "../../../../../csvprndata1/Workbook2.prn"
	afile,err := os.Open( prnfile )
	if err != nil {
		t.Errorf( "can't open prn test file: " + prnfile )
		return
	}
	defer afile.Close()
	html,err := prnFileContent2HTML( afile )
	if  err != nil {
		t.Errorf( "error parsing prn file: " + prnfile)
		return
	}
	check( html !=  "", "prnFileContent2HTML (1)", t )
}
