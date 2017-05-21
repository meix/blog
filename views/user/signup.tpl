<style type="text/css">

  a { text-decoration: none; }
  a:hover { text-decoration: none; }

  
  #signup {
    border-bottom: 1px solid #ddd;
    border-top: 1px solid #fff;
    padding: 200px 0;
    position: relative;
  }

</style>

<div id="signup" class="inner">
  <form class="form-horizontal" action="/signup" method="post" id="new-user"> 
    <div class="form-group">
      <label for="inputEmail3" class="col-sm-4 control-label"> 邮箱 </label>
      <div class="col-sm-6">
        <input class="form-control" id="email" name="email" placeholder="请输入邮箱">
      </div>
    </div>

    <div class="form-group">
      <label for="inputEmail3" class="col-sm-4 control-label"> 昵称 </label>
      <div class="col-sm-6">
        <input class="form-control" id="name" name="name" placeholder="请输入用户名">
      </div>
    </div>

    <div class="form-group">
      <label for="inputPassword3" class="col-sm-4 control-label"> 密码 </label>
      <div class="col-sm-6">
        <input type="password" class="form-control" id="password" name="password" placeholder="请输入密码">
      </div>
    </div>

    <div class="form-group">
      <label for="inputPassword3" class="col-sm-4 control-label"> 确认密码 </label>
      <div class="col-sm-6">
        <input type="password" class="form-control" id="password_confirmation" name="password_confirmation" placeholder="确认密码">
      </div>
    </div>

    <div class="form-group">
      <div class="col-sm-offset-4 col-sm-6">
        <button type="submit" class="btn btn-success"> 注册 </button>
      </div>
    </div>
  </form>
</div>
