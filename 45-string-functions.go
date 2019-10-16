package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("Contains:", s.Contains("test", "es"))
	p("Count:", s.Count("test", "t"))
	p("HasPrefix:", s.HasPrefix("test", "te"))
	p("HasSuffix:", s.HasSuffix("test", "st"))
	p("Index:", s.Index("test", "e"))
	p("Join:", s.Join([]string{"foo", "bar", "baz"}, ", "))
	p("Repeat:", s.Repeat("=", 10))
	p("Replace:", s.Replace("john", "joh", "evely", -1))
	p("Split:", s.Split("foo,bar,baz", ","))
	p("ToLower:", s.ToLower("FOOBAR"))
	p("ToUpper:", s.ToUpper("FOOBAR"))
	p()
	p("Len:", len("hello"))
	p("Char:", "hello"[1])
}
