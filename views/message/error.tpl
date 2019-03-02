
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width,initial-scale=1">
  <title>400 | Bad request.</title>

  <style type="text/css">
    body {
      padding: 30px 20px;
      font-family: -apple-system, BlinkMacSystemFont,
        "Segoe UI", "Roboto", "Oxygen", "Ubuntu", "Cantarell",
        "Fira Sans", "Droid Sans", "Helvetica Neue", sans-serif;
      color: #727272;
      line-height: 1.6;
    }

    .container {
      max-width: 500px;
      margin: 0 auto;
    }

    h1 {
      margin: 0 0 40px;
      font-size: 60px;
      line-height: 1;
      color: #252427;
      font-weight: 700;
    }

    h2 {
      margin: 100px 0 0;
      font-size: 20px;
      font-weight: 600;
      letter-spacing: 0.1em;
      color: #A299AC;
      text-transform: uppercase;
    }

    p {
      font-size: 16px;
      margin: 1em 0;
    }

    .go-back a {
      display: inline-block;
      margin-top: 3em;
      padding: 10px;
      color: #1B1A1A;
      font-weight: 700;
      border: solid 2px #e7e7e7;
      text-decoration: none;
      font-size: 16px;
      text-transform: uppercase;
      letter-spacing: 0.1em;
    }

    .go-back a:hover {
      border-color: #1B1A1A;
    }

    @media screen and (min-width: 768px) {
      body {
        padding: 50px;
      }
    }

    @media screen and (max-width: 480px) {
      h1 {
        font-size: 48px;
      }
    }
  </style>
</head>
<body>

<div class="container">
  <a href="/" class="brand"><img src="/static/img/golang.jpg" width="180" height="100"></a>

  <h1>{{.Data.Title}}</h1>

  <p>{{.Data.Msg}}</p>

  <span class="go-back"><a href="{{.Data.BackUrl}}">Go back</a></span>
</div>

</body>
</html>