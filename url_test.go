package url

import (
	"fmt"
	"testing"
)


func TestParser(t *testing.T){
	const test1 = "https://test.com/go"
	ur,err := Parse(test1)
	if err != nil { 
		t.Fatalf("the url shoud contain http")
	}
	if got,expected := ur.Sheme,"https" ; got != expected{
		t.Errorf("Parse(%q) , expected %q , we got %q",test1,expected,got)
	}
	if got,expected := ur.Host,"test.com" ; got != expected{
		t.Errorf("Parse(%q) , expected %q , we got %q",test1,expected,got)
	}
	if got,expected := ur.Path,"go" ; got != expected {
		t.Errorf("Parse(%q) , expected %q , we got %q",test1,expected,got)
	}

}
// func TestPort(t *testing.T){
// 	cases := map[string]struct{
// 		host string
// 		port string
// 	}{
// 		"With a port number":{"test.com:80","0",},
// 		"With an empty port number":{"test.com:","",},
// 		"Without a port number":{"test.com","",},
// 		"IP with a port number":{"01.2.3.3:90","9",},
// 		"IP without a port number":{"01.2.3.3","",},
// 	}
// 	for k,v := range cases{
// 		u := &URL{Host: v.host}
// 		if got,expected := u.port(),v.port ; expected != got{
// 			t.Errorf("the test %s faild , expected %q got %q",k,expected,got)
// 		}
// 	}
//
// }
var tests = map[string]struct{
		in string
		host string
		port string
	}{
		"With a port number" : {"test.com:8080","test.com","8080"},
		"With an empty port" : {"test.com:","test.com",""},
		"Without a port number" : {"test.com","test.com",""},
		"IP with a port number" : {"1.2.3:8080","1.2.3","8080"},
		"IP without a port number" : {"1.2.3","1.2.3",""},
	}

func TestHost(t *testing.T){
	for name,test := range tests{
		t.Run(fmt.Sprintf("%s %s",name,test.host),func(t *testing.T) {
			u := &URL{Host: test.host}
			if got,expected := u.hostName(),test.host ; got != expected{
				t.Errorf("got %q expected %q",u.hostName(),test.host)
			}
		})
	}
}
func TestPort1(t *testing.T){
	for name,test := range tests{
		t.Run(fmt.Sprintf("%s %s",name,test.in),func(t *testing.T) {
		u := &URL{Host: test.in}
		if got,expected := u.port(),test.port ; got != expected{
			t.Errorf("got %q expected %q",u.port(),test.port)
		}
		})
	}
}
var stringtests = map[string]struct{
		url string
		str string
		expectederr bool
	}{
		"full url" : {"https://test.com/path","Sheme:https Host:test.com Path:path",false},
		"empty Sheme" : {"://test.com/path","Sheme: Host:test.com Path:path",false},
		"full url with http" : {"http://test.com/path","Sheme:http Host:test.com Path:path",false},
		"no Sheme" : {"test.com/path","",true},
		"no Path" : {"https://test.com","Sheme:https Host:test.com Path:",false},
	}
func TestString(t *testing.T){
	for name,test := range stringtests{
		t.Run(fmt.Sprintf("%s %s",name,test.url),func(t *testing.T) {
		u,err := Parse(test.url)
		if test.expectederr{
			if u != nil{
				t.Fatal("expected url = nil")
			}
			if err == nil{
				t.Fatal("expected err != nil")
			}
			return
		}
		if test.expectederr == false{
			if err != nil{
				t.Fatal("expecte err = nil")
			}
		}
		if got,expected := u.String(),test.str ; got != expected{
			t.Errorf("\ngot %q\n expected %q",u.String(),test.str)
		}
		})
	}
}
