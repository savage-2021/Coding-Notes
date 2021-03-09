package main

import "fmt"

func main() {
	// placeOfInterest()
	forRange()
    // const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

    // fmt.Println("Println:")
    // fmt.Println(sample)

    // fmt.Println("Byte loop:")
    // for i := 0; i < len(sample); i++ {
    //     fmt.Printf("%x ", sample[i])
    // }
    // fmt.Printf("\n")

    // fmt.Println("Printf with %x:")
    // fmt.Printf("%x\n", sample)

    // fmt.Println("Printf with % x:")
    // fmt.Printf("% x\n", sample)

    // fmt.Println("Printf with %q:")
    // fmt.Printf("%q\n", sample)

    // fmt.Println("Printf with %+q:")
    // fmt.Printf("%+q\n", sample)
}

func placeOfInterest(){
	const placeOfInterest = `⌘`
	    fmt.Printf("plain string: ")
    fmt.Printf("%s", placeOfInterest)
    fmt.Printf("\n")

    fmt.Printf("quoted string: ")
    fmt.Printf("%+q", placeOfInterest)
    fmt.Printf("\n")

	fmt.Printf("Unicode string: ")
    fmt.Printf("%U", placeOfInterest)
    fmt.Printf("\n")

    fmt.Printf("hex bytes: ")
    for i := 0; i < len(placeOfInterest); i++ {
        fmt.Printf("%x ", placeOfInterest[i])
    }
    fmt.Printf("\n")
}

func forRange(){
	const nihongo = "日本語"
	const thing = "x0e697a5e69cace8aa9e"
	fmt.Printf("%x\n", nihongo)
	fmt.Printf("%U\n", string(thing))
    for index, runeValue := range nihongo {
        fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
    }
}