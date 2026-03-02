package url

import (
	"errors"
	"fmt"
	"strings"
)
type URL struct{
	Sheme string
	Host string
	Path string
}
func (ur *URL) hostName() string{
	host := ur.Host
	j := strings.Index(host,":")
	if j == -1{
		return host
	}
	return  host[:j]
}
func (ur *URL) port() string{
	host := ur.Host
	j := strings.Index(host,":")
	if j == -1{
		return ""
	}
	return host[j+1:]
}

func Parse(urlpath string) (*URL,error){
	i := strings.Index(urlpath,"://")
	if i == -1{
		return nil,errors.New("the http request should include https://")
	}
	j := strings.Index(urlpath[i+3:],"/")
	if j == -1 {
	sheme,host,path := urlpath[:i],urlpath[i+3:],""
	return &URL{sheme,host,path},nil
	}else{
	j = j + 3 + i
	sheme,host,path := urlpath[:i],urlpath[i+3:j],urlpath[j+1:]
	return &URL{sheme,host,path},nil}
}
func (u *URL) String()string{
	return fmt.Sprintf("Sheme:%s Host:%s Path:%s",u.Sheme,u.Host,u.Path)
}
