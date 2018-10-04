package main

import (
  "net/http"
  "log"
  "html/template"
)


type Button struct {
  Name       string
  Value      string
  IsDisabled bool
  IsChecked  bool
  Text       string
}

type PageVariables struct {
  PageTitle        string
  PageButtons []Button
  Answer           string
}


func main() {
  http.HandleFunc("/", DisplayButtons)
  http.HandleFunc("/selected", UserSelected)
  log.Fatal(http.ListenAndServe(":8080", nil))
}


func DisplayButtons(w http.ResponseWriter, r *http.Request){
 // Display some buttons to the user

   Title := "Are you a Macalester Student?"
   MyButtons := []Button{
     Button{"studentselect", "Yes", false, false, "Yes"},
     Button{"studentselect", "No", false, false, "No"},
   }

  MyPageVariables := PageVariables{
    PageTitle: Title,
    PageButtons : MyButtons,
    }

   t, err := template.ParseFiles("homepage.html") //parse the html file homepage.html
   if err != nil { // if there is an error
     log.Print("template parsing error: ", err) // log it
   }

   err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
   if err != nil { // if there is an error
     log.Print("template executing error: ", err) //log it
   }

}

func UserSelected(w http.ResponseWriter, r *http.Request){
  r.ParseForm()
  youridentity := r.Form.Get("studentselect")

  Title := "Your identity is"
  MyPageVariables := PageVariables{
    PageTitle: Title,
    Answer : youridentity,
    }

 // generate page by passing page variables into template
    t, err := template.ParseFiles("homepage.html") //parse the html file homepage.html
    if err != nil { // if there is an error
      log.Print("template parsing error: ", err) // log it
    }

    err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
    if err != nil { // if there is an error
      log.Print("template executing error: ", err) //log it
    }
}