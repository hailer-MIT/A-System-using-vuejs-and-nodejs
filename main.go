package main

import "fmt"

type Contact struct {
    Name    string
    Email   string
    Phone   string
}

func main() {
    contacts := []Contact{
        {"John Doe", "johndoe@example.com", "1234567890"},
        {"Jane Smith", "janesmith@example.com", "9876543210"},
    }

    // Add more contacts here...

    for _, contact := range contacts {
        fmt.Printf("Name: %s, Email: %s, Phone: %s\n", contact.Name, contact.Email, contact.Phone)
    }
}