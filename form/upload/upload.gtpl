<html>
<head>
    <title>上传文件</title>
</head>
<body>
<form action="/upload" method="post" enctype="multipart/form-data">
    <input type="file" name="uploadfile">
    <input type="hidden" name="token" value="{{.}}">
    <input type="submit" value="upload">
</form>
</body>
</html>