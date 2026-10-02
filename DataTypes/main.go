package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// 1. Two English Characters: "go"
	englishStr := "go"

	// 2. Two Hindi Characters: "नम" (Na and Ma)
	// Note: We are using just the first two base letters of "नमस्ते" to keep it to exactly 2 characters.
	hindiStr := "नम"

	fmt.Println("--- ENGLISH STRING COMPONENT ---")
	fmt.Println("String value:    ", englishStr)
	fmt.Println("len() [Bytes]:   ", len(englishStr))                    // Returns 2 because each English character takes 1 byte
	fmt.Println("Rune Count:      ", utf8.RuneCountInString(englishStr)) // Returns 2 characters
	fmt.Println("[]rune Conversion:", []rune(englishStr))                // Returns [103 111] (The 32-bit Unicode code points)
	fmt.Println("[]byte Conversion:", []byte(englishStr))                // Returns [103 111] (The raw 8-bit bytes)

	fmt.Println("\n--- HINDI STRING COMPONENT ---")
	fmt.Println("String value:    ", hindiStr)
	fmt.Println("len() [Bytes]:   ", len(hindiStr))                    // Returns 6! Because each Hindi character takes 3 bytes in UTF-8
	fmt.Println("Rune Count:      ", utf8.RuneCountInString(hindiStr)) // Returns 2 characters
	fmt.Println("[]rune Conversion:", []rune(hindiStr))                // Returns [2328 2350] (The 32-bit Unicode code points for न and म)
	fmt.Println("[]byte Conversion:", []byte(hindiStr))                // Returns [224 164 168 224 164 170] (6 raw underlying bytes)

	fmt.Println("\n--- STRINGS PACKAGE UTILITIES ---")
	str := "hello,world"
	fmt.Println("upperCase--", strings.ToUpper(str))
	fmt.Println("lowerCase--", strings.ToLower(str))
	fmt.Println("contains-- ", strings.Contains(str, "ll"))

	split := strings.Split(str, ",")
	fmt.Println("split--   ", split)
	fmt.Println("split len--", len(split))

	rep := strings.Replace("Git and Github", "Github", "Jenkins", 1)
	fmt.Println("rep--      ", rep)

	trim := strings.TrimSpace("   lakahyam     ")
	fmt.Println("trim--     ", trim)
}
