//--Summary:
//  Write a program to display server status.
//
//--Requirements:
//* Create a function to print server status, including:
//  - Number of servers
//  - Number of servers for each status (Online, Offline, Maintenance, Retired)
//* Store the existing slice of servers in a map
//* Default all of the servers to `Online`
//* Perform the following status changes and display server info:
//  - display server info
//  - change `darkstar` to `Retired`
//  - change `aiur` to `Offline`
//  - display server info
//  - change all servers to `Maintenance`
//  - display server info

package main

import "fmt"

// Works like an enum
const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func printSeverStatus(servers map[string]int) {
	fmt.Println("\nThere are", len(servers), "servers")

	stats := make(map[int]int)
	for _, status := range servers {
		switch status {
		case Online:
			stats[Online] += 1
		case Offline:
			stats[Offline] += 1
		case Maintenance:
			stats[Maintenance] += 1
		case Retired:
			stats[Retired] += 1
		}
	}

	fmt.Println(stats[Online], "server are online")
	fmt.Println(stats[Offline], "server are Offline")
	fmt.Println(stats[Maintenance], "server are undergoing Maintenance")
	fmt.Println(stats[Retired], "server are retired")
}



func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	serverStatus := make(map[string]int)
	for _, server := range servers {
		serverStatus[server] = Online
	}

	printSeverStatus(serverStatus)
	
	serverStatus["darkstar"] = Retired
	serverStatus["aiur"] = Offline

	printSeverStatus(serverStatus)

	for server, _ := range serverStatus {
		serverStatus[server] = Maintenance
	}

	printSeverStatus(serverStatus)	
}
