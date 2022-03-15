### JSON转换服务（v0.0.1版本，依赖数据库postgreSQL）


**json转换地址  GET /lgi/responseAdapter/json**

##### 请求参数

| 参数 | 含义                                       |
| ---- | ------------------------------------------ |
| name | 要使用的规则名称                           |
| data | 要处理的json（原始json格式，非字符串格式） |

##### 返回参数

| 参数 | 含义                                           |
| ---- | ---------------------------------------------- |
| code | 响应好（200表示成功，40x表示用户请求存在问题） |
| msg  | 响应描述                                       |
| data | json变化结果数据                               |

请求示例

```json
{
    "name": "demo2",
    "data": {
        "head":"head",
        "demo": {
            "key":"value"
        },
        "list": [
                {"d1":"d1"},
                {"d2":"d2"}
        ]
    }
}
```

返回示例

```json
{
    "code": 200,
    "msg": "成功",
    "data": {
        "demo": {
            "key": "value",
            "head": "value"
        },
        "list1": [
            {
                "d1": "d1"
            },
            {
                "d2": "d2"
            }
        ]
    }
}
```

**获取规则  GET /lgi/responseAdapter/rule**

请求示例

```
获取所有规则    GET 127.0.0.1:29989/lgi/responseAdapter/rule 
获取指定id规则  GET 127.0.0.1:29989/lgi/responseAdapter/rule?id=1 
```

响应示例

```json
{
    "code": 200,
    "msg": "成功",
    "data": [
        {
            "id": 11,
            "rule_name": "demo2",
            "rules": [
                {
                    "key": "demo.key",
                    "operation": "rename",
                    "content": "demo.newKey"
                }
            ],
            "stat": 1,
            "start_time": "",
            "end_time": ""
        },
        {
            "id": 12,
            "rule_name": "demo3",
            "rules": [
                {
                    "key": "demo.key",
                    "operation": "rename",
                    "content": "demo.newKey2"
                }
            ],
            "stat": 1,
            "start_time": "",
            "end_time": ""
        }
    ]
}
```

```json
{
    "code": 200,
    "msg": "成功",
    "data": {
        "id": 11,
        "rule_name": "demo2",
        "rules": [
            {
                "key": "demo.key",
                "operation": "rename",
                "content": "demo.newKey"
            }
        ],
        "stat": 1,
        "start_time": "",
        "end_time": ""
    }
}
```

**新增规则  POST /lgi/responseAdapter/rule**

请求体（json）参数

| 参数       | 说明                                                         |
| ---------- | ------------------------------------------------------------ |
| id         | post新增规则，id为0，必有参数                                |
| rule_name  | 规则名称，必有参数，唯一                                     |
| rules      | 详细规则数组                                                 |
| stat       | 状态，有效为1，无效为其他数字，必有字段                      |
| start_time | 规则有效起始时间，格式有两种: "2006-01-02 15:04:05" 和 "15:04:05"，，举例："2006-01-02 15:04:05"~"2006-01-03 15:04:05" 表示这段时间有效，"15:00:00"~"20:00:00" 表示每天这个时间段有效 |
| end_time   | 规则有效结束时间                                             |

详细规则参数（可参考规则转换说明）

| 参数      | 含义              |
| --------- | ----------------- |
| key       | 表示要操作的键    |
| operation | 对key（键）的操作 |
| content   | 操作的内容        |

请求示例

```json
        {
            "id":0,
            "rule_name": "demo3",
            "rules": [
                {
                    "key": "demo.key",
                    "operation": "rename",
                    "content": "demo.newKey"
                }
            ],
            "stat": 1,
            "start_time": "",
            "end_time": ""
        }
```

响应示例（响应的data中是当前所有的规则）

```json
{
    "code": 200,
    "msg": "成功",
    "data": [
        {
            "id": 11,
            "rule_name": "demo2",
            "rules": [
                {
                    "key": "demo.key",
                    "operation": "rename",
                    "content": "demo.newKey"
                }
            ],
            "stat": 1,
            "start_time": "",
            "end_time": ""
        },
        {
            "id": 13,
            "rule_name": "demo3",
            "rules": [
                {
                    "key": "demo.key",
                    "operation": "rename",
                    "content": "demo.newKey"
                }
            ],
            "stat": 1,
            "start_time": "",
            "end_time": ""
        }
    ]
}
```

**修改规则  PUT  /lgi/responseAdapter/rule**

修改规则和新增规则几乎一样，就是id需要明确

**删除规则  DELETE  /lgi/responseAdapter/rule**

请求示例

```
需指定规则id  DELETE 127.0.0.1:29989/lgi/responseAdapter/rule?id=1 
```

**重新加载规则  GET /lgi/responseAdapter/re**



### **JSON转换规则说明**

#### 1. **key**：表示要操作的键，表达式：

| 路径表达式 | 操作对象 | 说明                    |
| ---------- | -------- | ----------------------- |
| demo.k1    | json对象 | 表示键demo下的键k1      |
| demo[0]    | 数组对象 | 表示键demo下的键0号索引 |

**注**：key的表达式中不支持也不允许存在`[...]`和`@`这两种关键字

#### 2. **operation**：对key（键）的操作，包含 **键操作** 和 **值操作**

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

#### 3. **content**：操作的内容

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



#### 4. 部分规则示例

表示在键a下的b下的c的值插入一个键值对kv

```json
{
    "key": "a.b.c",
    "operation": "insert",
    "content": {"k":"v"}
}
```

表示在json的第一层下值插入一个键值对kv，其v的值取键d的值

```json
{
    "key": "",
    "operation": "insert",
    "content": {"k":"@d"}
}
```

表示把list值的最后一个元素添加到a下l的尾部

```json
{
    "key": "a.l",
    "operation": "append-tail",
    "content": "@list[-1]"
}
```

表示键demo的值替换成键a的值

```json
{
    "key": "demo",
    "operation": "replace",
    "content": "@a"
}
```

表示把list值的每一个元素添加到a下l的尾部

```json
{
    "key": "a.l",
    "operation": "append",
    "content": "@list[...]"
}
```

表示把键k1名称改为k2

```json
{
    "key": "k1",
    "operation": "rename",
    "content": "k2"
}
```

表示把键k1删除

```json
{
    "key": "k1",
    "operation": "delete",
    "content": ""
}
```

表示把键k1值的0号索引删除

```json
{
    "key": "k1[0]",
    "operation": "delete",
    "content": ""
}
```
