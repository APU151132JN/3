//+build ignore

package main

import (
	. "."
	"net/http"
	"time"
)

func main() {
	doc := NewDoc("/", testtempl)
	go func() {
		err := http.ListenAndServe(":7070", nil)
		if err != nil {
			panic(err)
		}
	}()
	for {
		time.Sleep(1 * time.Second)
		doc.Elem("e_time").SetValue(time.Now().Format("15:04:05"))
	}

}

const testtempl = `
<html>

<head>
	<style type="text/css">
		body      { margin: 20px; font-family: Ubuntu, Arial, sans-serif; }
		hr        { border-style: none; border-top: 1px solid #CCCCCC; }
		.ErrorBox { color: red; font-weight: bold; } 
		.TextBox  { border:solid; border-color:#BBBBBB; border-width:1px; padding-left:4px;}
	</style>
	{{.JS}}
</head>

<body>

	<h1> GUI test </h1>
	<p> {{.ErrorBox}} </p>
	<hr/>

	{{.Span "e_time" "time flies"}} <br/>
	{{.Button "e_hitme" "Hit me!"}} <br/>
	

	<hr/>
	

</body>
</html>
`