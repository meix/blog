<!DOCTYPE html>
<html>
<head>
  <title>Lin Li</title>
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css">
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-theme.min.css">

  <script src="/static/js/jquery.min.js"></script>

  <style type="text/css">
    *,body {
      margin: 0px;
      padding: 0px;
    }

    body {
      margin: 0px;
      font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
      font-size: 14px;
      line-height: 20px;
      overflow: auto;
    }

    a { text-decoration: none; }

    ul, ol {
      margin-top: 0;
      margin-bottom: 10px;
    }

    .main {
      margin-left: 300px;
      padding: 0 70px 0 40px;
      min-height: 100%;
      background: #fff;
      left: 0;
      position: absolute;
      right: 0;
    }
  </style>

</head>
<body>

  <div class="container-fluid">
    {{template "layout/sidebar.tpl" .}}
    <div class="main">
      {{.LayoutContent}}

      {{template "layout/footer.tpl" .}}
    </div>
  </div>
  
  <script src="/static/js/bootstrap.min.js"></script>
  <script src="/static/js/jquery.validate.js"></script>
  <!-- <script src="/static/js/user.js"></script> -->
  <script src="/static/js/auth.js"></script>
</body>
</html>