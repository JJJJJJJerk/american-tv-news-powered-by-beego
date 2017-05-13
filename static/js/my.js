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
let htText = $('body > div > div.content-wrapper > section.content-header > h1').text()

if (htText == "美剧资讯") {
        function loadResults() {
            var result = "";
            for (var i = 0; i < 10; i++) {
                result += "<li>Result " + i + "</li>";
            }
            $.ajax({
                url: "/echo/html/",
                type: "post",
                data: {
                    html: result,
                    delay: 3
                },
                beforeSend: function(xhr) {
                    $("#results").after($("<li class='loading'>Loading...</li>").fadeIn('slow')).data("loading", true);
                },
                success: function(data) {
                    var $results = $("#results");
                    $(".loading").fadeOut('fast', function() {
                        $(this).remove();
                    });
                    var $data = $(data);
                    $data.hide();
                    $results.append($data);
                    $data.fadeIn();
                    $results.removeData("loading");
                }
            });
        };

        $(function() {
            loadResults();

            $("section.content").scroll(function() {
                var $this = $(this);
                var $results = $("#results");

                if (!$results.data("loading")) {

                    if ($this.scrollTop() + $this.height() == $results.height()) {
                        loadResults();
                    }
                }
            });
        });

}





