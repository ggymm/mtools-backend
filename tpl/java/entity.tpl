package ${.BasePackageName}.${.PackageName}.entity;

/**
 * @author ${.Author}
 * @version 创建时间: ${.Now}
 */
@lombok.Data
@com.baomidou.mybatisplus.annotation.TableName("${.OriginTableName}")${if .UseParent}
@lombok.EqualsAndHashCode(callSuper = false)${- end}
@lombok.experimental.Accessors(chain = true)
public class ${.TableName} ${- if .UseParent} extends ${.ParentPackage} ${end -}{
    ${- range .Fields}
        ${- if $.UseParent -}
            ${- if not .Exclude}
                ${- if eq .IsKey "TRUE"}
                    ${- if eq .IsAuto "TRUE"}
    /**
     * ${.ColumnComment}
     */
    @com.baomidou.mybatisplus.annotation.TableId(value = "${.ColumnName}", type = com.baomidou.mybatisplus.annotation.IdType.AUTO)
                    ${- else}
    /**
     * ${.ColumnComment}
     */
    @com.baomidou.mybatisplus.annotation.TableId(value = "${.ColumnName}")
                    ${- end}
                ${- else}
    /**
     * ${.ColumnComment}
     */
                ${- end}
                ${- if .FormatDate}
    @com.fasterxml.jackson.annotation.JsonFormat(timezone = "GMT+8", pattern = "yyyy-MM-dd HH:mm:ss")
    @org.springframework.format.annotation.DateTimeFormat(pattern = "yyyy-MM-dd HH:mm:ss")
                ${- end}
    @com.baomidou.mybatisplus.annotation.TableField(value = "${.ColumnName}")
    private ${.DataType | FormatJavaDataType} ${- if $.UseOriginColumn} ${.ColumnName} ${else} ${.ColumnName | SmallCamel}${- end -};
            ${- end}
        ${- else}
            ${- if eq .IsKey "TRUE"}
                ${- if eq .IsAuto "TRUE"}
    /**
     * ${.ColumnComment}
     */
    @com.baomidou.mybatisplus.annotation.TableId(value = "${.ColumnName}", type = com.baomidou.mybatisplus.annotation.IdType.AUTO)
                ${- else}
    /**
     * ${.ColumnComment}
     */
    @com.baomidou.mybatisplus.annotation.TableId(value = "${.ColumnName}")
                ${- end}
            ${- else}
    /**
     * ${.ColumnComment}
     */
            ${- end}
            ${- if .FormatDate}
    @com.fasterxml.jackson.annotation.JsonFormat(timezone = "GMT+8", pattern = "yyyy-MM-dd HH:mm:ss")
    @org.springframework.format.annotation.DateTimeFormat(pattern = "yyyy-MM-dd HH:mm:ss")
            ${- end}
            ${- if .AutoFill}
    @com.baomidou.mybatisplus.annotation.TableField(value = "${.ColumnName}", fill = com.baomidou.mybatisplus.annotation.FieldFill.${.AutoFillType})
            ${- else}
    @com.baomidou.mybatisplus.annotation.TableField(value = "${.ColumnName}")
            ${- end}
    private ${.DataType | FormatJavaDataType} ${- if $.UseOriginColumn} ${.ColumnName} ${else} ${.ColumnName | SmallCamel}${- end -};
        ${- end}
    ${- end}
}
