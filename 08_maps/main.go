package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/golang/glog"
)

func usage() {
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	flag.Usage = usage
	flag.Set("logtostderr", "true")
	flag.Set("stderrthreshold", "WARNING")
	flag.Set("v", "2")
	// This is wa
	flag.Parse()
}

func main() {
	//pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	//	pflag.Parse()

	//	goflag.CommandLine.Parse([]string{})
	glog.Info("Starting transaction...")
	//Define map
	emails := make(map[string]string)

	emails["Rod"] = "rod@gmail.com"
	emails["Capito"] = "capito@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Rod"])
}
