package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// This tiny abstraction
	// This allows you to reutnr errors, as things go wrong
	// you can'd do that directly from main, so use this instead
	// like if you wanted to print errors a lot to stderror, this means you onlt have to write that code once instead o flots of time
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	db, dbtidy, err := setupDatabase()
	if err != nil {
		return error.Wrap(err, "setup database")
	}
	defer dbtidy()
	srv := &server {
		db: db
	}
	// more stuff 
}

// represents the component as as struct 
// shared dependenceies as fields
// no global state
// never use global variables in golang 
/*
	Constructor for server? 

	func newServer() *server{
		s := &server{}
		s.router() 
		return s
	}

	If you use a constructor, its hard to tell whats happening...
		-> you may be creating goroutines for isnteance
	however, with types, you know exactly whats hapening

*/
type server struct {
	db  *someDatabase
	router *someRouter
	email EmailSender
}

// Make server an http.Handler
// But pass execution to the router, don't try to handle it here, its a strnage place to do it
// If you want more logic here, use middleware 
// This essentially implements ServerHTTP to turn your server into an http.Handler
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServerHTTP(w, r)
}

// Have your own routes.go file
// one place for al routes 
// most code maintinance starts with a URL so its handy to have one place to look 
// In go, try to group things by responsibility
// very glanceable code 
package main 

func (s *server) routes(){
	s.router.Get("/api", s.handleAPI())
	s.router.Get("/about", s.handleAbout())
}

// every http request that comes in to your server gets its own goroutine 
// Handlers hang off the server 
// Handlers are methods on the server, which gives them access to the dependencies via s 
// Remember other handlers have access to s too, so be careful of dataraces
func (s *server) handelSomething() http.HandlerFunc {
	// put something here
}

/* Naming handlers...
since autocomplete lists and docs are alphabetically, you can group related funcitonality nicely

handleTasksCreate
handleTasksGet
handleTasksDelete 

handleAuthLogin
handleAuthLogout
*/

/*
	Also, notice, that we reutnr the handler, instead of passing it smoething 
	this allos for handler-specific setup, and allows closures 
	gives us an annoymous function
*/

func (s *server) handleSomething() http.HandlerFunc {
	thing := prepareThing()
	return func(w http.ResponseWriter, r *http.Request) {
		// use thing
	}
}

/*
	Specific arguments for handler responses
*/
func (s *server) handleGreeting(format string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, format, r.FormValue("name"))
	}
}

s.router.HandleFunc("/one", s.handleGreeting("Hello %s"))

// if you handler grows too big, and you have lots and lots of things grown on it, then have multiple server

// We should always use HandlerFunc over Handler 
// http.HandlerFunc implements http.Handler, so they are interchangeable 
// Pick whichever is easier to read in your case 

/*
Middleware are just Go function
take an http.HandlerFun and return a new one 
run code before.after the wrapped handler
or choose not to call teh wrapped handler at all 
has its own little closured environment too 
can access shared dependencies
*/

func (s *server) adminOnly(h http.HandlerFunc) http.HaandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		if !currentUser(r).IsAdmin {
			http.NotFound(w, r)
			return
		}
		h(w, r)
	}
}

/*
	Wire Middlewre up in routes.g
	let routes.go become a high level map of the service
*/
func (s *server) routes() {
	s.router.Get("/api", s.handleAPI())
	s.router.Get("/admin", s.adminOnly(s.handleAdminIndex()))
}

/*
Dealing with data
	try not to abstract too early
	we're obsessed with DRY code 
	solve things a few times first, and then try and find the abstraction
	instead of seeing it before you do it even once 
Respond Helper
*/
func (s *server) respond(w http.ResponseWriter, r *http.Request), data interface{}, status int) {
	w.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
	}
}
// Same for decoding 
func ( s *server) decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
// Future proof helpers, by always taking http.ReqponseWriter and *http.Request now, even if you don't need them
// 

/*
	Request and response data types
	Put them inside the handle functions 
	colocated stuff is easier to find
	creates its own closure
	declutters the package space 
	don't need long or unique names for them now
	you're storytelling here, as you see only whats relevant for this handler
	its great
*/

func (s *server) handleGreet() http.HandlerFunc {
	type request struct {
		Name string 
	}
	type response stuct {
		Greeting string `json:"greeting"`
	}
	return func(w http.ResponseWriter, r *http.Request){
		... 
	}
}

/*
Testing
Its a great tool, and you need it 
httptest is great apparently 
*/