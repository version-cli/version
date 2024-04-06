def on_page_markdown(markdown, *, config, **_):
    return markdown.replace("./docs/", config.docs_dir)
