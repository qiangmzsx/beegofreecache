
package beegofreecache

import (
    "testing"
    "github.com/coocood/freecache"

    "encoding/gob"
    "bytes"
)

func wTestFreeCacheMax(t *testing.T)  {
    cacheSize := 512*1024
    cache := freecache.NewCache(cacheSize)
    key:="12345678900987654321"
    val:="第二天王霞又打电话到公司请了半天的假，先是给儿子做了喜欢吃的皮蛋瘦肉粥，然后便带着儿子去商场花了几百块买了几款还算高档的营养品，便一起乘公交车来到西平三中。母子两人来到学校，已经是上午快放学的时候了。王霞先是给班主任黎曼娜打了电话，知道了她这时候正好没课，便问清了她在学校的宿舍号码，表示将要带着谢仁一起上门拜访。黎曼娜对于王霞母子的突然拜访还是有些意外的，她昨天得知谢仁失踪，曾委婉的表示让王霞夫妇去警局认尸，结果却被证实那个车祸中丧生的死便不是西平三中的学生，更不用说是活生生的谢仁了，这个误会让她心里很有些过意不去，原本她还想打个电话向王霞道歉呢。现在见她要来拜访，便决定待会儿当面向她道歉。王霞带着谢仁来到了西平三中的教师宿舍楼，敲门后，黎曼娜很快就迎接了出来，因她以前曾做过家访，对王霞还是有点印象的，便很快就认出了她来。将两人让进房里后，黎曼娜很是客气的请他们入座，又给他们端茶倒水。第二天王霞又打电话到公司请了半天的假，先是给儿子做了喜欢吃的皮蛋瘦肉粥，然后便带着儿子去商场花了几百块买了几款还算高档的营养品，便一起乘公交车来到西平三中。母子两人来到学校，已经是上午快放学的时候了。王霞先是给班主任黎曼娜打了电话，知道了她这时候正好没课，便问清了她在学校的宿舍号码，表示将要带着谢仁一起上门拜访。黎曼娜对于王霞母子的突然拜访还是有些意外的，她昨天得知谢仁失踪，曾委婉的表示让王霞夫妇去警局认尸，结果却被证实那个车祸中丧生的死便不是西平三中的学生，更不用说是活生生的谢仁了，这个误会让她心里很有些过意不去，原本她还想打个电话向王霞道歉呢。现在见她要来拜访，便决定待会儿当面向她道歉。王霞带着谢仁来到了西平三中的教师宿舍楼，敲门后，黎曼娜很快就迎接了出来，因她以前曾做过家访，对王霞还是有点印象的，便很快就认出了她来。将两人让进房里后，黎曼娜很是客气的请他们入座，又给他们端茶倒水。第二天王霞又打电话到公司请了半天的假，先是给儿子做了喜欢吃的皮蛋瘦肉粥，然后便带着儿子去商场花了几百块买了几款还算高档的营养品，便一起乘公交车来到西平三中。母子两人来到学校，已经是上午快放学的时候了。王霞先是给班主任黎曼娜打了电话，知道了她这时候正好没课，便问清了她在学校的宿舍号码，表示将要带着谢仁一起上门拜访。黎曼娜对于王霞母子的突然拜访还是有些意外的，她昨天得知谢仁失踪，曾委婉的表示让王霞夫妇去警局认尸，结果却被证实那个车祸中丧生的死便不是西平三中的学生，更不用说是活生生的谢仁了，这个误会让她心里很有些过意不去，原本她还想打个电话向王霞道歉呢。现在见她要来拜访，便决定待会儿当面向她道歉。王霞带着谢仁来到了西平三中的教师宿舍楼，敲门后，黎曼娜很快就迎接了出来，因她以前曾做过家访，对王霞还是有点印象的，便很快就认出了她来。将两人让进房里后，黎曼娜很是客气的请他们入座，又给他们端茶倒水。"
    err:=cache.Set([]byte(key),[]byte(val),100)
    if err!=nil {
        println(err.Error())
    }

}

func TestGobEncode(t *testing.T)  {
    var ii int = 123
    var ui uint= uint(23)
    bii,_:=GobEncode(ii)
    bui,_:=GobEncode(ui)
    println(bii,bui)
    var dii interface{}
    var dui uint
    e:=GobDecode(bii,&dii)
    //println(e.Error())
    e=GobDecode(bui,&dui)
    println(dii,dui,e)
    switch dii.(type) {
    case int:
        println("int")
    case int64:
        println("int64")
    case int32:
        println("int32")
    case uint:
        println("uint")
    case uint32:
        println("uint32")
    case uint64:
        println("uint64")
    default:
        println("default")
    }
    println(dii)

}


// Gob加密
func GobEncode(data interface{}) ([]byte, error) {

    buf := bytes.NewBuffer(nil)
    enc := gob.NewEncoder(buf)
    err := enc.Encode(data)
    if err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}

// Gob解密
func GobDecode(data []byte, to interface{}) error {

    buf := bytes.NewBuffer(data)
    dec := gob.NewDecoder(buf)
    return dec.Decode(to)
}