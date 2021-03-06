package graphpartitionlib

import (
	"errors"
	"fmt"
	"math/big"
	"sync"
	"log"

	bigintlib "github.com/Rakiiii/goBigIntLib"
	boolmatrixlib "github.com/Rakiiii/goBoolMatrix"
	graphlib "github.com/Rakiiii/goGraph"
)

//FindBestPartition function finds @graph optimal partition by bruteforce algorithm in range @start to @end in @amountOfGroups amount of groupse with @disbalance disbalance 
func FindBestPartion(graph *graphlib.Graph, start, end *big.Int, amountOfGroups int, disbalance float64) (Result, error) {
	var bestParameterValue int64 = int64(-1)
	bestMatrix := new(boolmatrixlib.BoolMatrixLinear)

	amountOfVertex := graph.AmountOfVertex()
	flag := true

	subMatrix := new(boolmatrixlib.BoolMatrixLinear)
	subMatrix.Init(amountOfGroups, amountOfVertex)

	for start.Cmp(end) < 0 {
		flag = true
		subMatrix.SetByNumber(start)

		/*if start.Cmp(big.NewInt(698709)) == 0{
			//log.Println("trues:",subMatrix.CountTrues()," disb is:",subMatrix.CheckDisbalance(disbalance))
			subMatrix.Print()
		}*/
		if subMatrix.CountTrues() == int64(amountOfVertex) && subMatrix.CheckDisbalance(disbalance) {
			log.Println("checking ",start.String())
			for i := 0; i < amountOfVertex; i++ {
				if subMatrix.CountTruesInLine(i) != 1 {
					flag = false
					//log.Println("Wrong trueth in line:",subMatrix.CountTruesInLine(i)," line number:",i+1)
					//subMatrix.Print()
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
					log.Println("BestResult CHanged")
					bestMatrix = subMatrix.Copy()
					bestParameterValue = (int64(graph.AmountOfEdges()) - subParameterValue/2)

					//debug
					//log.Println("bigInt:", start.String())
				}
			}
		}
		bigintlib.Inc(start)
	}

	return Result{bestMatrix, bestParameterValue}, nil
}

//CountParameter returns parameter of @graph parted as @matrix 
func CountParameter(graph *graphlib.Graph, matrix boolmatrixlib.IBoolMatrix) (int64, error) {
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
					if matrix.GetBool(edge, j) && i != edge {
						//log.Println("edge is:",i,"|",edge)
						result++
					}
				}
			}
		}
	}
	return result, nil
}

//Result struct represent of graph partition
type Result struct {
	Matrix boolmatrixlib.IBoolMatrix
	Value  int64
}

//AsyncFindBestPartion function finds @graph optimal partition by bruteforce algorithm in range @start to @end in @amountOfGroups amount of groupse with
//@disbalance disbalance can work assync returns res to @ch
func AsyncFindBestPartion(graph *graphlib.Graph, start, end string, amountOfGroups int, disbalance float64, wg *sync.WaitGroup, ch chan Result) {
	log.Println("Starrting new gorutin")
	defer wg.Done()
	
	newStart, _ := big.NewInt(0).SetString(start, 0)
	newEnd, _ := big.NewInt(0).SetString(end, 0)

	res, err := FindBestPartion(graph, newStart, newEnd, amountOfGroups, disbalance)
	if err != nil {
		fmt.Println(err)
		return
	}
	ch <- res
}
