<html>
    <head>
        <title>计算字符串的3DES ECB加密或者解密</title>
    </head>
    <body>
    <h4>计算一个字符串的3des加密或者解密结果</h4>
        <form method="POST">
        <div>
            <label>需要加密或者解密的字符串:</label>
            <textarea name="values"></textarea>
        </div>
            <div>
                <label>密钥(长度为24)<label>
                <input type="text" name="key">
            <div>
            <div>
                <label>加密或者解密<label>
                <select name="type">
                    <option value="1">加密</option>
                    <option value="0">解密</option>
                </select>
            <div>
            <div>
                <input type="submit" value="提交">            
            </div>
        </form>
    </body>
</html>