package ${.BasePackageName}.${.PackageName}.service;

import com.baomidou.mybatisplus.core.conditions.query.QueryWrapper;
import com.baomidou.mybatisplus.extension.plugins.pagination.Page;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import ${.BasePackageName}.${.PackageName}.entity.${.TableName};
import ${.BasePackageName}.${.PackageName}.mapper.${.TableName}Mapper;
import com.ninelock.core.response.ResultMsg;
import com.ninelock.core.toolkit.QueryGenerator;
import org.springframework.stereotype.Service;

import java.util.Map;

/**
 * @author ${.Author}
 * @version 创建时间: ${.Now}
 */
@Service
public class ${.TableName}Service extends ServiceImpl<${.TableName}Mapper, ${.TableName}> {

    /**
     * 查询分页列表数据
     * @param ${.OriginTableName | SmallCamel} 实体类对象，用于传递字段查询条件
     * @param parameterMap 其他请求参数，用于传递其他查询条件（比如LIKE，ORDER等）
     * @param page 页号
     * @param size 页大小
     * @return 同一返回值封装
     */
    public String getPage(${.TableName} ${.OriginTableName | SmallCamel}, Map<String, String[]> parameterMap, Integer page, Integer size) {
        QueryWrapper<${.TableName}> queryWrapper = QueryGenerator.initQueryWrapper(${.OriginTableName | SmallCamel}, parameterMap);
        Page<${.TableName}> pageable = new Page<>(page, size);
        return ResultMsg.success(this.page(pageable, queryWrapper));
    }

    /**
     * 保存数据
     * @param ${.OriginTableName | SmallCamel} 实体类对象
     * @return 同一返回值封装
     */
    public String create(${.TableName} ${.OriginTableName | SmallCamel}) {
        boolean save = this.save(${.OriginTableName | SmallCamel});
        if (save) {
            return ResultMsg.success();
        } else {
            return ResultMsg.fail("保存失败");
        }
    }
${if .HasID}
    /**
     * 根据ID获取数据
     * @param id ID值
     * @return 同一返回值封装
     */
    public String get(${.IDType} id) {
        ${.TableName} ${.OriginTableName | SmallCamel} = this.getById(id);
        return ResultMsg.success(${.OriginTableName | SmallCamel});
    }

    /**
     * 根据ID更新数据
     * @param {.OriginTableName | SmallCamel} 实体类对象
     * @return 同一返回值封装
     */
    public String update(${.TableName} ${.OriginTableName | SmallCamel}) {
        boolean update = this.updateById({.OriginTableName | SmallCamel});
        if (update) {
            return ResultMsg.success();
        } else {
            return ResultMsg.fail("更新失败");
        }
    }

    /**
     * 根据ID删除数据
     * @param id ID值
     * @return 同一返回值封装
     */
    public String delete(${.IDType} id) {
        boolean remove = this.removeById(id);
        if (remove) {
            return ResultMsg.success();
        } else {
            return ResultMsg.fail("删除失败");
        }
    }
${end -}
}
