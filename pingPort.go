package main
/*
	KK April 17 2018
*/

import (
	"os"
	"flag"
	"fmt"
	"net"
	"time"
	"strings"
	"path/filepath"
	"github.com/fatih/color"
)


func main(){

	// command line flags:
	targetPtr := flag.String("target", "127.0.0.1:80", "IP_Adderess:Port")
	protocolPtr := flag.String("protocol", "TCP", "TCP/UDP")
	timeoutPtr := flag.Int("timeout", 250, "timeout in milliseconds")
	silentPtr := flag.Bool("silent", false, "if set there is no output only exit code of 0 [success] or 1 [timeout]")
	nocolourPtr := flag.Bool("nocolour", false, "if set colour output is disabled")
	
    
    flag.Parse()
    protocol := strings.ToLower( *protocolPtr )


    // check if we are doing a simple default TCP 250ms timeout check:
    var connStr string
	if ( len( os.Args ) == 2 ){

		// default basic usage of pingPort.exe IP_Adderess:Port
		connStr = os.Args[1]

	}else if ( len( os.Args ) < 2 ){
		
		// too few arguments
		usage()
	}else{
		
		// we're doing an advanced flag setup
		// hope the flags are all set
		connStr = *targetPtr
	}

	// make the connection attempt
	_,err := net.DialTimeout(protocol, connStr, time.Millisecond * time.Duration(*timeoutPtr) )
	if err != nil {
		if ( *silentPtr == false ){ 
			if ( *nocolourPtr == false ){
				KPrintln("Connection Timeout", "red") 	
			}else{
				fmt.Println("Connection Timeout")
			}
		}
		os.Exit(1)
	}else{
		if ( *silentPtr == false ){ 
			if ( *nocolourPtr == false ){
				KPrintln("Connection Success", "green")
			}else{
				fmt.Println("Connection Success")
			}
		}
		os.Exit(0)
	}
}


// prints the command line usage
func usage() {


	// get only the file name from the absolute path in os.Args[0]
	_, file := filepath.Split(os.Args[0])

    fmt.Fprintf(os.Stderr, "\nBasic Usage: %s IP_Adderess:TCP_Port\n\n", file )
    fmt.Fprintf(os.Stderr, "Advanced Flag Usage: %s Flags:\n", file )
    flag.PrintDefaults()
    fmt.Fprintf(os.Stderr, "\n")
    os.Exit(2)
}



func KPrint( args ...string ){

	// validate args
	if len(args) < 1 {return}

	// populate local vars from args
	str := args[0]
	colour := "white"
	if len(args) == 2 { colour = args[1] }

	switch strings.ToLower(colour){
		case "green":
			str = color.HiGreenString(str)

		case "yellow":
			str = color.HiYellowString(str)

		case "red":
			str = color.HiRedString(str)

		case "magenta":
			str = color.HiMagentaString(str)

		case "cyan":
			str = color.HiCyanString(str)

		case "white":
		default:
			str = color.HiWhiteString(str)
	}

	// print the colour string
	fmt.Fprintf(color.Output, "%s", str)
}

//
// wrapper for KPrint, adds the new line
//
func KPrintln( args ...string ){

	args[0] += "\n" // add the new line to the str, arg 2 is the colour
	KPrint( args... )
}