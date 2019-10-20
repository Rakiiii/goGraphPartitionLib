package graphpartitionlib

import (
	"fmt"
	"log"
	"math/big"
	"runtime"
	"sync"
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

func TestAssync(t *testing.T) {
	var parser = new(graphlib.Parser)
	var g, err = parser.ParseUnweightedUndirectedGraphFromFile("testGraph1")
	if err != nil {
		fmt.Println(err)
		return
	}

	var result = Result{Matrix: nil, Value: -1}

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

	am := 2

	var wg sync.WaitGroup

	runtime.GOMAXPROCS(am)

	ch := make(chan Result, am)

	wg.Add(am)

	dif := big.NewInt(0)
	dif.Sub(end, start)

	dif.Div(dif, big.NewInt(int64(am)))

	subEnd := big.NewInt(0)
	subEnd.Add(subEnd, start)
	subEnd.Add(subEnd, dif)

	for i := 0; i < am; i++ {
		log.Println(start.String(), "|", subEnd.String())

		go AsyncFindBestPartion(g, start.String(), subEnd.String(), amountOfGroups, float64(1), &wg, ch)
		start.Add(start, dif)
		if i != am-2 {
			subEnd.Add(subEnd, dif)
		} else {
			subEnd = end
		}
	}

	wg.Wait()
	close(ch)

	for i := range ch {
		fmt.Println(i.Value)
		if result.Value < i.Value || result.Value == -1 {
			result = i
		}
	}

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
}
