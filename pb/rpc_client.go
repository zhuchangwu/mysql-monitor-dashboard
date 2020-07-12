package pb

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	rpc_task "mysql-monitor-dashboard/pb/monitor/task"
	flow_center "oa-flow-center/proto/oa/flow-center"
)

/**
 * mysql-moniotr作为rpc-clinet的实现
 * client和server交互 obj
 */
func sendMsgTo(task *rpc_task.RpcSysTask) {
	// 创建一个连接
	conn, err := grpc.Dial(":8082", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}
	defer conn.Close()

	// 基于第一个连接，创建客户端
	client :=	rpc_task.NewFlowServiceClient(conn)

	// 构建请求参数
	f := &flow_center.Flow{
		FlowType:  "reibersement",
		RecordId:  1,
		Applicate: "changwu",
		RoleMap: map[string]string{
			"manager": "tom",
			"boss":    "jerry",
		},
		DepartmentId: 2,
	}
	// 请求对象
	obj:=&rpc_task.RpcSysTask{
		ItemType:rpc_task.Type_Sys,
		Sysmsg:&rpc_task.SysMsg{
			ItemName: "cpu",
			Cycle: "10",
		},
		Appmsg:nil,
		Mysqlmsg:nil,
	}
	// todo 临时测试一下
	task = obj

	// 发送grpc请求
	response, err := client.SendTaskToMysqlMonitor(context.Background(), task)
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}
	// 解析响应结果
	fmt.Printf("flowId:[%v]\n", response.GetStatus())
	fmt.Printf("responseMsg:[%v]\n", response.GetResponseMsg())
}
