flamingo: {
	cmd: name:                   "readinglist"
	systemendpoint: serviceAddr: ":13210"
	opencensus: jaeger: enable:  true
}

core: {
	locale: date: location: "Europe/Berlin"
	gotemplate: engine: {
		templates: basepath: "templates"
		layout: dir:         "layouts"
	}
}
