package main

import "cloud_station_self/cli"

func main() {
	if err := cli.RootCmd.Execute(); err != nil {
		panic(err)
	}
}
