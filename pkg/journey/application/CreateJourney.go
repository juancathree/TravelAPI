package application

import (
	matrixDomain "TravelAPI/pkg/matrixCost/domain"
	placeDomain "TravelAPI/pkg/place/domain"
	"sort"
)

type Distance struct {
	Place    placeDomain.Place
	Distance float64
}

type ByDistance []Distance

func (b ByDistance) Len() int {
	return len(b)
}

func (b ByDistance) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByDistance) Less(i, j int) bool {
	return b[i].Distance < b[j].Distance
}

func CreateJourney(clusters *[][]placeDomain.Place, places *[]placeDomain.Place, matrixCost *matrixDomain.MatrixCost) *[]placeDomain.Place {

	initialJourney := InitialJourney(clusters, places, matrixCost)

	var journey []placeDomain.Place
	journey = append(journey, (*initialJourney)[0]...)
	(*initialJourney) = (*initialJourney)[1:]

	for range *initialJourney {

		min := 100.0
		cl := 0
		posAdd := 0
		posCloser := 0

		for j := range *initialJourney {

			if (*matrixCost).Costs[journey[0].Name][(*initialJourney)[j][0].Name] < min {
				min = (*matrixCost).Costs[journey[0].Name][(*initialJourney)[j][0].Name]
				cl = j
				posAdd = 0
				posCloser = 0
			}
			if (*matrixCost).Costs[journey[0].Name][(*initialJourney)[j][len((*initialJourney)[j])-1].Name] < min {
				min = (*matrixCost).Costs[journey[0].Name][(*initialJourney)[j][len((*initialJourney)[j])-1].Name]
				cl = j
				posAdd = 0
				posCloser = len((*initialJourney)[j]) - 1
			}
			if (*matrixCost).Costs[journey[len(journey)-1].Name][(*initialJourney)[j][0].Name] < min {
				min = (*matrixCost).Costs[journey[len(journey)-1].Name][(*initialJourney)[j][0].Name]
				cl = j
				posAdd = len(journey) - 1
				posCloser = 0
			}
			if (*matrixCost).Costs[journey[len(journey)-1].Name][(*initialJourney)[j][len((*initialJourney)[j])-1].Name] < min {
				min = (*matrixCost).Costs[journey[len(journey)-1].Name][(*initialJourney)[j][len((*initialJourney)[j])-1].Name]
				cl = j
				posAdd = len(journey) - 1
				posCloser = len((*initialJourney)[j]) - 1
			}
		}

		if posAdd == 0 {

			if posCloser == 0 {
				journey = append(reverse((*initialJourney)[cl]), journey...)
			} else {
				journey = append((*initialJourney)[cl], journey...)
			}

		} else {

			if posCloser == 0 {
				journey = append(journey, (*initialJourney)[cl]...)
			} else {
				journey = append(journey, reverse((*initialJourney)[cl])...)
			}
		}

		(*initialJourney) = append((*initialJourney)[:cl], (*initialJourney)[cl+1:]...)
	}

	return &journey
}

func reverse(s []placeDomain.Place) []placeDomain.Place {
	var newS []placeDomain.Place
	for i := len(s) - 1; i >= 0; i-- {
		newS = append(newS, s[i])
	}
	return newS
}

func InitialJourney(clusters *[][]placeDomain.Place, places *[]placeDomain.Place, matrixCost *matrixDomain.MatrixCost) *[][]placeDomain.Place {
	distance := CalculateClusterDistance(clusters, matrixCost)

	journey := make([][]placeDomain.Place, len(*distance))

	for i := range *distance {

		if len((*distance)[i]) == 2 {
			journey[i] = append(journey[i], (*distance)[i][0].Place, (*distance)[i][1].Place)
			continue
		} else if len((*distance)[i]) > 2 {
			journey[i] = append(journey[i], (*distance)[i][0].Place, (*distance)[i][1].Place)
		} else if len((*distance)[i]) == 1 {
			journey[i] = append(journey[i], (*distance)[i][0].Place)
		}

		for j := 2; j < len((*distance)[i]); j++ {

			min := 1000.0
			pos := -1
			p := (*distance)[i][j].Place

			for k := 0; k < len(journey[i])+1; k++ {
				cost := 0.0
				for d := 0; d < len(journey[i]); d++ {
					if k == d || k-d == 1 {
						cost += (*matrixCost).Costs[p.Name][journey[i][d].Name]
					} else if k-d < 0 {
						cost += (*matrixCost).Costs[journey[i][d].Name][journey[i][d-1].Name]
					} else {
						cost += (*matrixCost).Costs[journey[i][d].Name][journey[i][d+1].Name]
					}
				}
				if cost < min {
					min = cost
					pos = k
				}
			}

			if pos >= len(journey[i]) {
				journey[i] = append(journey[i], p)
			} else {
				journey[i] = append(journey[i][:pos+1], journey[i][pos:]...)
				journey[i][pos] = p
			}
		}
	}

	return &journey

}

func CalculateClusterDistance(clusters *[][]placeDomain.Place, matrixCost *matrixDomain.MatrixCost) *[][]Distance {
	distance := make([][]Distance, len(*clusters))

	for i := range *clusters {
		for j, p := range (*clusters)[i] {
			d := Distance{
				Place:    p,
				Distance: 0.0,
			}
			for k, p2 := range (*clusters)[i] {
				if j != k {
					d.Distance += matrixCost.Costs[p.Name][p2.Name]
				}
			}
			distance[i] = append(distance[i], d)
		}
		sort.Sort(ByDistance(distance[i]))
	}

	return &distance
}
