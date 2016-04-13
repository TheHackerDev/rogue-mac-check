package main

import (
    "os"
    "bufio"
    "fmt"
    "strings"
    "regexp"
) 

var (
    approvedMacMap map[string]string
    foundMacMap map[string]string
    unmatchedMacMap map[string]string
)

func init()  {
    // key = mac; value = full line in text file
    approvedMacMap = make(map[string]string)
    foundMacMap = make(map[string]string)
    unmatchedMacMap = make(map[string]string)
}

func main() {
    // Print welcome message
    fmt.Println(welcome())
    
    // Only continue if there are the right amount of command-line arguments
    if len(os.Args) == 3 {
        // Store approved MACs
        if err := storeInMap(os.Args[1], approvedMacMap); err != nil {
            panic(err)
        }
        
        // Store found MACs
        if err := storeInMap(os.Args[2], foundMacMap); err != nil {
            panic(err)
        }
        
        // Check for unmatched MAC addresses
        for mac, line := range foundMacMap {
            findUnmatched(mac, line)
        }
        
        // Print the unmatched MAC addresses to the screen
        fmt.Printf("Unmatched MAC address lines:\n\n")
        for _, line := range unmatchedMacMap {
            fmt.Printf("%s\n", line)
        }
        fmt.Printf("\n%d unmatched MAC addresses found.\n", len(unmatchedMacMap))
    }
}

// Prints a welcome message, with instructions for the use of the program.
func welcome() (msg string) {
    msg = fmt.Sprintf("Program usage:\n")
    msg += fmt.Sprintf("\t%s /path/to/allowed_macs.txt /path/to/found_macs.txt\n", os.Args[0])
    msg += fmt.Sprintf("\tInput: both text files must contain lines of newline-separated text with a valid MAC address on each line. Other text values on the line is allowed.\n")
    msg += fmt.Sprintf("\tOutput: the full text file line provided for each MAC addresses from the \"found_macs.txt\" file that were not present in the \"allowed_macs.txt\" file.\n")
    msg += fmt.Sprintf("\tNotes:\n")
    msg += fmt.Sprintf("\t\tThe name of the files do not matter, as long as they are standard plain-text files.\n")
    msg += fmt.Sprintf("\t\tThe text files must be in the same directory as the executable program.\n")
    
    return
}

// Store MAC addresses and their surrounding line text in the provided map
func storeInMap(fileLocation string, macMap map[string]string) error{
    // Open the first argument (file location) for reading.
    f, err := os.Open(fileLocation)
    if err != nil {
        return err
    }
    defer f.Close()
    // Create a reader to read the file
    scanner := bufio.NewScanner(f)
    // Iterate through each line
    for scanner.Scan() {
        line := scanner.Text()
        // Find the MAC address on the line
        regex, err := regexp.Compile("([a-zA-Z0-9][a-zA-Z0-9]:){5}[a-zA-Z0-9][a-zA-Z0-9]")
        if err != nil {
            return err
        }
        mac := regex.FindString(line)
        // Convert to uppercase
        mac = strings.ToUpper(mac)
        // Add the resulting strings to the provided MAC map.
        macMap[mac] = line
    }
    return nil
}

// Locate and check a MAC address against the approved MAC addresses,
// and adds it to the map of unmatched MAC addresses if unmatched.
func findUnmatched(foundMac string, line string) {
    // Look up the MAC string in the map of approved MACs.
    if _, exists := approvedMacMap[foundMac]; !exists {
        // MAC not found in approved mac list, add to unmatched MAC map
        unmatchedMacMap[foundMac] = line
    }
}