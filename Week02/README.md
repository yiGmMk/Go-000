学习笔记

0. panic & error & exception
- go 的panic当程序无法运行下去才应该触发，不能假设调用者会处理你发出的panic。例如索引越界、栈溢出等
-  error
-  exception 

1. 错误处理方式3种，最佳做法：only handle error once
-  sentinel error 预定义的特定错误。不够灵活
-  Error types实现了error接口的自定义类型,能够提供更多上下文信息，同时耦合性也更强
```
type MyError struct{
    Msg string
    Line int
}
func (e *MyError)Error()string{
    return e.Msg
}

err:=&MyError{}
switch err:=err.(type){
    case nil:
    case:*MyError:
    default:
}
```
-  Opaque Errors减少代码与调用者的耦合，不透明错误处理
```
x,err:=xxx_package.Foo()
if err!=nil{
    return err
}

//断言错误实现了特定的行为，而不是类型或值
type vaild interface{
    Vaild()bool
}
//可以在不导入定义错误的包或者实际上不了解 err 的底层类型的情况下实现
func IsVaild(err error)bool{
    val,ok:=err.(vaild)
    return ok && val.Vaild()
}
```
   

2. 错误处理，log记录：错误记录避免到处打日志，错误尽量在发生的地方wrap加上stack信息抛给上层记录日志
-  only handle error once
-  wrap errors，加上堆栈信息和其他上下文，github.com/pkg/errors
  ```
  errors.Wrap(err,"failed to do somethiing")

  errors.WithMessage(err,"some message for error")

  fmt.Printf("%T %v",errors.Cause(err),errors.Cause(err))
  fmt.Pringtln("stacktrace:%+v\n",err)
  ```
-    应用代码中使用errors.New or errors.Errorf
-   