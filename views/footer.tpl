<script>
$(".header-search").submit(function () {
    var term = encodeURIComponent($('input[name="keyword"]').val().trim());
    if(term === "") return false;
    location.pathname = "/s/" + term;
    return false;
})
</script>
<script data-no-instant>
InstantClick.init("mousedown")
</script>
<div class="footer d-flex justify-content-center">
    <div class="container-sm">
        <div class="d-flex justify-content-center pb-2">
            <a href="https://y.qq.com/" target="_blank" class="mx-2">QQ音乐</a>
            <a href="https://music.163.com/" target="_blank" class="mx-2">网易云音乐</a>
            <a href="https://www.kuwo.cn/" target="_blank" class="mx-2">酷我音乐</a>
            <a href="https://www.baidu.com/" target="_blank" class="mx-2">百度一下</a>
            <a href="https://github.com/liumingye/xiageba" target="_blank" class="mx-2">
                <img src="https://badgen.net/github/stars/liumingye/xiageba">
            </a>
        </div>
        <div class="text-center fs-7">
            Copyright © 2022-2024 xiageba
            <div class="text-black-50">
                本站所有数据均系网友搜集自互联网后分享.<br>
                音频版权来自各网站,本站不提供任何音频存储和贩卖服务,旨在音乐交流分享.<br>
                如涉及侵权请在反馈建议里,联系我们删除.
            </div>
        </div>
    </div>
</div>