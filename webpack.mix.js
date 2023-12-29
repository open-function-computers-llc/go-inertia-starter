let path = require("path");
let mix = require('laravel-mix');

mix.alias({
   "@": path.join(__dirname, "frontend"),
});

mix.js('frontend/app.js', 'dist')
   .version()
   .sourceMaps()
   .vue()
   .setPublicPath('dist');

mix.sass('frontend/scss/app.scss', 'dist')
   .version()
   .setPublicPath('dist');
