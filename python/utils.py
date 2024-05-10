import os
from torch.autograd import Variable
import torch

# Makes sure the files will end up in the same place relative to the base python path
def relative_file_path (file_path : str) :
    return os.path.join(os.path.dirname(__file__), file_path)

def parse_input(input):
    # parse a given string of inputs
    if len(input) == 1:
        formatted_input = Variable(torch.Tensor([int(i) for i in input[0].split(" ")]))
    else:
        formatted_input = Variable(torch.Tensor([int(i) for i in input]))

    return formatted_input
