package hangman_web

import (
	"fmt"
	hangman_web "hangman-web/hangman-classic"
	"html/template"
	"net/http"
)

type HangmanWeb struct {
	Word    string
	Life    int
	Display string
}

const port = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	cookie_Value, err := r.Cookie("Cookie")
	if err != nil {
		fmt.Println("Sans cookie")
		test := hangman_web.Dictionnary("./words/words.txt")
		cookie(w, r, test)
		fmt.Println("Word : ",test)
		fmt.Print(r.FormValue("inputText"))
		fmt.Println()

		renderTemplate(w, "./hangman-web/templates/home", HangmanWeb{
			Word: test,
		})
	}else {
		//code sans cookie
		fmt.Println("Avec cookie")
		fmt.Println("Word : ", cookie_Value.Value)
		fmt.Print(r.FormValue("inputText"))
		fmt.Println()

		renderTemplate(w, "./hangman-web/templates/home", HangmanWeb{
			Word: cookie_Value.Value,
		})
	}
}


// Take Value Word & save value
func cookie(w http.ResponseWriter, r *http.Request, Word string) {
	cookie := http.Cookie{
		Name:     "Cookie",
		Value:    Word,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)
	cookies := r.Cookies()
	fmt.Println("Cookies :", cookies)
	// r
}

func renderTemplate(w http.ResponseWriter, tmpl string, p HangmanWeb) {
	t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		fmt.Println(err)
	} else {
		t.Execute(w, p)
	}
}

// https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/04.1.html
// https://medium.com/@edwardpie/processing-form-request-data-in-golang-2dff4c2441be
// https://www.youtube.com/watch?v=GnLHI_nekm8
// https://golangcode.com/add-a-http-cookie/
//https://youtu.be/ONAnstqcEcA
// *** https://www.alexedwards.net/blog/working-with-cookies-in-go ***

//func ReadCookieServer(w http.ResponseWriter, req *http.Request) {
//	// read cookie
//	var cookie,err = req.Cookie("SessionID")
//	if err == nil {
//		var cookievalue = cookie.Value
//		w.Write([]byte(cookievalue))
//	}
//}
