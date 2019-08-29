# Widget Factory
A demo application making use of (all kinds of cool stuff)

# Running the Application

TODO: how to run

# Guidelines by Topic

Below is a concise explanation of how each major aspect of this application is organized.  Third party libraries are used where appropriate, in other cases it is merely convention.

To be clear, these steps have already been done on this project, but are documented here a guide to show how to rapidly construct your own application following this convention.

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
