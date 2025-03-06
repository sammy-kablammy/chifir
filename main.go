package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MEMORY_SIZE_IN_WORDS = 2_097_152

func main() {

	// initialize the VM's memory
	M := [MEMORY_SIZE_IN_WORDS]uint32{}
	PC := uint32(0)
	address := 0
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("--> Now entering instruction number", address)

		fmt.Printf("opcode: ")
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		value, err := strconv.Atoi(str[:len(str)-1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Expected integer input, try again\n")
			continue
		}
		M[address] = uint32(value)

		fmt.Printf("A: ")
		str, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		value, err = strconv.Atoi(str[:len(str)-1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Expected integer input, try again\n")
			continue
		}
		M[address + 1] = uint32(value)

		fmt.Printf("B: ")
		str, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		value, err = strconv.Atoi(str[:len(str)-1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Expected integer input, try again\n")
			continue
		}
		M[address + 2] = uint32(value)

		fmt.Printf("C: ")
		str, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		value, err = strconv.Atoi(str[:len(str)-1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Expected integer input, try again\n")
			continue
		}
		M[address + 3] = uint32(value)

		address += 4
	}

	fmt.Println()

	for {
		// fetch + decode
		opcode := M[PC]
		A := M[PC+1]
		B := M[PC+2]
		C := M[PC+3]

		// execute
		fmt.Println("PC:", PC)
		switch opcode {
		case 1:
			PC = M[A]
		case 2:
			if M[B] == 0 {
				PC = M[A]
			}
		case 3:
			M[A] = PC
		case 4:
			M[A] = M[B]
		case 5:
			M[A] = M[M[B]]
		case 6:
			M[M[B]] = M[A]
		case 7:
			M[A] = M[B] + M[C]
		case 8:
			M[A] = M[B] - M[C]
		case 9:
			M[A] = M[B] * M[C]
		case 10:
			M[A] = M[B] / M[C]
		case 11:
			M[A] = M[B] % M[C]
		case 12:
			if M[B] < M[C] {
				M[A] = 1
			} else {
				M[A] = 0
			}
		case 13:
			if !(M[B] != 0 && M[C] != 0) {
				M[A] = 1
			} else {
				M[A] = 0
			}
		default:
			fmt.Fprintf(os.Stderr, "Unexpected opcode: 0x%x\n", opcode)
			os.Exit(1)
		}

		if opcode != 1 && opcode != 2 {
			PC += 4
		}
	}
}
