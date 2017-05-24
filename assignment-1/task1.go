package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var (
	debug = false
)

type Line struct {
	Customers []Customer
	Size      int
}

type Customer struct {
	Time  int
	IsNew bool
}

type Server struct {
	TimeLeft int
	RateOld  int
	RateNew  int
	Breaks   map[int]int
	OnBreak  bool
}

func (q *Line) Pop() (cust Customer, err error) {
	if q.Size > 0 {
		cust = q.Customers[0]
		q.Customers = q.Customers[1:]
		q.Size--
		return cust, nil
	} else {
		return cust, errors.New("Empty queue, can't pop")
	}
}

func (q *Line) Peek() (cust Customer, err error) {
	if q.Size > 0 {
		cust = q.Customers[0]
		return cust, nil
	} else {
		return cust, errors.New("Empty queue, can't peek")
	}
}

func (q *Line) Push(cust Customer) {
	q.Customers = append(q.Customers, cust)
	q.Size++
}

type Sim struct {
	Replications     int
	ArrivalDist      int
	InterArrivalTime float64
	Start            int
	End              int
	RateOfNew        float64
	ScheduleNew      int
	Servers          []Server
	Queue            Line
}

func initSim() (s Sim) {

	fmt.Print("How many replications of the system: ")

	fmt.Scanf("%v", &s.Replications)

	fmt.Print("What arrival of distributions; 0) random 1) scheduled: ")

	fmt.Scanf("%v", &s.ArrivalDist)

	fmt.Print("What is mean time between arrivals: (in minutes) ")

	fmt.Scanf("%v", &s.InterArrivalTime)

	fmt.Print("What start time (in minutes): ")

	fmt.Scanf("%v %v", &s.Start)

	fmt.Print("What end time (in minutes): ")

	fmt.Scanf("%v %v", &s.End)

	fmt.Print("What percentage of new customers (0.00 - 1.00): ")

	fmt.Scanf("%v", &s.RateOfNew)

	fmt.Print("When will we schedule new customers? 0) first; 1) last; 2) randomly: ")

	fmt.Scanf("%v", &s.ScheduleNew)

	fmt.Print("How many servers working: ")

	var servs int
	fmt.Scanf("%v", &servs)

	for i := 1; i <= servs; i++ {
		var rateOld int
		var rateNew int
		var breaks int

		fmt.Printf("For worker %v:\n", i)
		fmt.Printf("  How many minutes for service (scheduled patient): ")

		fmt.Scanf("%v", &rateOld)

		fmt.Printf("  How many minutes for service (new patient): ")

		fmt.Scanf("%v", &rateNew)

		fmt.Printf("  How many breaks does this worker take: ")

		fmt.Scanf("%v", &breaks)

		br := make(map[int]int, breaks)
		for j := 1; j <= breaks; j++ {
			var start int
			var duration int

			fmt.Printf("  What is the start time and duration of break %v in minutes: ", j)

			fmt.Scanf("%v %v", &start, &duration)
			br[start] = duration
		}
		serv := Server{
			RateOld:  rateOld,
			RateNew:  rateNew,
			Breaks:   br,
			TimeLeft: 0,
			OnBreak:  false,
		}
		s.Servers = append(s.Servers, serv)
	}

	if debug {
		fmt.Println()
	}
	return s
}

func main() {

	rand.Seed(time.Now().UnixNano())

	s := initSim()

	ast, awt, mst, mwt, pit, acs, acw, serv := simulate(&s)

	fmt.Printf("Average system time: %v\n", ast)
	fmt.Printf("Average queue time: %v\n", awt)
	fmt.Printf("Max system time: %v\n", mst)
	fmt.Printf("Max queue time: %v\n", mwt)
	fmt.Printf("Proportion of time idle: %v\n", pit)
	fmt.Printf("Average customers in system: %v\n", acs)
	fmt.Printf("Average customers in queue %v\n", acw)
	fmt.Printf("Average customers served %v\n", serv)
}

func bleedArrivals(s *Sim) {
	for cust, err := s.Queue.Pop(); err == nil; cust, err = s.Queue.Pop() {
		fmt.Println(cust.Time)
	}
}

func printArrivals(s *Sim) {
	fmt.Println("Schedule of arrivals")
	for i, e := range s.Queue.Customers {
		fmt.Printf("%v: %v\n", i, e.Time)
	}
}

func simulate(s *Sim) (ast, awt, mst, mwt, pit, acs, acw, aveserv float64) {

	start := time.Now()
	for i := 0; i < s.Replications; i++ {
		var q Line

		sysTime := 0
		maxSys := 0
		totalSys := 0

		//servers := make([]int, s.Servers)

		waitTime := 0
		maxWait := 0
		totalWait := 0

		idleTime := 0
		custTotal := 0

		createArrivals(s)
		//printArrivals(s)

		idling := false

		for i := s.Start; i < s.End || !(q.Size <= 0) || !idling; i++ {
			if debug {
				fmt.Printf("Time: %v, Queue size: %v\n", i, q.Size)
			}
			idling = true

			// do we need to add a new arrival to the queue

			for next, err := s.Queue.Peek(); err == nil && next.Time <= i; next, err = s.Queue.Peek() {
				cust, err := s.Queue.Pop()
				if err == nil {
					if debug {
						fmt.Printf("A customer arrived\n")
					}
					q.Push(cust)
				}
			}

			// can our servers deal with someone
			for j, _ := range s.Servers {
				passMinute(j, &s.Servers[j], i)
				if s.Servers[j].TimeLeft == 0 {
					cust, err := q.Peek()
					if err != nil {
						// the server is idling
						idleTime++
					} else {
						idling = false
						q.Pop()
						if debug {
							fmt.Printf("Server %v took on a customer\n", j)
						}
						thisWait := i - cust.Time
						thisSys := thisWait
						if cust.IsNew {
							serviceTime := int(float64(s.Servers[j].RateNew) * rand.ExpFloat64())
							s.Servers[j].TimeLeft += serviceTime
							thisSys += serviceTime
						} else {
							serviceTime := int(float64(s.Servers[j].RateOld) * rand.ExpFloat64())
							s.Servers[j].TimeLeft += serviceTime
							thisSys += serviceTime
						}

						if thisWait > maxWait {
							maxWait = thisWait
						}
						if thisSys > maxSys {
							maxSys = thisSys
						}

						waitTime += thisWait
						sysTime += thisSys
						custTotal++
						totalSys++
					}
				} else if !s.Servers[j].OnBreak {
					// the server is with a patient
					idling = false
					totalSys++
				} else {
					// the server is on a break
					if debug {
						fmt.Printf("Server %v is on a break\n", j)
					}
				}
			}

			totalWait += q.Size
			totalSys += q.Size

		}
		elapsed := float64(s.End - s.Start)
		ast += float64(sysTime) / float64(custTotal)
		awt += float64(waitTime) / float64(custTotal)
		if float64(maxSys) > mst {
			mst = float64(maxSys)
		}
		if float64(maxWait) > mwt {
			mwt = float64(maxWait)
		}
		pit += float64(idleTime) / elapsed
		acs += float64(totalSys) / elapsed
		acw += float64(totalWait) / elapsed
		aveserv += float64(custTotal)
	}

	ast = ast / float64(s.Replications)
	awt = awt / float64(s.Replications)
	mst = mst
	mwt = mwt
	pit = pit / float64(s.Replications) / float64(len(s.Servers))
	acs = acs / float64(s.Replications)
	acw = acw / float64(s.Replications)
	aveserv = aveserv / float64(s.Replications)

	elapsed := time.Since(start)
	fmt.Printf("%v replications took %s\n", s.Replications, elapsed)
	return
}

func passMinute(id int, serv *Server, time int) {
	if serv.TimeLeft > 0 {
		serv.TimeLeft--
	}

	// account for breaks
	if serv.TimeLeft == 0 {
		if serv.OnBreak {
			serv.OnBreak = false
		}
		for j, b := range serv.Breaks {
			if j < time && time < j+b {
				if debug {
					fmt.Printf("Server %v will go on a break for %v minutes until %v\n", id, b, time+b)
				}
				serv.OnBreak = true
				serv.TimeLeft += b
				break
			}
		}
	}
}

func createArrivals(s *Sim) {
	if s.ArrivalDist == 0 { // random (poisson distributed)

		s.Queue.Customers = make([]Customer, 0)
		s.Queue.Size = 0
		var isNew bool
		duration := s.End - s.Start
		totalCusts := int(float64(duration) * s.RateOfNew / s.InterArrivalTime)
		sN := s.ScheduleNew
		for lastArrival := float64(s.Start); lastArrival < float64(s.End); {
			lastArrival = lastArrival + (s.InterArrivalTime * (rand.ExpFloat64()))
			lenCusts := len(s.Queue.Customers)
			isNew = (rand.Float64() < s.RateOfNew && sN == 2) || (sN == 0 && lenCusts < totalCusts) || (sN == 1 && lenCusts > totalCusts)
			cust := Customer{
				Time:  int(lastArrival),
				IsNew: isNew,
			}

			s.Queue.Push(cust)
		}

	} else { // scheduled, == 1

		s.Queue.Customers = make([]Customer, 0)
		s.Queue.Size = 0

		rand.Seed(time.Now().UnixNano())
		var isNew bool
		var arrivalTime float64
		duration := s.End - s.Start
		totalCusts := int(float64(duration) * s.RateOfNew / s.InterArrivalTime)
		sN := s.ScheduleNew
		for i := s.Start; i < s.End; i += int(s.InterArrivalTime) {
			arrivalTime = float64(i) + (s.InterArrivalTime * ((rand.Float64() - 0.5) * 0.5))
			lenCusts := len(s.Queue.Customers)
			isNew = (rand.Float64() < s.RateOfNew && sN == 2) || (sN == 0 && lenCusts < totalCusts) || (sN == 1 && lenCusts > totalCusts)
			cust := Customer{
				Time:  int(arrivalTime),
				IsNew: isNew,
			}
			s.Queue.Push(cust)
		}
	}
}
