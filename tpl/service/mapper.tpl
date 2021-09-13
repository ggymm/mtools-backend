package com.ninelock.api.${.PackageName}.mapper;

import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import com.ninelock.api.${.PackageName}.entity.${.TableName | formatBigCamel};
import org.apache.ibatis.annotations.Mapper;

/** @author ninelock-ai */
@Mapper
public interface ${.TableName | formatBigCamel}Mapper extends BaseMapper<${.TableName | formatBigCamel}> {}
