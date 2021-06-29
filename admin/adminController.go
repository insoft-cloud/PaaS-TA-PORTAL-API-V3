package admin

import (
	"PAAS-TA-PORTAL-V3/config"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func AdminHandleRequests(myRouter *mux.Router) {
	myRouter.HandleFunc("/clear_buildpack_cache", clearBuildpackCache).Methods("POST")
}

func clearBuildpackCache(w http.ResponseWriter, r *http.Request) {
	req, _ := http.NewRequest("POST", config.Config["cloudfoundry_domain"]+"/v3/admin/actions/clear_buildpack_cache", nil)
	req.Header.Set("Authorization", w.Header().Get("cf-Authorization"))
	_, err := config.Client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("succeed")
}
