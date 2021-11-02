package handdler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"restfull.com/rest/databasesetup"
)

type Students struct {
	Id       int
	Name     string
	Usn      string
	College  string
	Location string
}

func Welcomeresponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

/*func SelectDBPost(pg *pg.DB, post Details) (Details, error) {
	err := pg.Model(&post).Where("id = 1").Select()
	return post, err
}*/
func Getalldetails(w http.ResponseWriter, r *http.Request) { //it displays all the student details from the database
	db := databasesetup.Setup() //seting up connection with database
	//defer db.Close()
	//	fmt.Println("Print student")
	var mydet []Students
	err := db.Model(&mydet).Select() //Selects all details from table student and stores it to Student struct
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mydet)
	b, err := json.Marshal(mydet) //decoding struct to json
	str := string(b)              //converting  to string
	fmt.Fprintf(w, str)
	db.Close() //closing db conection
}

//Getadetails function we provide details of a particular details
func Getadetails(w http.ResponseWriter, r *http.Request) {
	db := databasesetup.Setup()
	vars := mux.Vars(r)
	key := vars["id"]
	intkey, _ := strconv.Atoi(key)
	stud := new(Students)
	err := db.Model(stud).Where("id=?", intkey).Select()
	if err != nil {
		panic(err)
	}
	b, err := json.Marshal(stud)
	str := string(b)
	fmt.Fprintf(w, str)
	db.Close()
}

//Postdetails func inserts details to the db
func Postdetails(w http.ResponseWriter, r *http.Request) {
	db := databasesetup.Setup()
	newpost, _ := ioutil.ReadAll(r.Body)
	var newdet Students
	json.Unmarshal(newpost, &newdet)
	_, err := db.Model(&newdet).Insert()
	if err != nil {
		panic(err)
	}
	Getalldetails(w, r)
	db.Close()
}

//Deletedetails func deletes the particular details from db
func Deletedetails(w http.ResponseWriter, r *http.Request) {
	db := databasesetup.Setup()
	vars := mux.Vars(r)
	val := vars["id"]
	intval, _ := strconv.Atoi(val)
	var det Students
	_, err := db.Model(&det).Where("id=?", intval).Delete()
	if err != nil {
		panic(err)
	}
	Getalldetails(w, r)
	db.Close()
}

//Updatedetails func will update a particular rows in db
func Updatedetails(w http.ResponseWriter, r *http.Request) {
	db := databasesetup.Setup()
	vars := mux.Vars(r)
	val := vars["id"]
	intval, _ := strconv.Atoi(val)
	updateval, _ := ioutil.ReadAll(r.Body)
	var updatedetail Students
	json.Unmarshal(updateval, &updatedetail)
	_, err := db.Model(&updatedetail).Where("id=?", intval).Update()
	if err != nil {
		panic(err)
	}
	Getalldetails(w, r)
	db.Close()
}
