# vim:shiftwidth=8:noexpandtab:

$env.config.hooks.pre_prompt = $env.config.hooks.pre_prompt? | default [] | append {
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
}
