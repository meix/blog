$(document).ready(function() {
    
  if ($('#create-article').length > 0)   

  var editor = new MediumEditor('.editable', {
    placeholder: {  
      text: '输入博客内容',
      hideOnClick: true
    }
  });


  // 数据验证
  $("#new-article").validate({
    onkeyup: false,
    success: "valid",
    errorPlacement: function (error, element) {
    },
    rules: {
      'title': {
        required: true
      },
      'keywords': {
        required: true
      },
      'summary': {
        required: true
      },
      'content': {
        required: true
      }
    },

    messages: {
      'title': {
        required: '文章标题不能为空'
      },
      'keywords': {
        required: '文章关键字不能为空'
      },
      'summary': {
        required: '文章简介不能为空'
      },
      'content': {
        required: '文章内容不能为空'
      }
    },

    submitHandler: function (form) {
      form.submit();
    }
  });


});