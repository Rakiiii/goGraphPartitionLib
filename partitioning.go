package graphpartitionlib

import (
	"errors"
	"fmt"
	"math/big"
	"sync"

	bigintlib "github.com/Rakiiii/goBigIntLib"
	boolmatrixlib "github.com/Rakiiii/goBoolMatrix"
	graphlib "github.com/Rakiiii/goGraph"
)

func FindBestPartion(graph *graphlib.Graph, start, end *big.Int, amountOfGroups int, disbalance float64) (Result, error) {
	var bestParameterValue int64 = int64(-1)
	bestMatrix := new(boolmatrixlib.BoolMatrix)

	amountOfVertex := graph.AmountOfVertex()
	flag := true

	subMatrix := new(boolmatrixlib.BoolMatrix)
	subMatrix.Init(amountOfGroups, amountOfVertex)

	for start.Cmp(end) < 0 {
		flag = true
		subMatrix.SetByNumber(start)

		if subMatrix.CountTrues() == int64(amountOfVertex) && subMatrix.CheckDisbalance(disbalance) {
			DebugLog("checking" + start.String())
			for i := 0; i < amountOfVertex; i++ {
				if subMatrix.CountTruesInLine(i) != 1 {
					flag = false
					break
				}
			}

			if flag {
				subParameterValue, err := CountParameter(graph, subMatrix)
				if err != nil {
					return Result{nil, bestParameterValue}, err
				}
				if (int64(graph.AmountOfEdges())-subParameterValue/2) < bestParameterValue || bestParameterValue == -1 {
					//DebugLog("BestResult Changed")
					fmt.Println("BestResult CHanged")
					bestMatrix = subMatrix.Copy()
					bestParameterValue = (int64(graph.AmountOfEdges()) - subParameterValue/2)
				}
			}
		}
		bigintlib.Inc(start)
	}

	return Result{bestMatrix, bestParameterValue}, nil
}

func CountParameter(graph *graphlib.Graph, matrix *boolmatrixlib.BoolMatrix) (int64, error) {
	result := int64(0)
	amV := graph.AmountOfVertex()

	if amV != matrix.Heigh() {
		return result, errors.New("BoolMatrix heigh is not equls to amount of graphs vertexes")
	}
	w := matrix.Width()

	for j := 0; j < w; j++ {
		for i := 0; i < amV; i++ {
			if matrix.GetBool(i, j) {
				edges := graph.GetEdges(i)
				for _, edge := range edges {
					if matrix.GetBool(edge, j) {
						result++
					}
				}
			}
		}
	}
	return result, nil
}

func DebugLog(str string) {
	//fmt.Println(str)
}

type Result struct {
	Matrix *boolmatrixlib.BoolMatrix
	Value  int64
}

func AsyncFindBestPartion(graph *graphlib.Graph, start, end *big.Int, amountOfGroups int, disbalance float64, wg *sync.WaitGroup, ch chan Result) {
	fmt.Println("Starrting new gorutin")
	defer wg.Done()

	/*var bestParameterValue int64 = math.MaxInt64
	bestMatrix := new(boolmatrixlib.BoolMatrix)

	amountOfVertex := graph.AmountOfVertex()
	flag := true

	subMatrix := new(boolmatrixlib.BoolMatrix)
	subMatrix.Init(amountOfGroups, amountOfVertex)
	for start.Cmp(end) < 0 {
		subMatrix.SetByNumber(start)
		//DebugLog("checking" + start.String())

		if subMatrix.CountTrues() == int64(amountOfVertex) && subMatrix.CheckDisbalance(disbalance) {
			for i := 0; i < amountOfVertex; i++ {
				if subMatrix.CountTruesInLine(i) != 1 {
					flag = false
					break
				}
			}

			if flag {
				//todo:: add func to count parameter and compare matrix
				subParameterValue, err := CountParameter(graph, subMatrix)
				if err != nil {
					panic(err)
				}
				if (int64(graph.AmountOfEdges())-subParameterValue/2) < (int64(graph.AmountOfEdges())-bestParameterValue/2) || bestParameterValue == 0 {
					DebugLog("BestResult Changed")
					bestMatrix = subMatrix.Copy()
					bestParameterValue = subParameterValue
				}
			}
			//bigintlib.Inc(start)
		}
		bigintlib.Inc(start)
	}*/

	res, err := FindBestPartion(graph, start, end, amountOfGroups, disbalance)
	if err != nil {
		fmt.Println(err)
		return
	}
	ch <- res
}
