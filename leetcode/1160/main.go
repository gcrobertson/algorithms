package main

import (
	"fmt"
)

func main() {

	ex1()
	ex2()
	ex3()
	ex4()
}

func ex1() {
	words := []string{"cat", "bt", "hat", "tree"}
	expect := 6
	chars := "atach"
	result := countCharacters(words, chars)

	fmt.Printf("Does result [%d] match expectation [%d]? [%v]\n", result, expect, result == expect)
}

func ex2() {
	words := []string{"hello", "world", "leetcode"}
	expect := 10
	chars := "welldonehoneyr"
	result := countCharacters(words, chars)

	fmt.Printf("Does result [%d] match expectation [%d]? [%v]\n", result, expect, result == expect)
}

func ex3() {
	words := []string{
		"qobxtxjzdngkrzamsxzdvbvkiwijokwdyndqllhqpaoxzwonriclsm",
		"fahtqqnuzhhhrcblquaosdfdcqoskxcsdnhtwvvvzsxkpjprytieieniafnltxmuzwkdnttofpibi",
		"xedhntgrqegfs",
		"wrssfvsbcehbahbvojnzgacbgveitirkjmmyiorwykynqddgydzgfvlvplfnyumgxturxraonpchd",
		"hpmdzhpgijirecxzkuyhptiytnuscu",
		"xljgysrjjukphjgzbpn",
		"gmwbirxt",
		"yhvpsvsnhfmbmcpihnqtodspbvexnpgcqrrughbakbufyjispkquvfppkaffypdpnvikjygdaarcigipfhuvzzzbgw",
		"gvvbgpjolobpbxcnhnzuqrsqgrkanwmnnkqynakkooailoafunopsrlijqhaqmszssxikftcfoqsw",
		"naktgxfyuvuoh",
		"muui",
		"ghzqskipqprrzeligdjyowypeerogxonvztsq",
		"onosezgutawtzteoagbftclsqpfsxtwetourxjxurdqevrir",
		"fcskqxwkzsldsjihahalnw",
		"lsstgzxjxabcbspjwelqmhtnurgfmdtpaihxnxad",
		"nwrwtwetgjwooevhxhkzlvsyghtkldioyjhkkniepktqs",
		"utohzbqvkqkq",
		"vdoqnhtidgicevprrrwlbtldtaxpsxdhxhgbwlkbeglkbhrujtafswjkozdmdpvwhsuskofmxxtprtpopacslinwozth",
		"zqlzxpemomnbzxlorvlkxt",
		"kubshvnldnkofitnnjere",
		"czxmqpowzzhdbhgtlqdokrzxowsvwyavfhcycctgdvrsfzmanrlktfaetnuayrqkvhcbxezfahkrxgaaztovrxuhnll",
		"rrcesnfbxglhjawldnnuiiepdsofbrsbjczlemusqwvenicxmtdmpwfrnizymfmqynvtkbrmablcugroshc",
		"thholqebekkysstqzlbbdapktxumygplqganucwnahmrihiraxdnvbiaqhykcti",
		"fagqnapzeeglbdzsbneosxmptmwopjcxztukhpjkqjmjbkpbzrfaqskctehyboeddmilkwlurcb",
		"wtjdykncubkduhxiwwusoxvzpnaxpnzdjmwddnonsmmvwmuaxghetgrwpoeqbprxgviwzagyqopfdakrqjgiy",
		"nsftehpgzstetbganbtozdopptseucdqkhzdmujnzjdvufqtikgeepttnrabb",
		"ozihktgnvljzhqwanxemtzxphzqvmoblvi",
		"iwdsrekqllbbyarzzmbqbvldvxctdeswiyahiwfoefhfsfwktdzaulnalewbusazjhcbtxjuck",
		"dylhds",
		"idnaddnzbodhjrpshhahnbbnrskruxszxeeyanumazvahktizectmmvjbhnhbrohsyqhrgq",
		"scyzsykrwzuozmbrbenfiqpxchtpmcxepjiglaeionsmbwrzeidupayusczcooudpcgkgspbuyzcdfymsejucb",
		"otpcfrhckxmnseayhwoyjjfkceaoznmmkikpdsuueyqmbsuelmhqnmdsjs",
		"xoghnhpypfiibqrpgtyux",
		"cxhpbcziptgandiwxtctdjpboiqwv",
		"gyjosuhwgbqpcdsdqadopdqozjxpzwxbqaffnxaddhgrxmunpidvpnijxnbawshcznkmprpgkhvzxmzbaftedgtfgjhaisdnva",
		"jlnxshfthqgzlnvjzztrnjswwsotpybxecyqhnfkbfwvmxpcs",
		"mtddspmqwbnoiajpexsuhxogarduzsoammqqelh",
		"sixkvxgnbalipwmkbcwpugpvolsvvlmaaeeobmoeogspbkbsskwjdqkite",
		"xhsutriuynfrldjbhexhxjgcfwcujvwsqfiaqwvjnkjkswdpaynelhryrzfeqjhajezmolk",
		"zfiredtxenzgtrkiamuoubetexzbnwkdtxbtihhtsfubnmryxq",
		"sot",
		"szgkkhuhkcnr",
		"firxkgvkzqlnallzwjispsizoawobemuhqrhpyvknasjzwctamfuonder",
		"vyei",
		"fidqszoicndwifns",
		"xkickycwzj",
		"gmybflbjunudxrwguzditaxmyadgdjeengut",
		"gvlwwnmrddhzwewugdtobauecealfhasyftgxkiqeluysxxmroukfziifpryvddggttojhcszeyjetbsldmorqnbuqreuxfw",
		"vbhuerxkcjlkamgruturkbrubbscmvzqruwvlrlyylcvuiothkhpznjzsfnyfoaqkziyjqzdreeygmtbdljwnaojexfgmjlup",
		"bxjvop",
		"aswdmctosieicucsjwxych",
		"lxbtlhnrfyaznrqgxrltmxkjmhdqjjgnvstxaygmuxznjfiiukm",
		"npncdabkmbgfhnedcmbfyjiplzwymgvsfrifvvjayobsygwolqbwkblqutakcshnlsqadtfiqmpauslvqd",
		"kuyemuuymacyvmahtagmcewkspverpjtjscccnrczchgkjzppdxcalxxcxrwnepk",
		"jxwouobfvzttz",
		"yucpsdhqvzboeezcqpxsepuuk",
		"iwbdmxdcbypnzqmgkrjhfivkrmnv",
		"lpqklgcwdvgbghfyoyejnpexrelywqfdtczflwetbxvzigtvisi",
		"dlasodatffcreungntdmezgfqklrabyymsnhdberufcrgpxgsziklznhdprbczhbxgtuslufycjronozdqumzlnidkvaydg",
		"oejzlmrrbdysqlezifcvgjnmvodfvmrnjmnfkejdbnnyljzjaxfecrfefdgarqbtwoijuktacywmsubtxtgzkbnstgrsrjpkdfe",
		"mlegvhxielwlfadlnqvnipcuizpdxgtvro",
		"vwgmwkbturocovaykdsjaftbtgmtwknnmhexfgcfchpwwgcgecfobbencotjizxbtdrijwfjxdsxanwfjyjamrxrdaiusgr",
		"xjmkcsekcorldyrjiavrhuhqtndujymc", "rmxwcaorznumwxgwnibnxwzvnxjhzwqzgmkgifnnnnzpgtsvprycjtdeirtpqbduursabbidzkfbexgthkoacagkb",
		"tkrsxhztwgvqxamjtexklnooaloydjhejlbasavskttwxiabarogvmfdfjttaxhgqljlbfnrvrwwbxkurhufiknoxfmemcv",
		"cojlyayladyrhofzetaddteqrjbyxtvyszgdorexqmgznqmuvemegbwki",
		"zktqnurtpttlcjgkmnprqeyeepqunfqqvjwuevwbvnaupyofwiqwhpyumyfwpjrruhleromrmpvczlkxqiuq",
		"gyxl",
		"qmhwlymfsjixvvjhkczylqcsnbjxliasetxciggtfl",
		"gizysqkqbyhzeagzsscvdigtcfiupyhqovaaioxfrphugxzrcjvwqwc",
		"hpgkulrmbaixnjiapmjiwhwsgromfqpxqkkiscjwpiicslwcijginxfiwqakeezeohhskxgvgdkezmqys",
		"vibswdyqaxjvqsffwmcismooheyhwzqvyzezluejztlgddshvwcryzcllaisxrygwqyyoiaivfvgtzicycyrkckv",
		"oyclwdejlifmocfjsbgmernsyitkfaahjxfnwnusrdypsziawlpsjgnavytdihpxcmugshnqbkyfts",
		"xwcbiporfbktpvqhznjuaxfcdykifuzwdsdhxirwwxinoffywgxscrtuwhvuwjejzqirsfijgguouapqpmfdnpsfm",
		"qffrjfkj",
		"rpynimubaxdgbrkdawgugnhobaowxdnzkiidworcsnejgqctftyksvbhiwkcnyffmsbxwptqn",
		"oeqdvldvfbklukstxwomapaauaozblhxudkdxihepqagndnlkvbqhluscvczhrsrhodnftoszhjdymuywdtjgzbmkrdf",
		"xviupppkeswkctwlqwyuhokbhqqjshmaeiouhununbhrkabacenkg",
		"habomjnlznqvckmowgelhnglfizogckplbymkdowfpicxydzgoskckraxpdysksvzezcpg",
		"zemawxwjeowraaqsqytsshtavnvoyiyollelxqabmjwhxtrdijiacbbjiiyzwkxboot",
		"jcnpsxnkbqdbh",
		"bftrbaurtzkftodotjguzjmwxrrmswxwclohotuanypmhtemmsaicwckowmcdmpnhcfzptekpgljfvwpqjgilxu",
		"bbaigjqxdmuchkkb",
		"effbptpwafzqbsebbjmvdcxdbmnlfqjklongafzkvaqmkehefedjxgxlbdhltihtgfqjjsdis",
		"lhuxgqpwcgpujfvvnlrxhgbiwxmxzhglyhkdkafafojtlkuxkmjwlxrd",
		"ussuyjqsxwsdhkjeipwycrkcxxjatoqhygzymgikqdnqiyyzlbcdlgtmneyickucbkpkza",
		"ykmsksjorovmdymlbgprvprnyxwppvwgmzfjsqouvgara",
		"wduqkbncayo",
		"qdfkyomjlhfozonwpdxllqdnvpohyijqmyymuclnydzmlqksntdfj",
		"ynzekkkljzrvnwclzcfgtvcmstxgadxpztofckypbgqgbnumnkeaqclaspzxjbygtkjznxeduhbksr",
		"aw", "jbremnahqoiqktpbkteefdkbrerrxmhqbbselpnfzapgmxidrhbixetaetjcroufa",
		"wtufuqgljvxzsurhviikdxyuythezjnvwrxqrykmotkhlphlyfljjsfugzwxxfqkc",
		"ytjuaagqnfxqwiopolnejmoxow",
		"oqoskpzkwxsffgeuuhdklidtmjobzayatyaqefgwgqbo",
		"qkhtpuhhkspffkpryvdjasbxhtfrmptpljszvtgvhfvqpxypxfdaphfqdmigzgfg",
		"goddugelwdvemxwureitzwqmbqeqtymrlrzahuxnpgufaagbpexwpoahvdnsyvgxw",
		"dfhxrctagxkuasofoejsalcerkbtsjwnbnoahsumzfyiiskhzzwyykeefszrzrbztbrzhxxgaajb",
		"tjsardsbhmhefysdqtbhzsxukbwdpioqaohloaancxdkkmofniojrvxj",
		"jzjozqfqiwyfadplibubtpcfxxfvjtbizxlmpaccjpihvnrtvfqtfiaztvfbednyemfmahbljdvykddawaujdksenm",
	}
	expect := 16
	chars := "fcxpzkzkqeyeijquyfixvlzjpzomujrqzxeoynjiflnmdrpdkrm"
	result := countCharacters(words, chars)
	fmt.Printf("Does result [%d] match expectation [%d]? [%v]\n", result, expect, result == expect)
}

func ex4() {
	words := []string{
		"dyiclysmffuhibgfvapygkorkqllqlvokosagyelotobicwcmebnpznjbirzrzsrtzjxhsfpiwyfhzyonmuabtlwin",
		"ndqeyhhcquplmznwslewjzuyfgklssvkqxmqjpwhrshycmvrb",
		"ulrrbpspyudncdlbkxkrqpivfftrggemkpyjl",
		"boygirdlggnh",
		"xmqohbyqwagkjzpyawsydmdaattthmuvjbzwpyopyafphx",
		"nulvimegcsiwvhwuiyednoxpugfeimnnyeoczuzxgxbqjvegcxeqnjbwnbvowastqhojepisusvsidhqmszbrnynkyop",
		"hiefuovybkpgzygprmndrkyspoiyapdwkxebgsmodhzpx",
		"juldqdzeskpffaoqcyyxiqqowsalqumddcufhouhrskozhlmobiwzxnhdkidr",
		"lnnvsdcrvzfmrvurucrzlfyigcycffpiuoo",
		"oxgaskztzroxuntiwlfyufddl",
		"tfspedteabxatkaypitjfkhkkigdwdkctqbczcugripkgcyfezpuklfqfcsccboarbfbjfrkxp",
		"qnagrpfzlyrouolqquytwnwnsqnmuzphne",
		"eeilfdaookieawrrbvtnqfzcricvhpiv",
		"sisvsjzyrbdsjcwwygdnxcjhzhsxhpceqz",
		"yhouqhjevqxtecomahbwoptzlkyvjexhzcbccusbjjdgcfzlkoqwiwue",
		"hwxxighzvceaplsycajkhynkhzkwkouszwaiuzqcleyflqrxgjsvlegvupzqijbornbfwpefhxekgpuvgiyeudhncv",
		"cpwcjwgbcquirnsazumgjjcltitmeyfaudbnbqhflvecjsupjmgwfbjo",
		"teyygdmmyadppuopvqdodaczob",
		"qaeowuwqsqffvibrtxnjnzvzuuonrkwpysyxvkijemmpdmtnqxwekbpfzs",
		"qqxpxpmemkldghbmbyxpkwgkaykaerhmwwjonrhcsubchs",
	}
	expect := 0
	chars := "usdruypficfbpfbivlrhutcgvyjenlxzeovdyjtgvvfdjzcmikjraspdfp"
	result := countCharacters(words, chars)

	fmt.Printf("Does result [%d] match expectation [%d]? [%v]\n", result, expect, result == expect)
}

/*
// Runtime: 64 ms, faster than 47.92% of Go online submissions for Find Words That Can Be Formed by Characters.
// Memory Usage: 6.7 MB, less than 16.67% of Go online submissions for Find Words That Can Be Formed by Characters.
func countCharacters(words []string, chars string) int {

	// step 1: create the hash map
	hash := make(map[string]int)

	for _, r := range chars {
		hash[string(r)]++
	}

	var result int

Loop:
	for _, word := range words {

		// NOTE: This is copying the map reference: t := hash
		// step 2: create copy of the hash map
		t := make(map[string]int)
		for i, v := range hash {
			t[i] = v
		}

		// step 3: range of unicode code points of string
		for _, r := range word {
			if v, ok := t[string(r)]; !ok || v < 0 {
				continue Loop
			}
			t[string(r)]--
			if t[string(r)] < 0 {
				continue Loop
			}
		}

		// fmt.Printf("%+v\n%+v\n%s\n\n", hash, t, word)

		result += len(word)
	}

	return result
}
*/

// Runtime: 12 ms, faster than 100.00% of Go online submissions for Find Words That Can Be Formed by Characters.
// Memory Usage: 6.1 MB, less than 97.92% of Go online submissions for Find Words That Can Be Formed by Characters.
func countCharacters(words []string, chars string) int {

	// hashtable using 0-25 character array mapped from ascii/unicode code point values
	charmap := [26]int{}
	for _, c := range []byte(chars) {
		charmap[c-'a']++
	}

	var result int

	// copy of hashtable populated for each new word
	mapcopy := [26]int{}

Loop:
	for _, word := range words {

		for i := range mapcopy {
			mapcopy[i] = charmap[i]
		}

		for _, c := range []byte(word) {

			if mapcopy[c-'a'] < 1 {
				continue Loop
			}

			mapcopy[c-'a']--
		}

		result += len(word)
	}

	return result
}
