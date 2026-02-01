<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">
  <xsl:output method="html" encoding="UTF-8"/>
  <xsl:template match="/">
    <html>
      <head>
        <title><xsl:value-of select="rss/channel/title"/> - RSS Feed</title>
        <style>
          * { box-sizing: border-box; margin: 0; padding: 0; }
          body {
            font-family: Georgia, serif;
            background: #fafaf9;
            color: #1c1917;
            line-height: 1.6;
            padding: 2rem;
            max-width: 42rem;
            margin: 0 auto;
          }
          @media (prefers-color-scheme: dark) {
            body { background: #0c0a09; color: #fafaf9; }
            a { color: #fca5a5; }
            .item { border-color: #292524; }
          }
          h1 { font-size: 1.5rem; margin-bottom: 0.5rem; }
          .meta { color: #57534e; font-size: 0.9rem; margin-bottom: 2rem; }
          .subscribe {
            display: inline-block;
            background: #991b1b;
            color: white;
            padding: 0.5rem 1rem;
            border-radius: 4px;
            text-decoration: none;
            font-size: 0.875rem;
            margin-bottom: 2rem;
          }
          .item {
            padding: 1rem 0;
            border-bottom: 1px solid #e7e5e4;
          }
          .item:last-child { border-bottom: none; }
          .item h2 { font-size: 1.1rem; margin-bottom: 0.25rem; }
          .item h2 a { color: inherit; text-decoration: none; }
          .item h2 a:hover { color: #991b1b; }
          .item time { font-size: 0.8rem; color: #57534e; }
          .item p { margin-top: 0.5rem; font-size: 0.95rem; color: #57534e; }
        </style>
      </head>
      <body>
        <h1><xsl:value-of select="rss/channel/title"/></h1>
        <p class="meta">
          This is an RSS feed. Copy the URL into your feed reader to subscribe.
        </p>
        <a class="subscribe" href="{rss/channel/link}">← Back to site</a>
        <div class="items">
          <xsl:for-each select="rss/channel/item">
            <div class="item">
              <h2><a href="{link}"><xsl:value-of select="title"/></a></h2>
              <time><xsl:value-of select="pubDate"/></time>
              <p><xsl:value-of select="description"/></p>
            </div>
          </xsl:for-each>
        </div>
      </body>
    </html>
  </xsl:template>
</xsl:stylesheet>
