# Lecturer GitHub

- <a href="https://github.com/GoesToEleven/golang-web-dev">GitHub Link</a>

# Templates

<a href="https://pkg.go.dev/text/template#Template.Funcs">Templates</a>

### Templating with Concatenation

- Merge file to index.html <br>
`go run main.go > index.html`

### Func New

`func New(name string) *Template`
- New allocates a new, undefined template with the given name.

### Parsing Templates

- Have a html file with extension `.gohtml`

- ParseFiles (Parses to a container with all the parsed files)<br>
`func ParseFiles(filenames ...string) (*Template, error)`

- Execute
`func (t *Template) Execute(wr io.Writer, data any) error`

- Execute template to stdout
`err = tpl.Execute(os.Stdout, nil)`

### Parsing multiple files in a template

- `ExecuteTemplate` executes the specific template (file name). <br>
- This is because there is more than one template in that pointer template.
`func (t *Template) ExecuteTemplate(wr io.Writer, name string, data any) error`

- Parse 2 more files to pointer to template tpl (like a container for all the templates)
```
tpl, err := template.ParseFiles("one.txt")
if err != nil {
    log.Fatalln(err)
}

tpl, err = tpl.ParseFiles("two.txt", "hello.txt")
if err != nil {
    log.Fatalln(err)
}
```

### ParseGlob

- `ParseGlob` creates a new Template and parses the template definitions from the file identified by the pattern. <br>
`func ParseGlob(pattern string) (*Template, error)`

- `Must` is a helper that wraps a call to a function returning (*Template, error) and panics if the error is non-nil.
`func Must(t *Template, err error) *Template`
```
func Must(t *Template, err error) *Template {
	if err != nil {
		panic(err)
	}
	return t
}
```

### Type template

- `template.Template`

### Passing data into templates

- Syntax of HTML data, use `{{.}}`
```
<body>
    <h1>The meaning of life: {{.}}</h1>
</body>
```

- Passing in a piece of data to the html file
```
err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 195)
```

### Variables in templates

- Set variables in HTML for parsing in Golang
```
<body>
    {{$wisdom := .}}
    <h1>The meaning of life: {{$wisdom}}</h1>
</body>
```
```
err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "Release self-focus; embrace other-focus")
```

### Passing composite data structures into templates

- Ranging over a piece of data in html
- `.` dot is the current piece of data
```
<body>
    <ul>
        {{range .}}
        <li>{{.}}</li>
        {{end}}
    </ul>
</body>
```
- Using variables to assign
```
<ul>
    {{range $index, $element := .}}
    <li>{{$index}} - {{$element}}</li>
    {{end}}
</ul>
```


- Passing in slices into templates
```
	sages := []string{"Leon", "James", "Jennifer", "Daniel"}

	// `func (t *Template) Execute(wr io.Writer, data any) error`
	err := tpl.Execute(os.Stdout, sages)
```

- For structs, the html template will be according to those defined in the struct
```
<body>
    <ul>
        <li>{{.Name}} -- {{.Motto}}</li>
    </ul>

    <ul>
        {{$x := .Name}}{{$y := .Motto}}
        <li>{{$x}} -- {{$y}}</li>
    </ul>
</body>
```
```
type sage struct {
	Name  string
	Motto string
}
```

### Functions in Templates

- `FuncMap`: FuncMap is the type of the map defining the mapping from names to functions.  <br>
`type FuncMap map[string]any`
```
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}
```
- Have to use FuncMap before parsing the file.
```
func init() {
	// Func takes a value of type funcmap
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}
```

- `Funcs` <br>
`func (t *Template) Funcs(funcMap FuncMap) *Template`

### Pipelines in Templates

- date formatting
- 01 is month, 02 is day, 06 is year
`Layout      = "01/02 03:04:05PM '06 -0700" // The reference time, in numerical order.`
```
func monthDayYear(t time.Time) string {
	// formatting string
	return t.Format("02-01-2006")
}
```

### Predefined Global Functions

#### INDEX
- Go
```
func main() {

	xs := []string{"zero", "one", "two", "three", "four", "five"}

	err := tpl.Execute(os.Stdout, xs)
	if err != nil {
		log.Fatalln(err)
	}
}
```
- HTML
```
{{index . 2}}
{{index . 0}}
{{index . 1}}
```
### Template Comments

- {{/* Comments */}}

### Nested Templates

- define:
```
{{define "TemplateName"}}
insert content here
{{end}}
```

- use:
```
{{template "TemplateName"}}
```



























