package main

import "os"
// import "strings"
import "fmt"
import "github.com/mailgun/mailgun-go"


func main() {

    // Setup MailGun environment variables. 
    var MAILGUN_API_KEY string
    var MAILGUN_PUBLIC_KEY string
    var MAILGUN_DOMAIN string

    MAILGUN_API_KEY = os.Getenv("MAILGUN_API_KEY")
    MAILGUN_PUBLIC_KEY = os.Getenv("MAILGUN_PUBLIC_KEY")
    MAILGUN_DOMAIN = os.Getenv("MAILGUN_DOMAIN")

    // Initialize MailGun API. 
    mg := mailgun.NewMailgun(MAILGUN_DOMAIN, MAILGUN_API_KEY, MAILGUN_PUBLIC_KEY)
    _ = mg


    fmt.Println("hello world")
}