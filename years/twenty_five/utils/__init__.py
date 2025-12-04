class Logger:
    def __init__(self, is_debug: bool):
        self.is_debug = is_debug

    def print(self, message: str):
        if not self.is_debug:
            return

        print(message)
