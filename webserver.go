package main

import (
  "go-simple-calculator"
  "flag"
	"fmt"
	"html/template"
	"net/http"
)

type apiHandler struct{}

func (*apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// validate method
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// read and validate the query string
	number := r.FormValue("number")
	if number == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := `
    <form method="post">
  		<input name="number" value="{{ .Number }}" required>
      <input type="submit" value="Calculate!">
		</form>
		{{ if .Number }}
      <h2>Result:</h2>
      <h3>{{ .Number }} = {{ .Result }}</h3>
    {{ end }}`

	// set the encoding
	w.Header().Add("Content-type", "text/html")

	// validate the method
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// prepare the data
	data := struct {
		Number string
    Result string
	}{
		Number: r.FormValue("number"),
    Result : calc.ChangeStringtoArray(r.FormValue("number")),
	}

	// parse the template
	t, err := template.New("form").Parse(tmpl)
	if err != nil {
		fmt.Println("Failed to parse template;", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	t.Execute(w, data)
}

func startServer(address string) {
	http.Handle("/api/calculate", &apiHandler{})
	http.HandleFunc("/", htmlHandler)

	fmt.Println("Starting server on http://" + address)
	http.ListenAndServe(address, nil)
}

func main() {
  var addr = flag.String("addr", "", "Interface and port to listen on")

	// parse the flags
	flag.Parse()

	startServer(*addr)
}
