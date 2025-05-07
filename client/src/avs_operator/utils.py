from torch.autograd import Variable
import torch


def field_element_to_input(input, scale=2):
    return input * 2 ** (-scale)


def parse_input(input, scale=2):
    # parse a given string of inputs
    if len(input) == 1:
        formatted_input = Variable(
            torch.Tensor(
                [field_element_to_input(float(i), scale) for i in input[0].split(" ")]
            )
        )
    else:
        formatted_input = Variable(
            torch.Tensor([field_element_to_input(float(i), scale) for i in input])
        )

    return formatted_input
