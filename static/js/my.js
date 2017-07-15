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



var current_url = window.location.href;


$(function () {
  //设置导航菜单高亮
  $('.nav-link').each(function (i, dom) {
    //在html data-uri设置uri
    //根据网站设置高亮菜单
    var node_uri = $(dom).attr('href');
    if (current_url.indexOf(node_uri) > 0) {
      $(dom).parent('li.nav-item').toggleClass('active')
    }
  });
  //设置面包屑高亮
});