<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>下歌吧 - 全网音乐 高品质MP3 在线免费下载 免费播放</title>
    <meta name="keywords" content="下歌吧,下歌网,音乐下载,无损音乐,歌曲下载,高品质音乐,歌曲搜索,音乐免费下载,MP3下载,收费音乐免费下载,付费音乐免费下载,在线mp3下载网站"/>
    <meta name="description" content="下歌吧在线音乐搜索，可以在线免费下载全网MP3付费歌曲、流行音乐、经典老歌等。曲库完整，更新迅速，试听流畅，支持高品质|无损音质">
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
        <div class="d-flex justify-content-between">
          <div class="d-flex align-items-center">
            <i class="iconfont icon-refresh fs-11 me-1"></i>
            <span class="fs-10 fw-bold">最新搜索</span>
          </div>
          <a href="/history">更多 &gt;</a>
        </div>
        <div class="d-flex flex-wrap">
          {{range .LatestSearchTerms}}
          <a class="key-tag key-tag-bg{{random 1 7}}" href="/s/{{.SearchTerm}}">{{.SearchTerm}}<span>{{.SecondsAgo}}</span></a>
          {{end}}
        </div>
      </div>

      <div class="row w-100 mt-3 g-0 box-sizing-border">
        <div class="col-lg-6 pe-lg-2 box-sizing-border">
          <div class="body-box p-3">
            <div class="d-flex justify-content-between align-items-center">
                <div class="d-flex align-items-center">
                  <i class="iconfont icon-rank fs-11 me-1"></i>
                  <span class="fs-10 fw-bold">昨日搜索排行</span>
                </div>
              <a href="/r/yesterday">更多 &gt;</a>
            </div>
            <div class="row rank-th">
              <div class="col-5">排名</div>
              <div class="col-7">关键词</div>
            </div>
            <div class="row rank-td">
              <div class="col-5">
                <i class="rank-no-0">1</i>
              </div>
              <div class="col-7">
                <a href="/s/周杰伦" class="song-title">周杰伦</a>
              </div>
            </div>
          </div>
        </div>
        <div class="col-lg-6 ps-lg-2 box-sizing-border mt-max-sm-3">
          <div class="body-box p-3">
            <div class="d-flex justify-content-between align-items-center">
              <div class="d-flex align-items-center">
                <i class="iconfont icon-rank fs-11 me-1"></i>
                <span class="fs-10 fw-bold">本周搜索排行</span>
              </div>
            </div>
            <div class="row rank-th">
              <div class="col-5">排名</div>
              <div class="col-7">关键词</div>
            </div>
            <div class="row rank-td">
              <div class="col-5">
                <i class="rank-no-0">1</i>
              </div>
              <div class="col-7">
                <a href="/s/周杰伦" class="song-title">周杰伦</a>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="row w-100 mt-3 g-0 box-sizing-border">
        <div class="col-lg-6 pe-lg-2 box-sizing-border">
          <div class="body-box p-3">
            <div class="d-flex justify-content-between align-items-center">
              <div class="d-flex align-items-center">
                <i class="iconfont icon-rank fs-11 me-1"></i>
                <span class="fs-10 fw-bold">本月搜索排行</span>
              </div>
            </div>
            <div class="row rank-th">
              <div class="col-5">排名</div>
              <div class="col-7">关键词</div>
            </div>
            <div class="row rank-td">
              <div class="col-5">
                <i class="rank-no-0">1</i>
              </div>
              <div class="col-7">
                <a href="/s/周杰伦" class="song-title">周杰伦</a>
              </div>
            </div>
          </div>
        </div>
        <div class="col-lg-6 ps-lg-2 box-sizing-border mt-max-sm-3">
          <div class="body-box p-3">
            <div class="d-flex justify-content-between align-items-center">
              <div class="d-flex align-items-center">
                <i class="iconfont icon-rank fs-11 me-1"></i>
                <span class="fs-10 fw-bold">上月搜索排行</span>
              </div>
            </div>
            <div class="row rank-th">
              <div class="col-5">排名</div>
              <div class="col-7">关键词</div>
            </div>
            <div class="row rank-td">
              <div class="col-5">
                <i class="rank-no-0">1</i>
              </div>
              <div class="col-7">
                <a href="/s/周杰伦" class="song-title">周杰伦</a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>

{{template "footer.tpl" .}}

</body>
</html>