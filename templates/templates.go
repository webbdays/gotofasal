package templates

// import (
// 	"html/template"
// 	"os"
// )

// func IndexPage() *template.Template {
// 	wd , err := os.Getwd()
// 	if err!=nil{
// 		panic("cannot get current working directory")
// 	}
// 	indexPage, err := template.ParseFiles(wd + "/index.html")
// 	if err != nil {
// 		panic("Cannot Parse index.html")
// 	}
// 	return indexPage
// }

// func SignInPage() *template.Template {
// 	signInPage, err := template.ParseFiles("signin.html")
// 	if err != nil {
// 		panic("Cannot Parse signin.html")
// 	}
// 	return signInPage
// }
