package graphpartitionlib

import (
	"fmt"
	"math/big"
	"testing"

	bigintlib "github.com/Rakiiii/goBigIntLib"
	boolmatrixlib "github.com/Rakiiii/goBoolMatrix"
	graphlib "github.com/Rakiiii/goGraph"
)

func TestFindBestPartition(t *testing.T) {
	var parser = new(graphlib.Parser)
	var g, err = parser.ParseUnweightedUndirectedGraphFromFile("testGraph1")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Graph parsed")

	amountOfGroups := 2
	if err != nil {
		fmt.Println(err)
		return
	}
	end := bigintlib.Pow(big.NewInt(2), int64(g.AmountOfVertex())*int64(amountOfGroups))
	fmt.Println(end.String())

	start := big.NewInt(0)

	//for i := 0; i < g.AmountOfVertex(); i++ {
	//	start.Add(start, bigintlib.Pow(big.NewInt(2), int64((amountOfGroups*g.AmountOfVertex()-1)-(amountOfGroups*(i+1)-i+1))))
	//}

	fmt.Println("Big int initialized")

	result, err := FindBestPartion(g, start, end, amountOfGroups, float64(0.7))

	if err != nil {
		t.Error(err)
	}

	if result.Matrix == nil {
		t.Error("result matrix nil")
	} else {
		result.Matrix.Print()
	}
	if result.Value == -1 {
		t.Error("result param is -1")
		if result.Matrix != nil {
			result.Matrix.Print()
		}
	} else {
		fmt.Println(result.Value)
	}

	//result.Matrix.Print()
	//fmt.Println(result.Value)

}

func TestCountParameter(t *testing.T) {
	var parser = new(graphlib.Parser)
	var g, err = parser.ParseUnweightedUndirectedGraphFromFile("testGraph1")
	if err != nil {
		fmt.Println("ERORRE", err)
		return
	}

	var b boolmatrixlib.BoolMatrix
	b.Init(2, 5)
	b.SetByNumber(big.NewInt(int64(681)))

	/*fmt.Println("Vertex:", g.AmountOfVertex())
	for i := 0; i < g.AmountOfVertex(); i++ {
		edges := g.GetEdges(i)
		fmt.Println("Edges len:", len(edges))
		for _, j := range edges {
			fmt.Println(j)
		}
	}*/

	result, err := CountParameter(g, &b)

	if result == 0 {
		t.Error("Wrong parameter:", result)
		b.Print()
	} else {
		fmt.Println("Param:", result)
	}

}
