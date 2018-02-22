package jpcontact

import ("github.com/mediocregopher/radix.v2/redis"
	"fmt"
       )
type ContactDetails struct {
	Name      string
	Age       string
	FavDrink1 string
	FavDrink2 string
}

func AddContact(newContact ContactDetails) {
	conn, err := redis.Dial("tcp", "10.1.1.21:32769")
	if err != nil {
		// log.Fatal(err)
	}
	defer conn.Close()
	resp := conn.Cmd("HMSET", "friend:"newContact.Name, "name", newContact.Name, "Age", newContact.Age, "favDrink1", newContact.FavDrink1, "favDrink2", newContact.FavDrink2)
	if resp.Err != nil {
		// log.Fatal(resp.Err)
	}
	resp2 := conn.Cmd("HMGET", "friend:"newContact.Name)
	fmt.Println(resp2)
}
