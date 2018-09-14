package store

import (
	"gopkg.in/mgo.v2"
	"fmt"
)

func getMondoDbSession() *mgo.Session{
	session, err := mgo.Dial("localhost:27017")
	if err != nil{
		fmt.Println("failed to establish mongo connection")
	}
	return session
}

func GetAllProducts() []Product{
	//session, err := mgo.Dial("localhost:27017")
	//if err != nil{
	//	fmt.Println("failed to establish mongo connection")
	//}
	session := getMondoDbSession()
	defer session.Close()
	p := []Product{}
	c := session.DB("dummyStore").C("store")
	if err := c.Find(nil).All(&p); err != nil{
		fmt.Println("not getting data")
	}
	return p
}


func FilterProductById(id interface{}) Product{
	session := getMondoDbSession()
	defer session.Close()
	c := session.DB("dummyStore").C("store")
	p := Product{}
	c.FindId(id).One(&p)
	return p
}

