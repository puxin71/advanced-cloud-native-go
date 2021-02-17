# github reference
https://github.com/JacobSNGoodwin/memrizr

# Using CURL to ad-hoc test web service

# Launch HTTP web server
PORT=3000 ./Gin-Web

## ping
curl -v http://localhost:3000/v1/ping

## get all books
curl -v http://localhost:3000/v1/books

## add a new book
curl -v \
 -X POST \
 -H "Content-Type:application/json" \
 -d '{"title":"GOOGLE INC","isbn":"2345UUx90"}' \
 http://localhost:3000/v1/books

curl -v \
 -X POST \
 -H "Content-Type:application/json" \
 -d '{"title":"TOY STORY","isbn":"2345UUx91"}' \
 http://localhost:3000/v1/books

## get a book by ISBN
curl -v \
 -X GET \
 http://localhost:3000/v1/books/2345UUx90

## get a slice of books by ISBNs
curl -v 'http://localhost:3000/v1/books?isbns=2345UUx90&isbns=2345UUx91'

# Personal helps
https://alvinalexander.com/git/

## remove files from the previous commit
git rm -r -f file

## push to github repo
```
git remote add origin https://github.com/puxin71/advanced-cloud-native-go.git
git branch -M main
git push -u origin main
```