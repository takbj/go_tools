# json
this is a go json tool like encode/json(golang 1.6) but allow you to use camelName/CamelName or under_score name style
and implement method map[key]interface{},the map's key more than "string",now added "bool","float64","float32","int","int8","int16","int32","int64","uint","uint8","uint16","uint32","uint64"
and implement `encoding.TextMarshaler`'s method MarshalText() (text []byte, err error) and `encoding.TextUnmarshaler`'s method UnmarshalText(text []byte) error

    package main

    import (
    	"fmt"

     	"github.com/zxfonline/json"
    )

    type Person struct {
    	Id        int
    	FirstName string
    	LastName  string `json:"name"`
    }

    func main() {
    	p := Person{
    		Id:        1,
    		FirstName: "FirstName",
    		LastName:  "LastName",
    	}
    	json.UnderScoreStyle()
    	b, _ := json.Marshal(&p)
    	fmt.Println(string(b))
    }

output

> {"id":1,"first_name":"FirstName","name":"LastName"}


    package main

     import (
     	"fmt"

     	"github.com/zxfonline/json"
     )

     type Person struct {
     	Id        int
     	FirstName string
     	LastName  string `json:"name"`
     }

     func main() {
     	p := Person{
     		Id:        1,
     		FirstName: "FirstName",
     		LastName:  "LastName",
     	}
     	b, _ := json.Marshal(&p)
     	fmt.Println(string(b))
     }

output

> {"id":1,"firstName":"FirstName","name":"LastName"}


	package main

	import (
	"fmt"

	"github.com/zxfonline/json"
	)


	func main() {
		m := make(map[*EE]bool)
		point := new(EE)
		//or
		//	point := &EE{}
		point.name = "aaa"
		m[point] = true
		b, err := json.Marshal(m)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
		a1 := make(map[*EE]bool)
		err = json.Unmarshal(b, &a1)
		if err != nil {
			panic(err)
		}
		b, err = json.Marshal(&a1)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
	}

	type EE struct {
		name string
	}

	func (n *EE) MarshalText() ([]byte, error) {
		return []byte(n.name), nil
	}

	func (n *EE) UnmarshalText(text []byte) error {
		n.name = string(text)
		return nil
	}

output

> {"aaa":true}
> {"aaa":true}

golang struct to json,save as mongodb document
and mongodb document parse to struct func test

	package main

	import (
		"fmt"
		"time"

		"github.com/zxfonline/json"
		"gopkg.in/mgo.v2"
		"gopkg.in/mgo.v2/bson"
	)

	type NodeA struct {
		Node1 int
		Node2 string
	}
	type NodeB struct {
		NodeAcode int
		NodeBcode string
		NodeASon  map[float64]*NodeA
		NodeBSon  []*NodeA
		NodeCcode time.Time
	}
	type NodeB1 struct {
		NodeAcode  int
		NodeBcode1 string
		NodeASon   map[float64]*NodeA
		NodeBSon   []*NodeA
	}

	func main() {
		node := &NodeB{NodeAcode: 21, NodeBcode: "21", NodeCcode: time.Now()}
		fmt.Println(node.NodeCcode.Format("2006-01-02 15:04:05"))
		sonst := make(map[float64]*NodeA)
		sonst[1.1] = &NodeA{Node1: 111, Node2: "111"}
		sonst[2.2] = &NodeA{Node1: 112, Node2: "112"}
		node.NodeASon = sonst
		sona := make([]*NodeA, 0)
		sona = append(sona, &NodeA{Node1: 211, Node2: "211"})
		sona = append(sona, &NodeA{Node1: 212, Node2: "212"})
		node.NodeBSon = sona

		b, err := json.Marshal(&node)
		if err != nil {
			panic(err)
		}
		fmt.Println("insert struct json=", string(b))

		var f interface{}
		err = json.Unmarshal([]byte(b), &f)
		if err != nil {
			panic(err)
		}
		var session *mgo.Session
		session, err = mgo.Dial("localhost:27017")
		if err != nil {
			panic(err)
		}
		defer session.Close()
		coll := session.DB("test").C("test")

		err = coll.Insert(f)
		if err != nil {
			panic(err)
		}

		var q []bson.M
		coll.Find(nil).All(&q)
		for _, info := range q {
			b, err = json.Marshal(info)
			if err != nil {
				panic(err)
			}
			fmt.Println("parse struct json=", string(b))
			nodeN := &NodeB{}
			err = json.Unmarshal([]byte(b), nodeN)
			if err != nil {
				panic(err)
			}
			//test content same?
			b, err = json.Marshal(&node)
			if err != nil {
				panic(err)
			}
			fmt.Println("new struct json=", string(b))
			fmt.Println(nodeN.NodeCcode.Format("2006-01-02 15:04:05"))
		}
	}
> 2016-04-16 16:10:30
> insert struct json= {"NodeAcode":21,"NodeBcode":"21","NodeASon":{"1.1":{"Node1":111,"Node2":"111"},"2.2":{"Node1":112,"Node2":"112"}},"NodeBSon":[{"Node1":211,"Node2":"211"},{"Node1":212,"Node2":"212"}],"NodeCcode":"2016-04-16T16:10:30.9692215+08:00"}

> new struct json= {"NodeAcode":21,"NodeBcode":"21","NodeASon":{"1.1":{"Node1":111,"Node2":"111"},"2.2":{"Node1":112,"Node2":"112"}},"NodeBSon":[{"Node1":211,"Node2":"211"},{"Node1":212,"Node2":"212"}],"NodeCcode":"2016-04-16T16:10:30.9692215+08:00"}
> 2016-04-16 16:10:30
