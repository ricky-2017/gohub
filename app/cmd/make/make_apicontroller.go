package make

import (
	"fmt"
	"gohub/pkg/console"
	"strings"

	"github.com/spf13/cobra"
)

var CmdMakeAPIController = &cobra.Command{
	Use:   "apicontroller",
	Short: "Create api controller,example : make apicontroller v1/user",
	Run:   runMakeAPIController,
	Args:  cobra.ExactArgs(1),
}

func runMakeAPIController(cmd *cobra.Command, args []string) {
	// 参数处理
	array := strings.Split(args[0], "/")
	if len(array) != 2 {
		console.Exit("api controller name format: {{api-version}}/{{api-name}}")
	}

	apiVersion, name := array[0], array[1]
	model := makeModelFromString(name)

	// 组建目标目录
	filePath := fmt.Sprintf("app/http/controllers/api/%s/%s_controller.go", apiVersion, model.TableName)

	createFileFromStub(filePath, "apicontroller", model)
}
