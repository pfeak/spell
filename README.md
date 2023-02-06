# Spell

Go implementation based on [Spell: Streaming Parsing of System Event Logs](https://ieeexplore.ieee.org/document/7837916).

`Spell` is an online real-time dynamic parsing method based on `LCS (Longest Common Subsequence)`, which is used to dynamically extract log data templates.

Documents：[English](https://github.com/pfeak/Spell/blob/master/README.md)
/[中文](https://github.com/pfeak/Spell/blob/master/docs/zh.md)

## Advantages and disadvantages

Combining the research and practical use in the paper, even if `Spell` is compared with offline methods, `Spell` is still very powerful in terms of efficiency and effectiveness.

However, the `Spell` algorithm also has disadvantages: when using the `Spell` algorithm in an actual scenario, if there are many types of logs, the performance of the `Spell` algorithm will decrease significantly as the number of extracted templates increases.

## Parameter

| Name       | Meaning                             | Description                                                                                 |
|------------|-------------------------------------|---------------------------------------------------------------------------------------------|
| splitRule  | log splitting rule                  | Regular rule string, default `[\s:=,]+`                                                     |
| label      | template placeholder wildcard       | default `<*>`, for example `I have <*> pen`                                                 |
| similarity | similarity between log and template | The lower the similarity, the more templates will be extracted, the value range `[0.01, 1]` |

## Usage

Prepare the log as follows:

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

Running `go run main.go` gives the following output:

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

## Reference

* [Spell: Streaming Parsing of System Event Logs](https://ieeexplore.ieee.org/document/7837916)

* [https://users.cs.utah.edu/~lifeifei/papers/spell.pdf](https://users.cs.utah.edu/~lifeifei/papers/spell.pdf)

* [Spell: Streaming Parsing of System Event Logs (paper reading)](https://saucer-man.com/information_security/388.html)

* [https://github.com/nbigaouette/spell-rs](https://github.com/nbigaouette/spell-rs)
