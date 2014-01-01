package main

import(
  "github.com/gorilla/mux"
  "net/http"
  "fmt"
  "regexp"
)

func routerHandlerFunc(router *mux.Router) http.HandlerFunc {
  return func(res http.ResponseWriter, req *http.Request) {
    router.ServeHTTP(res, req)
  }
}

func getStatus(projectUUID string, branch string) string {
  codeshipURL := fmt.Sprintf("https://www.codeship.io/projects/%s/status?branch=%s", projectUUID, branch)
  resp, _ := http.Get(codeshipURL)
  contentDisposition := resp.Header.Get("Content-Disposition")
  re, _ := regexp.Compile("inline; filename=\"status_(?P<status>.+).png\"")
  status := re.FindStringSubmatch(contentDisposition)[1]
  fmt.Println(status)
  return status
}

func getCodeshipStatus(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  vars := mux.Vars(r)
  projectUUID := vars["projectUUID"]
  branch := vars["branch"]
  status := getStatus(projectUUID, branch)
  fmt.Fprint(w, fmt.Sprintf("{\"status\": \"%s\"}", status))
}

func router() *mux.Router {
  r := mux.NewRouter()
  r.HandleFunc("/{projectUUID}/{branch:.+}", getCodeshipStatus).Methods("GET")
  return r
}
