package main

import (
	"fmt"
	"math/rand"
	"strconv"
	
	
)
type Match struct{
	Round string
	Player1 string
	Player2 string
	
	Winner string
}

func(M *Match)shuffleMatches( onGoingMatches []Match)([]Match){
	var matches []Match
    var players []string
	for _,player:=range onGoingMatches{
		players=append(players, player.Player1)
		players=append(players, player.Player2)
	}
	fmt.Println("list of players :",players,"total players ares",len(players))
	matches,_=M.logicForShuffling(players)
	return matches
	


}
func (M *Match)logicForShuffling(players []string)([]Match,error){
    rounds:=len(players)/2
	
	var matches []Match
	if(len(players)==1){
		fmt.Println("winner is ",players)
		return nil,nil
	}
	if(len(players)%2!=0){
		i:=rand.Intn(len(players))
		match:=&Match{
        Round: strconv.Itoa(rounds),
		Player1: players[i],
		Player2: "",
		Winner: players[i],
		}
		matches=append(matches, *match)
		players=append(players[:i], players[i+1:]...)
	}
	
	rand.Shuffle(len(players),func(i, j int) {
		players[i],players[j]=players[j],players[i]
	})

for i:=0;i<len(players);i=i+2{
	matches = append(matches, Match{
		Round: strconv.Itoa(rounds),
        Player1: players[i],
		Player2: players[i+1],
	})
}
return matches,nil
}
func main(){
	onGoingMatches := []Match{
		{Round:"", Player1: "Alice", Player2: "Bob", Winner: ""},
		{Round:"", Player1: "Charlie", Player2: "Dave", Winner: ""},
		{Round:"", Player1: "Eve", Player2: "Frank", Winner: ""},
		{Round:"", Player1: "Grace", Player2: "Heidi", Winner: ""},
		{Round:"", Player1: "Ivan", Player2: "Judy", Winner: ""},
		{Round:"", Player1: "Mallory", Player2: "Niaj", Winner: ""},
		{Round:"", Player1: "Oscar", Player2: "Peggy", Winner: ""},
		{Round:"", Player1: "Sybil", Player2: "Trent", Winner: ""},
		{Round:"", Player1: "Victor", Player2: "Walter", Winner: ""},
		{Round:"", Player1: "Xander", Player2: "Yvonne", Winner: ""},
		{Round:"", Player1: "Zara", Player2: "Amelia", Winner: ""},
		{Round:"", Player1: "Brian", Player2: "Clara", Winner: ""},
		{Round:"", Player1: "David", Player2: "Emma", Winner: ""},
		{Round:"", Player1: "Fred", Player2: "George", Winner: ""},
		{Round:"", Player1: "Henry", Player2: "Isabella", Winner: ""},
	}
	match:=&Match{}
	matches:=match.shuffleMatches(onGoingMatches)
    for i,liveMatch:=range matches{
       
		fmt.Println("round is ",liveMatch.Round,liveMatch.Player1,liveMatch.Player2,"match number:",i)
		fmt.Println("")
	   
		}
		
}