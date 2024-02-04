{{if gt .paginator.PageNums 1}}
<nav aria-label="..." class="mt-2 d-flex justify-content-center">
    <ul class="pagination pagination-sm">
    {{if .paginator.HasPrev}}
        <li class="page-item"><a class="page-link" href="/s/{{.Keyword}}/{{.paginator.PageLinkPrev}}">上页</a></li>
    {{else}}
        <li class="page-item disabled"><a class="page-link">上页</a></li>
    {{end}}
    {{range $index, $page := .paginator.Pages}}
        <li class="page-item{{if $.paginator.IsActive .}} active{{end}}">
            {{if ne $page 0}}
            <a class="page-link" href="/s/{{$.Keyword}}/{{$.paginator.PageLink $page}}">{{$page}}</a>
            {{else}}
            <a class="page-link">...</a>
            {{end}}
        </li>
    {{end}}
    {{if .paginator.HasNext}}
        <li class="page-item"><a class="page-link" href="/s/{{.Keyword}}/{{.paginator.PageLinkNext}}">下页</a></li>
    {{else}}
        <li class="page-item disabled"><a class="page-link">下页</a></li>
    {{end}}
    </ul>
</nav>
{{end}}