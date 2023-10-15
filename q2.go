package ds_hw_0

import (
	"bufio"
	"io"
	"strconv"
	"os"
	"sync"
)

// Sum numbers from channel `nums` and output sum to `out`.
// You should only output to `out` once.
// Do NOT modify function signature.
func sumWorker(nums chan int, out chan int) {
	sum := 0
	for num := range nums {
		sum += num
	}
	out <- sum
}

// Read integers from the file `fileName` and return sum of all values.
// This function must launch `num` go routines running
// `sumWorker` to find the sum of the values concurrently.
// You should use `checkError` to handle potential errors.
// Do NOT modify function signature.
func sum(num int, fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var wg sync.WaitGroup
	nums := make(chan int, num)
	out := make(chan int)

	// Launch 'num' sumWorker goroutines
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sumWorker(nums, out)
		}()
	}

	// Read integers from the file and send them to the channel 'nums'
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		nums <- val
	}
	close(nums)

	go func() {
		wg.Wait()
		close(out)
	}()

	totalSum := 0
	for partialSum := range out {
		totalSum += partialSum
	}

	return totalSum
}
