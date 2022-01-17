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
	words = append(words, fmt.Sprintf("现在的人很多都势利眼,看见出入豪车,佩戴名表,就以为这人多么多么的有钱,殊不知真正的有钱人其实非常低调,就拿我说吧,我平时虽然骑个破自行车,但是谁又能想到我还有一辆电动车呢? \n"))
	words = append(words, fmt.Sprintf("今天上午我陪一个女孩逛街,她看上了一条裙子可她没带钱,我就借钱给她.然后那个女孩就把裙子买了,穿上以后.女孩就问我,有女朋友没?我说没有,于是那女孩踮起脚尖,提着裙子转了一圈,还娇羞红着脸低着头问我,好看吗?我立即说,好看!女孩说嘻嘻,这裙子是我男朋友买的.我当时就一巴掌抽飞这个女孩,明明是老子借钱给你买的,你说是你男朋友买的,你是不是不想还钱?"))
	words = append(words, fmt.Sprintf("好兄弟天降车祸去世,我非常伤心.于是在兄弟墓前说:\"兄弟啊,我真的很舍不得你啊!我媳妇已经怀孕八个月了,很快就要生了,要不你直接投胎过来吧.\"一个多月后,媳妇生了个大胖小子,孩子一天天长大,越看孩子越觉得长得像我兄弟.心里很高兴,感觉自己的许愿成真了."))
	words = append(words, fmt.Sprintf("我花了6个月时间4万块钱培训了java开发,刚培训完两个星期我就收到了顺丰的offer 我承认我不是班里学习最好的 但我却是班里第一个找到工作的 而且还是个大厂 我一直相信 勤能补拙 只要有决心 什么事都是可以做到的 今天入职一个星期了 公司的人对我都很好 还给我配了 电动车 和头盔 还有冲锋衣 不说了 又要去送快递了. "))
	words = append(words, fmt.Sprintf("发现我挺能骗女孩子上床的,只要我一给她发消息,她就说她要睡觉了,我还真是个渣男. "))
	words = append(words, fmt.Sprintf("有没有懂车的朋友,大佬有没有熟人,帮我咨询一下价格在100万到1000万之间的车,动力3.0T以上,外观要好看,最好是白色,德国车系最好,内饰最好是全真皮的,带全景天窗. 最好是越野车,通过性好些,有认识的推荐我几辆. 我要作手机屏保,谢谢! \n"))
	words = append(words, fmt.Sprintf("某日,库克邀雷军上门做客,克煮酒招待.席间,克问:\"当今天下,手机界谁可称为英雄?\" 军:\"魅族蛟王,自命不凡;今虎踞南粤,声势浩大,部下能事者极多,又联高通可为英雄?\" 克:\"章奇技淫巧,不擅营销,喜闭门打磨联发科,最早布局却被后人居上,非英雄也.\" 军:\"有一人演讲卓绝,相声无敌,英语大成者,且定制系桶毒步天下,永浩可为英雄?\" 库克:\"锤子徒有虚名,沽名钓誉,空讲情怀大话,非英雄也.\" 军:\"有一人血气方刚,单挑BAT,战遍互联网,鸿祎乃英雄也?\" 克:\"红衣教主擅蛊惑人心,乃流氓出身,仇家遍地,实非英雄.\" 军:\"华为承东,出自名门,自研海思麒麟,麾下水军百万,可为英雄乎?\" 克:\"华为虽系名门,然大嘴实乃守户之 犬耳,空放大炮,偷换闪存,何足为英雄!\" 军:\"如一加作虎\乐视跃亭等辈皆何如?\" 克鼓掌大笑曰:\"此等碌碌小人,何足挂齿!\" 军:\"舍此之外,某实不知,谁能当之?\" 克以手指军,后自指,曰:\"今天下英雄,惟猴王与库耳! "))
	words = append(words, fmt.Sprintf("独坐空房手作妻,此事羞与外人提.\n面前电脑东瀛女,桌面卷纸铺整齐.\n一上一下渐入戏,忽快忽慢眼迷离.\n点点滴滴落在地,子子孙孙化作泥.\n事后惊觉无意义,决心不再手作妻.\n又是一天深夜里,再把卷纸铺整齐.\n昨日感慨已忘记,宅男废纸伤身体.\n还未娶妻先肾虚,你说可惜不可惜?"))
	words = append(words, fmt.Sprintf("友情提醒:通常别人问你忙不忙,最好说很忙.因为根据经验,如果你说不忙,十有八九对方都会让你忙起来. "))
	words = append(words, fmt.Sprintf("今天我跟妹子表白了,妹子拒绝了我,她狠心的对我说:你也不撒泡尿照照自己,你配得上我吗?于是 我脱下裤子,准备撒泡尿照照自己,我刚脱下裤子,妹子咽了一下口水,对我说:等等,也许……我们可以试一试 "))
	words = append(words, fmt.Sprintf("一部手机的寿命大约在三到五年,只有人类寿命的二十分之一. 手机只是你生活的一部分,而你却是它的全部. 请放下身边的杂事,多陪陪你的手机. "))
	words = append(words, fmt.Sprintf("看完后,心酸酸的.回想这几年,经历了太多起落,尝尽辛酸艰辛.从一开始什么都没有到30万,从30万到200万,从200万到500多万,从500万到800万,最后从800万到现在的1300万!不是炫耀, 我只是想通过我自己的经历告诉我的朋友们:手机像素越高,拍的照片越清晰. "))
	words = append(words, fmt.Sprintf("总是有这么一种人,惹得我们心情烦躁肝火上升.这里介绍一剂去火良方:陈皮4.5克,半夏3克,茯苓3克,甘草0.9克,加三碗水大火熬成一-碗,然后趁热泼对方脸上. "))
	words = append(words, fmt.Sprintf("刚才我老婆突然和我翻脸了,昨天她说要金饰品,我问不是饰品行不行,她说只要是金的就行.今天买了一箱金罐加多宝,本来以为可以一块喝的,没想到她不喜欢喝加多宝就和我翻脸了 "))
	words = append(words, fmt.Sprintf("少男与少女并排坐在那里,少男盯着少女 \n\"我来大姨妈了\"少女道,少年微微一叹气,但依然盯着少女,没有放弃的意思 \n\"我痔疮犯了\",少女道,少年微微一叹气,但依然盯着少女没有放弃的意思. \n\"我口腔溃疡了\"少女继续道,眼里满是委屈. \n少年似乎终于放弃了,他满脸遗憾的道,那好吧, 那好吧,来个鸳鸯锅. "))
	words = append(words, fmt.Sprintf("话说川普和希拉里完成电视辩论后很晚了,要吃东西补充能量,他们一起进一家面包店,可两人发现身上都没带钱.希拉里犹豫了一下,然后趁店员不注意时伸手取了3块糕点藏进口袋.\n\n她小声对川普说: :你看我多聪明,店主什么都没看到,而我甚至都不需要撒谎, 我一定会战胜你赢得美国总统宝座!\"川普说: \"为了达到目的,以欺骗和盗窃行事,难道这就是你的一贯作风吗?其实,用诚实的方式也能获得相同的结果.看我的!\"\n\n于是川普对面包店老板说:\"我是共和党总统候选人川普,现在我会表演一个魔术给你看,拜托你帮个忙把这块糕点递给我.\"\n\n店老板很好奇,立即按他吩咐给了他一块糕点.川普狼吞虎咽下去,并要求再来一块,也吞下肚.然后川普又要了第三块糕点,又一口顺到胃里.\n\n此时店主有些纳闷,就问川普到底魔术何时会出现.\n\n川普答曰:\"现在是见证奇迹的时刻了,刚才的三块蛋糕都在希拉里的口袋里.\"\n"))
	words = append(words, fmt.Sprintf("现在的女人真是不讲理,尤其是漂亮的女人,更是这样.今天我在路上走,一美女路过掉了盒口香糖.我就喊她,结果她上来就打人,真的是蛮不讲理!现在我东西也不想还给她了,对了,杰*邦这种口香糖好吃吗?"))
	words = append(words, fmt.Sprintf("如果你追一个人,长时间都追不到,那就问自己,\n第一,你是不是穷的只剩下一张床了.\n第二,你是不是胖的只剩下一堆肉了.\n第三,你是不是丑的只剩下背影了\n如果你都做的很好,你还是追不到那个人,就问自己最后一个问题,\n你是不是贱的只剩骨头了!\n"))
	words = append(words, fmt.Sprintf("昨天学校在查早恋,查到我女朋友头上来了,本以为要完蛋了,结果查出来她男朋友不是我……那群老师真傻哈哈哈!\n"))
	words = append(words, fmt.Sprintf("十年没联系的对象,今天突然给我电话,我激动万分,不知道如何是好,幻想着是不是要跟我重续前缘?结果她上来就是一句:\"你赶紧让你儿子转学,他现在在追你的女儿……\"\n"))
	words = append(words, fmt.Sprintf("今天在饭店吃饭,邻桌一对男女在秀恩爱.\n男:亲爱的,昨晚你太累了,多吃点鸡腿补补.\n女:不要啦,人家会胖的.\n男:我不在乎你胖的.\n女:我老公会在乎的……\n"))
	words = append(words, fmt.Sprintf("出差回来,上了一辆绿色的出租车,司机穿着绿色的衣服,放着苏打绿的歌,喝着绿茶,一路绿灯,回到家老婆赤裸的躺在床上.\n想想今天一路畅通,连脱衣的的时间都省了.\n"))
	words = append(words, fmt.Sprintf("小时候抓了一只老鼠,想着怎么弄死它.然后想到用电把它电焦……于是我找了两根铁丝,一根绑住老鼠头,一根绑住老鼠的腰.然后拿着铁丝另一头就往插座插了过去……要不是后来瓷闸保险丝烧断了,我可能会死在老鼠前面…\n"))
	words = append(words, fmt.Sprintf("我上拼多多去买了一个降落伞,价格出奇的便宜,而且惊奇地里面居然没有差评,我马上就买了一个,立马坐着飞机去跳伞.在我跳下去的那一刻,我貌似明白了为什么没有差评. \n"))
	words = append(words, fmt.Sprintf("向班级里喜欢的妹子表白,但被拒绝了.知道她喜欢吃甜点,于是我每天都去学校外的甜品店给她买点心.\n风雨无阻,从未间断.在我的努力下,她终于胖到了180斤,再也没有人喜欢她了.\n而我也终于追到了甜品店漂亮的收银员姐姐."))
	words = append(words, fmt.Sprintf("有一个叫波多野吉衣的女人命运很凄惨,在家被家人侵犯,在公司被同事侵犯,就连在公交车地铁上也被侵犯,我特别同情她,我想为她住持正义,奈何证据不足,所以请有视频证据的朋友能发给我一下,让我们一起为正义加油,让世界充满爱."))
	words = append(words,fmt.Sprintf("抽着中华 开着911 听着歌曲 女票在副驾驶露腿坐着 红灯下顺手搭在女票的腿上 这生活 啧啧 我在公交车上看的一清二楚.\n"))
	words = append(words,fmt.Sprintf("地铁上我带着耳机听这首歌,旁边有个小姑娘问我:你听的是沙漠骆驼吧?我惊讶的问你怎么知道?她说:你踩我脚的节奏我感觉出来的."))
	words = append(words,fmt.Sprintf("邹忌修八尺有余,而形貌昳丽.朝服衣冠,窥镜,谓其妻曰:"我孰与城北徐公美?"其妻曰:"君美甚,徐公何能及君也?"城北徐公,齐国之美丽者也.忌不自信,而复问其妾曰:"吾孰与徐公美?"妾曰:"徐公何能及君也?"旦日,客从外来,与坐谈,问之客曰:"吾与徐公孰美?"客曰:"忌你太美.""))
	words = append(words,fmt.Sprintf("昨晚跑顺风车接了一对小情侣回学校,差不多10:40了,听他们的意思好像11:00门禁.我开着车听男生说:"估计来不及了,要不今晚住外面吧?"女生害羞的嗯了一声,估计也有这意思.我顿时倦意全无,猛抬头:"你们放心,哥这450马力绝对来得及",然后一顿猛加速,准时把他们送到了学校……\n"))
	words = append(words,fmt.Sprintf("2019版相亲:准丈母娘提问流行款是:"小刘,你家车库物业费现在交多少哦? "这个问题非常有学问,既避免了物质化的庸俗,又包含了大量的信息,确保有房有车\有车位,而且能从物业费判断小区档次!标准答案是:阿姨,您问的是哪一套哦?\n"))
	words = append(words,fmt.Sprintf("我是我们村里第一个网购的人,我们这里大部分人都是用的小灵通,只有村长家有一部智能手机.当听说我要网购的时候,整个村都震惊了,村长跑到我家对我爸说,你家娃是不疯了.媳妇也跟我闹离婚这日子还过不过了.面对着重重压力,我依然坚持要买,我相信我这一个月工资不会白花.终于,快递小哥骑着摩托车出现了,小哥就像下达圣旨一样神圣.当快递小哥依依不舍的把快递交到我手上,我对他说: 只要你好好努力,不出十年你绝对可以拥有.在快递小哥泛着泪光的注视下,我颤抖着拆开包裹.那一刹那,感觉我的眼睛都要闪瞎了.oh my god ,此物应只有天上有,只恨我读书少,无法用华丽的词语来形容它.我拿着它骄傲的站在村口,顿时整个村都沸腾了,大姑娘小媳妇拼了命的向我涌来,更有人趁我不备的时候想摸摸它,我不给她们摸,她们就要跳井.就连村花都红着脸要跟我回家,看着隔壁老王杀人的目光,才想起这是他花了一麻袋地瓜换来的老婆.吓得我赶紧收起它,挤出人群落荒而逃. \n"))
	words = append(words,fmt.Sprintf("鲁迅姓周,文笔那么顺畅,为什么不叫周笔畅?\n周树人笔名鲁迅,他智慧过人,文笔又能深刻反映社会,那为什么不叫鲁智深?\n鲁迅姓周,他的思想在当时社会如同寒冬里的春雨洗涤人们心灵,为什么不叫周冬雨?\n鲁迅姓周,他的思想在那个时代如同流星划破黑夜飞驰而过给予人们希望,为什么不叫周星驰?\n鲁迅姓周,家门前有两颗枣树,他又是个人,那么为什么不叫周树人?\n鲁迅姓周,当年润笔费不少,也没发财在上海多买几套房,那么为什么他不叫周润发呢?\n鲁迅姓周,此文曲星斗转银河数十年,白马飞驰,时光无迹,那么为什么他不叫周星驰呢?\n鲁迅姓周,又是华人,写的杂文雄深雅健,那么为什么他不叫周华健呢?\n鲁迅姓周,留日海归,却从不谄媚外族,那么为什么他不叫周海媚呢?\n鲁迅姓周,人格顶天立地,在当年的文坛掀起阵阵波涛,那么为什么他不叫周立波呢?\n鲁迅姓周,其随笔,晓来如染霜林醉,鸥鸟惊闻四散飞,那么为什么他不叫周晓鸥呢?\n鲁迅姓周,关于他的传说不少,又是民族英雄,那么为什么他不叫周传雄呢?\n鲁迅姓周,也曾饮冰数十年,却不改倩士之俊,那么为什么他不叫周冰倩呢?\n鲁迅姓周,自小聪慧过人,敏而胜猹,那么为什么他不叫周慧敏呢?\n鲁迅姓周,博学中西,坐读马列,实乃识时务的俊杰,那么为什么他不叫周杰呢?\n鲁迅姓周,下笔迅疾如松间怒风,那么为什么不叫他周迅呢?\n鲁迅姓周,出生于人杰地灵的绍兴,拥有语妙绝伦的文采,那么为什么不叫他周杰伦呢?\n鲁迅姓周,矢志不渝,终其半生批判着麻木的国民,那么为什么不叫他周渝民呢?\n鲁迅姓周,阅其文,令人如感冬雷震震夏雨雪,那么为什么不叫他周冬雨呢?\n鲁迅姓周,写的论说文,鞭辟入里,颇有深度,那么为什么不叫他周深呢?\n"))
	words = append(words,fmt.Sprintf("雷震子有一双羽翼,能引天空的雷电,驾风穿越大海,世人称之为天海翼.\n刘备文武双全,献帝称之曰能,世人又叫他刘能.\n宋朝武力不足,被北方不断侵扰,宋仁宗想到抗击北方\威震华夏的关羽,不由得感叹:社稷为大宝,英雄为小宝,世人因此称关羽为宋小宝.\n张飞以武著称,努尔哈赤认为辽东亦属燕地,自己和张飞是同乡,他曾说过:文有孔子,登泰山而小天下;武有张飞,入辽东而小沈阳,后人因此称其为小沈阳.\n赵云与三人交好,年纪又最小,被后人称为赵四,又因其战为敌手,为战场一哥,也叫他四哥\n"))
	words = append(words,fmt.Sprintf(""))
	words = append(words,fmt.Sprintf(""))
	words = append(words,fmt.Sprintf(""))
	words = append(words,fmt.Sprintf(""))
	words = append(words, fmt.Sprintf("几年前被骗近了传销,很多人都试图逃跑,可门口有条大狼狗,逃跑的人没跑出20米远就被逮了回来.而我就比较聪明了,每次吃饭只吃一半留一半给看门狗.过了两个星期那条狗自然就跟我熟了.一天夜里,我感觉时机成熟了,趁他们不注意,我把狗给日了\n"))
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
