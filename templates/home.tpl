{{define "content"}}
<div class="container">
  <div class="row">
    {{ if .User }}
    <table class="table table-striped">
        <thead>
            <tr><th>ID</th><th>Application</th><th>Date</th><th>Try duration</th></tr>
        </thead>
        <tbody>
        {{range .Registrations}}
            <tr>
                <td>{{.ID}}</td>
                <td>{{.App}}</td>
                <td>{{.Date | formatDate}}</td>
                <td><span class="label {{if .HasExpired}}label-default{{else}}label-success{{end}}">{{.TryDuration}}</span></td>
            </tr>
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
