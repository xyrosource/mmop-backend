package mmop

/**
 * TODO:
 *   - add a favicon
 *   - update robots.txt, if needed
 *   - put a stylesheet in static/styles.css
 */

import (
	"fmt"
	"net/http"
)

func getHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, `<html>
	<head>
		<link rel="stylesheet" type="text/css" href="static/style.css" />
		<link rel="icon" type="image/png" href="favicon.ico" />
		<title>Page title</title>
	</head>
	<body>
		Under construction.
	</body>
</html>`)
}

func serveSingle(pattern string, filename string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}

func init() {
	http.HandleFunc("/", getHandler)

	serveSingle("/robots.txt", "./robots.txt")
	serveSingle("/favicon.ico", "./favicon.ico")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
}
