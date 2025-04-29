
M=$2  # Get the first command-line argument
C=$3
# Check if it's a valid integer
if ! [[ "$count" =~ ^[0-9]+$ ]]; then
    echo "Please provide a positive integer."
    exit 1
fi

# Loop from 1 to count
for ((i = 1; i <= count; i++)); do
    go run worker.go -C="$C" -config=config.txt &
done
