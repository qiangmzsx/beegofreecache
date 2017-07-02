#beegofreecache
## 什么是beegofreecache
```go
fatal error: concurrent map iteration and map write
```
更重要的是memory使用的是同步锁来确保数据一致性，这必然后导致性能有所损耗。
如果有熟悉Java的同学应该知道Map。HashMap中未进行同步考虑，而Hashtable则使用了synchronized，带来的直接影响就是可选择，我们可以在单线程时使用HashMap提高效率，而多线程时用Hashtable来保证安全。
beego内置的memory实现就相当于Hashtable。  
所以我们需要一个更好的memory的实现，经过测试决定使用freecache，大家可以去[freecache主页](https://github.com/coocood/freecache)查看详细信息。
freecache就相当于Java中的ConcurrentHashMap，性能也有很大的提高，官网上它的性能测试报告，在此不赘述了。  

但是freecache也会有一些缺点：
1. 当需要缓存的数据占用超过提前分配缓存的 1/1024 则不能缓存
2. 当分配内存过多则会导致内存占用高
3. key的长度需要小于65535字节
3. value只能是[]byte类型，使用很不友好
## 使用方法
为此我对freecache进行了简单的封装，让freecache可以集成搭配beego框架中，也可以不在beego框架中独立使用。  
### 安装
```
go get  git.oschina.net/qiangmzsx/beegofreecache
或
go get github.com/qiangmzsx/beegofreecache

```
导入包
```
import git.oschina.net/qiangmzsx/beegofreecache
```
### 集成到beego框架:
与beego使用其他的cache一样。
```go
var bm, err = cache.NewCache("beegofreecache", `{"size":64}`)
bm.Put("beegofreecache", 1, 10*time.Second)
bm.Get("beegofreecache")
bm.IsExist("beegofreecache")
bm.Delete("beegofreecache")
```
**注意：Get()函数返回的是[]byte这个需要业务使用beegofreecache内置GobDecode()转码**
```go
var bm, err = cache.NewCache("beegofreecache", `{"size":64}`)
func TestBeeFreeCacheGP(t *testing.T)  {

    if err!=nil {
        println(err.Error())
    }
    key:="key"
    bm.Put(key,1200,10*time.Second)
    val:=bm.Get(key)
    var ii int
    GobDecode(val.([]byte),&ii)
    println(val,ii)

}
```
### 独立使用
不依赖beego框架。  
```go
bm:=beegofreecache.NewFree(512)//size为M
bm.Put("beegofreecache", 1, 10*time.Second)
var ii int
bm.GetValue("beegofreecache",ii)
bm.IsExist("beegofreecache")
bm.Delete("beegofreecache")
```  

### 转化失败
在使用中如果出现了
```go
gob: type not registered for interface
```
别着急，可以在init()中加入：
```go
gob.Register(map[string]interface{}{})
gob.Register(map[string][]int{})
gob.Register(map[string][]int64{})
```
其中Register函数中的参数就是需要转化的类型了。
欢迎大家提供修改意见。