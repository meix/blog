<style type="text/css">

  a { text-decoration: none; }
  a:hover { text-decoration: none; }

  #content article {
    border-bottom: 1px solid #ddd;
    border-top: 1px solid #fff;
    padding: 10px 0 30px;
    position: relative;
  }

  article .title {
    color: #333;
    font-size: 2em;
    font-weight: 300;
    line-height: 35px;
  }

  article h1.title a {
    color: #333;
    transition: color .3s;
  }

  article .post-meta span {
    margin-right: 10px
  }

  article .post-meta .date {
    color: #333
  }
  
  article .post-meta .comments a {
    color: #999
  }

  article .entry-content {
    color: #444;
    font-size: 16px;
    font-family: Arial,'Hiragino Sans GB',冬青黑,'Microsoft YaHei',微软雅黑,SimSun,宋体,Helvetica,Tahoma,'Arial sans-serif';
    line-height: 1.8;
    word-wrap: break-word;
  }

  .page-navi {
    text-align: right;
    border-top: 1px solid #fff;
    border-bottom: 1px solid #ddd;
    line-height: 20px;
    overflow: hidden;
    padding: 20px 0;
    position: relative;
    width: 100%;
  }

  .blog-set a { margin-right: 15px }

  .create-blog { 
    padding: 0 0 10px; 
    border-bottom: 1px dashed #ddd; 
    text-align: right;
    margin-top: 25px
  }

</style>

<div id="content" class="inner">

  <div class="create-blog">
    <a type="button" class="btn btn-success" href="/article">（创建）Blog</a>
  </div>

  {{ range $index, $article := .articles }}
    <article>
      <h1 class="title"><a href="#">{{ $article.Title }}</a></h1>
      <p class="post-meta">
        <span class="date"><time>May 14, 2017</time></span>
        <span class="category"><a href="#">{{ $article.Keywords }}</a></span>
        <span class="comments"><a href="#">1 Comments</a></span>
      </p>
      <div class="entry-content">
        <p>{{ $article.Summary }}</p>
      </div>

      <div class="blog-set">
        <a type="button" class="btn btn-info" href="/article/{{ $article.Id }}">（查看）Blog</a>
        <a type="button" class="btn btn-warning" href="#">（编辑）Blog</a>
        <a type="button" class="btn btn-danger" href="#">（删除）Blog</a>
      </div>

    </article>  
  {{ end }}

  {{ template "layout/paginator.tpl" .}}
</div>


