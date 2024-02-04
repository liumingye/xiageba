<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.Music.Name}} - {{.Music.Singer}} - mp3免费在线下载播放 - 下歌吧 - 全网音乐 高品质MP3 在线免费下载 免费播放</title>
    <meta name="keywords" content="{{.Music.Name}},{{.Music.Singer}},下歌吧,下歌网,音乐下载,无损音乐,歌曲下载,高品质音乐,歌曲搜索,音乐免费下载,MP3下载,收费音乐免费下载,付费音乐免费下载,在线mp3下载网站"/>
    <meta name="description" content="{{.Music.Name}} - {{.Music.Singer}}.mp3免费在线下载播放,下歌吧在线音乐搜索，可以在线免费下载全网MP3付费歌曲、流行音乐、经典老歌等。曲库完整，更新迅速，试听流畅，支持高品质|无损音质">
    {{template "header.tpl" .}}
</head>

<nav class="navbar navbar-expand navbar-light">
    <div class="container-sm">
        <a class="navbar-brand ms-3" href="/">
            <strong>下歌吧</strong>
            <span class="navbar-text fs-10 ms-2">xiageba</span>
        </a>
    </div>
</nav>

<div class="w-100 d-flex justify-content-center">
    <div class="container-sm d-flex pb-3 flex-wrap">
        <form class="header-search d-flex mt-3 align-items-center" method="get">
            <input type="text" name="keyword" placeholder="输入歌名、歌手" class="ipt flex-grow-1">
            <button type="submit" class="iconfont icon-search"></button>
        </form>
        <div class="body-box mt-3 p-3">
            <div class="d-flex">
                {{if .Music.Pic}}
                <div class="me-2">
                    <img class="thumb-cover" src="{{.Music.Pic}}">
                </div>
                {{end}}
                <div class="d-flex flex-column justify-content-center">
                    <div class="fs-10 pb-2">{{.Music.Name}}</div>
                    <div class="">{{.Music.Singer}}</div>
                </div>
            </div>
            <div class="border-top mt-3 pt-3">
                <input type="text" id="oInput" style="display:none;">
                {{range .Links}}
                <div class="down-item">
                    <div>
                        <span class="me-2">下载链接:</span>
                        <a href="{{.url}}" target="_blank">{{.url}}</a>
                    </div>
                    <div>
                        <span class="me-2">提取码:</span>
                        <span onclick="copyText('{{.code}}')">{{.code}}</span>
                    </div>
                </div>
                {{end}}
            </div>
            {{if .Tags}}
            <div class="d-flex align-items-center border-top mt-3 pt-3">
                <span class="fs-10 fw-bold">标签</span>
            </div>
            <div class="d-flex flex-wrap">
                {{range .Tags}}
                    <a class="key-tag key-tag-bg{{.rand}}" href="/t/{{.name}}">{{.name}}</a>
                {{end}}
            </div>
            {{end}}
            <div class="border-top mt-3 pt-3">{{range .Lryic}}<p>{{.}}</p>{{end}}</div>
        </div>
    </div>
</div>

<script type="text/javascript">
function showToast (opt = {}) {
    let obj = $('#cToast');
    let titleObj = $('#cToastTitle');
    const type = opt.type || 'secondary'
    titleObj.html(opt.title || '提示');
    titleObj.attr('class', `me-auto text-${type}`);
    $('#cToastBody').html(opt.msg || '');
    opt.autohide = opt.autohide || true;
    opt.delay = opt.delay || 2000;
    obj.toast(opt);
    obj.toast('show');
};
function showToastSuccess (msg, delay) {
    showToast({
        title: '提示',
        msg: msg,
        type: 'success',
        delay
    });
};
function copyText(copyText) {
    // 获得到要复制的文本内容
    // 判断是否为ie浏览器，此方法只对IE浏览器有用
    if (window.clipboardData) {
        // 清除原有剪切板的数据
        window.clipboardData.clearData();
        // 将内容复制到剪切板
        window.clipboardData.setData("Text", copyText);
        // 其它浏览器,用别的方法
    } else {
        // 创建一个input对象
        let oInput = document.getElementById('oInput');
        oInput.style.display = '';
        // 赋值
        oInput.value = copyText;
        // 选择对象
        oInput.select();
        // 执行浏览器复制命令
        document.execCommand("Copy");
        oInput.style.display = 'none';
    }
    showToastSuccess('提取码复制完成', 2000);
}
</script>

<div class="position-fixed top-0 start-50 translate-middle-x p-3">
    <div id="cToast" class="toast" role="alert" aria-live="assertive" aria-atomic="true">
        <div class="toast-header">
            <strong class="me-auto" id="cToastTitle"></strong>
            <button type="button" class="btn-close" data-bs-dismiss="toast" aria-label="Close"></button>
        </div>
        <div class="toast-body" id="cToastBody"></div>
    </div>
</div>

{{template "footer.tpl" .}}

</body>
</html>