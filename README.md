# Widget Factory
A demo application making use of (all kinds of cool stuff)

# Running the Application

TODO: how to run

# Guidelines by Topic

Below is a concise explanation of how each major aspect of this application is organized.  Third party libraries are used where appropriate, in other cases it is merely convention.

This guide describes each aspect so you can easily make changes to modify it for your needs.

## Main and Setup with Wire

The application executable lives at `cmd/widgetfactoryd`.

To facilitate "wiring" your application as it grows, Google's [wire](https://github.com/google/wire) package is used.
A Setup function lives in wire.go and will look something like:

```go
func Setup() (*AppConfig, error) {
	wire.Build(
		NewMainStuff,
		NewDBConnString,
		NewDBDriverName,
	)
	return &MainStuff{}, nil
}
```

You can add NewNNN calls to the list to have them automatically wired.  Running `go generate` will produce wire_gen.go with the actual wiring code.

And then in main.go this Setup function is called to retrieve AppConfig and start the application:

```go
func main() {

	appConfig, err := Setup()
	if err != nil {
		log.Fatal(err)
	}

    // ...
}
```


## Vue UI

The Vue user interface is initialized in the "appui" folder using Vue's CLI tool (see https://cli.vuejs.org):

```bash
mkdir vue
cd vue
vue create .
```

### Run

You can then build and run the Vue UI using the usual tooling:

TODO

### Building for Distribution

TODO (use https://github.com/shurcooL/vfsgen to take the dist output and make it importable directly embedded in a Go package)

### Endpoints

TODO (describe how to wire the endpoints so the dev UI and prod UI can sensibly find the right URL to hit)

### Reusability in Other Vue Apps

Vue components that may be re-used by other Vue applications wishing to embed portions of relevant UI functionality can simple be imported at the appropriate path, i.e. in JS in another Vue project `import "path/to/appui/components/SomeComponent.Vue"`.


### Controllers &amp; Routing


Routing is done in the `routes.go` file.  Using `github.com/julienschmidt/httprouter`.  To add a new Route:


Controllers are 





### Model Generation

