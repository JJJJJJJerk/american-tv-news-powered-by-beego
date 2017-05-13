//在基本模板中已近加载了jq可以在任意的房使用jq
//有些方法需要在bootstrap.js 运行之后在运行
//所有页面都需要执行该js



//设置left side bar 高亮
let current_url = window.location.href;
let menus = $('body > div > aside.main-sidebar > section > ul > li').not('.header');
jQuery.each(menus, function (i, dom) {
    //在html data-uri设置uri
    //根据网站设置高亮菜单
    let node_uri = $(dom).data('uri');
    if (current_url.indexOf(node_uri)) {
        $(dom).toggleClass('active')
    }
});

//设置left side bar 二级菜单高亮
let sub_menus = $('body > div > aside.main-sidebar > section > ul > li > ul > li');
jQuery.each(sub_menus, function (i, sub_dom) {
    //在html data-uri设置uri
    //根据网站设置高亮菜单
    let sub_uri = $(sub_dom).find('a').attr('href');
    if (current_url.indexOf(sub_uri)) {
        $(sub_dom).toggleClass('active')
    }
});
//设置面包屑高亮

//资讯列表页面就执行以下js
let htText = $('span#page-ID').data('name')

if (htText == "article-index-page") {

    $('section.content').dropload({
    scrollArea : window,
    loadDownFn : function(me){
        $.ajax({
            type: 'GET',
            url: 'json/more.json',
            dataType: 'json',
            success: function(data){
                alert(data);
                // 每次数据加载完，必须重置
                me.resetload();
            },
            error: function(xhr, type){
                alert('Ajax error!');
                // 即使加载出错，也得重置
                me.resetload();
            }
        });
    }
});


}





