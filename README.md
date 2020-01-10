# fizzbuzz-go

Version Go of Java project : [FizzBuzz](https://github.com/julien-boulet/fizzbuzz)

#i18n
https://phrase.com/blog/posts/internationalisation-in-go-with-go-i18n/

#csrf
https://github.com/gorilla/csrf

#update import
go list -m -u all

#run all tests
go test $(go list ./... | grep -v /vendor/)


#test application

GET --> localhost:8083/fizzbuzz?int1=3&int2=5&limit=16&str1=fizz&str2=buzz
GET --> localhost:8083/oneTopStatistic