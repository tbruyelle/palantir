{{define "content"}}
<div class="container">
  <div class="row">
    <table class="table table-striped">
        <thead>
            <tr><th>Name</th><th>Try duration</th></tr>
        </thead>
        <tbody>
        {{range .Apps}}
            <tr>
                <td>{{.Name}}</td>
                <td>{{.TryDuration}}</td>
            </tr>
        {{end}}
        </tbody>
    </table>
  </div>
</div>
{{end}}
