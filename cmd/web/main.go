package main

import "github.com/kataras/iris"
import "os"
import "fmt"
import "github.com/heroku/herald"

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// PORT to bind to.
var PORT = getEnv("PORT", "8080")

func main() {
	app := iris.Default()

	// Method:   GET
	// Resource: http://localhost:8080/
	app.Handle("GET", "/", func(ctx iris.Context) {
		results := []string{}
		heraldBuildpacks := herald.GetBuildpacks()

		for _, bp := range heraldBuildpacks {
			results = append(results, bp.Name)
		}

		ctx.JSON(iris.Map{"buildpack": results})
	})

	app.Get("/buildpacks/{bp:string}", func(ctx iris.Context) {
		bp := ctx.Params().Get("bp")

		buildpack := herald.NewBuildpack(bp)
		// targets := buildpack.GetTargets()

		// ctx.JSON(iris.Map{"buildpack": buildpack.Name, "targets": targets})
		ctx.JSON(iris.Map{"buildpack": buildpack.Name})
	})

	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	// Method:   GET
	// Resource: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello iris web framework."})
	})

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(iris.Addr(fmt.Sprintf(":%s", PORT)))
}
