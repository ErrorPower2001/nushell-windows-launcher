# vim:shiftwidth=8:noexpandtab:

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
