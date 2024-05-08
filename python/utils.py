import os

# Makes sure the files will end up in the same place relative to the base python path
def relative_file_path (file_path : str) :
    return os.path.join(os.path.dirname(__file__), file_path)
