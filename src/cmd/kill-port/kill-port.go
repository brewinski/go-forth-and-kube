package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func getPIDProcessOnPort(port int) ([]string, error) {
	// Find the process ID of the process running on the specified port
	cmd := exec.Command("lsof", "-ti", fmt.Sprintf(":%d", port))
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("killProcessOnPort exec: %w", err)
	}

	fmt.Printf("output: %s\n", output)

	// Convert the process ID to an integer
	pid := strings.Split(string(output), "\n")
	if err != nil {
		return nil, fmt.Errorf("killProcessOnPort pid: %w", err)
	}

	return pid, nil
}

func killAllProcessesWithPID(pid []string) error {
	// join PID into a string
	pidStr := strings.Join(pid, " ")
	trimmedPID := strings.Trim(pidStr, " ")
	// Kill the process
	cmd := exec.Command("kill", "-9", trimmedPID)
	fmt.Println(cmd.String())
	return cmd.Run()
}

func main() {
	// get command line arguments
	port := flag.Int("port", 0, "port to kill")
	flag.Parse()

	if *port == 0 {
		log.Fatal(errors.New("port is required"))
	}

	pids, err := getPIDProcessOnPort(*port)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to kill process on port %d, error: %w", *port, err))
	}

	err = killAllProcessesWithPID(pids)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to kill process on port %d, error: %w", *port, err))
	}

	fmt.Printf("killed process on port %d\n", *port)
}
