package store

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var controller = &Controller{Handler: Handler{}}

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route

var routes = Routes{
	Route{
		"AddStory",
		"POST",
		"/story",
		controller.AddStory,
	},
	Route{
		"JoinStory",
		"POST",
		"/story/join",
		controller.JoinStory,
	},
	Route{
		"AddParagraph",
		"POST",
		"/story/paragraph",
		controller.AddParagraph,
	},
	Route{
		"GetParsedStory",
		"GET",
		"/story/pretty",
		controller.GetParsedStory,
	},
	Route{
		"GetStories",
		"GET",
		"/stories",
		controller.GetAllStories,
	},
	Route{
		"GetStory",
		"GET",
		"/story",
		controller.GetStoryByField,
	},
	Route{
		"WebSocket",
		"GET",
		"/ws",
		controller.handleConnections,
	},
}

// NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
