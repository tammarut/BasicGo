package main

//! Go Connect Mongodb | Tutorial
//* Pokemondb: only Go and mongo
import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Pokemon struct {
	ID      string `bson:"_id"`
	Name    string
	Element string
	Weight  string
}

func getPokemons() []Pokemon {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017/pokemondb")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("pokemondb").C("pokemons")
	pokemons := []Pokemon{}
	err = c.Find(nil).All(&pokemons)
	if err != nil {
		log.Fatal(err)
	}
	return pokemons
}

func getPokemon(id string) Pokemon {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017/pokemondb")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("pokemondb").C("pokemons")
	pokemon := Pokemon{}
	err = c.Find(bson.M{"_id": id}).One(&pokemon)
	if err != nil {
		log.Fatal(err)
	}
	return pokemon
}

func main() {
	pokemons := getPokemons()
	fmt.Println(pokemons)

	fmt.Println("-------------------------------------------")

	pokemon := getPokemon("002")
	fmt.Println(pokemon)
}
