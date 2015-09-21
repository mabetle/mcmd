// calc call cmd time
package main

import(
	"fmt"
	"os"
	"strings"
	"os/exec"
	"time"
	"flag"
)


func main(){
	startTime := time.Now()

	command := flag.String("cmd","", "Set the command.")
	args := flag.String("args", "", "Set the args. (separated by spaces)")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-cmd <command>] [-args <the arguments (separated by spaces)>]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()


	fmt.Println("Command: ", *command)
	fmt.Println("Arguments: ", *args)

	var argArray []string
	if *args != "" {
		argArray = strings.Split(*args, " ")
	} else {
		argArray = make([]string, 0)
	}
	cmd := exec.Command(*command, argArray...)
	buf, err := cmd.Output()


	if err != nil {
		fmt.Fprintf(os.Stderr, "The command failed to perform: %s (Command: %s, Arguments: %s)", err, *command, *args)
		return
	}
	fmt.Fprintf(os.Stdout, "Result: %s", buf)

	endTime := time.Now()

	duration:=endTime.Sub(startTime)

	fmt.Printf("Start at: %v \n" , startTime)
	fmt.Printf("End   at: %v \n" , endTime)
	fmt.Printf("Duration: %v \n", duration)

}




