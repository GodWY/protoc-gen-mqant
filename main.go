package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/liangdas/mqant/log"
	"sigs.k8s.io/yaml"
)

// 代码生成结构体
type CodeGenerator struct {
	// 生成go文件包名
	Package string
	// 生成go文件的导入包
	ImportsPath []string
	// 服务的名字
	Services string
	// 包含的协议
	Topic []string
	// 导出地址
	ExportPath string
	// 注册协议执行函数
	Methods      map[string]string
	TopicAndFunc map[string]string
}

// 获取包名
func GetPkg(v string) string {
	return strings.TrimSpace(strings.Split(v, "package")[1])
}

// 获取导入的包名
func GetImportsPath(v []string) []string {
	return []string{}
}

// 获取导出的包名
func GetExportPath(v string) string {
	return strings.TrimSpace(strings.Split(v, "=")[1])
}

// 获取协议的集合
func GetTopicFunction(options string) []string {
	oo := strings.Split(options, "{")[1]
	funcs := strings.Split(strings.Split(oo, "}")[0], "rpc")
	result := []string{}
	for _, v := range funcs {
		vs := strings.Split(v, "(")
		if len(vs) < 2 {
			continue
		}
		v = strings.TrimSpace(v)
		result = append(result, v)
	}
	return result
}

// 获取服务的名字
func GetService(options string) string {
	return strings.TrimSpace(strings.Split(options, "{")[0])
}

// 生成协议
func makeMethod(pkg, options string) string {
	// 生成规则暂定为 包名+函数名
	return pkg + "/" + options
}

// 生成结构体
func NewCodeGenerator(proto string) CodeGenerator {
	one, two := SplitCode(proto)
	var package_1, go_package string
	others := strings.Split(one, ";")
	var imports []string
	for _, v := range others {
		if strings.Contains(v, "package") && !strings.Contains(v, "option") {
			v = strings.TrimSpace(v)
			package_1 = strings.Split(v, "package")[1]
			package_1 = strings.TrimSpace(package_1)
		}
		if strings.Contains(v, "option") {
			go_package = strings.Split(v, "=")[1]
			go_package = strings.TrimSuffix(go_package, `"`)
			go_package = strings.Split(go_package, `"`)[1]
		}
		if strings.Contains(v, "import") {
			pp := strings.Split(v, ".proto")[0]
			imports = append(imports, strings.Split(pp, "import")[1])
		}
	}

	cod := CodeGenerator{
		Package:     package_1,
		ImportsPath: imports,
		Services:    GetService(two),
		Topic:       GetTopicFunction(two),
		ExportPath:  go_package,
	}
	method := make(map[string]string, len(cod.Topic))
	funcs := make(map[string]string, len(cod.Topic))
	for _, v := range cod.Topic {
		// 分解函数签名
		vs := strings.Split(v, "(")
		method[makeMethod(cod.Package, vs[0])] = vs[0]
		funcs[makeMethod(cod.Package, vs[0])] = v
	}
	cod.Methods = method
	cod.TopicAndFunc = funcs
	return cod
}

// 分割字符串为 服务和导包
func SplitCode(proto string) (one, two string) {
	topics := strings.Split(proto, "service")
	return topics[0], topics[1]
}

type conf struct {
	Input  string `yaml:"input"`
	Output string `yaml:"output"`
	Proto  string `yaml:"msg"`
	Rpc    string `yaml:"rpc"`
}

// mustLoadConfiguration 读取yaml配置
func mustLoadConfiguration() *conf {
	c := &conf{}
	f, err := ioutil.ReadFile("release.yaml")
	if err != nil {
		panic(err.Error())
	}
	if err := yaml.Unmarshal(f, &c); err != nil {
		panic("error")
	}
	return c
}
func main() {
	conf := mustLoadConfiguration()
	if conf == nil {
		return
	}
	// input := conf.Input
	output := conf.Output
	if output == "" {
		return
	}
	os.RemoveAll(output)
	err := os.MkdirAll(output, os.ModePerm)
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}
	input := conf.Input
	if input == "" {
		return
	}
	var cmd *exec.Cmd
	// 执行单个shell命令时, 直接运行即可

	// ls ../api/pb/message/*.proto | xargs protoc -I=../api/pb/ --go_out=../pkg/golang --go_opt=paths=source_relative
	pb := input + "/" + conf.Proto
	dirs, err := ioutil.ReadDir(pb)
	if err != nil {
		log.Error("read dir error: %v", err)
		return
	}

	// 生成结构pb文件文件
	for _, d := range dirs {
		sh := "protoc -I %s --go_out=pkg/golang %s"
		if !d.IsDir() {
			// 如果不是目录则使用上级目录
			sh = fmt.Sprintf(sh, input, input+"/"+d.Name())
			cmd = exec.Command(sh)
			cmd.Run()
			continue //
		}
		// 这里只允许二级嵌套
		chird := pb + "/" + d.Name()
		chirdDirs, err := ioutil.ReadDir(chird)
		if err != nil {
			log.Error("****************************", d.Name())
			continue
		}

		for _, f := range chirdDirs {
			cmd := exec.Command("protoc", "-I", chird, "--go_out", output, chird+"/"+f.Name())
			err := cmd.Run()
			if err != nil {
				log.Error("xxxxxxx", err)
			}
		}
	}
	rpc := input + "/" + conf.Rpc
	// 生成rpc消息
	rpcDirs, err := ioutil.ReadDir(rpc)
	if err != nil {
		log.Error("xxxx read_dir error", err)
		return
	}

	for _, dir := range rpcDirs {
		// 同样只允许两级嵌套
		if !dir.IsDir() {
			continue
		}
		netMsg := rpc + "/" + dir.Name()
		chirDirs, err := ioutil.ReadDir(netMsg)
		if err != nil {
			log.Error("xxxx read_dir error", err)
		}
		for _, fs := range chirDirs {
			filnae := netMsg + "/" + fs.Name()
			file, _ := ioutil.ReadFile(filnae)
			t, _ := template.New("template").Parse(tp0)
			conf := NewCodeGenerator(string(file))
			dirPath := output + "/" + conf.ExportPath
			out := dirPath + "/" + fs.Name() + ".go"
			os.Remove(dirPath)
			os.MkdirAll(dirPath, os.ModePerm)
			// os.Chdir(s)
			files, err := os.Create(out)
			if err != nil {
				fmt.Println("%v", err)
				return
			}
			t.Execute(files, conf)
			ExecGoFmt(out)
		}
	}

}

// 执行gofmt
func ExecGoFmt(file string) {
	cmd := exec.Command("gofmt", "-w", file)
	err := cmd.Run()
	if err != nil {
		log.Error("xxxxxxx", err)
	}
}

var tp0 = `
package {{.Package}}

import (
	"errors"
	"github.com/liangdas/mqant/gate"
	basemodule "github.com/liangdas/mqant/module/base"
	"github.com/liangdas/mqant/gate"
	client "github.com/liangdas/mqant/module"
	mqrpc "github.com/liangdas/mqant/rpc"
	"golang.org/x/net/context"
	{{- range $key, $value := .ImportsPath }}
  	{{$value}}
	{{- end}}
)


// var Register{{.Services}}TcpHandler = Register{{.Services}}TcpHandler

type {{.Services}} interface {
{{- range $key, $value := .Topic }}
  {{$value}}
{{- end}}
}

// 注册路由协议
func Register{{.Services}}TcpHandler(m *basemodule.BaseModule, ser {{.Services}}) {
{{- range $key, $value := .Methods}}
 m.GetServer().RegisterGO("{{$key}}", ser.{{$value}})
{{- end}}
}


// rpc 请求代理
type ClientProxyService struct {
	cli client.App
	name string
}

// 获取client实例
func NewLoginClient(cli client.App, name string) *ClientProxyService {
	return &ClientProxyService{
		cli: cli,
		name: name,
	}
}


{{- range $key, $value := .Methods}}
var {{$value}} = "{{$key}}"
{{- end}}

var ClientProxyIsNil = errors.New("proxy is nil")

{{- range $key, $value := .TopicAndFunc }}
func (proxy *ClientProxyService){{$value}} {
	if proxy == nil {
		return nil, ClientProxyIsNil
	}
	err = mqrpc.Proto(rsp, func() (reply interface{}, err interface{}) {
		return proxy.cli.Call(context.TODO(), proxy.name, "{{$key}}", mqrpc.Param(req))
	})
	return rsp, err
}
{{- end}}
`
