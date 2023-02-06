# Spell

基于 [Spell: Streaming Parsing of System Event Logs](https://ieeexplore.ieee.org/document/7837916) 的 go 实现。

`Spell` 是一种基于 `LCS(最长公共子序列)` 的在线实时动态解析日志的方法，用以动态提取日志数据模板。

文档：[English]((https://github.com/pfeak/Spell/blob/master/README.md))
/[中文](https://github.com/pfeak/Spell/blob/master/docs/zh.md)

## 优缺点

结合论文中的研究与实际使用，即便 `Spell` 与离线方法相比，`Spell` 在效率和有效性方面仍然非常厉害。

但 `Spell` 算法也有缺点：当在实际场景使用 `Spell` 算法，如果日志种类众多，随着提取模板数量的增加，`Spell`算法的性能会显著下降。

## 参数

| 参数名        | 含义       | 说明                            |
|------------|----------|-------------------------------|
| splitRule  | 日志分割规则   | 表现为正则规则字符串，默认 `[\s:=,]+`      |
| label      | 模板占位通配符  | 默认 `<*>`，例如 `I have <*> pen`  |
| similarity | 日志与模板相似度 | 相似度低，提取模板会越多，取值范围 `[0.01, 1]` |

## 使用

准备日志如下：

```
this is a pen
this is the pen
this is a pen

i am green
i am blue
i am yellow and red
i am grey and black

{"host":"192.168.1.23", "message":"logId=0000000013"}
{"host":"192.168.1.23", "message":"logId=0000000013", "id":"123"}
{"host":"192.168.1.25", "message":"logId=0000000015"}
{"host":"192.168.1.24", "message":"logId=0000000014", "id":"456"}
{"host":"192.168.1.25", "message":"logId=0000000013", "id":"123"}

{"host":"192.168.1.24", "message":"devName=FC020000067245 devId=FC020000067245 logId=0000000013"}
{"host":"192.168.1.23", "message":"devName=FC020000067245 devId=FC020000067242 logId=0000000014"}
{"host":"192.168.1.24", "message":"devName=FC020000067245 devId=FC020000067245 logId=0000000015"}
{"host":"192.168.1.26", "message":"devName=FC020000067245 devId=FC020000067245 logId=000000007466"}
{"host":"192.168.1.26", "message":"devName=FC020000067245 devId=FC020000067242 logId=0000000016", "time":"1234567890"}
{"host":"192.168.1.26", "message":"devName=FC020000067245 devId=FC020000067245 logId=0000000016", "time":"1234567890"}
```

运行 `go run main.go` 可得输出如下：

```shell
Template: [<*> is <*> pen]
Position: [0 2]

Template: [<*> am <*>]
Position: [0 2]

Template: [i am grey and black]
Position: []

Template: [<*> "logId <*>]
Position: [0 2]

Template: [<*> "message" <*>]
Position: [0 2]

Template: [<*> FC020000067245 devId <*> logId <*>]
Position: [0 3 5]
```

## 参考

* [Spell: Streaming Parsing of System Event Logs](https://ieeexplore.ieee.org/document/7837916)

* [https://users.cs.utah.edu/~lifeifei/papers/spell.pdf](https://users.cs.utah.edu/~lifeifei/papers/spell.pdf)

* [Spell: Streaming Parsing of System Event Logs (paper reading)](https://saucer-man.com/information_security/388.html)

* [https://github.com/nbigaouette/spell-rs](https://github.com/nbigaouette/spell-rs)
