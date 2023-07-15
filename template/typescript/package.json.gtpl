{
  "name": "{{.ApiName}}",
  "version": "{{.Info.Version}}",
  "description": "{{.Info.Desc}}",
  "main": "index.js",
  "scripts": {
    "build": "tsc",
    "test": "echo \"Error: no test specified\" && exit 1"
  },
  "author": "{{.Info.Author}}",
  "license": "ISC",
  "devDependencies": {
    "typescript": "^4.8.3"
  },
  "publishConfig": {
    "registry": "https://npm.cuishu.site"
  }
}