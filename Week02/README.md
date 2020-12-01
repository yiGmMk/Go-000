学习笔记

1.错误处理方式3种，最佳做法：only handle error once


2.log记录，错误记录避免到处打日志，错误尽量在发生的地方wrap加上stack信息抛给上层记录日志

3.