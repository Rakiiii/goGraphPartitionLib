package graphpartitionlib

import (
	"errors"
	"fmt"
	"math"
	"math/big"

	bigintlib "github.com/Rakiiii/goBigIntLib"
	boolmatrixlib "github.com/Rakiiii/goBoolMatrix"
	graphlib "github.com/Rakiiii/goGraph"
)

func FindBestPartion(graph *graphlib.Graph, start, end *big.Int, amountOfGroups int, disbalance float64) (*boolmatrixlib.BoolMatrix, int64, error) {
	var bestParameterValue int64 = math.MaxInt64
	bestMatrix := new(boolmatrixlib.BoolMatrix)

	amountOfVertex := graph.AmountOfVertex()
	flag := true

	subMatrix := new(boolmatrixlib.BoolMatrix)
	subMatrix.Init(amountOfGroups, amountOfVertex)
	for start.Cmp(end) < 0 {
		subMatrix.SetByNumber(start)
		DebugLog("checking" + start.String())

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
					return nil, bestParameterValue, err
				}
				if subParameterValue < bestParameterValue {
					bestMatrix = subMatrix.Copy()
					bestParameterValue = subParameterValue
				}
			}
			bigintlib.Inc(start)
		}
	}

	return bestMatrix, bestParameterValue, nil
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
	fmt.Println(str)
}
