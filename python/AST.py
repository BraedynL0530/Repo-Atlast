import tree_sitter_language_pack as ts
from dataclasses import dataclass

@dataclass
class response:
    language: str
    dependencies: list
    total_lines: int
    code_lines: int
    comment_lines: int
    blank_lines: int
    node_count: int
    error_count: int
    max_depth: int
    imports: list
    exports: list
    symbols: list


def detect_language(file_path):
    if not file_path:
        return None
    return ts.detect_language_from_path(file_path)

def ast_from_file(file_path:str) -> response | None:
    try:
        with open(file_path, "r", encoding="utf-8") as f:
            source = f.read()

        lang = detect_language(file_path)
        config = ts.ProcessConfig(
            language=lang,
            diagnostics=True,  # off by default, turn on for error detection
            symbols=True,  # off by default, useful for dead code later
        )
        result = ts.process(source, config)
        return result
    except Exception:
        return None


