package tpl

import _ "embed"

var (
	//go:embed views/vue.tpl
	VueTemplate string
)

var (
	//go:embed service/controller.tpl
	ControllerTemplate string
	//go:embed service/service.tpl
	ServiceTemplate string
	//go:embed service/mapper.tpl
	MapperTemplate string
	//go:embed service/mapper_xml.tpl
	MapperXmlTemplate string
	//go:embed service/entity.tpl
	EntityTemplate string
)
