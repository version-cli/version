import os
import requests

def on_page_markdown(markdown, **kwargs):
    if "version-cli/action@v" in markdown:
        try:
            # Check if VERSION_CLI_VERSION env var is set
            latest_version = os.environ.get("VERSION_CLI_VERSION")
            
            # Replace any version-cli/action@v... with the latest version
            import re
            markdown = re.sub(r'version-cli/action@v[0-9.]+', f'version-cli/action@{latest_version}', markdown)
        except Exception as e:
            print(f"Warning: Could not determine latest version: {e}")
    return markdown
