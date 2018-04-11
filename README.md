# Go / Gin / GORM Sample Application

To run this, you'll need the Go Programming Language, dep for dependency management, and a PostgreSQL database. Assuming you're using Homebrew on a Mac, this should install those for you:

```
brew install go dep postgresql
```

Start the database like this:

```
brew services start postgresql
```

Create a user and database named goweb:

```
createuser goweb --createdb
createdb goweb -U goweb
```

Use dep to install depenencies:

```
dep ensure
```

And you're finally ready to run the app:

```
go run main.go
```
