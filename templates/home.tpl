{{define "content"}}
<div class="container">
  <div class="row">
    {{ if .User }}
    <table class="table table-striped">
        <thead>
            <tr><th>ID</th><th>application</th><th>date</th></tr>
        </thead>
        <tbody>
        {{range .Registrations}}
            <tr><td>{{.ID}}</td><td>{{.App}}</td><td>{{.Date}}</td></tr>
        {{end}}
        </tbody>
    </table>
    {{else}}
    <div class="jumbotron">
      <h1>Welcome</h1>
      <p>Connect to your google account.</p>
    </div>
    {{end}}
  </div>
</div>
{{end}}
