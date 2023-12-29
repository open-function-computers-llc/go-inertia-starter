# Starter for a go and inertia app

This is a (pretty good?) starting point for a go/inertia/vue app. We like scss, so we use it. We like vue, so we use that too.

## How to use this

First, install all the node deps to build the frontend. We're using Laravel Mix to build and bundle up the assets. In a terminal, run `npm install` to download all the dependancies, then run `npx mix watch` from a terminal window to have node watch the vue, scss, and js files. It'll put an updated bundle in the `dist` folder on each file save.

Environment variables are necessary to tell the go app how to run, so copy `.env.example` to `.env`, and change the variables as you see fit. If you run the app with the default settings, the app will run at localhost on port 8900.

Next, build the go app. We have provided a little helper script `build.sh` to help build and embed the app into a binary that you can plop anywhere. If you need to build for a different OS or archtecture than your dev machine, just modify the build line to target your server's specs. I like to leave an empty terminal open where I can type in `./build.sh` anytime I need to rebuild the go app (which is anytime that we get a new js bundle or add any routes to the go app, so basically any changes at all). The script returns your cursor with the trailing `&` so you can always just hit the up arrow and enter to quickly kill the existing process, rebuild, and rerun the new build of the application. There's definitly room for better DX here.

## Roadmap

- Set up auth?
- Auth middleware?
- Better DX with rebuilds and hot reload?
- Sensible defaults?
- Another folder to dump other static assets that should be embedded?
- Do we want to involve the database? Do we need migrations? Is this a good enough spot to fork from and take each app in it's own direction?
