module.exports = {
	"skipr-api": {
		input: {
			target: `../.blink/open-api.yaml`,
		},
		output: {
			mode: "tags-split",
			target: "services",
			schemas: "types",
			override: {
				mutator: {
					path: "services/config.ts",
					name: "customInstance",
				},
			},
		},
	},
};
