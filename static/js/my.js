//在基本模板中已近加载了jq可以在任意的房使用jq
//有些方法需要在bootstrap.js 运行之后在运行
//所有页面都需要执行该js


(function (){
    //使用原生的js 来设置 导航栏高亮
let current_url = window.location.href;




jQuery.each(menus, function (i, dom) {
    //在html data-uri设置uri
    //根据网站设置高亮菜单
    let node_uri = $(dom).data('uri');
    if (current_url.indexOf(node_uri)) {
        $(dom).toggleClass('active')
    }
});

//设置面包屑高亮

}());




//去掉html tag 得到 plain string
function stripHtml(html)
{
   var tmp = document.createElement("DIV");
   tmp.innerHTML = html;
   return tmp.textContent || tmp.innerText || "";
}

//资讯列表页面就执行以下js

//The following solution works in Chrome, Firefox, Safari, IE9+ and also with iframes:


function humanTime(date) {

  var seconds = Math.floor((new Date() - date) / 1000);

  var interval = Math.floor(seconds / 31536000);

  if (interval > 1) {
    return interval + "年前";
  }
  interval = Math.floor(seconds / 2592000);
  if (interval > 1) {
    return interval + "月前";
  }
  interval = Math.floor(seconds / 86400);
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