package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	URL = "127.0.0.1:27017"
)

type Person struct {
	Id    int
	Name  string
	Phone string
}

type Men struct {
	Persons []Person
}

func main() {
	session, err := mgo.Dial(URL) //连接数据库
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	db := session.DB("test")     //数据库名称
	collection := db.C("person") //如果该集合已经存在的话，则直接返回

	//*****集合中元素数目********
	countNum, err := collection.Count()
	if err != nil {
		panic(err)
	}
	fmt.Println("Things objects count: ", countNum)

	//*******插入元素*******
	temp := &Person{
		Id:    2,
		Phone: "18811577546",
		Name:  "zhangzheHero",
	}
	//一次可以插入多个对象 插入两个Person对象
	err = collection.Insert(&Person{1, "Ale", "+55 53 8116 9639"}, temp)
	if err != nil {
		panic(err)
	}

	//*****查询单条数据*******
	result := Person{}
	err = collection.Find(bson.M{"phone": "456"}).One(&result)
	fmt.Println("Phone:", result.Name, result.Phone, result.Id)

	var persons []Person
	collection.Find(nil).All(&persons)
	fmt.Println("user:", persons)

	//*****查询多条数据*******
	var personAll Men //存放结果
	iter := collection.Find(nil).Iter()
	for iter.Next(&result) {
		fmt.Printf("Result: %v\n", result.Name)
		personAll.Persons = append(personAll.Persons, result)
	}

	//*******更新数据**********
	err = collection.Update(bson.M{"name": "ccc"}, bson.M{"$set": bson.M{"name": "ddd"}})
	err = collection.Update(bson.M{"name": "ddd"}, bson.M{"$set": bson.M{"phone": "12345678"}})
	err = collection.Update(bson.M{"name": "aaa"}, bson.M{"phone": "1245", "name": "bbb"})

	//******删除数据************
	_, err = collection.RemoveAll(bson.M{"name": "Ale"})

}

// func handlerErr(er) {
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
