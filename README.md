# Description / 描述

Nushell Windows Launcher is a wrapper for nushell on Windows, and by default, the environment variables

`XDG_CONFIG_HOME=%UserProfile%\.config\`

`XDG_CACHE_HOME=%UserProfile%\.cache\`

`XDG_DATA_HOME=%UserProfile%\.local\share\` 

to replace Nushell's default `$nu.default-config-dir`, `$nu.cache-dir`, `$nu.data-dir` paths.

>>> Nushell Windows Launcher 是一个 Windows 上的 nushell 的包装器，默认为 Nushell 设置环境变量
>>> 
>>> `XDG_CONFIG_HOME=%UserProfile%\.config\`
>>> 
>>> `XDG_CACHE_HOME=%UserProfile%\.cache\`
>>> 
>>> `XDG_DATA_HOME=%UserProfile%\.local\share\`
>>> 
>>> 以取代 Nushell 默认的 `$nu.default-config-dir`、`$nu.cache-dir`、`$nu.data-dir` 路径。



Paths can also be customized via command-line arguments

>>> 也可以通过命令行参数自定义路径。

# Download, Install, Setup / 下载，安装，设置

To download the prebuild executable, please read [Releases](https://github.com/ErrorPower2001/nushell-windows-launcher/releases/).

>>> 为下载预构建二进制可执行文件，请访问 [Releases](https://github.com/ErrorPower2001/nushell-windows-launcher/releases/)。



After download, move the executable to any directory and add it to the Path environment variable.

>>> 之后，可以将可执行文件移动到任意位置并将路径加入 Path 环境变量



If you use Scoop Installer, you can run:

>>> 如果你是 Scoop Installer 用户，也可以运行：

```
scoop install https://raw.githubusercontent.com/ErrorPower2001/nushell-windows-launcher/master/nushell-windows-launcher.json
```

After installation, if you don't want the XDG environment variables to pollute your other programs, you can configure the following hooks for Nushell:

>>> 在安装过后，如果你不想XDG环境变量污染你的其他程序，你可以为Nushell配置如下钩子：

```nu
# $nu.user-autoload-dirs/clear-nux-injected.nu
$env.config.hooks.pre_prompt = $env.config.hooks.pre_prompt? | default [] | append {
	# UUID: 4a8e9a31-86f3-4659-b830-69ce2401f829

	if ("NUX_INJECTED_CONFIG" in $env) {
		hide-env XDG_CONFIG_HOME
		hide-env NUX_INJECTED_CONFIG
	}
	if ("NUX_INJECTED_CACHE" in $env) {
		hide-env XDG_CACHE_HOME
		hide-env NUX_INJECTED_CACHE
	}
	if ("NUX_INJECTED_DATA" in $env) {
		hide-env XDG_DATA_HOME
		hide-env NUX_INJECTED_DATA
	}

	$env.config.hooks.pre_prompt = (
		$env.config.hooks.pre_prompt 
		| where { |it| (view source $it | str contains "4a8e9a31-86f3-4659-b830-69ce2401f829") == false }
	)
}
```

# Build / 构建

```
go build .
```
