package com.ninelock.api.${.PackageName}.controller;

import com.baomidou.mybatisplus.core.conditions.query.QueryWrapper;
import com.baomidou.mybatisplus.core.metadata.IPage;
import com.baomidou.mybatisplus.extension.plugins.pagination.Page;
import com.ninelock.core.response.ResultMsg;
import com.ninelock.core.toolkit.QueryGenerator;
import com.ninelock.api.${.PackageName}.entity.${.TableName | formatBigCamel};
import com.ninelock.api.${.PackageName}.service.${.TableName | formatBigCamel}Service;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.*;

import java.util.List;

import javax.servlet.http.HttpServletRequest;

/** @author ninelock-ai */
@Slf4j
@RestController
@RequestMapping("/${.TableName}/")
public class ${.TableName | formatBigCamel}Controller {
  final ${.TableName | formatBigCamel}Service ${.TableName | formatSmallCamel}Service;

  public ${.TableName | formatBigCamel}Controller(${.TableName | formatBigCamel}Service ${.TableName | formatSmallCamel}Service) {
    this.${.TableName | formatSmallCamel}Service = ${.TableName | formatSmallCamel}Service;
  }

  ${- if eq .ViewType "page"}
  @RequestMapping("get_page")
  public String getPage(
      ${.TableName | formatBigCamel} ${.TableName | formatSmallCamel},
      @RequestParam(name = "page", defaultValue = "1") Integer page,
      @RequestParam(name = "size", defaultValue = "10") Integer size,
      HttpServletRequest request) {
    QueryWrapper<${.TableName | formatBigCamel}> queryWrapper = QueryGenerator.initQueryWrapper(${.TableName | formatSmallCamel}, request.getParameterMap());
    Page<${.TableName | formatBigCamel}> pageable = new Page<>(page, size);
    IPage<${.TableName | formatBigCamel}> pageList = ${.TableName | formatSmallCamel}Service.page(pageable, queryWrapper);
    return ResultMsg.success(pageList);
  }
  ${- else if eq .ViewType "list"}

  @RequestMapping("get_list")
  public String getList(${.TableName | formatBigCamel} ${.TableName | formatSmallCamel}, HttpServletRequest request) {
    QueryWrapper<${.TableName | formatBigCamel}> queryWrapper = QueryGenerator.initQueryWrapper(${.TableName | formatSmallCamel}, request.getParameterMap());
    List<${.TableName | formatBigCamel}> list = ${.TableName | formatSmallCamel}Service.list(queryWrapper);
    return ResultMsg.success(list);
  }
  ${- end}

  @RequestMapping(value = "save", method = RequestMethod.POST)
  public String save(@RequestBody ${.TableName | formatBigCamel} ${.TableName | formatSmallCamel}) {
    return ResultMsg.success(${.TableName | formatSmallCamel}Service.save(${.TableName | formatSmallCamel}));
  }

  @RequestMapping(value = "update", method = RequestMethod.POST)
  public String update(@RequestBody ${.TableName | formatBigCamel} ${.TableName | formatSmallCamel}) {
    return ResultMsg.success(${.TableName | formatSmallCamel}Service.updateById(${.TableName | formatSmallCamel}));
  }

  @RequestMapping(value = "get")
  public String get(@RequestParam(name = "id") Long id) {
    return ResultMsg.success(${.TableName | formatSmallCamel}Service.getById(id));
  }

  @RequestMapping(value = "delete", method = RequestMethod.POST)
  public String delete(@RequestBody Long id) {
    return ResultMsg.success(${.TableName | formatSmallCamel}Service.removeById(id));
  }
}
