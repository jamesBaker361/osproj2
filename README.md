How to use:
Generate a filename.dat to extract primes from: `head -c 1K /dev/urandom > filename.dat` \\
Free up ports with `bash clean_ports.sh`
Call `go run main.go -N={N} -data=filename.dat  -config=config.txt` where {N} is in {1KB, 32KB, 64KB, 256KB, 1MB, 64MB} \\
Call `bash run_workers.sh M C` where M is an integer and C is in {64B, 1KB, 4KB, 8KB}
