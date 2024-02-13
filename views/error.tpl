<html>
<head>
<meta charset="utf-8">
<style type="text/css">@import 'https://fonts.googleapis.com/css?family=Inconsolata';html{min-height:100%}body{margin:0px;box-sizing:border-box;height:100%;background-color:#000000;background-image:radial-gradient(#333,#111);font-family:'Inconsolata',Helvetica,sans-serif;font-size:1.5rem;color:rgba(128,128,128,0.8);text-shadow:0 0 1ex rgba(51,51,51,1),0 0 2px rgba(255,255,255,0.8)}.overlay{pointer-events:none;position:absolute;width:100%;height:100%;background:repeating-linear-gradient(180deg,rgba(0,0,0,0) 0,rgba(0,0,0,0.3) 50%,rgba(0,0,0,0) 100%);background-size:auto 4px;z-index:99}.overlay::before{content:"";pointer-events:none;position:absolute;display:block;top:0;left:0;width:100%;height:100%;background-image:linear-gradient(0deg,transparent 0%,rgba(128,128,128,0.2) 2%,rgba(128,128,128,0.8) 3%,rgba(128,128,128,0.2) 3%,transparent 100%);background-repeat:no-repeat;animation:scan 7.5s linear 0s infinite}@keyframes scan{0%{background-position:0 -100vh}35%,100%{background-position:0 100vh}}.terminal{box-sizing:inherit;position:absolute;height:100%;width:1000px;max-width:100%;padding:4rem;text-transform:uppercase}.output{color:rgba(128,128,128,0.8);text-shadow:0 0 1ex rgba(51,51,51,0.8),0 0 2px rgba(255,255,255,0.8)}.output::before{content:"> "}a{color:#fff;text-decoration:none}a::before{content:"["}a::after{content:"]"}.errorcode{color:white}</style>
</head>
<body>
<div class="overlay"></div>
<div class="terminal">
<h1>{{.Title}} <span class="errorcode">{{.Error}}</span></h1>
<p class="output">{{.Content}}</p>
<p class="output">Please try <a href="javascript:location.reload()">Refresh</a> or <a href="/">Go Home</a></p>
<p class="output">Good luck</p>
<p class="output">beego {{.BeegoVersion}} (beego framework)</p>
<p class="output">golang version: {{.GoVersion}}</p>
</div>
</body>
</html>