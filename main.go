package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Read the input file.
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to open input file: %s", err)
	}
	defer f.Close()

	// Prepare the data structure to store the weights carried by each elf. We do not know how many elf we have, so we
	// start with 0 length.
	elfs := []int{0}

	// This is a pointer variable to move to the next elf.
	idx := 0

	// Create a buffered reader to read from the file.
	r := bufio.NewReader(f)
	for {
		// Read the current weight from the file.
		weight, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			// We received an error and it's not an EOF. Aborting.
			log.Fatalf("failed to read input file: %s", err)
		}

		// End of file reached. Exit from the loop.
		if err == io.EOF {
			break
		}

		// We remove all the spaces from the read line.
		weight = strings.TrimSpace(weight)
		if weight == "" {
			// It's an empty line. This means we have to move to the next elf.
			idx++
			elfs = append(elfs, 0)
			continue
		}

		// Convert the weight to integer. We ignore the error returned from Atoi as for this case the data in the
		// input file is always a valid integer.
		w, _ := strconv.Atoi(weight)

		// Add the weight to the total weight carried by this specific elf.
		elfs[idx] += w
	}

	// Sort the weight carried by the elf in ascending order.
	sort.Ints(elfs)

	// Maximum weight is the weight in the last element.
	fmt.Printf("Max weight: %d\n", elfs[idx])
	fmt.Printf("Top three total weights: %d\n", elfs[idx]+elfs[idx-1]+elfs[idx-2])
}
