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


function setUserAgent(window, userAgent) {
    if (window.navigator.userAgent != userAgent) {
        var userAgentProp = { get: function () { return userAgent; } };
        try {
            Object.defineProperty(window.navigator, 'userAgent', userAgentProp);
        } catch (e) {
            window.navigator = Object.create(navigator, {
                userAgent: userAgentProp
            });
        }
    }
}