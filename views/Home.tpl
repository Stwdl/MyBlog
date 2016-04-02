<!doctype html>
<html class="no-js">

	<head>
		<meta charset="utf-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="description" content="">
		<meta name="keywords" content="">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<title>醉翁二郎的猪窝</title>

		<!-- Set render engine for 360 browser -->
		<meta name="renderer" content="webkit">

		<!-- No Baidu Siteapp-->
		<meta http-equiv="Cache-Control" content="no-siteapp" />

		<link rel="icon" type="image/png" href="assets/i/favicon.png">

		<!-- Add to homescreen for Chrome on Android -->
		<meta name="mobile-web-app-capable" content="yes">
		<link rel="icon" sizes="192x192" href="../static/assets/i/app-icon72x72@2x.png">

		<!-- Add to homescreen for Safari on iOS -->
		<meta name="apple-mobile-web-app-capable" content="yes">
		<meta name="apple-mobile-web-app-status-bar-style" content="black">
		<meta name="apple-mobile-web-app-title" content="醉翁二郎" />
		<link rel="apple-touch-icon-precomposed" href="../static/assets/i/app-icon72x72@2x.png">

		<!-- Tile icon for Win8 (144x144 + tile color) -->
		<meta name="msapplication-TileImage" content="assets/i/app-icon72x72@2x.png">
		<meta name="msapplication-TileColor" content="#0e90d2">

		<link rel="stylesheet" href="../static/assets/css/amazeui.min.css">
		<link rel="stylesheet" href="../static/assets/css/app.css">
		<link rel="stylesheet" href="../static/assets/css/home.css">
	</head>

	<body>
		<header class="am-topbar am-topbar-inverse am-topbar-fixed-top" style="background-color: #222222;">
			<div class="am-container">
				<h1 class="am-topbar-brand">
      <a href="http://user.qzone.qq.com/405093360" class="am-text-ir">醉翁二郎</a>
    </h1>

				<button class="am-topbar-btn am-topbar-toggle am-btn am-btn-sm am-btn-success am-show-sm-only" data-am-collapse="{target: '#doc-topbar-collapse-4'}"><span class="am-sr-only">导航切换</span> <span class="am-icon-bars"></span></button>

				<div class="am-collapse am-topbar-collapse" id="doc-topbar-collapse-4">
					<ul class="am-nav am-nav-pills am-topbar-nav">
						<li class="am-active"><a href="/">主页</a></li>
						<li><a href="Article">文章</a></li>
						<li><a href="Photo">相册</a></li>
						<li><a href="Music">音乐</a></li>
						<li><a href="About">About Me</a></li>
						<li class="am-dropdown" data-am-dropdown>
							<a class="am-dropdown-toggle" data-am-dropdown-toggle href="javascript:;">
            其他 <span class="am-icon-caret-down"></span>
          </a>
							<ul class="am-dropdown-content">
								<li><a href="Message">给我留言</a></li>
								

							</ul>
						</li>
					</ul>
				</div>
			</div>
		</header>
		<div class="amz-banner">
			<div class="amz-container am-scrollspy-init am-scrollspy-inview am-animation-scale-up" data-am-scrollspy="{animation: 'scale-up', repeat: false}">
				<div class="am-g">
					<div class="am-u-sm-12">
						<a href="javascript:;">
							<img src="../static/img/logo.jpg" style="width:150px;height:150px">
							
						</a>
						<div style="position: absolute;line-height:1;top: 0px;left: 200px;">
						<h1>醉翁二郎</h1>
						<p>是父是夫是子 程序员量化交易爱好者 爱生活爱自由爱编程</p>
						<p>除去这些身份，仅仅是</p>
						<strong>一个孤独、碌碌无为的求索者</strong>
						
						</div>
					</div>
					
				</div>

			</div>
		</div>
		<div class="am-g am-g-fixed" style="padding-top: 10px;">
			<div class="imgslider">
				<div class="am-slider am-slider-default am-slider-carousel" data-am-flexslider="{itemWidth: 250, itemMargin: 5, slideshow: true}">
					<ul class="am-slides">
					{{range $key,$val :=.images}}
						<li><img src="{{$val.Url}}" /></li>
					{{end}}
						
					</ul>
				</div>
			</div>

		</div>
		<div class="am-g am-g-fixed" style="padding-top: 10px;">
			<div class="am-u-sm-3" style="padding: 0px;">
				<div class="am-panel am-panel-default">
					<div class="am-panel-hd">
						<h3 class="am-panel-title">最新文章</h3>
					</div>
					<div class="am-panel-bd">
						{{.TianQi}}
					</div>

					<ul class="am-list am-list-border">
					{{range $key,$val :=.articles}}
					<li><a href="../Article/{{$val.ID}}"><i class="am-icon-pencil am-icon-fw"></i>{{$val.Title}}</a></li>
					{{end}}
					</ul>

					<div class="am-panel-footer"><a href="../Article/"><i class="am-icon-pencil am-icon-fw"></i>更多...</a></div>
				</div>
			</div>
			<div class="am-u-sm-9" style="padding-right:0px;">
				<div class="am-panel-group" id="accordion">
				{{range $key,$val:=.DTS}}
  					<div class="am-panel am-panel-default">
   				 		<div class="am-panel-hd">
     			 			<h4 class="am-panel-title" data-am-collapse="{parent: '#accordion', target: '#do-not-say-{{add $key 1}}'}">
        					{{if eq $val.Type 1}}
								<a href="../Article/{{$val.ID}}" style="color:#0e90d2">{{$val.Title}}</a>
								{{else}}
								心情随笔
								{{end}}
								<span class="am-badge am-badge-success" style="position:absolute;right:5px">{{date $val.CreateTime "Y-m-d H:i:s"}}</span>
      			 			</h4>
    					</div>
    					<div id="do-not-say-{{add $key 1}}" class="am-panel-collapse am-collapse am-in">
      						<div class="am-panel-bd">
								{{toHtml $val.Content 200}}
      						</div>
    					</div>
  					</div>
				</div>
						
				{{end}}
 


			</div>
		</div>
		<div class="am-g am-g-fixed" style="padding-top: 10px;">
			<hr>
			<h1 id="comment-ping-lun-zu-jian">访客留言： <a href="../Message" style="float:right;font-size:20px;color:#ccc" >Shoe More...</a></h1>
			
			<ul class="am-comments-list am-comments-list-flip">
			{{range $key,$val:=.MSG}}
			{{if eq $val.TypeID 1}}
				<li class="am-comment">
					<a href="#link-to-user-home"><img src="../static/img/anymous.jpg" alt="" class="am-comment-avatar" width="48" height="48"></a>
					<div class="am-comment-main">
						<header class="am-comment-hd">
							<div class="am-comment-meta"><a href="#link-to-user" class="am-comment-author">{{$val.Name}}</a> 评论于
								<time datetime="2013-07-27T04:54:29-07:00" title="2013年7月27日 下午7:54 格林尼治标准时间+0800">{{date $val.Createtime "Y-m-d H:i:s"}}</time>
							</div>
						</header>
						<div class="am-comment-bd">
							{{str2html $val.Content}}
						</div>
					</div>
				</li>
				{{else}}
				
				<li class="am-comment am-comment-flip am-comment-secondary">
					<a href="#link-to-user-home"><img src="../static/assets/i/7.gif" alt="" class="am-comment-avatar" width="48" height="48"></a>
					<div class="am-comment-main">
						<header class="am-comment-hd">
							<div class="am-comment-meta"><a href="#link-to-user" class="am-comment-author">{{$val.Name}}</a> 评论于
								<time datetime="2013-07-27T04:54:29-07:00" title="2013年7月27日 下午7:54 格林尼治标准时间+0800">{{date $val.Createtime "Y-m-d H:i:s"}}</time>
							</div>
						</header>
						<div class="am-comment-bd">
						{{str2html $val.Content}}
						</div>
					</div>
				</li>
				{{end}}
			{{end}}
			</ul>
		</div>

		<!--[if (gte IE 9)|!(IE)]><!-->
		<script src="../static/assets/js/jquery.min.js"></script>
		<!--<![endif]-->
		<!--[if lte IE 8 ]>
<script src="http://libs.baidu.com/jquery/1.11.3/jquery.min.js"></script>
<script src="http://cdn.staticfile.org/modernizr/2.8.3/modernizr.js"></script>
<script src="assets/js/amazeui.ie8polyfill.min.js"></script>
<![endif]-->
		<script src="../static/assets/js/amazeui.min.js"></script>
	</body>

</html>