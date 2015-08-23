<!DOCTYPE html>
<html>
  <head>
    <title>Palantìr</title>

    <link rel="stylesheet" href="//maxcdn.bootstrapcdn.com/bootstrap/3.3.2/css/bootstrap.min.css">
    <link rel="stylesheet" href="/static/bootstrap/themes/flatly/bootstrap.min.css">

    <script src="/static/jquery/jquery-2.1.3.min.js"></script>
    <script src="/static/bootstrap/js/bootstrap.min.js"></script>
  </head>

  <body>
    <nav class="navbar navbar-default">
      <div class="container">
        <!-- Brand and toggle get grouped for better mobile display -->
        <div class="navbar-header">
          <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1">
            <span class="sr-only">Toggle navigation</span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </button>
          <a class="navbar-brand" href="/">Palantìr</a>
        </div>

        <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">

          <ul class="nav navbar-nav navbar-right">
            {{ if .User }}
            <li class="dropdown">
                    <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">{{ .User.Email }}<span class="caret"></span></a>

              <ul class="dropdown-menu" role="menu">
                <li><a href="javascript:void(0)" data-toggle="modal" data-target="#user-settings-modal">User Settings</a></li>

                <li class="divider"></li>

                <li><a href="/logout">Logout</a></li>
              </ul>
            </li>
            {{else}}
            <li>
                    <a href="/login">Login</a>
            </li>
            {{end}}
          </ul>
        </div><!-- /.navbar-collapse -->
      </div><!-- /.container-fluid -->
    </nav>

    {{template "content" .}}
  </body>
</html>
