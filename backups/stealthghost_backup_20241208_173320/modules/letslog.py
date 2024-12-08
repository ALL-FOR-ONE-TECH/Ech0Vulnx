import logging

class ErrorLogger:
    """
    The `ErrorLogger` class sets up a logger that captures error-level log messages 
    (and higher severity levels) and writes them to a file named `errorlog_heapdump.txt`.
    The logger uses Python's built-in `logging` module and includes timestamps, logger name, 
    log level, and the log message for each entry.
    """

    def __init__(self):
        pass
    def error_logger(self):
        self.logger = logging.getLogger(__name__)
        self.logger.setLevel(logging.ERROR)
        file_handler = logging.FileHandler('errorlog_heapdump.txt')
        file_handler.setLevel(logging.ERROR)
        formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s')
        file_handler.setFormatter(formatter)
        self.logger.addHandler(file_handler)
        self.logger.propagate = False
        return self.logger
    