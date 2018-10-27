package main

import (
	"fmt"

	"io/ioutil"
	"os"

	"github.com/dop251/goja"
)

func main() {
	if len(os.Args) > 1 {

		bs, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}

		opts := map[string]interface{}{
			"presets": []string{
				"es2015",
				"react",
				"stage-2",
				"babili",
			},
			"plugins": []string{
			//				"transform-react-jsx",
			//				"transform-es2015-block-scoping",
			//				"import",

			}}

		babelsrc, err := ioutil.ReadFile("./js/babel.js")
		if err != nil {
			panic(err)
		}
		polyfilljs, err := ioutil.ReadFile("./js/polyfill.js")
		if err != nil {
			panic(err)
		}
		babilijs, err := ioutil.ReadFile("./js/babili.js")
		if err != nil {
			panic(err)
		}
		vm := goja.New()
		_, err = vm.RunScript("babel.js", string(babelsrc))
		if err != nil {
			panic(fmt.Errorf("babel.js err: %s", err))
		}

		_, err = vm.RunScript("polyfill.js", string(polyfilljs))
		if err != nil {
			panic(fmt.Errorf("运行polyfill.js错误：%s", err))
		}

		_, err = vm.RunScript("运行babili.js错误：", string(babilijs))
		if err != nil {
			panic(fmt.Errorf("运行babili.js错误：%s", err))
		}

		var transform goja.Callable
		babel := vm.Get("Babel")
		if err := vm.ExportTo(babel.ToObject(vm).Get("transform"), &transform); err != nil {
			panic(fmt.Errorf("未找到”transform“函数，错误: %s", err))
		}

		v, err := transform(babel, vm.ToValue(string(bs)), vm.ToValue(opts))
		if err != nil {
			panic(err)
		}

		ioutil.WriteFile(os.Args[1]+".js", []byte(v.ToObject(vm).Get("code").String()), 0775)
	}
}
