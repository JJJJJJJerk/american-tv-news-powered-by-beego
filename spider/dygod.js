// Here You can type your custom JavaScript...
//base64
var Base64 = { _keyStr: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=", encode: function (e) { var t = ""; var n, r, i, s, o, u, a; var f = 0; e = Base64._utf8_encode(e); while (f < e.length) { n = e.charCodeAt(f++); r = e.charCodeAt(f++); i = e.charCodeAt(f++); s = n >> 2; o = (n & 3) << 4 | r >> 4; u = (r & 15) << 2 | i >> 6; a = i & 63; if (isNaN(r)) { u = a = 64 } else if (isNaN(i)) { a = 64 } t = t + this._keyStr.charAt(s) + this._keyStr.charAt(o) + this._keyStr.charAt(u) + this._keyStr.charAt(a) } return t }, decode: function (e) { var t = ""; var n, r, i; var s, o, u, a; var f = 0; e = e.replace(/[^A-Za-z0-9+/=]/g, ""); while (f < e.length) { s = this._keyStr.indexOf(e.charAt(f++)); o = this._keyStr.indexOf(e.charAt(f++)); u = this._keyStr.indexOf(e.charAt(f++)); a = this._keyStr.indexOf(e.charAt(f++)); n = s << 2 | o >> 4; r = (o & 15) << 4 | u >> 2; i = (u & 3) << 6 | a; t = t + String.fromCharCode(n); if (u != 64) { t = t + String.fromCharCode(r) } if (a != 64) { t = t + String.fromCharCode(i) } } t = Base64._utf8_decode(t); return t }, _utf8_encode: function (e) { e = e.replace(/rn/g, "n"); var t = ""; for (var n = 0; n < e.length; n++) { var r = e.charCodeAt(n); if (r < 128) { t += String.fromCharCode(r) } else if (r > 127 && r < 2048) { t += String.fromCharCode(r >> 6 | 192); t += String.fromCharCode(r & 63 | 128) } else { t += String.fromCharCode(r >> 12 | 224); t += String.fromCharCode(r >> 6 & 63 | 128); t += String.fromCharCode(r & 63 | 128) } } return t }, _utf8_decode: function (e) { var t = ""; var n = 0; var r = c1 = c2 = 0; while (n < e.length) { r = e.charCodeAt(n); if (r < 128) { t += String.fromCharCode(r); n++ } else if (r > 191 && r < 224) { c2 = e.charCodeAt(n + 1); t += String.fromCharCode((r & 31) << 6 | c2 & 63); n += 2 } else { c2 = e.charCodeAt(n + 1); c3 = e.charCodeAt(n + 2); t += String.fromCharCode((r & 15) << 12 | (c2 & 63) << 6 | c3 & 63); n += 3 } } return t } }



//创建迅雷地址的小米链接
function generateXiaomiLink(url, name) {
	return "https://d.miwifi.com/d2r/?url=" + Base64.encode(url) + "&src=demo" + "&name=" + encodeURIComponent(name);
}
//解析详情页面到json
function getDygodMeijuDetailJson() {
	//设置targe rel
	$('a').each(function (index, node) {
		$(node).attr('target', '_blank').attr('rel', 'nofollow');
	});
	//获取标签
	var h1 = $('h1').text();
	$('#Zoom > p:nth-child(27) > font > strong>font').text('【迅雷|小米】')
	//添加response image
	$('img').addClass('img-responsive');
	//删除垃圾dom
	$('#Zoom > div.play-list-box').remove();
	$('center').remove();
	$('hr').remove();
	$('script').remove();

	//遍历下载地址
	$('#Zoom > table> tbody > tr > td > anchor > a').each(function (index, value) {
		//根据地址生成迅雷地址
		var title = $(value).text();
		//这个js是迅雷页面自带的 还有一种方法可以生成按标签
		var href = ThunderEncode(title); 
		//强制鼓励垃圾字符
		var filted_node_txt = $(value).text().replace('[电影天堂www.dy2018.com]', '');

		//   /[\u4e00-\u9fa5]+\w+\.\w+\b/ig 中文
		var pattern = /[\u4e00-\u9fa5]+.+\b/ig;//正则表达式
		if (filted_node_txt) {
			var good_title = filted_node_txt.match(pattern)[0]
			var xiaoMiFuncClick = "openDownloadWindow('" + href + "','" + good_title + "')";
			$(value).text(good_title);

			thunerAImage = "<a target='_blank' rel='nofollow' title='迅雷下载' href='" + href + "'><img width='22px' alt='迅雷下载' src='/web-icon/thunder.png'></a>";
			$(value).append(thunerAImage);
			var xiaoMiDownloadLink = generateXiaomiLink(href, good_title)
			$(value).append('<a target="_blank" rel="nofollow" title="小米下载" href="' + xiaoMiDownloadLink + '"><img alt="小米下载" src="/web-icon/mi-icon.png" width="18px"></a>')
		}

		//过滤迅雷
		$(value).attr('href', href);//生成迅雷地址
		$(value).removeAttr('onclick');//
		$(value).removeAttr('target');//
		$(value).removeAttr('thundertype');//
		$(value).removeAttr('thunderrestitle');//
		$(value).removeAttr('oncontextmenu');//
		$(value).removeAttr('bqloxkcv');//
		$(value).removeAttr('mritqcam');//
		$(value).removeAttr('cdedkblh');//
		$(value).attr('target', '_blank');//生成迅雷地址

		//todo 匹配掉文件名字
	});
	var body = $('#Zoom').html();
	var json = { title: h1, content: body };
	return json;
}



//取电影天堂美剧列表信息
function GetDygodMeijuList(){

		var list = new Array();
		var host = location.host;
		$('ul > table > tbody > tr:nth-child(2) > td:nth-child(2) > b > a').each(function(index,item){
			var name = $(item).attr('title');
			var href = 'http://'+host + $(item).attr('href');
			list.push({name:name,href:href})
		});
		return list;
}