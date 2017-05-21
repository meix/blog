<style type="text/css">

  a { text-decoration: none; }
  a:hover { text-decoration: none; }

  
  #login {
    border-bottom: 1px solid #ddd;
    border-top: 1px solid #fff;
    padding: 200px 0;
    position: relative;
  }

  .signup-link {margin-left: 60px}
</style>

<div id="login" class="inner">
  <form class="form-horizontal" action="/login" method="post" id="login-user"> 
    
    <div class="form-group">
      <label for="inputEmail3" class="col-sm-4 control-label"> 邮箱 </label>
      <div class="col-sm-6">
        <input class="form-control" id="email" name="email" placeholder="请输入邮箱">
      </div>
    </div>
  
    <div class="form-group">
      <label for="inputPassword3" class="col-sm-4 control-label"> 密码 </label>
      <div class="col-sm-6">
        <input type="password" class="form-control" id="password" name="password" placeholder="请输入密码">
      </div>
    </div>

    <div class="form-group">
      <div class="col-sm-offset-4 col-sm-6">
        <button type="submit" class="btn btn-success"> 登陆 </button>
        <span class="signup-link"> 没有账号? <a href="/signup"> 立即注册 →</a></span>
      </div>
    </div>
  </form>

</div>