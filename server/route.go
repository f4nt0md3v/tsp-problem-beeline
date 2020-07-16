package server

import (
	"context"
	"sort"
	"sync"

	"go.uber.org/zap"

	pb "github.com/f4nt0md3v/tsp-problem-beeline/proto/route"
)

// RouteServer is a gRPC server it implements the methods defined by the RouteServer interface
type RouteServer struct {
	log *zap.SugaredLogger
}

// NewRouteServer - constructor function - creates a new RouteServer
func NewRouteServer(log *zap.SugaredLogger) *RouteServer {
	return &RouteServer{log: log}
}

// GetRoute implements the RouteServer GetRoute method and returns the paths
// for the given list of cities
func (s *RouteServer) GetRoute(_ context.Context, r *pb.RouteRequest) (*pb.RouteResponse, error) {
	startCity := r.GetStartCity()
	cities := r.GetCitiesList()
	s.log.Infow("[GetRoute]", "start_city", startCity, "cities_list", cities)

	if !contains(cities, startCity) {
		errMsg := "cities_list does not contain start_city"
		s.log.Error(errMsg)
		return &pb.RouteResponse{Error: errMsg}, nil
	}

	var (
		arr    []int
		routes []*pb.RoutePath
		wg     sync.WaitGroup
	)
	// exclude startCity to save time for calculations
	for i, city := range cities {
		if city != startCity {
			arr = append(arr, i+1)
		}
	}

	for _, perm := range permutations(arr) {
		var (
			p  = []string{startCity}
			rp pb.RoutePath
		)
		for _, c := range perm {
			p = append(p, cities[c-1])
		}
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			rp = prepareRoutePathResponse(p)
		}(&wg)
		wg.Wait()
		routes = append(routes, &rp)
	}

	// sorting by distance happens here
	sort.SliceStable(routes, func(i, j int) bool {
		return routes[i].Distance > routes[j].Distance
	})

	return &pb.RouteResponse{Routes: routes}, nil
}

func permutations(arr []int) [][]int {
	var (
		res     [][]int
		permute func([]int, int)
	)

	permute = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				permute(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	permute(arr, len(arr))
	return res
}

// Contains tells whether a contains x.
func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// Info taken from http://42-60-60.kz/pages/poleznaya-informatsiya/rasstoyaniya-mezhdu-gorodami-kazakhstana
var (
	distMtx = [17][17]int{
		{0, 2737, 1448, 1224, 2031, 196, 1155, 276, 694, 426, 476, 1359, 1326, 1926, 1038, 1505, 1315},
		{2737, 0, 1301, 3493, 893, 2948, 2365, 2500, 2059, 3184, 2509, 3732, 2986, 1388, 3796, 2810, 2441},
		{1448, 1301, 0, 2191, 595, 1659, 1063, 1201, 760, 1895, 1210, 2430, 1684, 461, 2507, 1508, 1353},
		{1224, 3493, 2191, 0, 2787, 1013, 1123, 1523, 1937, 1444, 1723, 249, 498, 2680, 1068, 677, 832},
		{2031, 893, 595, 2787, 0, 2242, 1659, 1796, 1355, 2478, 1805, 3026, 2280, 488, 3090, 2104, 1847},
		{196, 2948, 1659, 1013, 2242, 0, 929, 495, 909, 412, 695, 1148, 1104, 2137, 923, 1283, 1255},
		{1155, 2365, 1063, 1123, 1659, 929, 0, 1304, 1210, 1368, 1365, 1383, 617, 1553, 1887, 441, 289},
		{276, 2500, 1201, 1523, 1796, 495, 1304, 0, 430, 715, 188, 1658, 1625, 1698, 1327, 1804, 1525},
		{694, 2059, 760, 1937, 1355, 909, 1210, 430, 0, 1141, 439, 2072, 1834, 1250, 1753, 1658, 1506},
		{426, 3184, 1895, 1444, 2478, 412, 1368, 715, 1141, 0, 915, 1184, 1539, 2373, 592, 1710, 1667},
		{476, 2509, 1210, 1723, 1805, 695, 1365, 188, 439, 915, 0, 1858, 1825, 1700, 1527, 1813, 1711},
		{1359, 3732, 2430, 249, 3026, 1148, 1383, 1658, 2072, 1184, 1858, 0, 758, 1919, 809, 937, 1092},
		{1326, 2986, 1684, 498, 2280, 1104, 617, 1625, 1834, 1539, 1825, 758, 0, 2174, 1577, 167, 329},
		{1926, 1388, 461, 2680, 488, 2137, 1553, 1698, 1250, 2373, 1700, 1919, 2174, 0, 2985, 1998, 1827},
		{1038, 3796, 2507, 1068, 3090, 923, 1887, 1327, 1753, 592, 1527, 809, 1577, 2985, 0, 1756, 2115},
		{1505, 2810, 1508, 677, 2104, 1283, 441, 1804, 1658, 1710, 1813, 937, 167, 1998, 1756, 0, 167},
		{1315, 2441, 1353, 832, 1847, 1255, 289, 1525, 1506, 1667, 1711, 1092, 329, 1827, 2115, 167, 0},
	}
	cityDict = []string{
		"Нур-Султан", "Актау", "Актобе", "Алматы", "Атырау", "Караганда", "Кызылорда", "Кокшетау", "Костанай",
		"Павлодар", "Петропавловск", "Талдыкорган", "Тараз", "Уральск", "Усть-Каменогорск", "Шымкент", "Туркестан",
	}
)

type City int

func getCityIdx(city string) City {
	for i, c := range cityDict {
		if city == c {
			return City(i)
		}
	}
	return -1
}

func prepareRoutePathResponse(cityPath []string) pb.RoutePath {
	n := len(cityPath)
	var (
		distance int64
		path     string
	)
	for i, city := range cityPath {
		if i < n-1 {
			// get city indexes in distMtx
			cur := getCityIdx(city)
			next := getCityIdx(cityPath[i+1])
			// calculate overall distance
			distance += int64(distMtx[cur][next])

			path += city + "->"
		} else {
			path += city
		}
	}
	return pb.RoutePath{
		Path:     path,
		Distance: distance,
	}
}
