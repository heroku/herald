package main

import "os"
import "strings"
import "fmt"
import (
   _ "github.com/mailgun/mailgun-go"
)


func main() {

    // Setup Mailgun environment variables. 
    var MAILGUN_API_KEY string
    var MAILGUN_PUBLIC_KEY string

    MAILGUN_API_KEY = os.Getenv("MAILGUN_API_KEY")
    MAILGUN_PUBLIC_KEY = os.Getenv("MAILGUN_PUBLIC_KEY")


    fmt.Println("hello world")
}