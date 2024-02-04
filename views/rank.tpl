<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.Title}} - 下歌吧 - 全网音乐 高品质MP3 在线免费下载 免费播放</title>
    <meta name="keywords" content="{{.Title}},下歌吧,下歌网,音乐下载,无损音乐,歌曲下载,高品质音乐,歌曲搜索,音乐免费下载,MP3下载,收费音乐免费下载,付费音乐免费下载,在线mp3下载网站"/>
    <meta name="description" content="{{.Title}} - 下歌吧在线音乐搜索，可以在线免费下载全网MP3付费歌曲、流行音乐、经典老歌等。曲库完整，更新迅速，试听流畅，支持高品质|无损音质">
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
            <input type="text" name="keyword" value="{{.Keyword}}" placeholder="输入歌名、歌手" class="ipt flex-grow-1">
            <button type="submit" class="iconfont icon-search"></button>
        </form>
        
        <div class="body-box mt-3 p-3">
            <div class="d-flex justify-content-between align-items-center">
                <div class="fs-10 fw-bold">昨日搜索排行</div>
            </div>
            <div class="p-3">
                <div class="search-item pb-1 row bg-body-secondary">
                    <div class="col-7 fw-bold">关键词</div>
                    <div class="col-5 fw-bold">搜索次数</div>
                </div>
                {{range .Data}}
                <div class="search-item row">
                    <div class="col-7">
                        <a href="/s/{{.SearchTerm}}" class="link-blue">{{.SearchTerm}}</a>
                    </div>
                    <div class="col-5">{{.SearchCount}}</div>
                </div>
                {{end}}
            </div>
            {{template "paginator.tpl" .}}
        </div>
    </div>
</div>

{{template "footer.tpl" .}}
</body>
</html>