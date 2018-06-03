/*
Package snowflake generate distributed unique id
	1 bit：不使用，可以是 1 或 0
	41 bit：记录时间戳 (当前时间戳减去用户设置的初始时间，毫秒表示)，可记录最多 69 年的时间戳数据
	10 bit：用来记录分布式节点 ID，一般每台机器一个唯一 ID，也可以多进程每个进程一个唯一 ID，最大可部署 1024 个节点
	12 bit：序列号，用来记录不同 ID 同一毫秒时的序列号，最多可生成 4096 个序列号
*/
package snowflake
