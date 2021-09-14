package tpl

import _ "embed"

var (
	//go:embed views/vue.tpl
	VueTemplate string
)

var (
	//go:embed java/controller.tpl
	ControllerTemplate string
	//go:embed java/service.tpl
	ServiceTemplate string
	//go:embed java/mapper.tpl
	MapperTemplate string
	//go:embed java/mapper_xml.tpl
	MapperXmlTemplate string
	//go:embed java/entity.tpl
	EntityTemplate string
)

var (
	//go:embed go/handler.tpl
	HandlerTemplate string
	//go:embed go/model.tpl
	ModelTemplate string
)
