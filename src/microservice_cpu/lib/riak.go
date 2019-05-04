package lib

import (
	"fmt"
	"github.com/basho/riak-go-client"
	"os"
)

type dataType struct {
	Id          int64
	Time        int64
	CPU float64
}

func GetData(pid, startDate, endDate string) []dataType {
	var data []dataType
	nodeOpts := &riak.NodeOptions{
		RemoteAddress: "localhost:8087",
	}

	var node *riak.Node
	var err error
	if node, err = riak.NewNode(nodeOpts); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	nodes := []*riak.Node{node}
	opts := &riak.ClusterOptions{
		Nodes: nodes,
	}

	cluster, err := riak.NewCluster(opts)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer func() {
		if err := cluster.Stop(); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}()

	if err := cluster.Start(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	cmd, err := riak.NewTsQueryCommandBuilder().WithQuery(
		"select * from CPU " +
			"where " +
			"time > '" + startDate + "' and " +
			"time < '" + endDate + "' and " +
			"id = " + pid).Build()
	if err != nil {
		fmt.Print(err)
	}

	err = cluster.Execute(cmd)
	scmd, _ := cmd.(*riak.TsQueryCommand)
	for i := 0; i < len(scmd.Response.Rows); i++ {
		data = append(data, dataType{
			scmd.Response.Rows[i][0].GetSint64Value(),
			scmd.Response.Rows[i][2].GetTimestampValue(),
			scmd.Response.Rows[i][3].GetDoubleValue(),
		})
		fmt.Println(data)
	}
	return data
}
