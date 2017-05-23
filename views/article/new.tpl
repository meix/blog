<style type="text/css">

  a { text-decoration: none; }
  a:hover { text-decoration: none; }

  #create-article {
    border-bottom: 1px solid #ddd;
    border-top: 1px solid #fff;
    padding: 20px 0;
    position: relative;
  }

  .signup-link {margin-left: 60px}

  .editable {
    border: 1px solid #ddd;
    min-height: 500px;
    padding: 5px; 
  }

  textarea { resize: none; }

  #content { min-height: 500px; }

</style>

<div id="create-article" class="inner">
  <form class="form-horizontal" action="/article" method="post" id="new-article"> 
    
    <div class="form-group">
      <div class="col-sm-8">
        <input class="form-control" id="title" name="title" placeholder="请输入标题">
      </div>
      <div class="col-sm-4">
        <input class="form-control" id="keywords" name="keywords" placeholder="请输入关键字">
      </div>
    </div>

    <div class="form-group">
      <div class="col-sm-12">
        <textarea class="form-control" id="summary" name="summary" placeholder="请简短介绍文章内容 (20~50字)"></textarea>
      </div>
    </div>

    <div class="form-group">
      <div class="col-sm-12">    
        <textarea class="form-control" id="content" name="content" placeholder="请输入文章内容"></textarea>
      </div>
    </div>

<!--     <div class="form-group">
      <div class="col-sm-12">      
          <div class="editable"></div>          
      </div>
    </div> -->

    <div class="form-group">
      <div class="col-sm-12">
        <button type="submit" class="btn btn-success"> 创建 </button>
        <a type="submit" class="btn btn-default" href="/article/personal"> 返回 </a>
      </div>
    </div>
  </form>
</div>


