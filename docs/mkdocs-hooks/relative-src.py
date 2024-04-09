import re


def on_page_content(html: str, *, page, **_):
    if page.file.src_uri == "index.md":
        html = html.replace('data-is-relative="true" src="./docs/', 'data-is-relative="true" src="')
        return html
