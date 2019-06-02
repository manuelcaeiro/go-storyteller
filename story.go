package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var tpl *template.Template

var defaultHandlerTmpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title>Choose Your Own Adventure</title>
	</head>
	<body>
		<section class="page">
		<h1>{{.Title}}</h1>
		{{range .Paragraphs}}
			<p>{{.}}</p>
		{{end}}
		<ul>
		{{range .Options}}
			<li><a href="/{{.Chapter}}">{{.Text}}</a></li>
		{{end}}
		</ul>
		</section>
		<style>
			body {
				<!--font-family: helvetica, arial;-->
				font-family: "Lucida Console", Monaco, monospace;
				font-size: calc(12px + .5vw);
			}
			h1 {
				text-align: center;
				position: relative;
			}
			.page {
				width: 80%;
				max-width: 500px;
				margin: auto;
				margin-top: 40px;
				margin-bottom: 70px;
				padding: 40px;
				background: #FFFCF6;
				border: 1px solid #ffffff;
				box-shadow: -10px 10px 6px -6px #797979;
			}
			ul {
				border-top: 1px dotted #cccccc;
				padding: 10px 0 0 0;
				-webkit-padding-start: 0;
			}
			li {
				padding-top: 10px;
			}
			a,
			a: visited {
				text-decoration: none;
				color: #555555;
			}
			a: active,
			a: hover {
				color: #222222;
			}
			p {
				text-indent: 1em;
			}
		</style>
	</body>
</html>`

type HandlerOption func(h *handler)

type handler struct {
	s      Story
	t      *template.Template
	pathFn func(r *http.Request) string
}

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTmpl))
}

func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

func WithPathFunc(fn func(r *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.pathFn = fn
	}
}

func NewHandler(s Story, opts ...HandlerOption) http.Handler {
	h := handler{s, tpl, defaultPathFn}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

func defaultPathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	return path[1:]
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.pathFn(r)
	if chapter, ok := h.s[path]; ok {
		err := h.t.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found.", http.StatusNotFound)
}

func JsonStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func main() {
	port := flag.Int("port", 5016, "the port to start the application")
	filename := flag.String("file", "gideon.json", "the json file with CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := JsonStory(f)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", NewHandler(story))
	fmt.Printf("Starting the server on port: %d\n", *port)
	l, err := net.Listen("tcp", "localhost:5016")
	if err != nil {
		log.Fatal(err)
	}
	err = open("http://localhost:5016/")
	if err != nil {
		log.Println(err)
	}
	http.Serve(l, mux)
}
