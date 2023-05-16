package bufio

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
	With bufio, we can use the bufio.Writer method to accumulate data into a buffer before writing to IO.
	In the example below, we have demonstrated three likely situations that you may encounter:

		- The buffer is full
		- The buffer has space after a write
		- A write larger than buffer capacity is made

			1. The buffer is full

		As soon as the buffer is full, the write operation takes place.

			2. The buffer has space after write

		If the buffer still has space after the last write, it will not attempt to complete that write until specifically urged to do so by the Flush() method.

			3. A write larger than buffer capacity is made

		If a write is larger than buffer capacity,​ the buffer is skipped because there is no need to buffer.
*/

type Writer int

func (*Writer) Write(p []byte) (n int, err error) {
	fmt.Printf("Writing: %s\n", p)
	return len(p), nil
}

func BufioWriting() {
	// declare a buffered writer
	// with buffer size 4
	w1 := new(Writer)
	bw := bufio.NewWriterSize(w1, 4)

	// Case 1: Writing to buffer until full
	bw.Write([]byte{'1'})
	bw.Write([]byte{'2'})
	bw.Write([]byte{'3'})
	bw.Write([]byte{'4'}) // write - buffer is full

	fmt.Println(bw.Available())

	// Case 2: Buffer has space
	bw.Write([]byte{'5'})
	fmt.Println(bw.Available())
	err := bw.Flush() // forcefully write remaining
	if err != nil {
		panic(err)
	}

	// Case 3: (too) large write for buffer
	// Will skip buffer and write directly
	bw.Write([]byte("12345"))
}

func BufioWriteString() {
	data := []string{"an old falcon", "misty mountains",
		"a wise man", "a rainy morning"}

	f, err := os.Create("words.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	wr := bufio.NewWriter(f)

	for _, line := range data {

		wr.WriteString(line + "\n")
	}

	wr.Flush()

	fmt.Println("data written")
}

func BufioReading() {
	const singleLine string = "I'd love to have some coffee right about now"
	const multiLine string = "Reading is my...\r\n favourite"

	fmt.Println("Lenght of singleLine input is " + strconv.Itoa(len(singleLine)))
	str := strings.NewReader(singleLine)
	br := bufio.NewReaderSize(str, 25)

	fmt.Println("\n---Peek---")
	// Peek - Case 1: Simple peek implementation
	b, err := br.Peek(3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%q\n", b) // output: "I'd"

	// Peek - Case 2: Peek larger than buffer size
	_, err = br.Peek(30)
	if err != nil {
		fmt.Println(err) // output: "bufio: buffer full"
	}

	// Peek - Case 3: Buffer size larger than string
	br_large := bufio.NewReaderSize(str, 50)
	_, err = br_large.Peek(50)
	if err != nil {
		fmt.Println(err) // output: EOF
	}

	// ReadSlice
	fmt.Println("\n---ReadSlice---")
	str = strings.NewReader(multiLine)
	r := bufio.NewReader(str)
	for {
		token, err := r.ReadSlice('.')
		if len(token) > 0 {
			fmt.Printf("Token (ReadSlice): %q\n", token)
		}
		if err != nil {
			break
		}
	}

	// ReadLine
	fmt.Println("\n---ReadLine---")
	str = strings.NewReader(multiLine)
	r = bufio.NewReader(str)
	for {
		token, _, err := r.ReadLine()
		if len(token) > 0 {
			fmt.Printf("Token (ReadLine): %q\n", token)
		}
		if err != nil {
			break
		}
	}

	// ReadBytes
	fmt.Println("\n---ReadBytes---")
	str = strings.NewReader(multiLine)
	r.Reset(str)
	for {
		token, err := r.ReadBytes('\n')
		fmt.Printf("Token (ReadBytes): %q\n", token)
		if err != nil {
			break
		}
	}

	// Scanner
	fmt.Println("\n---Scanner---")
	str = strings.NewReader(multiLine)
	scanner := bufio.NewScanner(str)
	for scanner.Scan() {
		fmt.Printf("Token (Scanner): %q\n", scanner.Text())
	}

}

func BufioScanner() {
	f, err := os.Open("words.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func BufioWriteRune() {

	runes := "🐜🐬🐄🐘🦂🐫🐑🦍🐯🐞"

	f, err := os.Create("runes.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	wr := bufio.NewWriter(f)

	for _, _rune := range runes {

		wr.WriteRune(_rune)
		wr.WriteRune('\n')
	}

	wr.Flush()

	fmt.Println("runes written")
}

func BufioReadFromString() {
	words := []string{}

	data := "A foggy mountain.\nAn old falcon.\nA wise man."

	sc := bufio.NewScanner(strings.NewReader(data))

	sc.Split(bufio.ScanWords) // We tell the scanner to scan by words using Split.

	n := 0

	for sc.Scan() {
		words = append(words, sc.Text())
		n++
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("# of words: %d\n", n)

	for _, word := range words {

		fmt.Println(word)
	}

}
