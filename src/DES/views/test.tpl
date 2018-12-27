<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Document</title>
	<link rel="stylesheet" type="text/css" href="/static/css/common.css">
	<style type="text/css">
.main{width: 788px;height:420px;margin: 0 auto;padding:30px 45px 30px 15px;}
h1{font-size:25px;color:#330033;height:100px;
	margin-left: 235px;}
header{height:100px;position: relative;line-height: 100px;margin-bottom: 30px;}
header p{font-size:20px;position: absolute;left: 0;top:0px;height:100px; }
header p span{font-size:40px;color:#ff0033;font-family:'华文行楷';}
.put{height:32px;position: relative;line-height: 32px;}
.put p{font-size:15px;color:#999999;float: left;margin-right:10px;}
.put input{margin-right: 8px;float: left;border-radius:5px;}
.put input:nth-of-type(1){width: 123px;border: 1px #999999 solid;height: 25px;color: #999999;margin-top: 4px;}
.put input:nth-of-type(2){width: 70px;height: 32px;background-color: #ff0033;color:white;border:none;}
.put input:nth-of-type(3){width: 70px;height: 32px;background-color: #00ccff;color:white;border:none;}
.wrap{height:210px;margin-top: 10px;position: relative;}
.left,.right{width:350px;background-color:#eee;margin-right: 10px;border-radius:10px;float: left;height:210px;border: none;font-size: 25px;color:#ccc;line-height:50px;position:relative;}
.left input{background-color: white;width:100px;height:40px;font-size:15px;position: absolute;left:10px;bottom:10px;border: none;}
	</style>
</head>
<body>
	<div class="main">
		<header>
			<h1>一个免费即时的在线加密解密软件</h1>
			<p><span>快捷</span>加密解密</p>
		</header>
		<div class="put clearfix">
			<p>密钥</p>
			<input placeholder="请输入8位密钥" name="Key">
			<input value="加密" type="submit">
			<input value="解密" type="submit">
		</div>
		<div class="wrap">
			<div class="left">请输入你要加密/解密的文字
			<input type="button" value="上传文档" name="image">
			</div>
			<div class="right"></div>
		</div>
	</div>
</body>
</html>