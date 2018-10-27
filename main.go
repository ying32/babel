package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"path"
	"path/filepath"

	"strings"

	"io/ioutil"
	"os"

	"github.com/dop251/goja"
)

func main() {
	if len(os.Args) > 1 {

		var opts interface{}
		opts = defConfig()

		configbytes, err := ioutil.ReadFile("./babel.json")
		if err == nil {
			json.Unmarshal(configbytes, &opts)
		}

		vm := goja.New()
		_, err = vm.RunScript("babel.js", string(babelJsBytes))
		if err != nil {
			panic(fmt.Errorf("运行babel.js错误：%s", err))
		}

		_, err = vm.RunScript("polyfill.js", string(polyfillJsBytes))
		if err != nil {
			panic(fmt.Errorf("运行polyfill.js错误：%s", err))
		}

		_, err = vm.RunScript("运行babili.js错误：", string(babiliJsBytes))
		if err != nil {
			panic(fmt.Errorf("运行babili.js错误：%s", err))
		}

		var transformfn goja.Callable
		babel := vm.Get("Babel")
		if err := vm.ExportTo(babel.ToObject(vm).Get("transform"), &transformfn); err != nil {
			panic(fmt.Errorf("未找到”transform“函数，错误: %s", err))
		}

		// 指定目录下的

		if path.Ext(os.Args[1]) == "" {
			filepath.Walk(os.Args[1], func(filename string, info os.FileInfo, err error) error {
				if strings.ToLower(path.Ext(info.Name())) == ".jsx" {
					err = transformAndOutFile(filename, babel, vm, transformfn, opts)
					if err != nil {
						panic(err)
					}
				}
				return nil
			})
		} else {
			err = transformAndOutFile(os.Args[1], babel, vm, transformfn, opts)
			if err != nil {
				panic(err)
			}
		}
	}
}

func transformAndOutFile(filename string, babel goja.Value, vm *goja.Runtime, transformfn goja.Callable,
	opts interface{}) error {
	fmt.Println("翻译文件：", filename)
	codeBytes := readFile(filename)
	if len(codeBytes) == 0 {
		return errors.New("读源代码错误，文件：" + filename)
	}
	v, err := transformfn(babel, vm.ToValue(string(codeBytes)), vm.ToValue(opts))
	if err != nil {
		return err
	}
	newFileName := path.Base(filename)
	if strings.Count(newFileName, ".") > 0 {
		newFileName = newFileName[:strings.Index(newFileName, ".")]
	}
	newFileName += ".js"
	return ioutil.WriteFile(newFileName, []byte(v.ToObject(vm).Get("code").String()), 0775)
}

func readFile(filename string) []byte {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil
	}
	return bs
}

func defConfig() interface{} {
	return map[string]interface{}{
		"presets": []string{
			"es2015",
			"react",
			"stage-2",
			"babili",
		},
		"plugins": []string{
			"transform-react-jsx",
			"transform-es2015-block-scoping",
			//				"import",
		}}
}
