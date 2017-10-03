package main

import (
	"fmt"
	"net/http"
	"runtime"
)

const message = `
	         ,_---~~~~~----._         
	  _,,_,*^____      _____ *g*\"*, 
	 / __/ /'     ^.  /      \ ^@q   f 
	[  @f | @))    |  | @))   l  0 _/  
	 \ /   \~____ / __ \_____/    \   
	  |           _l__l_           I   
	  }          [______]           I  
	  ]            | | |            |  
	  ]             ~ ~             |  
	  |                            |   
	   |                           |   
---------------------------------------------

			HELLO FROM GOLANG !
GOOS=%v
GOARCH=%v
`

func main() {
	http.HandleFunc("/", func(response http.ResponseWriter, req *http.Request) {
		response.Header().Add("Content-Type", "text")
		response.Write([]byte(fmt.Sprintf(message, runtime.GOOS, runtime.GOARCH)))
	})
	http.ListenAndServe(":80", nil)
}
