$(document).ready(function() {

  // 注册
  if ( $('#signup').length > 0 ){
    
    // 数据验证
    $("#new-user").validate({
      onkeyup: false,
      success: "valid",
      errorPlacement: function (error, element) {
      },
      rules: {
        'email': {
          required: true
        },
        'password': {
          required: true,
          minlength: 6,
          maxlength: 12
        },
        'password_confirmation': {
          required: true,
          minlength: 6,
          maxlength: 12
        },
        'name': {
          required: true
        }
      },

      messages: {
        'email': {
          required: '邮箱不能为空'
        },
        'password': {
          required: '密码不能为空',
          minlength: '请输入至少6位字符的密码',
          maxlength: '密码不能超过12位字符'
        },
        'password_confirmation': {
          required: '确认密码不能为空',
          minlength: '请输入至少6位字符的密码',
          maxlength: '密码不能超过12位字符'
        },
        'name': {
          required: '用户名不能为空'
        }
      },

      submitHandler: function (form) {
        form.submit();
      }
    });
  }

    // 登陆
  if ( $('#login').length > 0 ){
    
    // 数据验证
    $("#login-user").validate({
      onkeyup: false,
      success: "valid",
      errorPlacement: function (error, element) {
      },
      rules: {
        'email': {
          required: true
        },
        'password': {
          required: true,
          minlength: 6,
          maxlength: 12
        }
      },

      messages: {
        'email': {
          required: '邮箱不能为空'
        },
        'password': {
          required: '密码不能为空',
          minlength: '请输入至少6位字符的密码',
          maxlength: '密码不能超过12位字符'
        }
      },

      submitHandler: function (form) {
        form.submit();
      }
    });
  }

});