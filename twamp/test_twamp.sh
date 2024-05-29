#!/bin/bash

# Compile the C++ program using make
make
if [ $? -ne 0 ]; then
    echo "Compilation failed"
    exit 1
fi
echo "Compilation successful"

# Define test cases
declare -a tests=(
    "10.41.101.94:8080 --count=5 --interval=50 --payload=64 --output=out1.txt"
    "10.41.101.94:8080 --count=1 --interval=10 --payload=12 --output=out2.txt"
    "10.41.101.94:8080 --count=1000 --interval=1000 --payload=65507 --output=out3.txt"
)

# Run test cases
for test in "${tests[@]}"
do
    echo "Running test: ./twamp_test $test"
    ./twamp_test $test
    if [ $? -ne 0 ]; then
        echo "Test failed: ./twamp_test $test"
        exit 1
    else
        echo "Test passed: ./twamp_test $test"
    fi
done

echo "All tests passed"

# Verify text outputs
for output in out1.txt out2.txt out3.txt
do
    if [ ! -f $output ]; then
        echo "Output file $output not found"
        exit 1
    fi
    echo "Output file $output found"

    # Verify the text content (example: checking the presence of certain keywords)
    if ! grep -q 'Packet Loss' $output || ! grep -q 'Min Latency' $output || ! grep -q 'Max Latency' $output || ! grep -q 'Avg Latency' $output; then
        echo "Output file $output does not contain expected content"
        exit 1
    fi
    echo "Output file $output contains expected content"
done

echo "All outputs verified successfully"

