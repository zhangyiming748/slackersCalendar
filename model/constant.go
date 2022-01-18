package model

import "fmt"

//anniversary
var (
	AS      AnniversarySlice
	QVOD    Anniversary //快播关闭
	Apple   Anniversary //iphone发布
	Nuclear Anniversary //中国第一枚实用氢弹试爆成功
)

type Anniversary struct {
	name   string //中文名
	subday int    //剩余日期
	date   string //日期
}

func (a *Anniversary) SetName(s string) {
	a.name = s
	return
}
func (a Anniversary) GetName() string {
	return a.name
}
func (a *Anniversary) SetSubDay(i int) {
	a.subday = i
	return
}
func (a Anniversary) GetSubDay() int {
	return a.subday
}
func (a *Anniversary) SetDate(s string) {
	a.date = s
	return
}
func (a Anniversary) GetDate() string {
	return a.date
}

type AnniversarySlice []Anniversary

func (s AnniversarySlice) Len() int { // 重写 Len() 方法
	return len(s)
}
func (s AnniversarySlice) Swap(i, j int) { // 重写 Swap() 方法
	s[i], s[j] = s[j], s[i]
}
func (s AnniversarySlice) Less(i, j int) bool { // 重写 Less() 方法, 从大到小排序
	return s[i].subday < s[j].subday
}

//Gift
var (
	words = []string{}
)

func init() {
	words = append(words, fmt.Sprintf("年轻的时候千万不要因为没钱而绝望,因为你要知道,你以后没钱的日子还很多"))
	words = append(words, fmt.Sprintf("铁打的身体,磁铁打的床"))
	words = append(words, fmt.Sprintf("现在的人很多都势利眼,看见出入豪车,佩戴名表,就以为这人多么多么的有钱,殊不知真正的有钱人其实非常低调,就拿我说吧,我平时虽然骑个破自行车,但是谁又能想到我还有一辆电动车呢?"))
	words = append(words, fmt.Sprintf("今天上午我陪一个女孩逛街,她看上了一条裙子可她没带钱,我就借钱给她.然后那个女孩就把裙子买了,穿上以后.女孩就问我,有女朋友没?我说没有,于是那女孩踮起脚尖,提着裙子转了一圈,还娇羞红着脸低着头问我,好看吗?我立即说,好看!女孩说嘻嘻,这裙子是我男朋友买的.我当时就一巴掌抽飞这个女孩,明明是老子借钱给你买的,你说是你男朋友买的,你是不是不想还钱?"))
	words = append(words, fmt.Sprintf("好兄弟天降车祸去世,我非常伤心.于是在兄弟墓前说:\"兄弟啊,我真的很舍不得你啊!我媳妇已经怀孕八个月了,很快就要生了,要不你直接投胎过来吧.\"一个多月后,媳妇生了个大胖小子,孩子一天天长大,越看孩子越觉得长得像我兄弟.心里很高兴,感觉自己的许愿成真了."))
	words = append(words, fmt.Sprintf("我花了6个月时间4万块钱培训了java开发,刚培训完两个星期我就收到了顺丰的offer,我承认我不是班里学习最好的,但我却是班里第一个找到工作的,而且还是个大厂,我一直相信,勤能补拙,只要有决心,什么事都是可以做到的,今天入职一个星期了,公司的人对我都很好,还给我配了,电动车,和头盔,还有冲锋衣,不说了,又要去送快递了"))
	words = append(words, fmt.Sprintf("发现我挺能骗女孩子上床的,只要我一给她发消息,她就说她要睡觉了,我还真是个渣男"))
	words = append(words, fmt.Sprintf("有没有懂车的朋友,大佬有没有熟人,帮我咨询一下价格在100万到1000万之间的车,动力3.0T以上,外观要好看,最好是白色,德国车系最好,内饰最好是全真皮的,带全景天窗,最好是越野车,通过性好些,有认识的推荐我几辆,我要作手机屏保,谢谢!"))
	words = append(words, fmt.Sprintf("某日,库克邀雷军上门做客,克煮酒招待.席间,克问:\"当今天下,手机界谁可称为英雄?\"军:\"魅族蛟王,自命不凡;今虎踞南粤,声势浩大,部下能事者极多,又联高通可为英雄?\"克:\"章奇技淫巧,不擅营销,喜闭门打磨联发科,最早布局却被后人居上,非英雄也.\"军:\"有一人演讲卓绝,相声无敌,英语大成者,且定制系桶毒步天下,永浩可为英雄?\"库克:\"锤子徒有虚名,沽名钓誉,空讲情怀大话,非英雄也.\"军:\"有一人血气方刚,单挑BAT,战遍互联网,鸿祎乃英雄也?\"克:\"红衣教主擅蛊惑人心,乃流氓出身,仇家遍地,实非英雄.\"军:\"华为承东,出自名门,自研海思麒麟,麾下水军百万,可为英雄乎?\"克:\"华为虽系名门,然大嘴实乃守户之犬耳,空放大炮,偷换闪存,何足为英雄!\"军:\"如一加作虎、乐视跃亭等辈皆何如?\"克鼓掌大笑曰:\"此等碌碌小人,何足挂齿!\"军:\"舍此之外,某实不知,谁能当之?\"克以手指军,后自指,曰:\"今天下英雄,惟猴王与库耳!"))
	words = append(words, fmt.Sprintf("独坐空房手作妻,此事羞与外人提.\n面前电脑东瀛女,桌面卷纸铺整齐.\n一上一下渐入戏,忽快忽慢眼迷离.\n点点滴滴落在地,子子孙孙化作泥.\n事后惊觉无意义,决心不再手作妻.\n又是一天深夜里,再把卷纸铺整齐.\n昨日感慨已忘记,宅男废纸伤身体.\n还未娶妻先肾虚,你说可惜不可惜?"))
	words = append(words, fmt.Sprintf("友情提醒:通常别人问你忙不忙,最好说很忙.因为根据经验,如果你说不忙,十有八九对方都会让你忙起来"))
	words = append(words, fmt.Sprintf("今天我跟妹子表白了,妹子拒绝了我,她狠心的对我说:你也不撒泡尿照照自己,你配得上我吗?于是,我脱下裤子,准备撒泡尿照照自己,我刚脱下裤子,妹子咽了一下口水,对我说:等等,也许……我们可以试一试"))
	words = append(words, fmt.Sprintf("一部手机的寿命大约在三到五年,只有人类寿命的二十分之一,手机只是你生活的一部分,而你却是它的全部,请放下身边的杂事,多陪陪你的手机."))
	words = append(words, fmt.Sprintf("看完后,心酸酸的.回想这几年,经历了太多起落,尝尽辛酸艰辛.从一开始什么都没有到30万,从30万到200万,从200万到500多万,从500万到800万,最后从800万到现在的1300万!不是炫耀,我只是想通过我自己的经历告诉我的朋友们:手机像素越高,拍的照片越清晰"))
	words = append(words, fmt.Sprintf("总是有这么一种人,惹得我们心情烦躁肝火上升.这里介绍一剂去火良方:陈皮4.5克,半夏3克,茯苓3克,甘草0.9克,加三碗水大火熬成一-碗,然后趁热泼对方脸上"))
	words = append(words, fmt.Sprintf("刚才我老婆突然和我翻脸了,昨天她说要金饰品,我问不是饰品行不行,她说只要是金的就行.今天买了一箱金罐加多宝,本来以为可以一块喝的,没想到她不喜欢喝加多宝就和我翻脸了"))
	words = append(words, fmt.Sprintf("少男与少女并排坐在那里,少男盯着少女\n\"我来大姨妈了\"少女道,少年微微一叹气,但依然盯着少女,没有放弃的意思\n\"我痔疮犯了\",少女道,少年微微一叹气,但依然盯着少女没有放弃的意思\n\"我口腔溃疡了\"少女继续道,眼里满是委屈\n少年似乎终于放弃了,他满脸遗憾的道,那好吧,那好吧,来个鸳鸯锅"))
	words = append(words, fmt.Sprintf("话说川普和希拉里完成电视辩论后很晚了,要吃东西补充能量,他们一起进一家面包店,可两人发现身上都没带钱.希拉里犹豫了一下,然后趁店员不注意时伸手取了3块糕点藏进口袋\n她小声对川普说:你看我多聪明,店主什么都没看到,而我甚至都不需要撒谎,我一定会战胜你赢得美国总统宝座!\"川普说:\"为了达到目的,以欺骗和盗窃行事,难道这就是你的一贯作风吗?其实,用诚实的方式也能获得相同的结果.看我的!\"\n\n于是川普对面包店老板说:\"我是共和党总统候选人川普,现在我会表演一个魔术给你看,拜托你帮个忙把这块糕点递给我.\"\n\n店老板很好奇,立即按他吩咐给了他一块糕点.川普狼吞虎咽下去,并要求再来一块,也吞下肚.然后川普又要了第三块糕点,又一口顺到胃里.\n\n此时店主有些纳闷,就问川普到底魔术何时会出现.\n\n川普答曰:\"现在是见证奇迹的时刻了,刚才的三块蛋糕都在希拉里的口袋里.\""))
	words = append(words, fmt.Sprintf("现在的女人真是不讲理,尤其是漂亮的女人,更是这样.今天我在路上走,一美女路过掉了盒口香糖.我就喊她,结果她上来就打人,真的是蛮不讲理!现在我东西也不想还给她了,对了,杰*邦这种口香糖好吃吗?"))
	words = append(words, fmt.Sprintf("如果你追一个人,长时间都追不到,那就问自己,\n第一,你是不是穷的只剩下一张床了.\n第二,你是不是胖的只剩下一堆肉了.\n第三,你是不是丑的只剩下背影了\n如果你都做的很好,你还是追不到那个人,就问自己最后一个问题,\n你是不是贱的只剩骨头了!"))
	words = append(words, fmt.Sprintf("昨天学校在查早恋,查到我女朋友头上来了,本以为要完蛋了,结果查出来她男朋友不是我……那群老师真傻哈哈哈!"))
	words = append(words, fmt.Sprintf("十年没联系的对象,今天突然给我电话,我激动万分,不知道如何是好,幻想着是不是要跟我重续前缘?结果她上来就是一句:\"你赶紧让你儿子转学,他现在在追你的女儿……\""))
	words = append(words, fmt.Sprintf("今天在饭店吃饭,邻桌一对男女在秀恩爱.\n男:亲爱的,昨晚你太累了,多吃点鸡腿补补.\n女:不要啦,人家会胖的.\n男:我不在乎你胖的.\n女:我老公会在乎的……"))
	words = append(words, fmt.Sprintf("出差回来,上了一辆绿色的出租车,司机穿着绿色的衣服,放着苏打绿的歌,喝着绿茶,一路绿灯,回到家老婆赤裸的躺在床上.\n想想今天一路畅通,连脱衣的的时间都省了."))
	words = append(words, fmt.Sprintf("小时候抓了一只老鼠,想着怎么弄死它.然后想到用电把它电焦……于是我找了两根铁丝,一根绑住老鼠头,一根绑住老鼠的腰.然后拿着铁丝另一头就往插座插了过去……要不是后来瓷闸保险丝烧断了,我可能会死在老鼠前面…"))
	words = append(words, fmt.Sprintf("我上拼多多去买了一个降落伞,价格出奇的便宜,而且惊奇地里面居然没有差评,我马上就买了一个,立马坐着飞机去跳伞.在我跳下去的那一刻,我貌似明白了为什么没有差评"))
	words = append(words, fmt.Sprintf("向班级里喜欢的妹子表白,但被拒绝了.知道她喜欢吃甜点,于是我每天都去学校外的甜品店给她买点心.\n风雨无阻,从未间断.在我的努力下,她终于胖到了180斤,再也没有人喜欢她了.\n而我也终于追到了甜品店漂亮的收银员姐姐."))
	words = append(words, fmt.Sprintf("有一个叫波多野吉衣的女人命运很凄惨,在家被家人侵犯,在公司被同事侵犯,就连在公交车地铁上也被侵犯,我特别同情她,我想为她住持正义,奈何证据不足,所以请有视频证据的朋友能发给我一下,让我们一起为正义加油,让世界充满爱."))
	words = append(words, fmt.Sprintf("抽着中华,开着911,听着歌曲,女票在副驾驶露腿坐着,红灯下顺手搭在女票的腿上,这生活,啧啧,我在公交车上看的一清二楚."))
	words= append(words, fmt.Sprintf("地铁上我带着耳机听这首歌,旁边有个小姑娘问我:你听的是沙漠骆驼吧?我惊讶的问你怎么知道?她说:你踩我脚的节奏我感觉出来的."))
	words = append(words, fmt.Sprintf("邹忌修八尺有余,而形貌昳丽.朝服衣冠,窥镜,谓其妻曰:\"我孰与城北徐公美?\"其妻曰:\"君美甚,徐公何能及君也?\"城北徐公,齐国之美丽者也.忌不自信,而复问其妾曰:\"吾孰与徐公美?\"妾曰:\"徐公何能及君也?\"旦日,客从外来,与坐谈,问之客曰:\"吾与徐公孰美?\"客曰:\"忌你太美\""))
	words = append(words, fmt.Sprintf("昨晚跑顺风车接了一对小情侣回学校,差不多10:40了,听他们的意思好像11:00门禁.我开着车听男生说:\"估计来不及了,要不今晚住外面吧?\"女生害羞的嗯了一声,估计也有这意思.我顿时倦意全无,猛抬头:\"你们放心,哥这450马力绝对来得及\",然后一顿猛加速,准时把他们送到了学校……"))
	words = append(words, fmt.Sprintf("2019版相亲:准丈母娘提问流行款是:\"小刘,你家车库物业费现在交多少哦?\"这个问题非常有学问,既避免了物质化的庸俗,又包含了大量的信息,确保有房有车有车位,而且能从物业费判断小区档次!标准答案是:阿姨,您问的是哪一套哦?"))
	words = append(words, fmt.Sprintf("我是我们村里第一个网购的人,我们这里大部分人都是用的小灵通,只有村长家有一部智能手机.当听说我要网购的时候,整个村都震惊了,村长跑到我家对我爸说,你家娃是不疯了.媳妇也跟我闹离婚这日子还过不过了.面对着重重压力,我依然坚持要买,我相信我这一个月工资不会白花.终于,快递小哥骑着摩托车出现了,小哥就像下达圣旨一样神圣.当快递小哥依依不舍的把快递交到我手上,我对他说:只要你好好努力,不出十年你绝对可以拥有.在快递小哥泛着泪光的注视下,我颤抖着拆开包裹.那一刹那,感觉我的眼睛都要闪瞎了.oh my god ,此物应只有天上有,只恨我读书少,无法用华丽的词语来形容它.我拿着它骄傲的站在村口,顿时整个村都沸腾了,大姑娘小媳妇拼了命的向我涌来,更有人趁我不备的时候想摸摸它,我不给她们摸,她们就要跳井.就连村花都红着脸要跟我回家,看着隔壁老王杀人的目光,才想起这是他花了一麻袋地瓜换来的老婆.吓得我赶紧收起它,挤出人群落荒而逃."))
	words = append(words, fmt.Sprintf("鲁迅姓周,文笔那么顺畅,为什么不叫周笔畅?\n周树人笔名鲁迅,他智慧过人,文笔又能深刻反映社会,那为什么不叫鲁智深?\n鲁迅姓周,他的思想在当时社会如同寒冬里的春雨洗涤人们心灵,为什么不叫周冬雨?\n鲁迅姓周,他的思想在那个时代如同流星划破黑夜飞驰而过给予人们希望,为什么不叫周星驰?\n鲁迅姓周,家门前有两颗枣树,他又是个人,那么为什么不叫周树人?\n鲁迅姓周,当年润笔费不少,也没发财在上海多买几套房,那么为什么他不叫周润发呢?\n鲁迅姓周,此文曲星斗转银河数十年,白马飞驰,时光无迹,那么为什么他不叫周星驰呢?\n鲁迅姓周,又是华人,写的杂文雄深雅健,那么为什么他不叫周华健呢?\n鲁迅姓周,留日海归,却从不谄媚外族,那么为什么他不叫周海媚呢?\n鲁迅姓周,人格顶天立地,在当年的文坛掀起阵阵波涛,那么为什么他不叫周立波呢?\n鲁迅姓周,其随笔,晓来如染霜林醉,鸥鸟惊闻四散飞,那么为什么他不叫周晓鸥呢?\n鲁迅姓周,关于他的传说不少,又是民族英雄,那么为什么他不叫周传雄呢?\n鲁迅姓周,也曾饮冰数十年,却不改倩士之俊,那么为什么他不叫周冰倩呢?\n鲁迅姓周,自小聪慧过人,敏而胜猹,那么为什么他不叫周慧敏呢?\n鲁迅姓周,博学中西,坐读马列,实乃识时务的俊杰,那么为什么他不叫周杰呢?\n鲁迅姓周,下笔迅疾如松间怒风,那么为什么不叫他周迅呢?\n鲁迅姓周,出生于人杰地灵的绍兴,拥有语妙绝伦的文采,那么为什么不叫他周杰伦呢?\n鲁迅姓周,矢志不渝,终其半生批判着麻木的国民,那么为什么不叫他周渝民呢?\n鲁迅姓周,阅其文,令人如感冬雷震震夏雨雪,那么为什么不叫他周冬雨呢?\n鲁迅姓周,写的论说文,鞭辟入里,颇有深度,那么为什么不叫他周深呢?"))
	words = append(words, fmt.Sprintf("雷震子有一双羽翼,能引天空的雷电,驾风穿越大海,世人称之为天海翼.\n刘备文武双全,献帝称之曰能,世人又叫他刘能.\n宋朝武力不足,被北方不断侵扰,宋仁宗想到抗击北方威震华夏的关羽,不由得感叹:社稷为大宝,英雄为小宝,世人因此称关羽为宋小宝.\n张飞以武著称,努尔哈赤认为辽东亦属燕地,自己和张飞是同乡,他曾说过:文有孔子,登泰山而小天下;武有张飞,入辽东而小沈阳,后人因此称其为小沈阳.\n赵云与三人交好,年纪又最小,被后人称为赵四,又因其战为敌手,为战场一哥,也叫他四哥"))
	words = append(words, fmt.Sprintf("六小龄童姓章,他演的孙悟空受泽于天,他为什么不叫章泽天呢?"))
	words = append(words, fmt.Sprintf("刘德华姓刘,他是实力很强的东方演艺界名人,他为什么不叫刘强东呢?"))
	words = append(words, fmt.Sprintf("大家好,我姓彭,因为出生的时候家里面在摆鱼宴,所以我叫彭摆鱼"))
	words = append(words, fmt.Sprintf("我姓马,我出生的时候天上有好多云,所以我叫马好多"))
	words = append(words, fmt.Sprintf("大家好,我姓吴,因为出生的时候天现奇龙,所以我叫吴天谴"))
	words = append(words, fmt.Sprintf("我姓孙,我出生的时候天空响起了红色的雷声,所以我叫孙炸天"))
	words = append(words, fmt.Sprintf("我复姓东方,因为我爹和别人猜拳从未败过,所以我叫东方老赢"))
	words = append(words, fmt.Sprintf("我姓马,因为家里是卖啤酒的,所以给他起名叫马卖啤"))
	words = append(words, fmt.Sprintf("我姓白,出生的时候医院旁边池塘里有一百株荷花,所以我叫白池塘"))
	words = append(words, fmt.Sprintf("我姓范,因为我的身体凉冰冰,像凉水一样,所以我叫范凉水"))
	words = append(words, fmt.Sprintf("我姓成,因为出生的时候有条龙从家门口飞过去,所以我叫成有条"))
	words = append(words, fmt.Sprintf("我姓李,因为出生的时候看见了条很小的龙,所以我叫李很小"))
	words = append(words, fmt.Sprintf("我姓朱,出生时天空飞过一条龙,所以我叫朱一条"))
	words = append(words, fmt.Sprintf("我姓李,出生的时候接连出现了很多杰出的人物,所以我的名字叫做李出多"))
	words = append(words, fmt.Sprintf("我姓胡,我出生的时候我妈妈痛的大声唱歌,所以我叫胡大声"))
	words = append(words, fmt.Sprintf("我姓吴,家人希望我成为一个像先祖那样的文人彦士,所以我叫吴先祖"))
	words = append(words, fmt.Sprintf("我姓古,因为出生的时候家人希望我每天快乐,所以我名字叫古每快"))
	words = append(words, fmt.Sprintf("我姓黄,因为是在秋天出生,所以起名黄天出"))
	words = append(words, fmt.Sprintf("我姓叶,因为家人觉得我出生的时间算得上良辰,所以取名叫叶吉日"))
	words = append(words, fmt.Sprintf("我姓马,我爸说我出生的时候白雾化龙腾空而起,所以我叫马而起"))
	words = append(words, fmt.Sprintf("我姓刘,我出生时,家里希望我德高望重,华丽光鲜,所以给我取名叫刘高丽"))
	words = append(words, fmt.Sprintf("我姓玛,出生在苏州,我妈希望我长大后很美丽,所以叫我玛美丽"))
	words = append(words, fmt.Sprintf("我姓刘,因为出生的时候家里人希望我能平步青云.所以我叫刘平步"))
	words = append(words, fmt.Sprintf("几年前被骗近了传销,很多人都试图逃跑,可门口有条大狼狗,逃跑的人没跑出20米远就被逮了回来.而我就比较聪明了,每次吃饭只吃一半留一半给看门狗.过了两个星期那条狗自然就跟我熟了.一天夜里,我感觉时机成熟了,趁他们不注意,我把狗给日了"))
	words = append(words,fmt.Sprintf("今天我终于狠心买了一杯星巴克尝尝,能喝上星巴克一直是我的梦想.每当我在大街上看到喝星巴克的人,我都会偷偷瞄一眼.真羡慕那些有钱人.昨天我发了工钱, 在经过激烈的思想斗争之后,我终于下定决心明天去买杯星巴克喝,我没有告诉我的父母,他们知道后一定会阻拦我的,我激动的一宿没睡.第二天,我花了两个小时来到市中心的泉水专卖店,我完全被那种气势震住了,豪华的装修...我怀着忐忑不安的心情走进了咖啡店,面积达80平米的营业大厅布置和外面一样豪华,友善的服务员热情的接见了我,可能是太过紧张了,我开口要求试喝,刚说完我就脸红,可是服务员并没有感到诧异,对了,可能来买星巴克的人都得要求先试喝吧,毕竟30元/杯的价格不是说买就买的,我激动地接过星巴克,小心翼翼地.这时的我已经忘了要注意绅士风度,贪婪地喝了一口,爽滑的口感,不是我以前喝的自来水所能比拟的,美丽漂亮的女服务员热情的向我讲解星巴克的配料和制作流程,在谈话中我得知他刚从韩国首尔大学毕业主修酒店管理双硕士学位,回国后她拒绝了去喜来登当店长的机会,来到星巴克专卖店工作,我也为她的明智选择而感到欣慰,毕竟进入星巴克专卖店是多少代人的梦想与追求.临走时我掏出带着体温的100元钱买了两杯, 刚要走出店门时服务员朝我喊:\"你的星巴克!\"我回眸一笑百媚生回应到:\"是你的星巴克.\"服务员当时就脸红了,我随即以迅雷不及掩耳盗铃儿响叮当之势走出门口.我想,服务员肯定会把所有的星巴克从柜子上都拿下来,一瓶一瓶地数\"他明天会来,他明天不会来...\"走在大街上我喝着星巴克,大家都向我投来羡慕的目光,我也感觉对人生更加有信心了,回到家,我兴奋得不能自己,还是忍不住和我的星巴克合照了,因为星巴克实在太过耀眼,我怕有些不法分子会盯上我,所以我用我美丽的签名挡住了我俊朗的面庞.嘘!记得不要告诉别人我喝过星巴克,这年头,像我这样低调的人不多了")
	words = append(words,fmt.Sprintf("现在的女生真TM会装，和几个同学KTV小聚，最后走的时候女神说她喝醉了，要我送她回家，我说你TMD喝果汁也醉? 幸亏老子机灵，要不又耽误老子回去玩贪玩蓝月了")
	words = append(words,fmt.Sprintf("有一次，朋友们在我家聚会，大家都走之后妹子说要留下来看我玩贪玩蓝月，我打完一只野猪之后问她要不要玩，她说她不想玩这个，让我先去洗澡。当时我就给她撵出家门。TM想趁着老子洗澡把老子装备全卖了? 当老子傻? 多亏我TM机智过人")
	words = append(words,fmt.Sprintf("那天，我在网吧玩贪玩蓝月，一个妹子在旁边看韩剧。接近中午12点了，妹子说她不想看韩剧了，她家里只有她一个人，她希望我能和她一起去她家吃饭玩游戏，她还补充说她要和我玩一个世界上最好玩的游戏，这个游戏无数男人都梦寐以求! 我一听就呵呵了，果断拒绝了她。尼玛这个世界上难道还有比贪玩蓝月更好玩的游戏?")
	words = append(words,fmt.Sprintf("我家世世代代务农为生，我不知道外面是什么样的。在我眼中，世界只有毛坯房，渣土路，一望无际的田野和天空，还有墙上的计划生育标语。直到有一天，我家墙上出现了一个\"拆\"字，之后，我的生活就被彻底颠覆了。我搬进了城里，有了几辈子赚不到的钱。原以为，我可以无忧无虑度过后半生了，可是，一些莫名其妙的人又出现在我面前，他们带我见识了这座城市的纸醉金迷，我迷失了自己，那段日子真是浑浑噩噩。谁知有一天，之前跟我称兄道弟的人却对我威胁恐吓，我看了看我赌博欠下的几十万的高利贷，猛然惊醒，可是我能做什么?现在的我除吃喝嫖赌，一无所长。 我茫然无助的走在大街上，一条一条看着那些催债的短信，我兜里仅剩十块钱，又瞥见路边的网吧\"把钱花掉，然后自我了断吧\"我冒出了这样的想法。是啊，对于一个走投无路的人来说，还有什么选择呢?对着电脑，我漫无目的的游荡，因为我也不知道我还能做什么。突然，屏幕右下角弹出窗口\"张家辉邀请你玩贪玩蓝月\" \"张家辉?影帝啊\"虽然不知道贪玩蓝月是什么，但张家辉出面，我也不好说什么，反正我也没有什么可失去的了。进入游戏，游戏中的人物便开始打怪，突然，屏幕发出一阵刺眼的光芒，定睛一看，**戒指，屠龙宝刀!这时我收到张家辉的私聊，要收购那两件装备，我想也没想同意了，因为我没当成什么大事。过了一会，我收到了到账通知，我得到了142.68万元，我激动的屏住了呼吸，我知道这意味着什么。夜里，我叫来了给我高利贷的人，把一把把现金甩在他们脸上，叫他们滚，从此断绝往来。现在我每天和张家辉一起玩贪玩蓝月，每天收入上百万，逐渐走上了成功之路。 多年以后，我拥有美满的生活，却不经意间想起那个下午，那个我一生难忘之人的面影，还有他那萦绕在耳边的话\"我是渣渣灰，贪玩蓝月，是兄弟就来砍我\"那一天不经意间做出的决定，不经意间听到的话语，却能成为一个人矢志不渝的信条。")
	words = append(words,fmt.Sprintf("有位年轻人在外5年，终于要回家了。 老板问：你是要5年的工资还是3句忠告? 年轻人思考了片刻说：明早给您答复。 老板说：可以。 当晚年轻人未眠... 第二天早晨，他对老板说：要3个忠告。 于是老板给他3句话： 一不要试图寻找不可能的捷径，世上没有便宜的事，只有脚踏实地才是最好的方法。 二不要对明知不是好事的事过分好奇，有可能你会因此而丧命。 三不要在冲动时做任何决定，否则这个决定有可能成为你一辈子的遗憾。 说完老板给年轻人了一些钱和一个精美盒子装的礼物，并说礼物在到家后才能开。 年轻人上路了... 他开了好几天车，没过多久遇到一个路口，他打听：请问到**走哪条路近? 路人甲：走小路吧，近。 路人乙：走大路吧，安全。 年轻人因迫不急待要与家中父母见面，便走了小路。 走了没多久，就听到有路人说附近闹劫匪。年轻人忽然想起了老板的第一个忠告:不要试图寻找捷径。 于是他回头去走那条大路。 又走了几天，年轻人找到一家便直的客栈投宿。半夜听到有女子哭，他睡不下，于是决定出门看看。 这时他想到了第二个忠告:不要对明知不是好事的事过分好奇。 于是他又睡下了。 次日起程时，店家惊道：你还活着! 年轻人不解，遂问原由。 店家说，他有个疯女儿，发病时用哭声引人出来再将其杀害，昨晚投宿的客人只有你他一个活着。 年轻人长叹：啊... 又走了几天离家不远矣。他更加激动了。天刚黑下来没多久，他就走到了自己的村子。到家时他刚准备敲门，却听到屋里有男人争吵的声音。 年轻人疑惑，开了门却见几个凶神恶煞的男子在威胁恐吓母亲。 他生气极了，跑到厨房里拿起菜刀，但这时他想到了第三个忠告：不要在冲动时做任何决定，可能这个决定会让你后悔一辈子。 于是他冷静下来，到屋里详细了解才得知父亲欠了几十万的赌债不敢回家，好说歹说才送走讨债鬼。 年轻人又安抚母亲：一切有我，放心。 其实他此时心烦到了极点，几十万对他来说数目太大了，不可能凑得够。 此时年轻人想起了老板那叫他回家再打开的礼物。 他的心怦怦的跳，缓慢的打开盒子，希望里面真的有可以帮助自己的东西，不过他失望了，里面只有一个U盘。 年轻人不解，他把U盘插到笔记本上，打开后发现里边只有一个叫贪玩蓝月的游戏。他迅速点开，看到游戏封面后惊呼：张老板。 从此之后，他就陪着老板刷戒指，才两个星期不到就把父亲几十万赌债还清了，还为自己换了辆四驱车。")
	words = append(words,fmt.Sprintf("这是一个人人手机不离手的时代，她要回你信息早就回了。你苦苦等待一个答案，殊不知，不回复已经是答案了———来自《鲁迅没说过的心里话》")
	words = append(words,fmt.Sprintf("情人节早上，我想给女神发个52块的红包，不小心点成了520块。还好我用的手机比较智能，输入密码后，它提醒我可用的余额不足")
	words = append(words,fmt.Sprintf("")
	words = append(words,fmt.Sprintf("")
	words = append(words,fmt.Sprintf("")
	words = append(words,fmt.Sprintf("")
	words = append(words,fmt.Sprintf("")
	words = append(words,fmt.Sprintf("")
	words = append(words,fmt.Sprintf("")
	words = append(words,fmt.Sprintf("")
	words = append(words,fmt.Sprintf("")
	words = append(words,fmt.Sprintf("")
	words = append(words,fmt.Sprintf("")
	words = append(words,fmt.Sprintf("")
	words = append(words,fmt.Sprintf("")
	words = append(words,fmt.Sprintf("")
	words = append(words,fmt.Sprintf("")
	words = append(words,fmt.Sprintf("")
	words = append(words,fmt.Sprintf("")
}

//timer
var HappyWeekDayMap = map[string]string{
	"Monday":    "星期一,还有四天周末",
	"Tuesday":   "星期二,还有三天周末",
	"Wednesday": "星期三,工作日过去一半了",
	"Thursday":  "肯德基疯狂星期四",
	"Friday":    "周五就应该有周五的态度,有什么事下周一再说",
	"Saturday":  "星期六",
	"Sunday":    "星期日",
}
var SadWeekDayMap = map[string]string{
	"Monday":    "星期一",
	"Tuesday":   "星期二",
	"Wednesday": "星期三",
	"Thursday":  "星期四",
	"Friday":    "星期五不能松懈,明天一定要去公司加班哦!",
	"Saturday":  "星期六",
	"Sunday":    "星期日",
}

//Happy

type HappySlice []Festival

func (c HappySlice) Len() int { // 重写 Len() 方法
	return len(c)
}
func (c HappySlice) Swap(i, j int) { // 重写 Swap() 方法
	c[i], c[j] = c[j], c[i]
}
func (c HappySlice) Less(i, j int) bool { // 重写 Less() 方法, 从大到小排序
	return c[i].subDay < c[j].subDay
}

type Festival struct {
	chineseName string //节日中文名
	subDay      int    //节日剩余日期
	date        string //节日日期
}

func (f *Festival) SetChineseName(s string) {
	f.chineseName = s
	return
}
func (f Festival) GetChineseName() string {
	return f.chineseName
}
func (f *Festival) SetSubDay(i int) {
	f.subDay = i
	return
}
func (f Festival) GetSubDay() int {
	return f.subDay
}
func (f *Festival) SetDate(s string) {
	f.date = s
	return
}
func (f Festival) GetDate() string {
	return f.date
}

var (
	Countdown []Festival
	//公历节日
	NewYear                  Festival // 元旦
	ValentinesDay            Festival //情人节
	WomensDay                Festival //妇女节
	ArborDay                 Festival //植树节
	AprilFoolsDay            Festival //愚人节
	LaborDay                 Festival //劳动节
	YouthDay                 Festival //青年节
	ChildrensDay             Festival //儿童节
	NationalDay              Festival //国庆节
	NationalPovertyReliefDay Festival //国家扶贫日
	Halloween                Festival //万圣夜
	Double11                 Festival //
	NationalConstitutionDay  Festival ////国家宪法日
	ThanksgivingDay          Festival //感恩节
	Double12                 Festival //
	NationalMemorialDay      Festival //国家公祭日
	ChristmasEve             Festival //平安夜
	Christmas                Festival //圣诞节
	//农历节日
	ChineseNewYear      Festival //春节
	LanternFestival     Festival //元宵节
	ChingMingFestival   Festival //清明节
	MothersDay          Festival //母亲节
	DragonBoatFestival  Festival //端午节
	SingleDogDay        Festival //七夕节
	MidyearFestival     Festival //中元节
	MidAutumnFestival   Festival //中秋节
	DoubleNinthFestival Festival //重阳节
	LabaFestival        Festival //腊八节
	//其他
	WinterOlympics Festival //冬奥会
)

const (
	Year = 365 //一年多少天
)

//Sad
type SadSlice []Involution

func (s SadSlice) Len() int { // 重写 Len() 方法
	return len(s)
}
func (s SadSlice) Swap(i, j int) { // 重写 Swap() 方法
	s[i], s[j] = s[j], s[i]
}
func (s SadSlice) Less(i, j int) bool { // 重写 Less() 方法, 从大到小排序
	return s[i].subDay < s[j].subDay
}

type Involution struct { //内卷
	chineseName string //日期中文名
	subDay      int    //日期倒计时
	date        string //当天
}

func (i *Involution) SetChineseName(s string) {
	i.chineseName = s
	return
}
func (i Involution) GetChineseName() string {
	return i.chineseName
}
func (i *Involution) SetSubDay(n int) {
	i.subDay = n
	return
}
func (i Involution) GetSubDay() int {
	return i.subDay
}
func (i *Involution) SetDate(s string) {
	i.date = s
	return
}
func (i Involution) GetDate() string {
	return i.date
}

var (
	Lyingflat                  []Involution
	CollegeEntranceExamination Involution //高考

)
