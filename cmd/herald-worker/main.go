package main

import "os"
import "fmt"
import "time"
import "github.com/gocelery/gocelery"
import "github.com/hashicorp/go-getter"


func get_buildpacks() [2]string {
    var buildpacks [2]string

    buildpacks[0] = "https://github.com/heroku/heroku-buildpack-python/archive/master.zip"
    buildpacks[1] = "https://github.com/heroku/heroku-buildpack-php/archive/master.zip"

    return buildpacks
}

func download_buildpacks() {
    buildpacks := get_buildpacks()
    _ = buildpacks
    getter.Get("example", "https://github.com/heroku/heroku-buildpack-python/archive/master.zip")
}

func add(a int, b int) int {
    fmt.Println("hello world")
    return a + b
}


func main() {

    // Setup MailGun environment variables. 
    // var MAILGUN_API_KEY string
    // var MAILGUN_PUBLIC_KEY string
    // var MAILGUN_DOMAIN string

    // MAILGUN_API_KEY = os.Getenv("MAILGUN_API_KEY")
    // MAILGUN_PUBLIC_KEY = os.Getenv("MAILGUN_PUBLIC_KEY")
    // MAILGUN_DOMAIN = os.Getenv("MAILGUN_DOMAIN")

    download_buildpacks()
    // Setup AMPQ environment variables.
    var RABBITMQ_BIGWIG_URL string

    RABBITMQ_BIGWIG_URL = os.Getenv("RABBITMQ_BIGWIG_URL")

    // Configure Celery Broker and Client. 
    celeryBroker := gocelery.NewAMQPCeleryBroker(RABBITMQ_BIGWIG_URL)
    celeryBackend := gocelery.NewAMQPCeleryBackend(RABBITMQ_BIGWIG_URL)
    celeryClient, _ := gocelery.NewCeleryClient(celeryBroker, celeryBackend, 2)

    // Configure Celery tasks. 
    celeryClient.Register("worker.add", add)

    // Start the worker.
    go celeryClient.StartWorker()

    asyncResult, err := celeryClient.Delay("worker.add", 3, 5)
    if err != nil {
        panic(err)
    }

    _ = asyncResult


    // TODO: Main loop here, of checking buildpacks here for status updates.
    time.Sleep(10 * time.Second)


}