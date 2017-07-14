//在基本模板中已近加载了jq可以在任意的房使用jq
//有些方法需要在bootstrap.js 运行之后在运行
//所有页面都需要执行该js




//去掉html tag 得到 plain string
function stripHtml(html) {
  var tmp = document.createElement("DIV");
  tmp.innerHTML = html;
  return tmp.textContent || tmp.innerText || "";
}

//资讯列表页面就执行以下js

//The following solution works in Chrome, Firefox, Safari, IE9+ and also with iframes:




//格式化是时间

function changeTimeTagToHumanTime(dateString) {
  var string = dateString.substring(0, 19);
  var date = new Date(string);
  var seconds = Math.floor((new Date() - date) / 1000);

  interval = Math.floor(seconds / 3600 / 24 / 30);
  if (interval > 1) {
    return dateString.substring(0, 10);
  }
  interval = Math.floor(seconds / 3600 / 24);
  if (interval > 1) {
    return interval + "天前";
  }
  interval = Math.floor(seconds / 3600);
  if (interval > 1) {
    return interval + "小时前";
  }
  interval = Math.floor(seconds / 60);
  if (interval > 1) {
    return interval + "分前";
  }
  return Math.floor(seconds) + "秒前";
}

function changeTimeTagToMonthDate(dateString) {
  return dateString.substring(5, 10);
}

var current_url = window.location.href;

//js 转换实践格式
$('span.time.human-date').each(function (idx, ele) {
  var dataString = $(ele).data('time');
  var res = changeTimeTagToHumanTime(dataString);
  $(ele).text(res);
});
$('span.time.month-date').each(function (idx, ele) {
  var dataString = $(ele).data('time');
  var res = changeTimeTagToMonthDate(dataString);
  $(ele).text(res);
});
$(function () {
  //设置导航菜单高亮
  $('.nav-link').each(function (i, dom) {
    //在html data-uri设置uri
    //根据网站设置高亮菜单
    let node_uri = $(dom).attr('href');
    if (current_url.indexOf(node_uri) > 0) {
      $(dom).parent('li.nav-item').toggleClass('active')
    }
  });
  //设置面包屑高亮
});