package com.ninelock.api.${.PackageName}.service;

import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import com.ninelock.api.${.PackageName}.entity.${.TableName | formatBigCamel};
import com.ninelock.api.${.PackageName}.mapper.${.TableName | formatBigCamel}Mapper;
import org.springframework.stereotype.Service;

/** @author ninelock-ai */
@Service
public class ${.TableName | formatBigCamel}Service extends ServiceImpl<${.TableName | formatBigCamel}Mapper, ${.TableName | formatBigCamel}> {}
