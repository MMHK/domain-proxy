// http
package lib

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Service struct {
	Listen string
	config *Config
}

type Service_Status struct {
	Status bool
}

type ServiceError struct {
	Error string
}

func NewService(listen string, config *Config) (s *Service) {
	s = &Service{
		Listen: listen,
		config: config,
	}

	return
}

func (s *Service) Start() {
	r := mux.NewRouter()

	r.HandleFunc("/", s.RedirectSwagger)
	r.HandleFunc("/reload", s.Reload)
	r.HandleFunc("/add", s.AddEntry)
	r.HandleFunc("/remove", s.RemoveEntry)
	r.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/",
		http.FileServer(http.Dir(fmt.Sprintf("%s/swagger", s.config.WebRoot)))))
	r.NotFoundHandler = http.HandlerFunc(s.NotFoundHandle)

	http.ListenAndServe(s.Listen, r)
}

func (this *Service) RedirectSwagger(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "/swagger/index.html", 301)
}

func (s *Service) ResponseJSON(source interface{}, writer http.ResponseWriter) {
	json_str, err := json.Marshal(source)
	if err != nil {
		s.ResponseError(err, writer)
	}
	writer.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(writer, string(json_str))
}

func (s *Service) NotFoundHandle(writer http.ResponseWriter, request *http.Request) {
	server_error := &ServiceError{Error: "handle not found!"}
	json_str, _ := json.Marshal(server_error)
	log.Info(string(json_str))
	http.Error(writer, string(json_str), 404)
}

func (s *Service) ResponseError(err error, writer http.ResponseWriter) {
	server_error := ServiceError{Error: err.Error()}
	json_str, _ := json.Marshal(server_error)
	http.Error(writer, string(json_str), 500)
}

//
// swagger:operation GET /reload reloadService
//
// 触发web service 服务重载
//
// ---
// produces:
//   - application/json
// responses:
//   200:
//     description: OK
//   500:
//     description: Error
//
//
func (s *Service) Reload(writer http.ResponseWriter, request *http.Request) {
	err := ReloadService(s.config.ReloadCMD)
	if err != nil {
		s.ResponseError(err, writer)
		return
	}
	s.ResponseJSON(Service_Status{Status: true}, writer)
	return
}

//
// swagger:operation POST /add addDomain
//
// 添加域名配置
//
// ---
// consumes:
//   - application/x-www-form-urlencoded
// produces:
//   - application/json
// parameters:
// - name: domain
//   type: string
//   in: formData
//   required: true
//   description: 需要增加的三级域名
// - name: ip
//   type: string
//   in: formData
//   required: true
//   description: 绑定的内网IP地址（IPv4）
// responses:
//   200:
//     description: OK
//   500:
//     description: Error
//
//
func (s *Service) AddEntry(writer http.ResponseWriter, request *http.Request) {
	domain := request.FormValue("domain")
	ip := request.FormValue("ip")

	err := createNewDomainConfig(domain, ip, s.config.DomainCfgSaveDir, s.config.EntryTemplate, s.config.DomainCfgFileNameFormat)
	if err != nil {
		s.ResponseError(err, writer)
		return
	}
	s.ResponseJSON(Service_Status{Status: true}, writer)
	return
}

//
// swagger:operation POST /remove removeDomain
//
// 移除域名配置
//
// ---
// consumes:
//   - application/x-www-form-urlencoded
// produces:
//   - application/json
// parameters:
// - name: domain
//   type: string
//   in: formData
//   required: true
//   description: 需要移除的三级域名
// - name: ip
//   type: string
//   in: formData
//   required: true
//   description: 绑定的内网IP地址（IPv4）
// responses:
//   200:
//     description: OK
//   500:
//     description: Error
//
//
func (s *Service) RemoveEntry(writer http.ResponseWriter, request *http.Request) {
	domain := request.FormValue("domain")
	ip := request.FormValue("ip")

	err := RemoveDomainConfig(domain, ip, s.config.DomainCfgSaveDir, s.config.DomainCfgFileNameFormat)
	if err != nil {
		s.ResponseError(err, writer)
		return
	}
	s.ResponseJSON(Service_Status{Status: true}, writer)
	return
}
