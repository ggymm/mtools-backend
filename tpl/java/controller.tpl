package ${.BasePackageName}.${.PackageName}.controller;

import ${.BasePackageName}.${.PackageName}.entity.${.TableName};
import ${.BasePackageName}.${.PackageName}.service.${.TableName}Service;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;

/**
 * @author ${.Author}
 * @version 创建时间: ${.Now}
 */
@RestController
@RequestMapping("${.OriginTableName | SmallCamel}")
public class ${.TableName}Controller {

    final ${.TableName}Service ${.OriginTableName | SmallCamel}Service;

    public ${.TableName}Controller(${.TableName}Service ${.OriginTableName | SmallCamel}Service) {
        this.${.OriginTableName | SmallCamel}Service = ${.OriginTableName | SmallCamel}Service;
    }

    @GetMapping(value = "getPage", produces = "application/json;charset=UTF-8")
    public String getPage(HttpServletRequest request, ${.TableName} ${.OriginTableName | SmallCamel},
            @RequestParam(name = "page", defaultValue = "1") Integer page,
            @RequestParam(name = "size", defaultValue = "10") Integer size) {
        return ${.OriginTableName | SmallCamel}Service.getPage(${.OriginTableName | SmallCamel}, request.getParameterMap(), page, size);
    }

    @PostMapping(value = "create", produces = "application/json;charset=UTF-8")
    public String create(@RequestBody ${.TableName} ${.OriginTableName | SmallCamel}) {
        return ${.OriginTableName | SmallCamel}Service.create(${.OriginTableName | SmallCamel});
    }
${if .HasID}
    @GetMapping(value = "get", produces = "application/json;charset=UTF-8")
    public String get(@RequestParam(name = "id") ${.IDType} id) {
        return ${.OriginTableName | SmallCamel}Service.get(id);
    }

    @PostMapping(value = "update", produces = "application/json;charset=UTF-8")
    public String update(@RequestBody ${.TableName} ${.OriginTableName | SmallCamel}) {
        return ${.OriginTableName | SmallCamel}Service.update(${.OriginTableName | SmallCamel});
    }

    @PostMapping(value = "delete", produces = "application/json;charset=UTF-8")
    public String delete(@RequestBody ${.IDType} id) {
        return ${.OriginTableName | SmallCamel}Service.delete(id);
    }
${end -}
}
