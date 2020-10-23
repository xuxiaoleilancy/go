<!DOCTYPE html>

<html>
<head>
  <title>RWorld</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="shortcut icon" href="/static/img/R.png" type="image/x-icon" />

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
      background-color: #fff;
    }

    header,
    footer {
      width: auto;
      margin-left: auto;
      margin-right: auto;
    }

    .logo {
      background-image: url(/static/img/R.png);
      background-repeat: no-repeat;
      -webkit-background-size: 100px 100px;
      background-size: 100px 100px;
      background-position: center center;
      text-align: center;
      font-size: 42px;
      padding: 250px 0 70px;
      font-weight: normal;
      text-shadow: 0px 1px 2px #ddd;
    }

    header {
      padding: 20px 0;
    }

    footer {
      line-height: 1.8;
      text-align: center;
      padding: 50px 0;
      color: #999;
    }

    .description {
      text-align: center;
      font-size: 16px;
    }

    a {
      color: #444;
      text-decoration: none;
    }

    .backdrop {
      position: absolute;
      width: 100%;
      height: 100%;
      box-shadow: inset 0px 0px 100px #ddd;
      z-index: -1;
      top: 0px;
      left: 0px;
    }
  </style>
</head>

<body>
  <header>
    <h1 class="logo">欢迎来到RWorld</h1>
    <div class="description">
	我的英文名字叫做“RWorld",中文的意思就是“锐世界”.<br /><br />
      我是帮助您进行软件测试、部门沟通、技术文档发布的一系列工具的集合.<br /><br />
	
    </div>
  </header>
  <header>
    <h2 class="logo">
      <a href="http://{{<.Website>}}/rautomatic">RAutomatic</a> </p></h2>
  </header>
<header>
    <h2 class="logo">
      <a href="http://{{<.Website>}}/todo">TODO</a> </p></h2>
  </header>

<header>
    <h2 class="logo">
      <a href="http://{{<.Website>}}/rtcvideo">RTCVideo</a> </p></h2>
  </header>

<header>
    <h2 class="logo">
      <a href="http://{{<.Website>}}/rtest">RTest</a> </p></h2>
  </header>

<header>
    <h2 class="logo">
      <a href="http://{{<.Website>}}/rdoc">RDoc</a> </p></h2>
  </header>

<div class="author">
      
    </div>
  <footer>
    <div class="author">
      您可以通过这个地址访问:
      <a href="http://{{<.Website>}}">http://{{<.Website>}}</a> </p>
      联系我:
      <a class="email" href="mailto:{{<.Email>}}">{{<.Email>}}</a>
    </div>
  </footer>
  <div class="backdrop"></div>

  <script src="/static/js/reload.min.js"></script>
</body>
</html>
