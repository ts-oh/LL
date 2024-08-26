package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome!!!</h1><p>Site made in golang すごいね。。。</p>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch contact me by email<a href=\"mailto:asdf@asdf.com\">asdf@asdf.com</a></p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
	<ul>
	  <li>
		<b>Is there a free version?</b>
		Yes! We offer a free trial for 30 days on any paid plans.
	  </li>
	  <li>
		<b>What are your support hours?</b>
		We have support staff answering emails 24/7, though response
		times may be a bit slower on weekends.
	  </li>
	  <li>
		<b>How do I contact support?</b>
		Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>
	  </li>
	</ul>
	`)
}

func galleriesIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	userID := chi.URLParam(r, "userID")
	w.Write([]byte(fmt.Sprintf("gallery id is %v", userID)))
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.With(middleware.Logger).Get("/galleries/{userID}", galleriesIdHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on: 3000...")
	http.ListenAndServe(":3000", r)
}
