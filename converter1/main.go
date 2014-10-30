package main

import ("fmt";"bufio";"os";"path/filepath";"encoding/gob";"encoding/json";"strconv")

func iiiiiiiiiiii(i interface{}) {
fmt.Println("");i=bufio.Writer{};i=os.File{};filepath.IsAbs("");i=json.Decoder{}
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
	X int	`json:"x"`
	Y int	`json:"y"`
	Size int	`json:"size"`
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
//	used := make(map[string]bool)

	var v Fun

	for{
	    er1 := dec.Decode(&v)
	    if er1 != nil {
		fmt.Println("decode:", er1)
		break
	    }
		m[v.Xargs[0]]++
	}

	s.Close()
///////////////////////

	var graph Graph

	eid := 0
	did := 1
	fid := 0
	for cmd, cnt := range m {

		p := filepath.Dir(cmd)
		if dirzid[p] == 0 {
		for {
			if dirzid[p] == 0 {
				dir := Node{
					Id: "d" + strconv.Itoa(did),
					Label: p,
					Size: 1,
				}

				dirzid[p] = did
				graph.Nodes = append(graph.Nodes, dir)
				did++
			} else {
				break
			}

			if p == "/" {
				break
			}

			p = filepath.Dir(p)
		}
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
			Size: cnt,
		}
		graph.Nodes = append(graph.Nodes, binfile)
		fid++
	}

q, er1 := json.Marshal(graph)
if er1 != nil {
    fmt.Println("error:", er1)
}
os.Stdout.Write(q)
//	fmt.Println("**********")



//	_= enc.Encode(nodes)

//	os.Stdout.Write(b)

}
