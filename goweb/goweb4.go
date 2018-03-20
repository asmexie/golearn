package main

import (
	"net/http"
	"text/template"
)

func main()  {
	http.HandleFunc("/", myHandlerFunc)
	http.ListenAndServe(":8080",nil)
}

func myHandlerFunc(w http.ResponseWriter, r *http.Request)  {
	w.Header().Add("Content Type","text/html")
	tmpl, err := template.New("anyNameForTemplate").Parse(doc)
	if err != nil{
		tmpl.Execute(w, nil)
	}
}
const doc =`
<!DOCTYPE html>
<html>
<head lang="en">
	<meta charset="UTF-8">
	<title>First Template</title>
</head>
<body>
	<h1>Hello Jamaica</h1>
<body>
</html>
`