/*
*
When enum is going to mutually exclusive then the typical sequence numbering is good.
However, when the nature of enum is multi-select then numbering in 2's power helps.
For ex:
First enum  =  1    =     1
Second enum =  2    =    10
Third enum  =  4    =   100
Fourth enum =  8    =  1000
Fifth enum  =  16   = 10000

Based on which bit is set, by making use of bitwise operators you can decide the type
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")

	fmt.Println("(Raw | Field) is Raw:", (Raw | Field).is(Raw))
	fmt.Println("(Raw | Field) is Field:", (Raw | Field).is(Field))
	fmt.Println("(Raw | Field) is Graph:", (Raw | Field).is(Graph))
	fmt.Println("(Raw | Graph) is Graph:", (Raw | Graph).is(Graph))

	fmt.Println("(All) is Raw:", All.is(Raw))
	fmt.Println("(All) is Aggregate:", All.is(Aggregate))
	fmt.Println("(All) is Graph:", All.is(Graph))
	fmt.Println("(All) is Field:", All.is(Field))
	fmt.Println("end")
}

type QueryType int

const (
	Raw QueryType = 1 << iota
	Aggregate
	Field
	Graph
	All = Raw | Aggregate | Field | Graph
)

func (q QueryType) String() string {
	switch q {
	case Raw:
		return "raw"
	case Aggregate:
		return "aggregate"
	case Field:
		return "field"
	case Graph:
		return "graph"
	default:
		return "unknown"
	}
}

func (q QueryType) is(t QueryType) bool {
	return q&t == t
}
