package main

import "fmt"

type Contact struct {
    Name    string
    Email   string
    Phone   string
}

// ... rest of your code

func editContact(contacts []Contact, name string, newContact Contact) []Contact {
    for i, contact := range contacts {
        if contact.Name == name {
            contacts[i] = newContact
            return contacts
        }
    }
    return contacts
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