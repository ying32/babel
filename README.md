# babel
jsx编译和压缩，无需Node.js，仅单exe，简单方便。 

----

做这个东东的原因是以前在学习了阿里的antDesign感觉nodejs跟npm用得很不爽，也不喜欢用这两东西，后来就放弃了，直到昨天重新捡起来，看了下antDesign和vue的网站又想起来折腾这事，看到vue有不需要nodejs的方法，就搜索了下antDesign有没有同样的，还真找到了，不过很少而且没几个说明，缺失的太多了，就自己研究起来了，折腾各种库，新的es语法（原antDesign网站上的例子直接来用不了，还折腾了一番语法方面的事情），最后还好都成功了。   

但又遇到了新的问题，有些浏览器又不支持es新语法。又寻找方法发现用babel可以转换，但又涉及到了nodejs，一时间很无奈，太不喜欢nodejs这东西了，最后在github上检索了下有没有go方面的，倒是发现有了，用了几个要不就是编译不过，要不就是想添加新的功能导致无法运行，最后只好都放弃了。从[goja-babel](https://github.com/jvatic/goja-babel)这个项目中了解到babel就是通过一个函数转的而已，于是干脆自己写了（放弃goja-babel是因为我本想给它添加个js压缩功能的插件babili，无奈运行中各莫名奇妙的错）。

# 获取方法

```
go get github.com/ying32/babel
cd xxxxx
go build

```

# 使用方法
```bat
rem 单一文件编译
babel.exe test.jsx

rem 编译目录下所有jsx文件, 后面接目录名就行了
babel.exe ./  

```

# 配置文件 babel.json
可有可无的，如果没有这个会默认使用可执行文件中内置的配置，有则使用这个。
```json
{
	"presets": [
		"es2015",
		"react",
		"stage-2",
		"babili"
	],
	"plugins": [
		"transform-react-jsx",
		"transform-es2015-block-scoping"
	]
}
```

# 第三方库：
```
github.com/dop251/goja

```

# 说明

仓库中js目录下的代码来自 https://www.bootcdn.cn/ 网站。这个里面的js只是在开发调试的时候用了下，用目录下的genres可将js生成.go源码直接编译进可执行文件中，这样就不需要带这些东东了。