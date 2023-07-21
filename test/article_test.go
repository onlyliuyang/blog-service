package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/blog-service/initialization"
	"github.com/blog-service/pkg/app"
	"github.com/blog-service/pkg/util"
	"testing"
)

func init() {
	err := initialization.SetupSetting("/Users/admin/go/src/github.com/go-programing-tour-book/blog-service/configs/app")
	if err != nil {
		//t.Fatalf("init.setupSetting err: %v", err)
		fmt.Printf("init.setupSetting err: %v\n", err)
		return
	}

	err = initialization.SetupRedis()
	if err != nil {
		//t.Fatalf("init.setupRedis err: %v", err)
		fmt.Printf("init.setupRedis err: %v\n", err)
		return
	}
}

func TestArticleCreate(t *testing.T) {
	//var params map[string]interface{}
	params := []map[string]interface{}{
		{
			"url":       "http://www.baidu.com",
			"show":      1,
			"sort":      1,
			"title":     "晨会小故事1",
			"content":   "1960年，哈佛大学的罗森塔尔博士曾在加州一所学校做过一个著名的实验。\n\n　　新学年开始时，罗森塔尔博士让校长把三位教师叫进办公室，对他们说：“根据你们过去的教学表现，你们是本校最优秀的老师。因此，我们特意挑选了100名全校最聪明的学生组成三个班让你们教。这些学生的智商比其他孩子都高，希望你们能让他们取得更好的成绩。”\n\n　　三位老师都高兴地表示一定尽力。校长又叮嘱他们，对待这些孩子，要像平常一样，不要让孩子或孩子的家长知道他们是被特意挑选出来的，老师们都答应了。\n\n　　一年之后，这三个班的学生成绩果然排在整个学区的前列。这时，校长告诉了老师们真相：这些学生并不是刻意选出的最优秀的学生，只不过是随机抽调的最普通的学生。老师们没想到会是这样，都认为自己的教学水平确实高。这时校长又告诉了他们另一个真相，那就是，他们也不是被特意挑选出的全校最优秀的教师，也不过是随机抽调的普通老师罢了。\n\n　　这个结果正是博士所料到的：这三位教师都认为自己是最优秀的，并且学生又都是高智商的，因此对教学工作充满了信心，工作自然非常卖力，结果肯定非常好了。\n\n　　在做任何事情以前，如果能够充分肯定自我，就等于已经成功了一半。当你面对挑战时，你不妨告诉自己：你就是最优秀的和最聪明的，那么结果肯定是另一种模样。",
			"author_id": 10000,
		},
		{
			"url":       "http://www.baidu.com",
			"show":      1,
			"sort":      1,
			"title":     "晨会小故事2",
			"content":   "两个青年一同开山，一个把石块儿砸成石子运到路边，卖给建房人，一个直接把石块运到码头，卖给杭州的花鸟商人。因为这儿的石头总是奇形怪状，他认为卖重量不如卖造型。三年后，卖怪石的青年成为村里第一个盖起瓦房的人。\n\n　　后来，不许开山，只许种树，于是这儿成了果园。每到秋天，漫山遍野的鸭儿梨招来八方商客。他们把堆积如山的梨子成筐成筐地运往北京、上海，然后再发往韩国和日本。因为这儿的梨汁浓肉脆，香甜无比，就在村上的人为鸭儿梨带来的小康日子欢呼雀跃时，曾卖过怪石的人卖掉果树，开始种柳。因为他发现，来这儿的客商不愁挑不上好梨，只愁买不到盛梨的筐。五年后，他成为第一个在城里买房的人。\n\n　　再后来，一条铁路从这儿贯穿南北，这儿的人上车后，可以北到北京，南抵九龙。小村对外开放，果农也由单一的卖果开始发展果品加工及市场开发。就在一些人开始集资办厂的时候，那个人又在他的地头砌了一道三米高百米长的墙。这道墙面向铁路，背依翠柳，两旁是一望无际的万亩梨园。坐火车经过这里的人，在欣赏盛开的梨花时，会醒目地看到四个大字：可口可乐。据说这是五百里山川中惟一的一个广告，那道墙的主人仅凭这座墙，每年又有四万元的额外收入。\n\n　　90年代末，日本一著名公司的人士来华考察，当他坐火车经过这个小山村的时候，听到这个故事，马上被此人惊人的商业化头脑所震惊，当即决定下车寻找此人。\n\n　　当日本人找到这个人时，他正在自己的店门口与对门的店主吵架。原来，他店里的西装标价800元一套，对门就把同样的西装标价750元；他标750元，对门就标700元。一个月下来，他仅批发出８套，而对门的客户却越来越多，一下子发出了800套。\n\n　　日本人一看这情形，对此人失望不已。但当他弄清真相后，又惊喜万分，当即决定以百万年薪聘请他。原来，对面那家店也是他的。当你在马路上散步的时候，当你坐在火车上向外眺望的时候，假如有一个相貌平平的人，说赚钱是一件很容易的事，仅需要一点点智慧就够了，你千万不要侧目，说不定他就是一个身价百万的人。",
			"author_id": 10001,
		},
		{
			"url":       "http://www.baidu.com",
			"show":      1,
			"sort":      1,
			"title":     "晨会小故事3",
			"content":   "某大公司准备以高薪雇用一名小车司机，经过层层筛选和考试之后，只剩下三名技术最优良的竞争者。主考者问他们：“悬崖边有块金子，你们开着车去拿，觉得能距离悬崖多近而又不至于掉落呢？”\n\n　　“二公尺。”第一位说。\n\n　　“半公尺。”第二位很有把握地说。\n\n　　“我会尽量远离悬崖，愈远愈好。”第三位说。\n\n　　结果这家公司录取了第三位。\n\n　　中年以前不要怕，中年以后不要悔。\n\n　　３０年前，一个年轻人离开故乡，开始创造自己的前途。他动身的第一站，是去拜访本族的族长，请求指点。老族长正在练字，他听说本族有位后辈开始踏上人生的旅途，就写了３个字：不要怕。然后抬起头来，望着年轻人说：“孩子，人生的秘诀只有６个字，今天先告诉你３个，供你半生受用。”\n\n　　３０年后，这个从前的年轻人已是人到中年，有了一些成就，也添了很多伤心事。归程漫漫，到了家乡，他又去拜访那位族长。他到了族长家里，才知道老人家几年前已经去世，家人取出一个密封的信封对他说：\n\n　　“这是族长生前留给你的，他说有一天你会再来。”还乡的游子这才想起来，３０年前他在这里听到人生的一半秘诀，拆开信封，里面赫然又是３个大字：不要悔。\n\n　　打开失败旁边的窗户，也许你就看到了希望\n\n　　一个小女孩趴在窗台上，看窗外的人正埋葬她心爱的小狗，不禁泪流满面，悲恸不已。她的外祖父见状，连忙引她到另一个窗口，让她欣赏他的玫瑰花园。果然小女孩的心情顿时明朗。老人托起外孙女的下巴说：“孩子，你开错了窗户。”",
			"author_id": 10002,
		},
		{
			"url":       "http://www.baidu.com",
			"show":      1,
			"sort":      1,
			"title":     "用人之道",
			"content":   "去过庙的人都知道，一进庙门，首先是弥陀佛，笑脸迎客，而在他的北面，则是黑口黑脸的韦陀。但相传在很久以前，他们并不在同一个庙里，而是分别掌管不同的庙。\n",
			"author_id": 10001,
		},
		{
			"url":       "http://www.baidu.com",
			"show":      1,
			"sort":      1,
			"title":     "所长无用",
			"content":   "有个鲁国人擅长编草鞋，他妻子擅长织白绢。他想迁到越国去。友人对他说：“你到越国去，一定会贫穷的。”“为什么？”“草鞋，是用来穿着走路的，但越国人习惯于赤足走路；白绢，是用来做帽子的，但越国人习惯于披头散发。凭着你的长处，到用不到你的地方去，这样，要使自己不贫穷，难道可能吗？”\n心得：一个人要发挥其专长，就必须适合社会环境需要。如果脱离社会环境的需要，其专长也就失去了价值。因此，我们要根据社会得需要，决定自己的行动，更好去发挥自己的专长。\n",
			"author_id": 10002,
		},
	}
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json "

	for i := 0; i < len(params); i++ {
		param := params[i]
		body, err := util.Post(context.Background(), "http://127.0.0.1:8080/api/v1/articles", param, headers)
		if err != nil {
			t.Fatalf("创建文章失败, err: %v", err)
		}

		var response app.CommonResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			t.Fatalf(err.Error())
		}

		if response.Code != 0 {
			t.Fatalf(response.Msg)
		}
		t.Log("ok")
	}
}

func TestArticleList(t *testing.T) {
	var params map[string]interface{}
	params = map[string]interface{}{
		"url":       "some",
		"show":      "1",
		"sort":      "1",
		"title":     "我测试的文章",
		"content":   "首先，我们定义一个url变量，它表示我们要请求的URL。\n使用http.Get方法发送请求，并将响应结果赋值给resp变量。\n对于发送HTTP请求可能会发生的错误，我们使用err变量进行捕捉。\n在函数返回前，使用defer语句关闭响应的Body流。\n从响应的Body流中读取内容并将结果赋值给body变量。\n最后，将响应的Body转换成字符串并输出到控制台中。",
		"author_id": "10000",
	}
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json "
	body, err := util.Get(context.Background(), "http://127.0.0.1:8080/api/v1/articles", params, headers)
	if err != nil {
		t.Fatalf("获取文章失败, err: %v", err)
	}

	var response app.CommonResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if response.Code != 0 {
		t.Fatalf(response.Msg)
	}
	fmt.Println(string(body))
}

func TestArticleDetail(t *testing.T) {
	id := "1000004"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	body, err := util.Get(context.Background(), "http://127.0.0.1:8080/api/v1/articles/"+id, nil, headers)
	if err != nil {
		t.Fatalf("获取文章失败, err: %v", err)
	}

	var response app.CommonResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if response.Code != 0 {
		t.Fatalf(response.Msg)
	}
	fmt.Println(string(body))
}

func TestArticleDelete(t *testing.T) {
	id := "1000403"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	body, err := util.Delete(context.Background(), "http://127.0.0.1:8080/api/v1/articles/"+id, nil, headers)
	if err != nil {
		t.Fatalf("删除文章失败, err: %v", err)
	}

	var response app.CommonResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if response.Code != 0 {
		t.Fatalf(response.Msg)
	}
	fmt.Println(string(body))
}
