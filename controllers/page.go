package controllers

import (
	"net/http"
	"html/template"
	"build_web/GoPractice/dlog"
)

var templates map[string]*template.Template

/*
加载这个包时就收集模板
 */
func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	templates["login"] = template.Must(template.ParseFiles("templates/login.html", "templates/login_base.html"))
	templates["overview"] = template.Must(template.ParseFiles("templates/overview.html", "templates/base.html"))
	templates["users"] = template.Must(template.ParseFiles("templates/users.html", "templates/base.html"))
	templates["channel"] = template.Must(template.ParseFiles("templates/channel.html", "templates/base.html"))
	templates["rule"] = template.Must(template.ParseFiles("templates/rule.html", "templates/base.html"))
}


/*
渲染模板
收集模板页面提供名称和对应的模板渲染
 */

func renderTemplate(w http.ResponseWriter, name string, template string, data interface{}) {
	// Ensure the template exists in the map.
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, "The template does not exist.", http.StatusInternalServerError)
	}
	err := tmpl.ExecuteTemplate(w, template, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/*
渲染总览页面
 */
func Overview(w http.ResponseWriter, r *http.Request) {
	var log = dlog.DcLog()
	log.Println(r.Method, r.URL.Path)
	if r.Method == "GET" {
		renderTemplate(w, "overview", "base", nil)
	}
}

/*
用户页面
 */

func UserView(w http.ResponseWriter, r *http.Request) {
	var log = dlog.DcLog()
	log.Println(r.Method, r.URL.Path)
	if r.Method == "GET" {
		renderTemplate(w, "users", "base", nil)
	}
}

/*
渠道页面
 */

func ChannelView(w http.ResponseWriter, r *http.Request) {
	var log = dlog.DcLog()
	log.Println(r.Method, r.URL.Path)
	if r.Method == "GET" {
		renderTemplate(w, "channel", "base", nil)
	}
}

/*
套餐页面
 */

func RuleView(w http.ResponseWriter, r *http.Request) {
	var log = dlog.DcLog()
	log.Println(r.Method, r.URL.Path)
	if r.Method == "GET" {
		renderTemplate(w, "rule", "base", nil)
	}
}