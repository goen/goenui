package main

import ("fmt";"bufio";"os";"path/filepath";"encoding/gob";"encoding/json";"strconv";"io";"math/rand";"math";"html/template")

func iiiiiiiiiiii(i interface{}) {
fmt.Println("");i=bufio.Writer{};i=os.File{};filepath.IsAbs("");i=json.Decoder{};i=io.EOF;i=rand.Rand{}
math.Log(0)
}

type Fun struct {
	Tid        uint32 //thread specific id
	Sysfun     uint16 // platform specific syscall id
	Reterr     int64  // return error code /child's Tid
	Sig        uint8
	Sec        uint64
	NSec       uint32
	Unfinished bool //beginning
	Resumed    bool //end
	Signal     bool
	Xargs      []string //exec & args
}

func getgbin() (s string, e error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("No current wd")
	}
	for {
		wd = filepath.Dir(wd)
		s = wd + "/" + mpoint_gbin
		//				fmt.Println(s)
		_, err := os.Stat(s)
		if err == nil {
			return s, nil
		}
		if "/" == wd {
			break
		}
	}
	return "", fmt.Errorf("goen dir not found")
}

func lddl() {
	fmt.Println("making the pipeline")

	gb, er4 := getgbin()
	if er4 != nil {
		fmt.Println(er4)
		return
	}

	s, er0 := os.Open(gb+"/loop")
	if er0 != nil {
		fmt.Println(er0)
		return
	}

	t, er1 := os.Open(gb+"/trace")
	if er1 != nil {
		fmt.Println(er1)
		return
	}

	s.Close()
	t.Close()
}

type Node struct {
	Id string	`json:"id"`
	Label string	`json:"label"`
	X float64	`json:"x"`
	Y float64	`json:"y"`
	Size float64	`json:"size"`
	Color string	`json:"color"`
}
type Edge struct {
	Id string	`json:"id"`
	Source string	`json:"source"`
	Target string	`json:"target"`
}
type Graph struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge  `json:"edges"`
}

func rootnode() Node {
	return Node{
		Id: "d0",
		Label: "/",
		Size: 2,
		X: 0,
		Y: 0,
	}
}

func main() {
	// the reading part
	s, er0 := os.Open("trace.gob")
	if er0 != nil {
		fmt.Println(er0)
		return
	}

	r := bufio.NewReader(s)
	dec := gob.NewDecoder(r)


	m := make(map[string]int)
	dirzid := make(map[string]int)
	dirid2chld := make(map[int]int)
//	used := make(map[string]bool)

	var v Fun

	for{
	    er1 := dec.Decode(&v)
	    if er1 != nil {
		if er1 == io.EOF { break }
		fmt.Println("decode:", er1)
		return
	    }
		m[v.Xargs[0]]++
	}

	s.Close()
///////////////////////

	var graph Graph

	eid := 0
	did := 1
	fid := 0

	_ = fid
	_ = eid

	for cmd, cnt := range m {
		_ = cnt

		p := filepath.Dir(cmd)
		for {
			if dirzid[p] == 0 {

//				fmt.Fprintln(os.Stderr, "Here we have a new dir:",did,":", p)


				dir := Node{
					Id: "d" + strconv.Itoa(did),
					Label: p,
					Size: 0.2,
					X: 0.,
					Y: 0.,
					Color: "#f00",
				}

				dirzid[p] = did
				graph.Nodes = append(graph.Nodes, dir)
				did++
			} else {
				break
			}

			if len(p) <= 1 {
				break
			}

			p = filepath.Dir(p)

			if p == "" {
				p = "."
			}

			to := dirzid[p]
			if to == 0 {
				to = did
			}

			dirid2chld[to]++

//			fmt.Fprintln(os.Stderr, "Making a connection from:", did-1, "to:", to)
			edge := Edge{
				Id: "e" + strconv.Itoa(eid),
				Source: "d" + strconv.Itoa(did-1),
				Target: "d" + strconv.Itoa(to),
			}
			graph.Edges = append(graph.Edges, edge)
			eid++
		}


		p = filepath.Dir(cmd)
		if dirzid[p] != 0 {
			//make a connection 
			edge := Edge{
				Id: "e" + strconv.Itoa(eid),
				Source: "d" + strconv.Itoa(dirzid[p]),
				Target: "f" + strconv.Itoa(fid),
			}
			graph.Edges = append(graph.Edges, edge)
			eid++
		}


		binfile := Node{
			Id: "f" + strconv.Itoa(fid),
			Label: cmd,
			Size: math.Log(float64(cnt)),
			X: (rand.Float64() - 0.5)*0.5,
			Y: (rand.Float64() - 0.5)*0.5,
			Color: "#00f",
		}
		graph.Nodes = append(graph.Nodes, binfile)
		fid++

	}

q, er1 := json.Marshal(graph)
if er1 != nil {
    fmt.Println("error:", er1)
}
os.Stdout.Write(q)


}
