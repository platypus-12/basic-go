package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type Element struct {
	name  string
	id    string
	class string
}

// ../../ch01/fetch http://www.w3.org/TR/2006/REC-xml11-20060816 | go run xmlselect.go div class=back div h2
func main() {
	var arg []Element
	for _, v := range os.Args[1:] {
		if strings.Contains(v, "=") {
			str := strings.Split(v, "=")
			switch str[0] {
			case "id":
				arg[len(arg)-1].id = str[1]
			case "class":
				arg[len(arg)-1].class = str[1]
			}
		} else {
			arg = append(arg, Element{name: v})
		}
	}

	dec := xml.NewDecoder(os.Stdin)
	var stack []Element
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			element := Element{name: tok.Name.Local}
			for _, attr := range tok.Attr {
				switch attr.Name.Local {
				case "id":
					element.id = attr.Value
				case "class":
					element.class = attr.Value
				}
			}
			stack = append(stack, element)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if containsAll(stack, arg) {
				for _, v := range stack {
					fmt.Printf("%s ", v.name)
					if v.id != "" {
						fmt.Printf("id=%s ", v.id)
					}
					if v.class != "" {
						fmt.Printf("class=%s  ", v.class)
					}
				}
				fmt.Printf(":%s\n", tok)
			}
		}
	}
}

func containsAll(x, y []Element) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0].name == y[0].name {
			if y[0].id != "" {
				if x[0].id != y[0].id {
					return false
				}
			}
			if y[0].class != "" {
				if x[0].class != y[0].class {
					return false
				}
			}
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
