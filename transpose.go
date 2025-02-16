package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "flag"
)

// transposeWord swaps the middle letters of a word according to a deterministic rule
func transposeWord(word string) string {
    runes := []rune(word) // Convert string to rune slice (for Unicode safety)
    length := len(runes)

    if length <= 2 {
        // Words with 2 or fewer letters remain unchanged
        return string(runes)
    }

    // Swap the last letter with the second for 3-letter words
    if length == 3 {
        runes[1], runes[2] = runes[2], runes[1]
        return string(runes)
    }

    // For longer words: first and last letters remain unchanged
    middle := runes[1 : length-1]

    // Deterministic swapping: swap each even pair with the odd one
    for i := 0; i < len(middle)-1; i += 2 {
        middle[i], middle[i+1] = middle[i+1], middle[i]
    }

    // Reconstruct the new string
    return string(runes[0]) + string(middle) + string(runes[length-1])
}

// detransposeWord reverses the deterministic rule applied by transposeWord
func detransposeWord(word string) string {
    runes := []rune(word) // Convert string to rune slice (for Unicode safety)
    length := len(runes)

    if length <= 2 {
        // Words with 2 or fewer letters remain unchanged
        return string(runes)
    }

    // Swap the last letter with the second for 3-letter words
    if length == 3 {
        runes[1], runes[2] = runes[2], runes[1]
        return string(runes)
    }

    // For longer words: first and last letters remain unchanged
    middle := runes[1 : length-1]

    // Reverse the deterministic swapping: swap each odd pair with the even one
    for i := 0; i < len(middle)-1; i += 2 {
        middle[i], middle[i+1] = middle[i+1], middle[i]
    }

    // Reconstruct the new string
    return string(runes[0]) + string(middle) + string(runes[length-1])
}

// encodeText transposes all words in a sentence
func encodeText(text string) string {
    words := strings.Fields(text) // Split sentence into words
    for i, word := range words {
        words[i] = transposeWord(word)
    }
    return strings.Join(words, " ") // Join words back into a sentence
}

// decodeText detransposes all words in a sentence
func decodeText(encodedText string) string {
    words := strings.Fields(encodedText) // Split sentence into words
    for i, word := range words {
        words[i] = detransposeWord(word)
    }
    return strings.Join(words, " ") // Join words back into a sentence
}

// printUsage prints the usage instructions
func printUsage() {
    fmt.Println("Usage: transpose [-d]")
    fmt.Println("  - Default: Transposes text from stdin and outputs it to stdout.")
    fmt.Println("  - -d: Detransposes text from stdin and outputs it to stdout.")
    fmt.Println("Example:")
    fmt.Println("  echo 'This is a Test.' | transpose")
    fmt.Println("  echo 'Thsi si a Tsset.' | transpose -d")
}

// main function
func main() {
    // Flag for decoding (-d)
    decodeFlag := flag.Bool("d", false, "Decode instead of encode")
    flag.Parse()

    // Scanner for stdin
    scanner := bufio.NewScanner(os.Stdin)
    var result strings.Builder

    // Process the input text
    for scanner.Scan() {
        line := scanner.Text()

        if *decodeFlag {
            result.WriteString(decodeText(line) + "\n")
        } else {
            result.WriteString(encodeText(line) + "\n")
        }
    }

    // Handle errors when reading from stdin
    if err := scanner.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
        os.Exit(1)
    }

    // Output the result to stdout
    fmt.Print(result.String())
}