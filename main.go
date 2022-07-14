package main

import "fmt"

func main() {
	//taking input from user
	var input string
	fmt.Printf("Enter input string : ")
	fmt.Scanln(&input)
	length := len(input)

	//calling the encoder function to convert the string
	encodedString := Base64Encoder(input, length)
	fmt.Println(encodedString)
}

//function to encode the string to base64 encoded format
func Base64Encoder(input string, length int) string {
	//character set of base64 encoding scheme
	charSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

	var (
		encodedString string
		padding       int
	)

	//taking 3 characters at a time for conversion to base64
	for i := 0; i < length; i += 3 {

		var value, count, noOfBits, index int

		//binary representation of string
		for j := i; j < length && j <= i+2; j++ {
			value <<= 8
			value |= int(input[j])
			count++
		}

		noOfBits = count * 8   //number of bits
		padding = noOfBits % 3 //number of paddings required to successfully encode into base64

		for noOfBits != 0 {
			if noOfBits >= 6 {
				temp := noOfBits - 6
				index = (value >> temp) & 63 //taking 6 bits at a time for encoding
				noOfBits -= 6
			} else {
				temp := 6 - noOfBits
				index = (value << temp) & 63
				noOfBits = 0
			}

			encodedString += string(charSet[index]) //appending the base64 characters to the resultant string
		}

	}

	for i := 0; i < padding; i++ {
		encodedString += "="
	}

	return encodedString
}
