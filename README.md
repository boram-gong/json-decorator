### **json-decorator**

json 结构转化器

#### **JSON转换规则说明**

##### 规则结构

```
type Rule struct {
	Key           string                 `json:"key"`
	Operation     string                 `json:"operation"`
	Content       interface{}            `json:"content"`
}
```

##### 规则参数说明

**key**：表示要操作的键，表达式：

| 路径表达式 | 操作对象 | 说明                    |
| ---------- | -------- | ----------------------- |
| demo.k1    | json对象 | 表示键demo下的键k1      |
| demo[0]    | 数组对象 | 表示键demo下的键0号索引 |

**operation**：对key（键）的操作，包含 **键操作** 和 **值操作**

**键操作**

| 表达式 | 说明                   |
| ------ | ---------------------- |
| rename | 更改键名称             |
| delete | 删除键（键和值都删除） |

**值操作**

| 表达式      | 说明                                 | 适用类型        |
| ----------- | ------------------------------------ | --------------- |
| replace     | 替换值                               | 任意类型        |
| append-head | 头部增加值                           | list、string    |
| append-tail | 尾部增加值                           | list、string    |
| move-head   | 移动值到目标头部                     | list            |
| move-tail   | 移动值到目标尾部                     | list            |
| move        | 移动值（移动后，移动前值的键将删除） | map(json键值对) |
| insert      | 插入值                               | map(json键值对) |
| delete      | 删除指定索引的值（key必须指定索引）  | list            |

**content**：操作的内容

- 当operation是**键操作**时：

  表达式和key（json对象）一样，当operation为delete时，此参数可以不填

- 当operation是**值操作**时：

  可填写任意类型，但需要注意类型与key不冲突，参数内容会依照operation的操作方式赋予给key

  支持关键字，包含

  | 关键字          | 说明                         |
  | :-------------- | ---------------------------- |
  | @               | 取值                         |
  | [...]           | 数组中每一个值               |
  | [-1]            | 数组中最后一个值             |
  | {"key":"value"} | {}内存在可放键值对，可放多个 |
  
   示例：
  
  | 参数内容        | 说明                                                     |
  | --------------- | -------------------------------------------------------- |
  | @demo           | 表示取源json键demo的值                                   |
  | @demo[...]      | 表示取源json键demo（此键的值必须是数组类型）的每一个值   |
  | @demo[-1]       | 表示取源json键demo（此键的值必须是数组类型）的最后一个值 |
  | {"key":"@demo"} | 插入键值对，其值取键demo的值                             |

##### 举几个规则示例

```
// 表示在键a下的b下的c的值插入一个键值对kv
{
    "key": "a.b.c",
    "operation": "insert",
    "content": {"k":"v"}
},
// 表示在json的第一层下值插入一个键值对kv，其v的值取键d的值
{
    "key": "",
    "operation": "insert",
    "content": {"k":"@d"}
},
// 表示把list值的最后一个元素添加到a下l的尾部
{
    "key": "a.l",
    "operation": "append-tail",
    "content": "@list[-1]"
},
// 表示键demo的值替换成键a的值
{
    "key": "demo",
    "operation": "replace",
    "content": "@a"
},
// 表示把list值的每一个元素添加到a下l的尾部
{
    "key": "a.l",
    "operation": "append",
    "content": "@list[...]"
},
// 表示把键k1名称改为k2
{
    "key": "k1",
    "operation": "rename",
    "content": "k2"
},
// 表示把键k1删除
{
    "key": "k1",
    "operation": "delete",
    "content": ""
},
// 表示把键k1值的0号索引删除
{
    "key": "k1[0]",
    "operation": "delete",
    "content": ""
}
```

