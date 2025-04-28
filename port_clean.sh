#!/bin/bash

# List of ports to check
ports=("5001" "5002" "5003")

# Loop over each port
for port in "${ports[@]}"; do
    # Find the processes running on the specified port
    pids=$(lsof -t -i :$port)
    
    # If there are no processes found
    if [ -z "$pids" ]; then
        echo "No processes found on port $port."
    else
        # Kill each process found on the port
        for pid in $pids; do
            kill -9 $pid
            if [ $? -eq 0 ]; then
                echo "Successfully killed process with PID $pid on port $port."
            else
                echo "Failed to kill process with PID $pid on port $port."
            fi
        done
    fi
done
