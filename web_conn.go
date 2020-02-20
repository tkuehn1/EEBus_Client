package Studienarbeit_src

import "net/http"

func web_conn() {

	http.HandleFunc("/index", IoT_Index)
	http.ListenAndServe(":3000", nil)
}
