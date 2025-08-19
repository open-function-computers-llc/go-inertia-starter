# Starter for a go and inertia app

This is a (pretty good?) starting point for a go/inertia/vue app. We like scss, so we use it. We like vue, so we use that too. Vite is rad, so we use that for hot reload, bundles, and all that other JS stuff.

## How to use this

Before you do anything else, copy the `.env.example` file to `.env`. Change anything that you want to change, but note that the `APP_PORT` value must match the same port that is used in the `APP_URL` key.

Next, install all the node deps to build the frontend. In a terminal, run `npm install` to download all the dependancies. The built in vite dev server will be used to serve the frontend assets, but it needs to have the go server running to get the initial HTML (and all future JSON responses) to render the page.

We have included a really simple `build.sh` script which is used for building a production build. You can use it for local dev, but the steps below will give you a better workflow.

## Local Dev Process

First, we need to get a go server built and running. Run this from the project root:

```
go build -o cmd/go-inertia-app
./cmd/go-inertia-app
```

Now you should have a built server and it should be running. Woot!

Next, run `npx vite` from a new terminal window to have node watch the vue, scss, and js files. When your go program is running in `development` mode, the app.scss and app.js files will be served from vite, so you'll get hot reload!

## Prod builds

The provided a little helper script `build.sh` can help build and embed the app into a binary that you can plop anywhere. If you need to build for a different OS or archtecture than your dev machine, just modify the build line to target your server's specs.

Running the build script will turn off CGO, set the app version to the latest git-commit hash, build the frontend assets with vite for a production build, and then embed that bundle into the go binary.

When you run the binary with an `APP_ENV` that isn't `development` (prod, staging, whatever) the code will look for assets that are bundled into the binary, meaning you won't have to push anything except this binary out to your prod servers.

## Roadmap

- Set up auth?
- Auth middleware?
- Sensible defaults?
- Another folder to dump other static assets that should be embedded?
- Do we want to involve the database? Do we need migrations? Is this a good enough spot to fork from and take each app in it's own direction?
