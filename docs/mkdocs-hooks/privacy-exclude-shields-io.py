import logging
import os
import re
from urllib.parse import urlparse

import material.plugins.privacy.plugin as privacy
from unittest.mock import patch

from urllib.parse import ParseResult as URL
from mkdocs.structure.files import File
from colorama import Fore, Style

# Set up logging
log = logging.getLogger("mkdocs.material.privacy")


def _is_excluded(self, url: URL, initiator: File | None = None):
    if not self._is_external(url):
        return True

    # Skip if external assets must not be processed
    if not self.config.assets:
        return True

    # If initiator is given, format for printing
    via = ""
    if initiator:
        via = "".join([
            Fore.WHITE, Style.DIM,
            f"in '{initiator.src_uri}' ",
            Style.RESET_ALL
        ])

    # Print warning if fetching is not enabled
    if not self.config.assets_fetch:
        log.warning(f"External file: {url.geturl()} {via}")
        return True

    if "img.shields.io" in url.geturl():
        return True

    # File is not excluded
    return False


if os.getenv("EXCLUDE_SHIELDS_IO_PRIVACY") == "true":
    # Only exclude shields.io from the privacy plugin if desired (i.e. for the "main" version of docs)
    privacy.PrivacyPlugin._is_excluded = _is_excluded
