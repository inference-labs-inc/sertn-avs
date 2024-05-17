import os
from torch.autograd import Variable
import torch

# Makes sure the files will end up in the same place relative to the base python path
def relative_file_path (file_path : str) :
    return os.path.join(os.path.dirname(__file__), file_path)

def input_to_field_element(input, scale = 2):
    return int(round(input * (2**scale), 0))

def field_element_to_input(input, scale = 2):
    return input * 2**(-scale)

def parse_input(input, scale = 2):
    # parse a given string of inputs
    if len(input) == 1:
        formatted_input = Variable(torch.Tensor([field_element_to_input(float(i), scale) for i in input[0].split(" ")]))
    else:
        formatted_input = Variable(torch.Tensor([field_element_to_input(float(i), scale) for i in input]))

    return formatted_input
