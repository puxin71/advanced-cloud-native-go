# github reference
https://github.com/JacobSNGoodwin/memrizr

# Launch HTTP web server
PORT=3000 ./Gin-Web

# CURL test endpoints 
## ping
curl -v http://localhost:3000/v1/ping

## get all books
curl -v http://localhost:3000/v1/books

# Personal helps
## remove files from the previous commit
git rm -r -f file

## push to github repo
```
git remote add origin https://github.com/puxin71/advanced-cloud-native-go.git
git branch -M main
git push -u origin main
```