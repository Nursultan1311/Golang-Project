package controllers

import (
	"Goland/database"
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"time"
)

var tpl *template.Template
var db *sql.DB
var err error
const SecretKey = "Hello"

// With by this page we can add a new users
// Firstly you should have users table on your database with "username", "email", "name", "password" fields.

// Function for registration 
func Signup(res http.ResponseWriter, req *http.Request) {
	db = database.ConnectToDB() // we connect to the database
	if req.Method != "POST" { // if method GET we return html file
		tpl.ExecuteTemplate(res, "signup.html", nil)
		return
	}
	// if method is POST then we should get new user data from post request and insert it to the database
	//getting users data from post request by field name
	username := req.FormValue("username")
	password1 := req.FormValue("password1")
	password2 := req.FormValue("password2")
	email := req.FormValue("email")
	name := req.FormValue("name")

	var user string
	
	// comparing two passwords
	if password1 == password2 {
		// getting data from database with giving username
		err := db.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&user)
		switch {
			// if username doesn't have in database then we create a new user with this username
		case err == sql.ErrNoRows:
			// we hash the password 
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password1), bcrypt.DefaultCost)
			if err != nil {
				http.Error(res, "Server error, unable to create your account.", 500)
				return
			}
			// we add new datas to the database
			_, err = db.Exec("INSERT INTO users(username, password, email, name) VALUES(?, ?, ?, ?)", username, hashedPassword, email, name)
			if err != nil {
				http.Error(res, "Server error, unable to create your account.", 500)
				return
			}

			res.Write([]byte("User created!"))
			return
		case err != nil:
			// if user with this username exists, then we show error with this text
			http.Error(res, "Server error, unable to create your account.", 500)
			return
		default:
			http.Redirect(res, req, "/", 301)
		}
	} else {
		http.Error(res, "Password doesn't match. Both passwords should be same!", 500)
		return
	}
}

var Logerror string
// function for logging
func Login(res http.ResponseWriter, req *http.Request) {
	db = database.ConnectToDB() // connecting to the database 

	if req.Method != "POST" && Logerror != "" { // if request method id GET then we show login.html
		tpl.ExecuteTemplate(res, "login.html", Logerror)
		return
	}
	// if request method is POST then we are checking the user by password and username
	username := req.FormValue("username")
	password := req.FormValue("password")

	var databaseUsername string
	var databasePassword string
	// Getting user by username from database
	err := db.QueryRow("SELECT username, password FROM users WHERE username=?", username).Scan(&databaseUsername, &databasePassword)

	if err != nil { // id user doesn't exist we return the value "logger" with text below
		http.Redirect(res, req, "/login", 301)
		Logerror = "Dont have any user"
		return
	}
	// if user exist, then we hash the password and compare with password from database
	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
	if err != nil { // id password not same we return password is incorrect value
		http.Redirect(res, req, "/login", 301)
		Logerror = "Password is incorrect"
		return
	}
	// if user exist and password also correct we generate the new token and add it to http cookies
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    databaseUsername,
		ExpiresAt: time.Now().Add(time.Minute * 60).Unix(), //1 day
	})
	token, err := claims.SignedString([]byte(SecretKey))

	newCookie := http.Cookie{ // creating new cookie for this user
		Name:     "jar",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
		MaxAge:   99,
	}

	http.SetCookie(res, &newCookie) // add the cookie for http
	http.Redirect(res, req, "/", 301) // then redirect to main page
}
// Logout for removing cookie for user
func Logout(res http.ResponseWriter, req *http.Request) {
	newCookie := http.Cookie{ // creating new cookie with expired date
		Name:     "jar",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(res, &newCookie) // add this cookie for changing old cookie 
	http.Redirect(res, req, "/", 301)
}

func Home(res http.ResponseWriter, req *http.Request) { // home page gets user and put it home.html file 
	tpl, _ = tpl.ParseGlob("templates/*.html")

	username, _ := GetUser(req)
	tpl.ExecuteTemplate(res, "index.html", username)
}

func GetUser(req *http.Request) (string, error) { // function which return user if its has or empty value
	cookie, err := req.Cookie("jar") // we get the cookie from http for checking authentication of user
	if err != nil {
		fmt.Println("Something was wrong")
		return "", err
	}
	// if http has token then we take user name and return it for home function

	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		fmt.Println("Something was wrong")
		return "", err
	}

	claims := token.Claims.(*jwt.StandardClaims)
	fmt.Println(claims.Issuer)

	return claims.Issuer, nil
}
